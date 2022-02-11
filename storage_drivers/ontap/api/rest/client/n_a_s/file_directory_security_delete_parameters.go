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
)

// NewFileDirectorySecurityDeleteParams creates a new FileDirectorySecurityDeleteParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewFileDirectorySecurityDeleteParams() *FileDirectorySecurityDeleteParams {
	return &FileDirectorySecurityDeleteParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewFileDirectorySecurityDeleteParamsWithTimeout creates a new FileDirectorySecurityDeleteParams object
// with the ability to set a timeout on a request.
func NewFileDirectorySecurityDeleteParamsWithTimeout(timeout time.Duration) *FileDirectorySecurityDeleteParams {
	return &FileDirectorySecurityDeleteParams{
		timeout: timeout,
	}
}

// NewFileDirectorySecurityDeleteParamsWithContext creates a new FileDirectorySecurityDeleteParams object
// with the ability to set a context for a request.
func NewFileDirectorySecurityDeleteParamsWithContext(ctx context.Context) *FileDirectorySecurityDeleteParams {
	return &FileDirectorySecurityDeleteParams{
		Context: ctx,
	}
}

// NewFileDirectorySecurityDeleteParamsWithHTTPClient creates a new FileDirectorySecurityDeleteParams object
// with the ability to set a custom HTTPClient for a request.
func NewFileDirectorySecurityDeleteParamsWithHTTPClient(client *http.Client) *FileDirectorySecurityDeleteParams {
	return &FileDirectorySecurityDeleteParams{
		HTTPClient: client,
	}
}

/* FileDirectorySecurityDeleteParams contains all the parameters to send to the API endpoint
   for the file directory security delete operation.

   Typically these are written to a http.Request.
*/
type FileDirectorySecurityDeleteParams struct {

	/* AccessControl.

	   Remove all SLAG ACLs. Currently bulk deletion of file-directory ACLs is not supported.
	*/
	AccessControlQueryParameter *string

	/* Path.

	   target path
	*/
	PathPathParameter string

	/* SvmUUID.

	   UUID of the SVM to which this object belongs.
	*/
	SVMUUIDPathParameter string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the file directory security delete params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *FileDirectorySecurityDeleteParams) WithDefaults() *FileDirectorySecurityDeleteParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the file directory security delete params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *FileDirectorySecurityDeleteParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the file directory security delete params
func (o *FileDirectorySecurityDeleteParams) WithTimeout(timeout time.Duration) *FileDirectorySecurityDeleteParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the file directory security delete params
func (o *FileDirectorySecurityDeleteParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the file directory security delete params
func (o *FileDirectorySecurityDeleteParams) WithContext(ctx context.Context) *FileDirectorySecurityDeleteParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the file directory security delete params
func (o *FileDirectorySecurityDeleteParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the file directory security delete params
func (o *FileDirectorySecurityDeleteParams) WithHTTPClient(client *http.Client) *FileDirectorySecurityDeleteParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the file directory security delete params
func (o *FileDirectorySecurityDeleteParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAccessControlQueryParameter adds the accessControl to the file directory security delete params
func (o *FileDirectorySecurityDeleteParams) WithAccessControlQueryParameter(accessControl *string) *FileDirectorySecurityDeleteParams {
	o.SetAccessControlQueryParameter(accessControl)
	return o
}

// SetAccessControlQueryParameter adds the accessControl to the file directory security delete params
func (o *FileDirectorySecurityDeleteParams) SetAccessControlQueryParameter(accessControl *string) {
	o.AccessControlQueryParameter = accessControl
}

// WithPathPathParameter adds the path to the file directory security delete params
func (o *FileDirectorySecurityDeleteParams) WithPathPathParameter(path string) *FileDirectorySecurityDeleteParams {
	o.SetPathPathParameter(path)
	return o
}

// SetPathPathParameter adds the path to the file directory security delete params
func (o *FileDirectorySecurityDeleteParams) SetPathPathParameter(path string) {
	o.PathPathParameter = path
}

// WithSVMUUIDPathParameter adds the svmUUID to the file directory security delete params
func (o *FileDirectorySecurityDeleteParams) WithSVMUUIDPathParameter(svmUUID string) *FileDirectorySecurityDeleteParams {
	o.SetSVMUUIDPathParameter(svmUUID)
	return o
}

// SetSVMUUIDPathParameter adds the svmUuid to the file directory security delete params
func (o *FileDirectorySecurityDeleteParams) SetSVMUUIDPathParameter(svmUUID string) {
	o.SVMUUIDPathParameter = svmUUID
}

// WriteToRequest writes these params to a swagger request
func (o *FileDirectorySecurityDeleteParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.AccessControlQueryParameter != nil {

		// query param access_control
		var qrAccessControl string

		if o.AccessControlQueryParameter != nil {
			qrAccessControl = *o.AccessControlQueryParameter
		}
		qAccessControl := qrAccessControl
		if qAccessControl != "" {

			if err := r.SetQueryParam("access_control", qAccessControl); err != nil {
				return err
			}
		}
	}

	// path param path
	if err := r.SetPathParam("path", o.PathPathParameter); err != nil {
		return err
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