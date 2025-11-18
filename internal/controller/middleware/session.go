package middleware

import (
	"context"
	"errors"
	"fmt"
	"go-microservice/config"
	"os"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/redis/go-redis/v9"
)

// SessionManager provides JWT creation/validation and Redis-backed revocation.
// Usage:
//   rdb := redis.NewClient(&redis.Options{Addr: os.Getenv("REDIS_ADDR")})
//   jwtSecret := os.Getenv("JWT_SECRET")
//   sm := NewSessionManager(rdb, []byte(jwtSecret))
//   e.Use(sm.Middleware())
// Then handlers can read claims with ctx.Get("claims") or user id via ctx.Get("user_id").

var (
	ErrNoTokenProvided = errors.New("no token provided")
)

// SessionManager holds dependencies for token operations.
type SessionManager struct {
	redis     *redis.Client
	secret    []byte
	issuer    string
	ttl       time.Duration // token TTL when creating tokens
	clockSkew time.Duration
}

// NewSessionManager creates a manager with sensible defaults. Configure via env vars
func NewSessionManager(rdb *redis.Client, config config.Configuration) *SessionManager {
	issuer := os.Getenv("JWT_ISSUER")
	if issuer == "" {
		issuer = "go-microservice"
	}
	ttlMin := 60 // default 60 minutes
	if s := os.Getenv(config.JwtTtl); s != "" {
		if v, err := strconv.Atoi(s); err == nil {
			ttlMin = v
		}
	}
	return &SessionManager{
		redis:     rdb,
		secret:    []byte(config.JwtSecret),
		issuer:    issuer,
		ttl:       time.Duration(ttlMin) * time.Minute,
		clockSkew: 30 * time.Second,
	}
}

// CustomClaims extends registered claims with application fields.
type CustomClaims struct {
	UserID string   `json:"user_id"`
	Roles  []string `json:"roles,omitempty"`
	jwt.RegisteredClaims
}

// CreateToken creates a signed JWT and returns token string and its expiry.
func (s *SessionManager) CreateToken(userID string, roles []string) (string, time.Time, error) {
	now := time.Now().UTC()
	exp := now.Add(s.ttl)
	jti := uuid.New().String()
	claims := CustomClaims{
		UserID: userID,
		Roles:  roles,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    s.issuer,
			Subject:   userID,
			ID:        jti,
			IssuedAt:  jwt.NewNumericDate(now),
			NotBefore: jwt.NewNumericDate(now),
			ExpiresAt: jwt.NewNumericDate(exp),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signed, err := token.SignedString(s.secret)
	if err != nil {
		return "", time.Time{}, err
	}
	// Optionally store allow-list in redis (here we store jti -> 1 with expiry = token TTL)
	ctx := context.Background()
	_ = s.redis.Set(ctx, s.revokedKey(jti), "0", s.ttl).Err() // value 0 means active
	return signed, exp, nil
}

// ParseAndValidate parses a token string and returns claims if valid and not revoked.
func (s *SessionManager) ParseAndValidate(tokenStr string) (*CustomClaims, error) {
	parser := jwt.NewParser(jwt.WithLeeway(s.clockSkew))
	var claims CustomClaims
	token, err := parser.ParseWithClaims(tokenStr, &claims, func(t *jwt.Token) (interface{}, error) {
		// ensure signing method
		if t.Method.Alg() != jwt.SigningMethodHS256.Alg() {
			return nil, fmt.Errorf("unexpected signing method: %v", t.Header["alg"])
		}
		return s.secret, nil
	})
	if err != nil {
		return nil, err
	}
	if !token.Valid {
		return nil, errors.New("invalid token")
	}
	// check revocation in Redis (if jti absent or value != "0" considered revoked)
	ctx := context.Background()
	if claims.ID == "" {
		return nil, errors.New("token missing jti")
	}
	val, err := s.redis.Get(ctx, s.revokedKey(claims.ID)).Result()
	if err == redis.Nil {
		// absent -> treat as revoked (safer), or choose to allow. We treat absent as revoked.
		return nil, errors.New("token revoked or unknown")
	} else if err != nil {
		return nil, err
	}
	if val != "0" {
		return nil, errors.New("token revoked")
	}
	return &claims, nil
}

// RevokeToken marks a token (by jti) as revoked in Redis. If expiry is 0, will use the remaining token TTL.
func (s *SessionManager) RevokeToken(jti string, expiry time.Duration) error {
	ctx := context.Background()
	if expiry <= 0 {
		expiry = s.ttl
	}
	// set to 1 indicating revoked
	return s.redis.Set(ctx, s.revokedKey(jti), "1", expiry).Err()
}

func (s *SessionManager) revokedKey(jti string) string {
	return fmt.Sprintf("jwt:jti:%s", jti)
}

// Middleware returns an echo middleware that validates JWT and attaches claims to the context.
// If a route is unauthenticated, you can skip the middleware or set it to optional by allowing
// a query param or header. This middleware requires token in Authorization: Bearer <token>.
func (s *SessionManager) Middleware() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			// Extract token from header
			auth := c.Request().Header.Get("Authorization")
			if auth == "" {
				return echo.NewHTTPError(401, ErrNoTokenProvided.Error())
			}
			var tokenStr string
			_, err := fmt.Sscanf(auth, "Bearer %s", &tokenStr)
			if err != nil || tokenStr == "" {
				return echo.NewHTTPError(401, "invalid authorization header")
			}
			claims, err := s.ParseAndValidate(tokenStr)
			if err != nil {
				return echo.NewHTTPError(401, err.Error())
			}
			// Attach to context for handlers
			c.Set("claims", claims)
			c.Set("user_id", claims.UserID)
			c.Set("jti", claims.ID)
			return next(c)
		}
	}
}

// Helper to retrieve claims from echo.Context
func GetClaims(c echo.Context) (*CustomClaims, bool) {
	v := c.Get("claims")
	if v == nil {
		return nil, false
	}
	cl, ok := v.(*CustomClaims)
	return cl, ok
}

// Helper to revoke current request token. Returns error if none attached.
func RevokeCurrentToken(c echo.Context, s *SessionManager) error {
	jtiIfc := c.Get("jti")
	if jtiIfc == nil {
		return errors.New("no jti in context")
	}
	jti, ok := jtiIfc.(string)
	if !ok || jti == "" {
		return errors.New("invalid jti in context")
	}
	// revoke using remaining TTL derived from token expiry if available
	claimsIfc := c.Get("claims")
	if claimsIfc == nil {
		return s.RevokeToken(jti, 0)
	}
	claims, ok := claimsIfc.(*CustomClaims)
	if !ok || claims.ExpiresAt == nil {
		return s.RevokeToken(jti, 0)
	}
	remain := time.Until(claims.ExpiresAt.Time)
	if remain < time.Second {
		remain = time.Second
	}
	return s.RevokeToken(jti, remain)
}
