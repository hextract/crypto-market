// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/loads"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/runtime/security"
	"github.com/go-openapi/spec"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// NewMarketMainAPI creates a new MarketMain instance
func NewMarketMainAPI(spec *loads.Document) *MarketMainAPI {
	return &MarketMainAPI{
		handlers:            make(map[string]map[string]http.Handler),
		formats:             strfmt.Default,
		defaultConsumes:     "application/json",
		defaultProduces:     "application/json",
		customConsumers:     make(map[string]runtime.Consumer),
		customProducers:     make(map[string]runtime.Producer),
		PreServerShutdown:   func() {},
		ServerShutdown:      func() {},
		spec:                spec,
		useSwaggerUI:        false,
		ServeError:          errors.ServeError,
		BasicAuthenticator:  security.BasicAuth,
		APIKeyAuthenticator: security.APIKeyAuth,
		BearerAuthenticator: security.BearerAuth,

		JSONConsumer: runtime.JSONConsumer(),

		JSONProducer: runtime.JSONProducer(),
		TxtProducer:  runtime.TextProducer(),

		GetAccountBalanceHandler: GetAccountBalanceHandlerFunc(func(params GetAccountBalanceParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation GetAccountBalance has not yet been implemented")
		}),
		GetMetricsHandler: GetMetricsHandlerFunc(func(params GetMetricsParams) middleware.Responder {
			return middleware.NotImplemented("operation GetMetrics has not yet been implemented")
		}),
		GetTransactionsPurchaseHandler: GetTransactionsPurchaseHandlerFunc(func(params GetTransactionsPurchaseParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation GetTransactionsPurchase has not yet been implemented")
		}),
		GetTransactionsTransfersHandler: GetTransactionsTransfersHandlerFunc(func(params GetTransactionsTransfersParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation GetTransactionsTransfers has not yet been implemented")
		}),
		PostTransactionsDepositHandler: PostTransactionsDepositHandlerFunc(func(params PostTransactionsDepositParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation PostTransactionsDeposit has not yet been implemented")
		}),
		PostTransactionsWithdrawHandler: PostTransactionsWithdrawHandlerFunc(func(params PostTransactionsWithdrawParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation PostTransactionsWithdraw has not yet been implemented")
		}),
		CancelBidHandler: CancelBidHandlerFunc(func(params CancelBidParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation CancelBid has not yet been implemented")
		}),
		CreateBidHandler: CreateBidHandlerFunc(func(params CreateBidParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation CreateBid has not yet been implemented")
		}),
		GetBidByIDHandler: GetBidByIDHandlerFunc(func(params GetBidByIDParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation GetBidByID has not yet been implemented")
		}),
		GetBidsHandler: GetBidsHandlerFunc(func(params GetBidsParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation GetBids has not yet been implemented")
		}),

		// Applies when the "api_key" header is set
		APIKeyAuth: func(token string) (interface{}, error) {
			return nil, errors.NotImplemented("api key auth (api_key) api_key from header param [api_key] has not yet been implemented")
		},
		// default authorizer is authorized meaning no requests are blocked
		APIAuthorizer: security.Authorized(),
	}
}

/*MarketMainAPI Continuous market API for cryptocurrency trading and account management */
type MarketMainAPI struct {
	spec            *loads.Document
	context         *middleware.Context
	handlers        map[string]map[string]http.Handler
	formats         strfmt.Registry
	customConsumers map[string]runtime.Consumer
	customProducers map[string]runtime.Producer
	defaultConsumes string
	defaultProduces string
	Middleware      func(middleware.Builder) http.Handler
	useSwaggerUI    bool

	// BasicAuthenticator generates a runtime.Authenticator from the supplied basic auth function.
	// It has a default implementation in the security package, however you can replace it for your particular usage.
	BasicAuthenticator func(security.UserPassAuthentication) runtime.Authenticator

	// APIKeyAuthenticator generates a runtime.Authenticator from the supplied token auth function.
	// It has a default implementation in the security package, however you can replace it for your particular usage.
	APIKeyAuthenticator func(string, string, security.TokenAuthentication) runtime.Authenticator

	// BearerAuthenticator generates a runtime.Authenticator from the supplied bearer token auth function.
	// It has a default implementation in the security package, however you can replace it for your particular usage.
	BearerAuthenticator func(string, security.ScopedTokenAuthentication) runtime.Authenticator

	// JSONConsumer registers a consumer for the following mime types:
	//   - application/json
	JSONConsumer runtime.Consumer

	// JSONProducer registers a producer for the following mime types:
	//   - application/json
	JSONProducer runtime.Producer
	// TxtProducer registers a producer for the following mime types:
	//   - text/plain
	TxtProducer runtime.Producer

	// APIKeyAuth registers a function that takes a token and returns a principal
	// it performs authentication based on an api key api_key provided in the header
	APIKeyAuth func(string) (interface{}, error)

	// APIAuthorizer provides access control (ACL/RBAC/ABAC) by providing access to the request and authenticated principal
	APIAuthorizer runtime.Authorizer

	// GetAccountBalanceHandler sets the operation handler for the get account balance operation
	GetAccountBalanceHandler GetAccountBalanceHandler
	// GetMetricsHandler sets the operation handler for the get metrics operation
	GetMetricsHandler GetMetricsHandler
	// GetTransactionsPurchaseHandler sets the operation handler for the get transactions purchase operation
	GetTransactionsPurchaseHandler GetTransactionsPurchaseHandler
	// GetTransactionsTransfersHandler sets the operation handler for the get transactions transfers operation
	GetTransactionsTransfersHandler GetTransactionsTransfersHandler
	// PostTransactionsDepositHandler sets the operation handler for the post transactions deposit operation
	PostTransactionsDepositHandler PostTransactionsDepositHandler
	// PostTransactionsWithdrawHandler sets the operation handler for the post transactions withdraw operation
	PostTransactionsWithdrawHandler PostTransactionsWithdrawHandler
	// CancelBidHandler sets the operation handler for the cancel bid operation
	CancelBidHandler CancelBidHandler
	// CreateBidHandler sets the operation handler for the create bid operation
	CreateBidHandler CreateBidHandler
	// GetBidByIDHandler sets the operation handler for the get bid by id operation
	GetBidByIDHandler GetBidByIDHandler
	// GetBidsHandler sets the operation handler for the get bids operation
	GetBidsHandler GetBidsHandler

	// ServeError is called when an error is received, there is a default handler
	// but you can set your own with this
	ServeError func(http.ResponseWriter, *http.Request, error)

	// PreServerShutdown is called before the HTTP(S) server is shutdown
	// This allows for custom functions to get executed before the HTTP(S) server stops accepting traffic
	PreServerShutdown func()

	// ServerShutdown is called when the HTTP(S) server is shut down and done
	// handling all active connections and does not accept connections any more
	ServerShutdown func()

	// Custom command line argument groups with their descriptions
	CommandLineOptionsGroups []swag.CommandLineOptionsGroup

	// User defined logger function.
	Logger func(string, ...interface{})
}

// UseRedoc for documentation at /docs
func (o *MarketMainAPI) UseRedoc() {
	o.useSwaggerUI = false
}

// UseSwaggerUI for documentation at /docs
func (o *MarketMainAPI) UseSwaggerUI() {
	o.useSwaggerUI = true
}

// SetDefaultProduces sets the default produces media type
func (o *MarketMainAPI) SetDefaultProduces(mediaType string) {
	o.defaultProduces = mediaType
}

// SetDefaultConsumes returns the default consumes media type
func (o *MarketMainAPI) SetDefaultConsumes(mediaType string) {
	o.defaultConsumes = mediaType
}

// SetSpec sets a spec that will be served for the clients.
func (o *MarketMainAPI) SetSpec(spec *loads.Document) {
	o.spec = spec
}

// DefaultProduces returns the default produces media type
func (o *MarketMainAPI) DefaultProduces() string {
	return o.defaultProduces
}

// DefaultConsumes returns the default consumes media type
func (o *MarketMainAPI) DefaultConsumes() string {
	return o.defaultConsumes
}

// Formats returns the registered string formats
func (o *MarketMainAPI) Formats() strfmt.Registry {
	return o.formats
}

// RegisterFormat registers a custom format validator
func (o *MarketMainAPI) RegisterFormat(name string, format strfmt.Format, validator strfmt.Validator) {
	o.formats.Add(name, format, validator)
}

// Validate validates the registrations in the MarketMainAPI
func (o *MarketMainAPI) Validate() error {
	var unregistered []string

	if o.JSONConsumer == nil {
		unregistered = append(unregistered, "JSONConsumer")
	}

	if o.JSONProducer == nil {
		unregistered = append(unregistered, "JSONProducer")
	}
	if o.TxtProducer == nil {
		unregistered = append(unregistered, "TxtProducer")
	}

	if o.APIKeyAuth == nil {
		unregistered = append(unregistered, "APIKeyAuth")
	}

	if o.GetAccountBalanceHandler == nil {
		unregistered = append(unregistered, "GetAccountBalanceHandler")
	}
	if o.GetMetricsHandler == nil {
		unregistered = append(unregistered, "GetMetricsHandler")
	}
	if o.GetTransactionsPurchaseHandler == nil {
		unregistered = append(unregistered, "GetTransactionsPurchaseHandler")
	}
	if o.GetTransactionsTransfersHandler == nil {
		unregistered = append(unregistered, "GetTransactionsTransfersHandler")
	}
	if o.PostTransactionsDepositHandler == nil {
		unregistered = append(unregistered, "PostTransactionsDepositHandler")
	}
	if o.PostTransactionsWithdrawHandler == nil {
		unregistered = append(unregistered, "PostTransactionsWithdrawHandler")
	}
	if o.CancelBidHandler == nil {
		unregistered = append(unregistered, "CancelBidHandler")
	}
	if o.CreateBidHandler == nil {
		unregistered = append(unregistered, "CreateBidHandler")
	}
	if o.GetBidByIDHandler == nil {
		unregistered = append(unregistered, "GetBidByIDHandler")
	}
	if o.GetBidsHandler == nil {
		unregistered = append(unregistered, "GetBidsHandler")
	}

	if len(unregistered) > 0 {
		return fmt.Errorf("missing registration: %s", strings.Join(unregistered, ", "))
	}

	return nil
}

// ServeErrorFor gets a error handler for a given operation id
func (o *MarketMainAPI) ServeErrorFor(operationID string) func(http.ResponseWriter, *http.Request, error) {
	return o.ServeError
}

// AuthenticatorsFor gets the authenticators for the specified security schemes
func (o *MarketMainAPI) AuthenticatorsFor(schemes map[string]spec.SecurityScheme) map[string]runtime.Authenticator {
	result := make(map[string]runtime.Authenticator)
	for name := range schemes {
		switch name {
		case "api_key":
			scheme := schemes[name]
			result[name] = o.APIKeyAuthenticator(scheme.Name, scheme.In, o.APIKeyAuth)

		}
	}
	return result
}

// Authorizer returns the registered authorizer
func (o *MarketMainAPI) Authorizer() runtime.Authorizer {
	return o.APIAuthorizer
}

// ConsumersFor gets the consumers for the specified media types.
// MIME type parameters are ignored here.
func (o *MarketMainAPI) ConsumersFor(mediaTypes []string) map[string]runtime.Consumer {
	result := make(map[string]runtime.Consumer, len(mediaTypes))
	for _, mt := range mediaTypes {
		switch mt {
		case "application/json":
			result["application/json"] = o.JSONConsumer
		}

		if c, ok := o.customConsumers[mt]; ok {
			result[mt] = c
		}
	}
	return result
}

// ProducersFor gets the producers for the specified media types.
// MIME type parameters are ignored here.
func (o *MarketMainAPI) ProducersFor(mediaTypes []string) map[string]runtime.Producer {
	result := make(map[string]runtime.Producer, len(mediaTypes))
	for _, mt := range mediaTypes {
		switch mt {
		case "application/json":
			result["application/json"] = o.JSONProducer
		case "text/plain":
			result["text/plain"] = o.TxtProducer
		}

		if p, ok := o.customProducers[mt]; ok {
			result[mt] = p
		}
	}
	return result
}

// HandlerFor gets a http.Handler for the provided operation method and path
func (o *MarketMainAPI) HandlerFor(method, path string) (http.Handler, bool) {
	if o.handlers == nil {
		return nil, false
	}
	um := strings.ToUpper(method)
	if _, ok := o.handlers[um]; !ok {
		return nil, false
	}
	if path == "/" {
		path = ""
	}
	h, ok := o.handlers[um][path]
	return h, ok
}

// Context returns the middleware context for the market main API
func (o *MarketMainAPI) Context() *middleware.Context {
	if o.context == nil {
		o.context = middleware.NewRoutableContext(o.spec, o, nil)
	}

	return o.context
}

func (o *MarketMainAPI) initHandlerCache() {
	o.Context() // don't care about the result, just that the initialization happened
	if o.handlers == nil {
		o.handlers = make(map[string]map[string]http.Handler)
	}

	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/account/balance"] = NewGetAccountBalance(o.context, o.GetAccountBalanceHandler)
	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/metrics"] = NewGetMetrics(o.context, o.GetMetricsHandler)
	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/transactions/purchase"] = NewGetTransactionsPurchase(o.context, o.GetTransactionsPurchaseHandler)
	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/transactions/transfers"] = NewGetTransactionsTransfers(o.context, o.GetTransactionsTransfersHandler)
	if o.handlers["POST"] == nil {
		o.handlers["POST"] = make(map[string]http.Handler)
	}
	o.handlers["POST"]["/transactions/deposit"] = NewPostTransactionsDeposit(o.context, o.PostTransactionsDepositHandler)
	if o.handlers["POST"] == nil {
		o.handlers["POST"] = make(map[string]http.Handler)
	}
	o.handlers["POST"]["/transactions/withdraw"] = NewPostTransactionsWithdraw(o.context, o.PostTransactionsWithdrawHandler)
	if o.handlers["DELETE"] == nil {
		o.handlers["DELETE"] = make(map[string]http.Handler)
	}
	o.handlers["DELETE"]["/market/{bid_id}"] = NewCancelBid(o.context, o.CancelBidHandler)
	if o.handlers["POST"] == nil {
		o.handlers["POST"] = make(map[string]http.Handler)
	}
	o.handlers["POST"]["/bid"] = NewCreateBid(o.context, o.CreateBidHandler)
	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/market/{bid_id}"] = NewGetBidByID(o.context, o.GetBidByIDHandler)
	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/bid"] = NewGetBids(o.context, o.GetBidsHandler)
}

// Serve creates a http handler to serve the API over HTTP
// can be used directly in http.ListenAndServe(":8000", api.Serve(nil))
func (o *MarketMainAPI) Serve(builder middleware.Builder) http.Handler {
	o.Init()

	if o.Middleware != nil {
		return o.Middleware(builder)
	}
	if o.useSwaggerUI {
		return o.context.APIHandlerSwaggerUI(builder)
	}
	return o.context.APIHandler(builder)
}

// Init allows you to just initialize the handler cache, you can then recompose the middleware as you see fit
func (o *MarketMainAPI) Init() {
	if len(o.handlers) == 0 {
		o.initHandlerCache()
	}
}

// RegisterConsumer allows you to add (or override) a consumer for a media type.
func (o *MarketMainAPI) RegisterConsumer(mediaType string, consumer runtime.Consumer) {
	o.customConsumers[mediaType] = consumer
}

// RegisterProducer allows you to add (or override) a producer for a media type.
func (o *MarketMainAPI) RegisterProducer(mediaType string, producer runtime.Producer) {
	o.customProducers[mediaType] = producer
}

// AddMiddlewareFor adds a http middleware to existing handler
func (o *MarketMainAPI) AddMiddlewareFor(method, path string, builder middleware.Builder) {
	um := strings.ToUpper(method)
	if path == "/" {
		path = ""
	}
	o.Init()
	if h, ok := o.handlers[um][path]; ok {
		o.handlers[um][path] = builder(h)
	}
}
