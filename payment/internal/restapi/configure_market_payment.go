// This file is safe to edit. Once it exists it will not be overwritten

package restapi

import (
	"crypto/tls"
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"

	"github.com/h4x4d/crypto-market/payment/internal/restapi/operations"
)

//go:generate swagger generate server --target ../../internal --name MarketPayment --spec ../../api/swagger/payment.yaml --principal interface{}

func configureFlags(api *operations.MarketPaymentAPI) {
	// api.CommandLineOptionsGroups = []swag.CommandLineOptionsGroup{ ... }
}

func configureAPI(api *operations.MarketPaymentAPI) http.Handler {
	// configure the api here
	api.ServeError = errors.ServeError

	// Set your custom logger if needed. Default one is log.Printf
	// Expected interface func(string, ...interface{})
	//
	// Example:
	// api.Logger = log.Printf

	api.UseSwaggerUI()
	// To continue using redoc as your UI, uncomment the following line
	// api.UseRedoc()

	api.JSONConsumer = runtime.JSONConsumer()

	api.JSONProducer = runtime.JSONProducer()

	// Applies when the "api_key" header is set
	if api.APIKeyAuth == nil {
		api.APIKeyAuth = func(token string) (interface{}, error) {
			return nil, errors.NotImplemented("api key auth (api_key) api_key from header param [api_key] has not yet been implemented")
		}
	}

	// Set your custom authorizer if needed. Default one is security.Authorized()
	// Expected interface runtime.Authorizer
	//
	// Example:
	// api.APIAuthorizer = security.Authorized()

	if api.GetAccountBalanceHandler == nil {
		api.GetAccountBalanceHandler = operations.GetAccountBalanceHandlerFunc(func(params operations.GetAccountBalanceParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation operations.GetAccountBalance has not yet been implemented")
		})
	}
	if api.GetMetricsHandler == nil {
		api.GetMetricsHandler = operations.GetMetricsHandlerFunc(func(params operations.GetMetricsParams) middleware.Responder {
			return middleware.NotImplemented("operation operations.GetMetrics has not yet been implemented")
		})
	}
	if api.GetTransactionsPurchaseHandler == nil {
		api.GetTransactionsPurchaseHandler = operations.GetTransactionsPurchaseHandlerFunc(func(params operations.GetTransactionsPurchaseParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation operations.GetTransactionsPurchase has not yet been implemented")
		})
	}
	if api.GetTransactionsTransfersHandler == nil {
		api.GetTransactionsTransfersHandler = operations.GetTransactionsTransfersHandlerFunc(func(params operations.GetTransactionsTransfersParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation operations.GetTransactionsTransfers has not yet been implemented")
		})
	}
	if api.PostTransactionsWithdrawHandler == nil {
		api.PostTransactionsWithdrawHandler = operations.PostTransactionsWithdrawHandlerFunc(func(params operations.PostTransactionsWithdrawParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation operations.PostTransactionsWithdraw has not yet been implemented")
		})
	}

	api.PreServerShutdown = func() {}

	api.ServerShutdown = func() {}

	return setupGlobalMiddleware(api.Serve(setupMiddlewares))
}

// The TLS configuration before HTTPS server starts.
func configureTLS(tlsConfig *tls.Config) {
	// Make all necessary changes to the TLS configuration here.
}

// As soon as server is initialized but not run yet, this function will be called.
// If you need to modify a config, store server instance to stop it individually later, this is the place.
// This function can be called multiple times, depending on the number of serving schemes.
// scheme value will be set accordingly: "http", "https" or "unix".
func configureServer(s *http.Server, scheme, addr string) {
}

// The middleware configuration is for the handler executors. These do not apply to the swagger.json document.
// The middleware executes after routing but before authentication, binding and validation.
func setupMiddlewares(handler http.Handler) http.Handler {
	return handler
}

// The middleware configuration happens before anything, this middleware also applies to serving the swagger.json document.
// So this is a good place to plug in a panic handling middleware, logging and metrics.
func setupGlobalMiddleware(handler http.Handler) http.Handler {
	return handler
}
