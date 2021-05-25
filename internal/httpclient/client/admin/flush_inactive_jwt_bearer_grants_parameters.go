// Code generated by go-swagger; DO NOT EDIT.

package admin

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"

	"github.com/ory/hydra/internal/httpclient/models"
)

// NewFlushInactiveJwtBearerGrantsParams creates a new FlushInactiveJwtBearerGrantsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewFlushInactiveJwtBearerGrantsParams() *FlushInactiveJwtBearerGrantsParams {
	return &FlushInactiveJwtBearerGrantsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewFlushInactiveJwtBearerGrantsParamsWithTimeout creates a new FlushInactiveJwtBearerGrantsParams object
// with the ability to set a timeout on a request.
func NewFlushInactiveJwtBearerGrantsParamsWithTimeout(timeout time.Duration) *FlushInactiveJwtBearerGrantsParams {
	return &FlushInactiveJwtBearerGrantsParams{
		timeout: timeout,
	}
}

// NewFlushInactiveJwtBearerGrantsParamsWithContext creates a new FlushInactiveJwtBearerGrantsParams object
// with the ability to set a context for a request.
func NewFlushInactiveJwtBearerGrantsParamsWithContext(ctx context.Context) *FlushInactiveJwtBearerGrantsParams {
	return &FlushInactiveJwtBearerGrantsParams{
		Context: ctx,
	}
}

// NewFlushInactiveJwtBearerGrantsParamsWithHTTPClient creates a new FlushInactiveJwtBearerGrantsParams object
// with the ability to set a custom HTTPClient for a request.
func NewFlushInactiveJwtBearerGrantsParamsWithHTTPClient(client *http.Client) *FlushInactiveJwtBearerGrantsParams {
	return &FlushInactiveJwtBearerGrantsParams{
		HTTPClient: client,
	}
}

/* FlushInactiveJwtBearerGrantsParams contains all the parameters to send to the API endpoint
   for the flush inactive jwt bearer grants operation.

   Typically these are written to a http.Request.
*/
type FlushInactiveJwtBearerGrantsParams struct {

	// Body.
	Body *models.FlushInactiveJwtBearerGrantsParams

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the flush inactive jwt bearer grants params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *FlushInactiveJwtBearerGrantsParams) WithDefaults() *FlushInactiveJwtBearerGrantsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the flush inactive jwt bearer grants params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *FlushInactiveJwtBearerGrantsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the flush inactive jwt bearer grants params
func (o *FlushInactiveJwtBearerGrantsParams) WithTimeout(timeout time.Duration) *FlushInactiveJwtBearerGrantsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the flush inactive jwt bearer grants params
func (o *FlushInactiveJwtBearerGrantsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the flush inactive jwt bearer grants params
func (o *FlushInactiveJwtBearerGrantsParams) WithContext(ctx context.Context) *FlushInactiveJwtBearerGrantsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the flush inactive jwt bearer grants params
func (o *FlushInactiveJwtBearerGrantsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the flush inactive jwt bearer grants params
func (o *FlushInactiveJwtBearerGrantsParams) WithHTTPClient(client *http.Client) *FlushInactiveJwtBearerGrantsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the flush inactive jwt bearer grants params
func (o *FlushInactiveJwtBearerGrantsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the flush inactive jwt bearer grants params
func (o *FlushInactiveJwtBearerGrantsParams) WithBody(body *models.FlushInactiveJwtBearerGrantsParams) *FlushInactiveJwtBearerGrantsParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the flush inactive jwt bearer grants params
func (o *FlushInactiveJwtBearerGrantsParams) SetBody(body *models.FlushInactiveJwtBearerGrantsParams) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *FlushInactiveJwtBearerGrantsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
