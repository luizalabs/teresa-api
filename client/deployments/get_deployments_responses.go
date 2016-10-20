package deployments

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/luizalabs/tapi/models"
)

// GetDeploymentsReader is a Reader for the GetDeployments structure.
type GetDeploymentsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the recieved o.
func (o *GetDeploymentsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetDeploymentsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 401:
		result := NewGetDeploymentsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewGetDeploymentsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		result := NewGetDeploymentsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	}
}

// NewGetDeploymentsOK creates a GetDeploymentsOK with default headers values
func NewGetDeploymentsOK() *GetDeploymentsOK {
	return &GetDeploymentsOK{}
}

/*GetDeploymentsOK handles this case with default header values.

Get app deployments
*/
type GetDeploymentsOK struct {
	Payload GetDeploymentsOKBodyBody
}

func (o *GetDeploymentsOK) Error() string {
	return fmt.Sprintf("[GET /teams/{team_id}/apps/{app_id}/deployments][%d] getDeploymentsOK  %+v", 200, o.Payload)
}

func (o *GetDeploymentsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetDeploymentsUnauthorized creates a GetDeploymentsUnauthorized with default headers values
func NewGetDeploymentsUnauthorized() *GetDeploymentsUnauthorized {
	return &GetDeploymentsUnauthorized{}
}

/*GetDeploymentsUnauthorized handles this case with default header values.

User not authorized
*/
type GetDeploymentsUnauthorized struct {
	Payload *models.Unauthorized
}

func (o *GetDeploymentsUnauthorized) Error() string {
	return fmt.Sprintf("[GET /teams/{team_id}/apps/{app_id}/deployments][%d] getDeploymentsUnauthorized  %+v", 401, o.Payload)
}

func (o *GetDeploymentsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Unauthorized)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetDeploymentsForbidden creates a GetDeploymentsForbidden with default headers values
func NewGetDeploymentsForbidden() *GetDeploymentsForbidden {
	return &GetDeploymentsForbidden{}
}

/*GetDeploymentsForbidden handles this case with default header values.

User does not have the credentials to access this resource

*/
type GetDeploymentsForbidden struct {
	Payload *models.Unauthorized
}

func (o *GetDeploymentsForbidden) Error() string {
	return fmt.Sprintf("[GET /teams/{team_id}/apps/{app_id}/deployments][%d] getDeploymentsForbidden  %+v", 403, o.Payload)
}

func (o *GetDeploymentsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Unauthorized)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetDeploymentsDefault creates a GetDeploymentsDefault with default headers values
func NewGetDeploymentsDefault(code int) *GetDeploymentsDefault {
	return &GetDeploymentsDefault{
		_statusCode: code,
	}
}

/*GetDeploymentsDefault handles this case with default header values.

Error
*/
type GetDeploymentsDefault struct {
	_statusCode int

	Payload *models.GenericError
}

// Code gets the status code for the get deployments default response
func (o *GetDeploymentsDefault) Code() int {
	return o._statusCode
}

func (o *GetDeploymentsDefault) Error() string {
	return fmt.Sprintf("[GET /teams/{team_id}/apps/{app_id}/deployments][%d] getDeployments default  %+v", o._statusCode, o.Payload)
}

func (o *GetDeploymentsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GenericError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*GetDeploymentsOKBodyBody get deployments o k body body

swagger:model GetDeploymentsOKBodyBody
*/
type GetDeploymentsOKBodyBody struct {
	models.Pagination

	/* items

	Required: true
	*/
	Items []*models.Deployment `json:"items"`
}

// Validate validates this get deployments o k body body
func (o *GetDeploymentsOKBodyBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.Pagination.Validate(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateItems(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetDeploymentsOKBodyBody) validateItems(formats strfmt.Registry) error {

	if err := validate.Required("getDeploymentsOK"+"."+"items", "body", o.Items); err != nil {
		return err
	}

	for i := 0; i < len(o.Items); i++ {

		if swag.IsZero(o.Items[i]) { // not required
			continue
		}

		if o.Items[i] != nil {

			if err := o.Items[i].Validate(formats); err != nil {
				return err
			}
		}

	}

	return nil
}
