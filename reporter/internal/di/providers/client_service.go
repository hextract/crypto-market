package providers

import "reporter/internal/infrastructure/adapters/delivery/web/client/operations"

func ClientServiceProvider(config MatchingEngineConfig) operations.ClientService {
	return operations.NewClientWithBasicAuth(
		config.Host,
		config.BasePath,
		config.Scheme,
		config.User,
		config.Password,
	)
}
