package apps

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/luizalabs/tapi/models"
)

// NewUpdateAppAutoScaleParams creates a new UpdateAppAutoScaleParams object
// with the default values initialized.
func NewUpdateAppAutoScaleParams() *UpdateAppAutoScaleParams {
	var ()
	return &UpdateAppAutoScaleParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateAppAutoScaleParamsWithTimeout creates a new UpdateAppAutoScaleParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewUpdateAppAutoScaleParamsWithTimeout(timeout time.Duration) *UpdateAppAutoScaleParams {
	var ()
	return &UpdateAppAutoScaleParams{

		timeout: timeout,
	}
}

/*UpdateAppAutoScaleParams contains all the parameters to send to the API endpoint
for the update app auto scale operation typically these are written to a http.Request
*/
type UpdateAppAutoScaleParams struct {

	/*AppName
	  App name

	*/
	AppName string
	/*Body*/
	Body *models.AutoScale

	timeout time.Duration
}

// WithAppName adds the appName to the update app auto scale params
func (o *UpdateAppAutoScaleParams) WithAppName(appName string) *UpdateAppAutoScaleParams {
	o.AppName = appName
	return o
}

// WithBody adds the body to the update app auto scale params
func (o *UpdateAppAutoScaleParams) WithBody(body *models.AutoScale) *UpdateAppAutoScaleParams {
	o.Body = body
	return o
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateAppAutoScaleParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	r.SetTimeout(o.timeout)
	var res []error

	// path param app_name
	if err := r.SetPathParam("app_name", o.AppName); err != nil {
		return err
	}

	if o.Body == nil {
		o.Body = new(models.AutoScale)
	}

	if err := r.SetBodyParam(o.Body); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
