// This file is safe to edit. Once it exists it will not be overwritten

package restapi

import (
	"crypto/tls"
	"fmt"
	"github.com/h4x4d/crypto-market/main/internal/restapi/handlers"
	"github.com/h4x4d/crypto-market/pkg/client"
	"log"
	"net/http"
	"os"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"

	"github.com/h4x4d/crypto-market/main/internal/models"
	"github.com/h4x4d/crypto-market/main/internal/restapi/operations"
)

//go:generate swagger generate server --target ../../internal --name MarketMain --spec ../../api/swagger/main.yaml --principal models.User --exclude-main

func configureFlags(api *operations.MarketMainAPI) {
	// api.CommandLineOptionsGroups = []swag.CommandLineOptionsGroup{ ... }
}

func configureAPI(api *operations.MarketMainAPI) http.Handler {
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
	api.TxtProducer = runtime.TextProducer()

	// Applies when the "api_key" header is set
	manager, err := client.NewClient()
	if err != nil {
		log.Fatal(err)
	}
	// Applies when the "api_key" header is set
	api.APIKeyAuth = func(token string) (*models.User, error) {
		user, err := manager.CheckToken(token)
		if err != nil {
			return nil, err
		}
		return (*models.User)(user), nil
	}

	connStr := fmt.Sprintf("postgres://%s:%s@%s:%s/%s", os.Getenv("POSTGRES_USER"),
		os.Getenv("POSTGRES_PASSWORD"), "db", os.Getenv("POSTGRES_PORT"), os.Getenv("HOTEL_DB_NAME"))
	handler, makeErr := handlers.NewHandler(connStr)
	for makeErr != nil {
		handler, makeErr = handlers.NewHandler(connStr)
	}

	api.GetAccountBalanceHandler = operations.GetAccountBalanceHandlerFunc(handler.GetBalanceHandler)

	if api.GetMetricsHandler == nil {
		api.GetMetricsHandler = operations.GetMetricsHandlerFunc(func(params operations.GetMetricsParams) middleware.Responder {
			return middleware.NotImplemented("operation operations.GetMetrics has not yet been implemented")
		})
	}
	if api.GetTransactionsPurchaseHandler == nil {
		api.GetTransactionsPurchaseHandler = operations.GetTransactionsPurchaseHandlerFunc(func(params operations.GetTransactionsPurchaseParams, principal *models.User) middleware.Responder {
			return middleware.NotImplemented("operation operations.GetTransactionsPurchase has not yet been implemented")
		})
	}
	if api.GetTransactionsTransfersHandler == nil {
		api.GetTransactionsTransfersHandler = operations.GetTransactionsTransfersHandlerFunc(func(params operations.GetTransactionsTransfersParams, principal *models.User) middleware.Responder {
			return middleware.NotImplemented("operation operations.GetTransactionsTransfers has not yet been implemented")
		})
	}
	if api.PostTransactionsDepositHandler == nil {
		api.PostTransactionsDepositHandler = operations.PostTransactionsDepositHandlerFunc(func(params operations.PostTransactionsDepositParams, principal *models.User) middleware.Responder {
			return middleware.NotImplemented("operation operations.PostTransactionsDeposit has not yet been implemented")
		})
	}
	if api.PostTransactionsWithdrawHandler == nil {
		api.PostTransactionsWithdrawHandler = operations.PostTransactionsWithdrawHandlerFunc(func(params operations.PostTransactionsWithdrawParams, principal *models.User) middleware.Responder {
			return middleware.NotImplemented("operation operations.PostTransactionsWithdraw has not yet been implemented")
		})
	}
	if api.CancelBidHandler == nil {
		api.CancelBidHandler = operations.CancelBidHandlerFunc(func(params operations.CancelBidParams, principal *models.User) middleware.Responder {
			return middleware.NotImplemented("operation operations.CancelBid has not yet been implemented")
		})
	}
	if api.CreateBidHandler == nil {
		api.CreateBidHandler = operations.CreateBidHandlerFunc(func(params operations.CreateBidParams, principal *models.User) middleware.Responder {
			return middleware.NotImplemented("operation operations.CreateBid has not yet been implemented")
		})
	}
	if api.GetBidByIDHandler == nil {
		api.GetBidByIDHandler = operations.GetBidByIDHandlerFunc(func(params operations.GetBidByIDParams, principal *models.User) middleware.Responder {
			return middleware.NotImplemented("operation operations.GetBidByID has not yet been implemented")
		})
	}
	if api.GetBidsHandler == nil {
		api.GetBidsHandler = operations.GetBidsHandlerFunc(func(params operations.GetBidsParams, principal *models.User) middleware.Responder {
			return middleware.NotImplemented("operation operations.GetBids has not yet been implemented")
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
