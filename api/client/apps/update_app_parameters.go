package apps

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/luizalabs/paas/api/models"
)

// NewUpdateAppParams creates a new UpdateAppParams object
// with the default values initialized.
func NewUpdateAppParams() *UpdateAppParams {
	var ()
	return &UpdateAppParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateAppParamsWithTimeout creates a new UpdateAppParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewUpdateAppParamsWithTimeout(timeout time.Duration) *UpdateAppParams {
	var ()
	return &UpdateAppParams{

		timeout: timeout,
	}
}

/*UpdateAppParams contains all the parameters to send to the API endpoint
for the update app operation typically these are written to a http.Request
*/
type UpdateAppParams struct {

	/*AppID
	  App ID

	*/
	AppID int64
	/*Body*/
	Body *models.App
	/*TeamID
	  Team ID

	*/
	TeamID int64

	timeout time.Duration
}

// WithAppID adds the appId to the update app params
func (o *UpdateAppParams) WithAppID(AppID int64) *UpdateAppParams {
	o.AppID = AppID
	return o
}

// WithBody adds the body to the update app params
func (o *UpdateAppParams) WithBody(Body *models.App) *UpdateAppParams {
	o.Body = Body
	return o
}

// WithTeamID adds the teamId to the update app params
func (o *UpdateAppParams) WithTeamID(TeamID int64) *UpdateAppParams {
	o.TeamID = TeamID
	return o
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateAppParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	r.SetTimeout(o.timeout)
	var res []error

	// path param app_id
	if err := r.SetPathParam("app_id", swag.FormatInt64(o.AppID)); err != nil {
		return err
	}

	if o.Body == nil {
		o.Body = new(models.App)
	}

	if err := r.SetBodyParam(o.Body); err != nil {
		return err
	}

	// path param team_id
	if err := r.SetPathParam("team_id", swag.FormatInt64(o.TeamID)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
