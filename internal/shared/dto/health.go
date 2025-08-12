package dto

type HealthCheckResponse struct {
	AppStatus   string `json:"app_status"`
	RedisStatus string `json:"redis_status"`
	DbStatus    string `json:"db_status"`
}
