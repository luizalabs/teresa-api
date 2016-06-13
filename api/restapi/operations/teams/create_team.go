package teams

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// CreateTeamHandlerFunc turns a function with the right signature into a create team handler
type CreateTeamHandlerFunc func(CreateTeamParams, interface{}) middleware.Responder

// Handle executing the request and returning a response
func (fn CreateTeamHandlerFunc) Handle(params CreateTeamParams, principal interface{}) middleware.Responder {
	return fn(params, principal)
}

// CreateTeamHandler interface for that can handle valid create team params
type CreateTeamHandler interface {
	Handle(CreateTeamParams, interface{}) middleware.Responder
}

// NewCreateTeam creates a new http.Handler for the create team operation
func NewCreateTeam(ctx *middleware.Context, handler CreateTeamHandler) *CreateTeam {
	return &CreateTeam{Context: ctx, Handler: handler}
}

/*CreateTeam swagger:route POST /teams teams createTeam

Create teams

Create a team.

*/
type CreateTeam struct {
	Context *middleware.Context
	Handler CreateTeamHandler
}

func (o *CreateTeam) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, _ := o.Context.RouteInfo(r)
	var Params = NewCreateTeamParams()

	uprinc, err := o.Context.Authorize(r, route)
	if err != nil {
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}
	var principal interface{}
	if uprinc != nil {
		principal = uprinc
	}

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params, principal) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
