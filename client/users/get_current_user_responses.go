package users

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/luizalabs/tapi/models"
)

// GetCurrentUserReader is a Reader for the GetCurrentUser structure.
type GetCurrentUserReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the recieved o.
func (o *GetCurrentUserReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetCurrentUserOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 404:
		result := NewGetCurrentUserNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		result := NewGetCurrentUserDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	}
}

// NewGetCurrentUserOK creates a GetCurrentUserOK with default headers values
func NewGetCurrentUserOK() *GetCurrentUserOK {
	return &GetCurrentUserOK{}
}

/*GetCurrentUserOK handles this case with default header values.

User info
*/
type GetCurrentUserOK struct {
	Payload *models.User
}

func (o *GetCurrentUserOK) Error() string {
	return fmt.Sprintf("[GET /users/me][%d] getCurrentUserOK  %+v", 200, o.Payload)
}

func (o *GetCurrentUserOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.User)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCurrentUserNotFound creates a GetCurrentUserNotFound with default headers values
func NewGetCurrentUserNotFound() *GetCurrentUserNotFound {
	return &GetCurrentUserNotFound{}
}

/*GetCurrentUserNotFound handles this case with default header values.

Resource not found
*/
type GetCurrentUserNotFound struct {
	Payload *models.NotFound
}

func (o *GetCurrentUserNotFound) Error() string {
	return fmt.Sprintf("[GET /users/me][%d] getCurrentUserNotFound  %+v", 404, o.Payload)
}

func (o *GetCurrentUserNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.NotFound)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCurrentUserDefault creates a GetCurrentUserDefault with default headers values
func NewGetCurrentUserDefault(code int) *GetCurrentUserDefault {
	return &GetCurrentUserDefault{
		_statusCode: code,
	}
}

/*GetCurrentUserDefault handles this case with default header values.

Error
*/
type GetCurrentUserDefault struct {
	_statusCode int

	Payload *models.GenericError
}

// Code gets the status code for the get current user default response
func (o *GetCurrentUserDefault) Code() int {
	return o._statusCode
}

func (o *GetCurrentUserDefault) Error() string {
	return fmt.Sprintf("[GET /users/me][%d] getCurrentUser default  %+v", o._statusCode, o.Payload)
}

func (o *GetCurrentUserDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GenericError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
