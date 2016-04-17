package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"net/http"
	"strings"

	httpkit "github.com/go-swagger/go-swagger/httpkit"
	middleware "github.com/go-swagger/go-swagger/httpkit/middleware"
	security "github.com/go-swagger/go-swagger/httpkit/security"
	loads "github.com/go-swagger/go-swagger/loads"
	spec "github.com/go-swagger/go-swagger/spec"
	strfmt "github.com/go-swagger/go-swagger/strfmt"
	"github.com/go-swagger/go-swagger/swag"

	"github.com/go-swagger/go-swagger/examples/generated/restapi/operations/pet"
	"github.com/go-swagger/go-swagger/examples/generated/restapi/operations/store"
	"github.com/go-swagger/go-swagger/examples/generated/restapi/operations/user"
)

// NewPetstoreAPI creates a new Petstore instance
func NewPetstoreAPI(spec *loads.Document) *PetstoreAPI {
	o := &PetstoreAPI{
		spec:            spec,
		handlers:        make(map[string]map[string]http.Handler),
		formats:         strfmt.Default,
		defaultConsumes: "application/json",
		defaultProduces: "application/json",
		ServerShutdown:  func() {},
	}

	return o
}

/*PetstoreAPI This is a sample server Petstore server.

[Learn about Swagger](http://swagger.wordnik.com) or join the IRC channel '#swagger' on irc.freenode.net.

For this sample, you can use the api key 'special-key' to test the authorization filters
*/
type PetstoreAPI struct {
	spec            *loads.Document
	context         *middleware.Context
	handlers        map[string]map[string]http.Handler
	formats         strfmt.Registry
	defaultConsumes string
	defaultProduces string
	// XMLConsumer registers a consumer for a "application/xml" mime type
	XMLConsumer httpkit.Consumer
	// UrlformConsumer registers a consumer for a "application/x-www-form-urlencoded" mime type
	UrlformConsumer httpkit.Consumer
	// JSONConsumer registers a consumer for a "application/json" mime type
	JSONConsumer httpkit.Consumer

	// XMLProducer registers a producer for a "application/xml" mime type
	XMLProducer httpkit.Producer
	// JSONProducer registers a producer for a "application/json" mime type
	JSONProducer httpkit.Producer

	// APIKeyAuth registers a function that takes a token and returns a principal
	// it performs authentication based on an api key api_key provided in the header
	APIKeyAuth func(string) (interface{}, error)

	// PetAddPetHandler sets the operation handler for the add pet operation
	PetAddPetHandler pet.AddPetHandler
	// UserCreateUserHandler sets the operation handler for the create user operation
	UserCreateUserHandler user.CreateUserHandler
	// UserCreateUsersWithArrayInputHandler sets the operation handler for the create users with array input operation
	UserCreateUsersWithArrayInputHandler user.CreateUsersWithArrayInputHandler
	// UserCreateUsersWithListInputHandler sets the operation handler for the create users with list input operation
	UserCreateUsersWithListInputHandler user.CreateUsersWithListInputHandler
	// StoreDeleteOrderHandler sets the operation handler for the delete order operation
	StoreDeleteOrderHandler store.DeleteOrderHandler
	// PetDeletePetHandler sets the operation handler for the delete pet operation
	PetDeletePetHandler pet.DeletePetHandler
	// UserDeleteUserHandler sets the operation handler for the delete user operation
	UserDeleteUserHandler user.DeleteUserHandler
	// PetFindPetsByStatusHandler sets the operation handler for the find pets by status operation
	PetFindPetsByStatusHandler pet.FindPetsByStatusHandler
	// PetFindPetsByTagsHandler sets the operation handler for the find pets by tags operation
	PetFindPetsByTagsHandler pet.FindPetsByTagsHandler
	// StoreGetOrderByIDHandler sets the operation handler for the get order by Id operation
	StoreGetOrderByIDHandler store.GetOrderByIDHandler
	// PetGetPetByIDHandler sets the operation handler for the get pet by Id operation
	PetGetPetByIDHandler pet.GetPetByIDHandler
	// UserGetUserByNameHandler sets the operation handler for the get user by name operation
	UserGetUserByNameHandler user.GetUserByNameHandler
	// UserLoginUserHandler sets the operation handler for the login user operation
	UserLoginUserHandler user.LoginUserHandler
	// UserLogoutUserHandler sets the operation handler for the logout user operation
	UserLogoutUserHandler user.LogoutUserHandler
	// StorePlaceOrderHandler sets the operation handler for the place order operation
	StorePlaceOrderHandler store.PlaceOrderHandler
	// PetUpdatePetHandler sets the operation handler for the update pet operation
	PetUpdatePetHandler pet.UpdatePetHandler
	// PetUpdatePetWithFormHandler sets the operation handler for the update pet with form operation
	PetUpdatePetWithFormHandler pet.UpdatePetWithFormHandler
	// UserUpdateUserHandler sets the operation handler for the update user operation
	UserUpdateUserHandler user.UpdateUserHandler

	// ServeError is called when an error is received, there is a default handler
	// but you can set your own with this
	ServeError func(http.ResponseWriter, *http.Request, error)

	// ServerShutdown is called when the HTTP(S) server is shut down and done
	// handling all active connections and does not accept connections any more
	ServerShutdown func()

	// Custom command line argument groups with their descriptions
	CommandLineOptionsGroups []swag.CommandLineOptionsGroup
}

// SetDefaultProduces sets the default produces media type
func (o *PetstoreAPI) SetDefaultProduces(mediaType string) {
	o.defaultProduces = mediaType
}

// SetDefaultConsumes returns the default consumes media type
func (o *PetstoreAPI) SetDefaultConsumes(mediaType string) {
	o.defaultConsumes = mediaType
}

// DefaultProduces returns the default produces media type
func (o *PetstoreAPI) DefaultProduces() string {
	return o.defaultProduces
}

// DefaultConsumes returns the default consumes media type
func (o *PetstoreAPI) DefaultConsumes() string {
	return o.defaultConsumes
}

// Formats returns the registered string formats
func (o *PetstoreAPI) Formats() strfmt.Registry {
	return o.formats
}

// RegisterFormat registers a custom format validator
func (o *PetstoreAPI) RegisterFormat(name string, format strfmt.Format, validator strfmt.Validator) {
	o.formats.Add(name, format, validator)
}

// Validate validates the registrations in the PetstoreAPI
func (o *PetstoreAPI) Validate() error {
	var unregistered []string

	if o.XMLConsumer == nil {
		unregistered = append(unregistered, "XMLConsumer")
	}

	if o.UrlformConsumer == nil {
		unregistered = append(unregistered, "UrlformConsumer")
	}

	if o.JSONConsumer == nil {
		unregistered = append(unregistered, "JSONConsumer")
	}

	if o.XMLProducer == nil {
		unregistered = append(unregistered, "XMLProducer")
	}

	if o.JSONProducer == nil {
		unregistered = append(unregistered, "JSONProducer")
	}

	if o.APIKeyAuth == nil {
		unregistered = append(unregistered, "APIKeyAuth")
	}

	if o.PetAddPetHandler == nil {
		unregistered = append(unregistered, "pet.AddPetHandler")
	}

	if o.UserCreateUserHandler == nil {
		unregistered = append(unregistered, "user.CreateUserHandler")
	}

	if o.UserCreateUsersWithArrayInputHandler == nil {
		unregistered = append(unregistered, "user.CreateUsersWithArrayInputHandler")
	}

	if o.UserCreateUsersWithListInputHandler == nil {
		unregistered = append(unregistered, "user.CreateUsersWithListInputHandler")
	}

	if o.StoreDeleteOrderHandler == nil {
		unregistered = append(unregistered, "store.DeleteOrderHandler")
	}

	if o.PetDeletePetHandler == nil {
		unregistered = append(unregistered, "pet.DeletePetHandler")
	}

	if o.UserDeleteUserHandler == nil {
		unregistered = append(unregistered, "user.DeleteUserHandler")
	}

	if o.PetFindPetsByStatusHandler == nil {
		unregistered = append(unregistered, "pet.FindPetsByStatusHandler")
	}

	if o.PetFindPetsByTagsHandler == nil {
		unregistered = append(unregistered, "pet.FindPetsByTagsHandler")
	}

	if o.StoreGetOrderByIDHandler == nil {
		unregistered = append(unregistered, "store.GetOrderByIDHandler")
	}

	if o.PetGetPetByIDHandler == nil {
		unregistered = append(unregistered, "pet.GetPetByIDHandler")
	}

	if o.UserGetUserByNameHandler == nil {
		unregistered = append(unregistered, "user.GetUserByNameHandler")
	}

	if o.UserLoginUserHandler == nil {
		unregistered = append(unregistered, "user.LoginUserHandler")
	}

	if o.UserLogoutUserHandler == nil {
		unregistered = append(unregistered, "user.LogoutUserHandler")
	}

	if o.StorePlaceOrderHandler == nil {
		unregistered = append(unregistered, "store.PlaceOrderHandler")
	}

	if o.PetUpdatePetHandler == nil {
		unregistered = append(unregistered, "pet.UpdatePetHandler")
	}

	if o.PetUpdatePetWithFormHandler == nil {
		unregistered = append(unregistered, "pet.UpdatePetWithFormHandler")
	}

	if o.UserUpdateUserHandler == nil {
		unregistered = append(unregistered, "user.UpdateUserHandler")
	}

	if len(unregistered) > 0 {
		return fmt.Errorf("missing registration: %s", strings.Join(unregistered, ", "))
	}

	return nil
}

// ServeErrorFor gets a error handler for a given operation id
func (o *PetstoreAPI) ServeErrorFor(operationID string) func(http.ResponseWriter, *http.Request, error) {
	return o.ServeError
}

// AuthenticatorsFor gets the authenticators for the specified security schemes
func (o *PetstoreAPI) AuthenticatorsFor(schemes map[string]spec.SecurityScheme) map[string]httpkit.Authenticator {

	result := make(map[string]httpkit.Authenticator)
	for name, scheme := range schemes {
		switch name {

		case "api_key":

			result[name] = security.APIKeyAuth(scheme.Name, scheme.In, func(tok string) (interface{}, error) { return o.APIKeyAuth(tok) })

		}
	}
	return result

}

// ConsumersFor gets the consumers for the specified media types
func (o *PetstoreAPI) ConsumersFor(mediaTypes []string) map[string]httpkit.Consumer {

	result := make(map[string]httpkit.Consumer)
	for _, mt := range mediaTypes {
		switch mt {

		case "application/xml":
			result["application/xml"] = o.XMLConsumer

		case "application/x-www-form-urlencoded":
			result["application/x-www-form-urlencoded"] = o.UrlformConsumer

		case "application/json":
			result["application/json"] = o.JSONConsumer

		}
	}
	return result

}

// ProducersFor gets the producers for the specified media types
func (o *PetstoreAPI) ProducersFor(mediaTypes []string) map[string]httpkit.Producer {

	result := make(map[string]httpkit.Producer)
	for _, mt := range mediaTypes {
		switch mt {

		case "application/xml":
			result["application/xml"] = o.XMLProducer

		case "application/json":
			result["application/json"] = o.JSONProducer

		}
	}
	return result

}

// HandlerFor gets a http.Handler for the provided operation method and path
func (o *PetstoreAPI) HandlerFor(method, path string) (http.Handler, bool) {
	if o.handlers == nil {
		return nil, false
	}
	um := strings.ToUpper(method)
	if _, ok := o.handlers[um]; !ok {
		return nil, false
	}
	h, ok := o.handlers[um][path]
	return h, ok
}

func (o *PetstoreAPI) initHandlerCache() {
	if o.context == nil {
		o.context = middleware.NewRoutableContext(o.spec, o, nil)
	}

	if o.handlers == nil {
		o.handlers = make(map[string]map[string]http.Handler)
	}

	if o.handlers["POST"] == nil {
		o.handlers[strings.ToUpper("POST")] = make(map[string]http.Handler)
	}
	o.handlers["POST"]["/pets"] = pet.NewAddPet(o.context, o.PetAddPetHandler)

	if o.handlers["POST"] == nil {
		o.handlers[strings.ToUpper("POST")] = make(map[string]http.Handler)
	}
	o.handlers["POST"]["/users"] = user.NewCreateUser(o.context, o.UserCreateUserHandler)

	if o.handlers["POST"] == nil {
		o.handlers[strings.ToUpper("POST")] = make(map[string]http.Handler)
	}
	o.handlers["POST"]["/users/createWithArray"] = user.NewCreateUsersWithArrayInput(o.context, o.UserCreateUsersWithArrayInputHandler)

	if o.handlers["POST"] == nil {
		o.handlers[strings.ToUpper("POST")] = make(map[string]http.Handler)
	}
	o.handlers["POST"]["/users/createWithList"] = user.NewCreateUsersWithListInput(o.context, o.UserCreateUsersWithListInputHandler)

	if o.handlers["DELETE"] == nil {
		o.handlers[strings.ToUpper("DELETE")] = make(map[string]http.Handler)
	}
	o.handlers["DELETE"]["/stores/order/{orderId}"] = store.NewDeleteOrder(o.context, o.StoreDeleteOrderHandler)

	if o.handlers["DELETE"] == nil {
		o.handlers[strings.ToUpper("DELETE")] = make(map[string]http.Handler)
	}
	o.handlers["DELETE"]["/pets/{petId}"] = pet.NewDeletePet(o.context, o.PetDeletePetHandler)

	if o.handlers["DELETE"] == nil {
		o.handlers[strings.ToUpper("DELETE")] = make(map[string]http.Handler)
	}
	o.handlers["DELETE"]["/users/{username}"] = user.NewDeleteUser(o.context, o.UserDeleteUserHandler)

	if o.handlers["GET"] == nil {
		o.handlers[strings.ToUpper("GET")] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/pets/findByStatus"] = pet.NewFindPetsByStatus(o.context, o.PetFindPetsByStatusHandler)

	if o.handlers["GET"] == nil {
		o.handlers[strings.ToUpper("GET")] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/pets/findByTags"] = pet.NewFindPetsByTags(o.context, o.PetFindPetsByTagsHandler)

	if o.handlers["GET"] == nil {
		o.handlers[strings.ToUpper("GET")] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/stores/order/{orderId}"] = store.NewGetOrderByID(o.context, o.StoreGetOrderByIDHandler)

	if o.handlers["GET"] == nil {
		o.handlers[strings.ToUpper("GET")] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/pets/{petId}"] = pet.NewGetPetByID(o.context, o.PetGetPetByIDHandler)

	if o.handlers["GET"] == nil {
		o.handlers[strings.ToUpper("GET")] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/users/{username}"] = user.NewGetUserByName(o.context, o.UserGetUserByNameHandler)

	if o.handlers["GET"] == nil {
		o.handlers[strings.ToUpper("GET")] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/users/login"] = user.NewLoginUser(o.context, o.UserLoginUserHandler)

	if o.handlers["GET"] == nil {
		o.handlers[strings.ToUpper("GET")] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/users/logout"] = user.NewLogoutUser(o.context, o.UserLogoutUserHandler)

	if o.handlers["POST"] == nil {
		o.handlers[strings.ToUpper("POST")] = make(map[string]http.Handler)
	}
	o.handlers["POST"]["/stores/order"] = store.NewPlaceOrder(o.context, o.StorePlaceOrderHandler)

	if o.handlers["PUT"] == nil {
		o.handlers[strings.ToUpper("PUT")] = make(map[string]http.Handler)
	}
	o.handlers["PUT"]["/pets"] = pet.NewUpdatePet(o.context, o.PetUpdatePetHandler)

	if o.handlers["POST"] == nil {
		o.handlers[strings.ToUpper("POST")] = make(map[string]http.Handler)
	}
	o.handlers["POST"]["/pets/{petId}"] = pet.NewUpdatePetWithForm(o.context, o.PetUpdatePetWithFormHandler)

	if o.handlers["PUT"] == nil {
		o.handlers[strings.ToUpper("PUT")] = make(map[string]http.Handler)
	}
	o.handlers["PUT"]["/users/{username}"] = user.NewUpdateUser(o.context, o.UserUpdateUserHandler)

}

// Serve creates a http handler to serve the API over HTTP
// can be used directly in http.ListenAndServe(":8000", api.Serve(nil))
func (o *PetstoreAPI) Serve(builder middleware.Builder) http.Handler {
	if len(o.handlers) == 0 {
		o.initHandlerCache()
	}

	return o.context.APIHandler(builder)
}
