// This file is safe to edit. Once it exists it will not be overwritten

package restapi

import (
	"crypto/tls"
	"github.com/h4x4d/crypto-market/auth/internal/restapi/handlers"
	"github.com/rs/cors"
	"log"
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"

	"github.com/h4x4d/crypto-market/auth/internal/restapi/operations"
)

//go:generate swagger generate server --target ../../internal --name MarketAuth --spec ../../api/swagger/auth.yaml --principal interface{} --exclude-main

func configureFlags(api *operations.MarketAuthAPI) {
	// api.CommandLineOptionsGroups = []swag.CommandLineOptionsGroup{ ... }
}

func configureAPI(api *operations.MarketAuthAPI) http.Handler {
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

	handler, err := handlers.NewHandler()
	if err != nil {
		log.Fatalln(err)
	}

	api.PostAuthLoginHandler = operations.PostAuthLoginHandlerFunc(handler.LoginHandler)
	api.PostAuthRegisterHandler = operations.PostAuthRegisterHandlerFunc(handler.RegisterHandler)
	api.PostAuthChangePasswordHandler = operations.PostAuthChangePasswordHandlerFunc(handler.ChangePasswordHandler)
	api.GetMetricsHandler = operations.GetMetricsHandlerFunc(handlers.MetricsHandler)

	if api.GetMetricsHandler == nil {
		api.GetMetricsHandler = operations.GetMetricsHandlerFunc(func(params operations.GetMetricsParams) middleware.Responder {
			return middleware.NotImplemented("operation operations.GetMetrics has not yet been implemented")
		})
	}
	if api.PostAuthChangePasswordHandler == nil {
		api.PostAuthChangePasswordHandler = operations.PostAuthChangePasswordHandlerFunc(func(params operations.PostAuthChangePasswordParams) middleware.Responder {
			return middleware.NotImplemented("operation operations.PostAuthChangePassword has not yet been implemented")
		})
	}
	if api.PostAuthLoginHandler == nil {
		api.PostAuthLoginHandler = operations.PostAuthLoginHandlerFunc(func(params operations.PostAuthLoginParams) middleware.Responder {
			return middleware.NotImplemented("operation operations.PostAuthLogin has not yet been implemented")
		})
	}
	if api.PostAuthRegisterHandler == nil {
		api.PostAuthRegisterHandler = operations.PostAuthRegisterHandlerFunc(func(params operations.PostAuthRegisterParams) middleware.Responder {
			return middleware.NotImplemented("operation operations.PostAuthRegister has not yet been implemented")
		})
	}
	if api.PostAuthValidateTokenHandler == nil {
		api.PostAuthValidateTokenHandler = operations.PostAuthValidateTokenHandlerFunc(func(params operations.PostAuthValidateTokenParams) middleware.Responder {
			return middleware.NotImplemented("operation operations.PostAuthValidateToken has not yet been implemented")
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
	// Настройки CORS
	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"*"},
		AllowCredentials: true,
		Debug:            false,
	})

	return c.Handler(handler)
}
