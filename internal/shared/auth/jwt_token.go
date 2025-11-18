package auth

import (
	"fmt"
	"strconv"
	"time"

	"go-microservice/config"

	"github.com/golang-jwt/jwt/v5"
)

// Session holds minimal information about the user and token times.
type Session struct {
	UserId   string `json:"sub" validate:"required"`
	Expired  int64  `json:"exp"`
	IssuedAt int64  `json:"iat"`
}

// AuthTokenPayload represents access and refresh tokens returned to client.
type AuthTokenPayload struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

// Valid implements jwt.Claims when needed.
func (ss *Session) Valid() error { return nil }

// IsSessionExpired checks if the session Expired timestamp is in the past.
func (ss *Session) IsSessionExpired() error {
	if time.Now().After(time.Unix(ss.Expired, 0)) {
		return fmt.Errorf("session expired")
	}
	return nil
}

// GenerateAuthToken creates HS256 signed access and refresh tokens using JwtSecret and JwtTtl (minutes).
func GenerateAuthToken(ss *Session, c *config.Configuration) (AuthTokenPayload, error) {
	var resp AuthTokenPayload
	if ss == nil || ss.UserId == "" {
		return resp, fmt.Errorf("invalid session: missing user id")
	}

	// Parse TTL from config (minutes).
	tlMin, err := strconv.Atoi(c.JwtTtl)
	if err != nil || tlMin <= 0 {
		tlMin = 60
	}
	secret := []byte(c.JwtSecret)
	now := time.Now().UTC()

	// Access token
	accessExp := now.Add(time.Duration(tlMin) * time.Minute)
	accessClaims := jwt.RegisteredClaims{
		Subject:   ss.UserId,
		IssuedAt:  jwt.NewNumericDate(now),
		NotBefore: jwt.NewNumericDate(now),
		ExpiresAt: jwt.NewNumericDate(accessExp),
	}
	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, accessClaims)
	signedAccess, err := accessToken.SignedString(secret)
	if err != nil {
		return resp, err
	}

	// Refresh token (2x ttl of access)
	refreshExp := now.Add(2 * time.Duration(tlMin) * time.Minute)
	refreshClaims := jwt.RegisteredClaims{
		Subject:   ss.UserId,
		IssuedAt:  jwt.NewNumericDate(now),
		NotBefore: jwt.NewNumericDate(now),
		ExpiresAt: jwt.NewNumericDate(refreshExp),
	}
	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshClaims)
	signedRefresh, err := refreshToken.SignedString(secret)
	if err != nil {
		return resp, err
	}

	resp.AccessToken = signedAccess
	resp.RefreshToken = signedRefresh
	ss.IssuedAt = now.Unix()
	ss.Expired = accessExp.Unix()
	return resp, nil
}

// ParseToken parses and validates a token string with the provided secret.
func ParseToken(tokenString string, secret string) (*jwt.Token, error) {
	parser := jwt.NewParser()
	return parser.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if token.Method.Alg() != jwt.SigningMethodHS256.Alg() {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(secret), nil
	})
}
