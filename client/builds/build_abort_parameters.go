// Code generated by go-swagger; DO NOT EDIT.

package builds

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/slapec93/bitrise-api-client/models"
)

// NewBuildAbortParams creates a new BuildAbortParams object
// with the default values initialized.
func NewBuildAbortParams() *BuildAbortParams {
	var ()
	return &BuildAbortParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewBuildAbortParamsWithTimeout creates a new BuildAbortParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewBuildAbortParamsWithTimeout(timeout time.Duration) *BuildAbortParams {
	var ()
	return &BuildAbortParams{

		timeout: timeout,
	}
}

// NewBuildAbortParamsWithContext creates a new BuildAbortParams object
// with the default values initialized, and the ability to set a context for a request
func NewBuildAbortParamsWithContext(ctx context.Context) *BuildAbortParams {
	var ()
	return &BuildAbortParams{

		Context: ctx,
	}
}

// NewBuildAbortParamsWithHTTPClient creates a new BuildAbortParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewBuildAbortParamsWithHTTPClient(client *http.Client) *BuildAbortParams {
	var ()
	return &BuildAbortParams{
		HTTPClient: client,
	}
}

/*BuildAbortParams contains all the parameters to send to the API endpoint
for the build abort operation typically these are written to a http.Request
*/
type BuildAbortParams struct {

	/*AppSlug
	  App slug

	*/
	AppSlug string
	/*BuildAbortParams
	  Build abort parameters

	*/
	BuildAbortParams *models.V0BuildAbortParams
	/*BuildSlug
	  Build slug

	*/
	BuildSlug string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the build abort params
func (o *BuildAbortParams) WithTimeout(timeout time.Duration) *BuildAbortParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the build abort params
func (o *BuildAbortParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the build abort params
func (o *BuildAbortParams) WithContext(ctx context.Context) *BuildAbortParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the build abort params
func (o *BuildAbortParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the build abort params
func (o *BuildAbortParams) WithHTTPClient(client *http.Client) *BuildAbortParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the build abort params
func (o *BuildAbortParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAppSlug adds the appSlug to the build abort params
func (o *BuildAbortParams) WithAppSlug(appSlug string) *BuildAbortParams {
	o.SetAppSlug(appSlug)
	return o
}

// SetAppSlug adds the appSlug to the build abort params
func (o *BuildAbortParams) SetAppSlug(appSlug string) {
	o.AppSlug = appSlug
}

// WithBuildAbortParams adds the buildAbortParams to the build abort params
func (o *BuildAbortParams) WithBuildAbortParams(buildAbortParams *models.V0BuildAbortParams) *BuildAbortParams {
	o.SetBuildAbortParams(buildAbortParams)
	return o
}

// SetBuildAbortParams adds the buildAbortParams to the build abort params
func (o *BuildAbortParams) SetBuildAbortParams(buildAbortParams *models.V0BuildAbortParams) {
	o.BuildAbortParams = buildAbortParams
}

// WithBuildSlug adds the buildSlug to the build abort params
func (o *BuildAbortParams) WithBuildSlug(buildSlug string) *BuildAbortParams {
	o.SetBuildSlug(buildSlug)
	return o
}

// SetBuildSlug adds the buildSlug to the build abort params
func (o *BuildAbortParams) SetBuildSlug(buildSlug string) {
	o.BuildSlug = buildSlug
}

// WriteToRequest writes these params to a swagger request
func (o *BuildAbortParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param app-slug
	if err := r.SetPathParam("app-slug", o.AppSlug); err != nil {
		return err
	}

	if o.BuildAbortParams != nil {
		if err := r.SetBodyParam(o.BuildAbortParams); err != nil {
			return err
		}
	}

	// path param build-slug
	if err := r.SetPathParam("build-slug", o.BuildSlug); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
