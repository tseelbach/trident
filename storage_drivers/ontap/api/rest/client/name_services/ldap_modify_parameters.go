// Code generated by go-swagger; DO NOT EDIT.

package name_services

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

	"github.com/netapp/trident/storage_drivers/ontap/api/rest/models"
)

// NewLdapModifyParams creates a new LdapModifyParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewLdapModifyParams() *LdapModifyParams {
	return &LdapModifyParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewLdapModifyParamsWithTimeout creates a new LdapModifyParams object
// with the ability to set a timeout on a request.
func NewLdapModifyParamsWithTimeout(timeout time.Duration) *LdapModifyParams {
	return &LdapModifyParams{
		timeout: timeout,
	}
}

// NewLdapModifyParamsWithContext creates a new LdapModifyParams object
// with the ability to set a context for a request.
func NewLdapModifyParamsWithContext(ctx context.Context) *LdapModifyParams {
	return &LdapModifyParams{
		Context: ctx,
	}
}

// NewLdapModifyParamsWithHTTPClient creates a new LdapModifyParams object
// with the ability to set a custom HTTPClient for a request.
func NewLdapModifyParamsWithHTTPClient(client *http.Client) *LdapModifyParams {
	return &LdapModifyParams{
		HTTPClient: client,
	}
}

/* LdapModifyParams contains all the parameters to send to the API endpoint
   for the ldap modify operation.

   Typically these are written to a http.Request.
*/
type LdapModifyParams struct {

	/* Info.

	   Info specification
	*/
	Info *models.LdapService

	/* SvmUUID.

	   UUID of the SVM to which this object belongs.
	*/
	SVMUUIDPathParameter string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the ldap modify params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *LdapModifyParams) WithDefaults() *LdapModifyParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the ldap modify params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *LdapModifyParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the ldap modify params
func (o *LdapModifyParams) WithTimeout(timeout time.Duration) *LdapModifyParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the ldap modify params
func (o *LdapModifyParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the ldap modify params
func (o *LdapModifyParams) WithContext(ctx context.Context) *LdapModifyParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the ldap modify params
func (o *LdapModifyParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the ldap modify params
func (o *LdapModifyParams) WithHTTPClient(client *http.Client) *LdapModifyParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the ldap modify params
func (o *LdapModifyParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithInfo adds the info to the ldap modify params
func (o *LdapModifyParams) WithInfo(info *models.LdapService) *LdapModifyParams {
	o.SetInfo(info)
	return o
}

// SetInfo adds the info to the ldap modify params
func (o *LdapModifyParams) SetInfo(info *models.LdapService) {
	o.Info = info
}

// WithSVMUUIDPathParameter adds the svmUUID to the ldap modify params
func (o *LdapModifyParams) WithSVMUUIDPathParameter(svmUUID string) *LdapModifyParams {
	o.SetSVMUUIDPathParameter(svmUUID)
	return o
}

// SetSVMUUIDPathParameter adds the svmUuid to the ldap modify params
func (o *LdapModifyParams) SetSVMUUIDPathParameter(svmUUID string) {
	o.SVMUUIDPathParameter = svmUUID
}

// WriteToRequest writes these params to a swagger request
func (o *LdapModifyParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Info != nil {
		if err := r.SetBodyParam(o.Info); err != nil {
			return err
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