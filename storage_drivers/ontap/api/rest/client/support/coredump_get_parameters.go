// Code generated by go-swagger; DO NOT EDIT.

package support

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

// NewCoredumpGetParams creates a new CoredumpGetParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewCoredumpGetParams() *CoredumpGetParams {
	return &CoredumpGetParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewCoredumpGetParamsWithTimeout creates a new CoredumpGetParams object
// with the ability to set a timeout on a request.
func NewCoredumpGetParamsWithTimeout(timeout time.Duration) *CoredumpGetParams {
	return &CoredumpGetParams{
		timeout: timeout,
	}
}

// NewCoredumpGetParamsWithContext creates a new CoredumpGetParams object
// with the ability to set a context for a request.
func NewCoredumpGetParamsWithContext(ctx context.Context) *CoredumpGetParams {
	return &CoredumpGetParams{
		Context: ctx,
	}
}

// NewCoredumpGetParamsWithHTTPClient creates a new CoredumpGetParams object
// with the ability to set a custom HTTPClient for a request.
func NewCoredumpGetParamsWithHTTPClient(client *http.Client) *CoredumpGetParams {
	return &CoredumpGetParams{
		HTTPClient: client,
	}
}

/*
CoredumpGetParams contains all the parameters to send to the API endpoint

	for the coredump get operation.

	Typically these are written to a http.Request.
*/
type CoredumpGetParams struct {

	/* Fields.

	   Specify the fields to return.
	*/
	FieldsQueryParameter []string

	/* Name.

	   Core name
	*/
	NamePathParameter string

	/* NodeUUID.

	   Node UUID
	*/
	NodeUUIDPathParameter string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the coredump get params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CoredumpGetParams) WithDefaults() *CoredumpGetParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the coredump get params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CoredumpGetParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the coredump get params
func (o *CoredumpGetParams) WithTimeout(timeout time.Duration) *CoredumpGetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the coredump get params
func (o *CoredumpGetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the coredump get params
func (o *CoredumpGetParams) WithContext(ctx context.Context) *CoredumpGetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the coredump get params
func (o *CoredumpGetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the coredump get params
func (o *CoredumpGetParams) WithHTTPClient(client *http.Client) *CoredumpGetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the coredump get params
func (o *CoredumpGetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithFieldsQueryParameter adds the fields to the coredump get params
func (o *CoredumpGetParams) WithFieldsQueryParameter(fields []string) *CoredumpGetParams {
	o.SetFieldsQueryParameter(fields)
	return o
}

// SetFieldsQueryParameter adds the fields to the coredump get params
func (o *CoredumpGetParams) SetFieldsQueryParameter(fields []string) {
	o.FieldsQueryParameter = fields
}

// WithNamePathParameter adds the name to the coredump get params
func (o *CoredumpGetParams) WithNamePathParameter(name string) *CoredumpGetParams {
	o.SetNamePathParameter(name)
	return o
}

// SetNamePathParameter adds the name to the coredump get params
func (o *CoredumpGetParams) SetNamePathParameter(name string) {
	o.NamePathParameter = name
}

// WithNodeUUIDPathParameter adds the nodeUUID to the coredump get params
func (o *CoredumpGetParams) WithNodeUUIDPathParameter(nodeUUID string) *CoredumpGetParams {
	o.SetNodeUUIDPathParameter(nodeUUID)
	return o
}

// SetNodeUUIDPathParameter adds the nodeUuid to the coredump get params
func (o *CoredumpGetParams) SetNodeUUIDPathParameter(nodeUUID string) {
	o.NodeUUIDPathParameter = nodeUUID
}

// WriteToRequest writes these params to a swagger request
func (o *CoredumpGetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.FieldsQueryParameter != nil {

		// binding items for fields
		joinedFields := o.bindParamFields(reg)

		// query array param fields
		if err := r.SetQueryParam("fields", joinedFields...); err != nil {
			return err
		}
	}

	// path param name
	if err := r.SetPathParam("name", o.NamePathParameter); err != nil {
		return err
	}

	// path param node.uuid
	if err := r.SetPathParam("node.uuid", o.NodeUUIDPathParameter); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindParamCoredumpGet binds the parameter fields
func (o *CoredumpGetParams) bindParamFields(formats strfmt.Registry) []string {
	fieldsIR := o.FieldsQueryParameter

	var fieldsIC []string
	for _, fieldsIIR := range fieldsIR { // explode []string

		fieldsIIV := fieldsIIR // string as string
		fieldsIC = append(fieldsIC, fieldsIIV)
	}

	// items.CollectionFormat: "csv"
	fieldsIS := swag.JoinByFormat(fieldsIC, "csv")

	return fieldsIS
}
