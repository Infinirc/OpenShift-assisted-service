// Code generated by go-swagger; DO NOT EDIT.

package events

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// V2TriggerEventHandlerFunc turns a function with the right signature into a v2 trigger event handler
type V2TriggerEventHandlerFunc func(V2TriggerEventParams, interface{}) middleware.Responder

// Handle executing the request and returning a response
func (fn V2TriggerEventHandlerFunc) Handle(params V2TriggerEventParams, principal interface{}) middleware.Responder {
	return fn(params, principal)
}

// V2TriggerEventHandler interface for that can handle valid v2 trigger event params
type V2TriggerEventHandler interface {
	Handle(V2TriggerEventParams, interface{}) middleware.Responder
}

// NewV2TriggerEvent creates a new http.Handler for the v2 trigger event operation
func NewV2TriggerEvent(ctx *middleware.Context, handler V2TriggerEventHandler) *V2TriggerEvent {
	return &V2TriggerEvent{Context: ctx, Handler: handler}
}

/*
	V2TriggerEvent swagger:route POST /v2/events events v2TriggerEvent

Add new assisted installer event.
*/
type V2TriggerEvent struct {
	Context *middleware.Context
	Handler V2TriggerEventHandler
}

func (o *V2TriggerEvent) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewV2TriggerEventParams()
	uprinc, aCtx, err := o.Context.Authorize(r, route)
	if err != nil {
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}
	if aCtx != nil {
		*r = *aCtx
	}
	var principal interface{}
	if uprinc != nil {
		principal = uprinc.(interface{}) // this is really a interface{}, I promise
	}

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params, principal) // actually handle the request
	o.Context.Respond(rw, r, route.Produces, route, res)

}
