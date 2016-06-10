package deployments

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/luizalabs/paas/api/models"
)

/*GetDeploymentsOK Get app deployments

swagger:response getDeploymentsOK
*/
type GetDeploymentsOK struct {

	// In: body
	Payload GetDeploymentsOKBodyBody `json:"body,omitempty"`
}

// NewGetDeploymentsOK creates GetDeploymentsOK with default headers values
func NewGetDeploymentsOK() *GetDeploymentsOK {
	return &GetDeploymentsOK{}
}

// WithPayload adds the payload to the get deployments o k response
func (o *GetDeploymentsOK) WithPayload(payload GetDeploymentsOKBodyBody) *GetDeploymentsOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get deployments o k response
func (o *GetDeploymentsOK) SetPayload(payload GetDeploymentsOKBodyBody) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetDeploymentsOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if err := producer.Produce(rw, o.Payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}

}

/*GetDeploymentsUnauthorized User not authorized

swagger:response getDeploymentsUnauthorized
*/
type GetDeploymentsUnauthorized struct {

	// In: body
	Payload *models.Unauthorized `json:"body,omitempty"`
}

// NewGetDeploymentsUnauthorized creates GetDeploymentsUnauthorized with default headers values
func NewGetDeploymentsUnauthorized() *GetDeploymentsUnauthorized {
	return &GetDeploymentsUnauthorized{}
}

// WithPayload adds the payload to the get deployments unauthorized response
func (o *GetDeploymentsUnauthorized) WithPayload(payload *models.Unauthorized) *GetDeploymentsUnauthorized {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get deployments unauthorized response
func (o *GetDeploymentsUnauthorized) SetPayload(payload *models.Unauthorized) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetDeploymentsUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(401)
	if o.Payload != nil {
		if err := producer.Produce(rw, o.Payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

/*GetDeploymentsForbidden User does not have the credentials to access this resource


swagger:response getDeploymentsForbidden
*/
type GetDeploymentsForbidden struct {

	// In: body
	Payload *models.Unauthorized `json:"body,omitempty"`
}

// NewGetDeploymentsForbidden creates GetDeploymentsForbidden with default headers values
func NewGetDeploymentsForbidden() *GetDeploymentsForbidden {
	return &GetDeploymentsForbidden{}
}

// WithPayload adds the payload to the get deployments forbidden response
func (o *GetDeploymentsForbidden) WithPayload(payload *models.Unauthorized) *GetDeploymentsForbidden {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get deployments forbidden response
func (o *GetDeploymentsForbidden) SetPayload(payload *models.Unauthorized) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetDeploymentsForbidden) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(403)
	if o.Payload != nil {
		if err := producer.Produce(rw, o.Payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

/*GetDeploymentsDefault Error

swagger:response getDeploymentsDefault
*/
type GetDeploymentsDefault struct {
	_statusCode int

	// In: body
	Payload *models.Error `json:"body,omitempty"`
}

// NewGetDeploymentsDefault creates GetDeploymentsDefault with default headers values
func NewGetDeploymentsDefault(code int) *GetDeploymentsDefault {
	if code <= 0 {
		code = 500
	}

	return &GetDeploymentsDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the get deployments default response
func (o *GetDeploymentsDefault) WithStatusCode(code int) *GetDeploymentsDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the get deployments default response
func (o *GetDeploymentsDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the get deployments default response
func (o *GetDeploymentsDefault) WithPayload(payload *models.Error) *GetDeploymentsDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get deployments default response
func (o *GetDeploymentsDefault) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetDeploymentsDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		if err := producer.Produce(rw, o.Payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
