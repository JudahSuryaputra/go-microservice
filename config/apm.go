package config

type ApmConfiguration struct {
	// existing fields...
	ApmServiceName string `mapstructure:"go-microservice"`
	ApmServerUrl   string `mapstructure:"apm_server_url"`
	ApmSecretToken string `mapstructure:"apm_secret_token"`
	ApmEnvironment string `mapstructure:"apm_environment"`
}
