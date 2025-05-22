package providers

type MatchingEngineConfig struct {
	Host     string
	BasePath string
	Scheme   string
	User     string
	Password string
}

func MatchingEngineConfigProvider() MatchingEngineConfig {
	config := MatchingEngineConfig{}
	config.Host = "0.0.0.0"
	config.BasePath = "/"
	config.Scheme = "http"
	config.User = "user"
	config.Password = "password"
	return config
}
