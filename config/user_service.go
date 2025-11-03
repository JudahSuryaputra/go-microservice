package config

type ServiceConfig struct {
	UserServiceAddress  string `envconfig:"USER_SERVICE_ADDRESS" default:":50051"`
	OrderServiceAddress string `envconfig:"ORDER_SERVICE_ADDRESS" default:":50052"`
}
