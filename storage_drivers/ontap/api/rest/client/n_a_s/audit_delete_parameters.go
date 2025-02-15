// Code generated by go-swagger; DO NOT EDIT.

package n_a_s

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
	"github.com/go-openapi/swag"
)

// NewAuditDeleteParams creates a new AuditDeleteParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewAuditDeleteParams() *AuditDeleteParams {
	return &AuditDeleteParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewAuditDeleteParamsWithTimeout creates a new AuditDeleteParams object
// with the ability to set a timeout on a request.
func NewAuditDeleteParamsWithTimeout(timeout time.Duration) *AuditDeleteParams {
	return &AuditDeleteParams{
		timeout: timeout,
	}
}

// NewAuditDeleteParamsWithContext creates a new AuditDeleteParams object
// with the ability to set a context for a request.
func NewAuditDeleteParamsWithContext(ctx context.Context) *AuditDeleteParams {
	return &AuditDeleteParams{
		Context: ctx,
	}
}

// NewAuditDeleteParamsWithHTTPClient creates a new AuditDeleteParams object
// with the ability to set a custom HTTPClient for a request.
func NewAuditDeleteParamsWithHTTPClient(client *http.Client) *AuditDeleteParams {
	return &AuditDeleteParams{
		HTTPClient: client,
	}
}

/*
AuditDeleteParams contains all the parameters to send to the API endpoint

	for the audit delete operation.

	Typically these are written to a http.Request.
*/
type AuditDeleteParams struct {

	/* Force.

	   Indicates to force delete audit configuration.
	*/
	ForceQueryParameter *bool

	/* ReturnTimeout.

	   The number of seconds to allow the call to execute before returning. When doing a POST, PATCH, or DELETE operation on a single record, the default is 0 seconds.  This means that if an asynchronous operation is started, the server immediately returns HTTP code 202 (Accepted) along with a link to the job.  If a non-zero value is specified for POST, PATCH, or DELETE operations, ONTAP waits that length of time to see if the job completes so it can return something other than 202.
	*/
	ReturnTimeoutQueryParameter *int64

	/* SvmUUID.

	   UUID of the SVM to which this object belongs.
	*/
	SVMUUIDPathParameter string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the audit delete params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *AuditDeleteParams) WithDefaults() *AuditDeleteParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the audit delete params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *AuditDeleteParams) SetDefaults() {
	var (
		returnTimeoutQueryParameterDefault = int64(0)
	)

	val := AuditDeleteParams{
		ReturnTimeoutQueryParameter: &returnTimeoutQueryParameterDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the audit delete params
func (o *AuditDeleteParams) WithTimeout(timeout time.Duration) *AuditDeleteParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the audit delete params
func (o *AuditDeleteParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the audit delete params
func (o *AuditDeleteParams) WithContext(ctx context.Context) *AuditDeleteParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the audit delete params
func (o *AuditDeleteParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the audit delete params
func (o *AuditDeleteParams) WithHTTPClient(client *http.Client) *AuditDeleteParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the audit delete params
func (o *AuditDeleteParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithForceQueryParameter adds the force to the audit delete params
func (o *AuditDeleteParams) WithForceQueryParameter(force *bool) *AuditDeleteParams {
	o.SetForceQueryParameter(force)
	return o
}

// SetForceQueryParameter adds the force to the audit delete params
func (o *AuditDeleteParams) SetForceQueryParameter(force *bool) {
	o.ForceQueryParameter = force
}

// WithReturnTimeoutQueryParameter adds the returnTimeout to the audit delete params
func (o *AuditDeleteParams) WithReturnTimeoutQueryParameter(returnTimeout *int64) *AuditDeleteParams {
	o.SetReturnTimeoutQueryParameter(returnTimeout)
	return o
}

// SetReturnTimeoutQueryParameter adds the returnTimeout to the audit delete params
func (o *AuditDeleteParams) SetReturnTimeoutQueryParameter(returnTimeout *int64) {
	o.ReturnTimeoutQueryParameter = returnTimeout
}

// WithSVMUUIDPathParameter adds the svmUUID to the audit delete params
func (o *AuditDeleteParams) WithSVMUUIDPathParameter(svmUUID string) *AuditDeleteParams {
	o.SetSVMUUIDPathParameter(svmUUID)
	return o
}

// SetSVMUUIDPathParameter adds the svmUuid to the audit delete params
func (o *AuditDeleteParams) SetSVMUUIDPathParameter(svmUUID string) {
	o.SVMUUIDPathParameter = svmUUID
}

// WriteToRequest writes these params to a swagger request
func (o *AuditDeleteParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.ForceQueryParameter != nil {

		// query param force
		var qrForce bool

		if o.ForceQueryParameter != nil {
			qrForce = *o.ForceQueryParameter
		}
		qForce := swag.FormatBool(qrForce)
		if qForce != "" {

			if err := r.SetQueryParam("force", qForce); err != nil {
				return err
			}
		}
	}

	if o.ReturnTimeoutQueryParameter != nil {

		// query param return_timeout
		var qrReturnTimeout int64

		if o.ReturnTimeoutQueryParameter != nil {
			qrReturnTimeout = *o.ReturnTimeoutQueryParameter
		}
		qReturnTimeout := swag.FormatInt64(qrReturnTimeout)
		if qReturnTimeout != "" {

			if err := r.SetQueryParam("return_timeout", qReturnTimeout); err != nil {
				return err
			}
		}
	}

	// path param svm.uuid
	if err := r.SetPathParam("svm.uuid", o.SVMUUIDPathParameter); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
