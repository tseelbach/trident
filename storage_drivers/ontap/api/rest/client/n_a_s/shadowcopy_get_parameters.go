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

// NewShadowcopyGetParams creates a new ShadowcopyGetParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewShadowcopyGetParams() *ShadowcopyGetParams {
	return &ShadowcopyGetParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewShadowcopyGetParamsWithTimeout creates a new ShadowcopyGetParams object
// with the ability to set a timeout on a request.
func NewShadowcopyGetParamsWithTimeout(timeout time.Duration) *ShadowcopyGetParams {
	return &ShadowcopyGetParams{
		timeout: timeout,
	}
}

// NewShadowcopyGetParamsWithContext creates a new ShadowcopyGetParams object
// with the ability to set a context for a request.
func NewShadowcopyGetParamsWithContext(ctx context.Context) *ShadowcopyGetParams {
	return &ShadowcopyGetParams{
		Context: ctx,
	}
}

// NewShadowcopyGetParamsWithHTTPClient creates a new ShadowcopyGetParams object
// with the ability to set a custom HTTPClient for a request.
func NewShadowcopyGetParamsWithHTTPClient(client *http.Client) *ShadowcopyGetParams {
	return &ShadowcopyGetParams{
		HTTPClient: client,
	}
}

/*
ShadowcopyGetParams contains all the parameters to send to the API endpoint

	for the shadowcopy get operation.

	Typically these are written to a http.Request.
*/
type ShadowcopyGetParams struct {

	/* ClientUUID.

	   client shadowcopy ID
	*/
	ClientUUIDPathParameter string

	/* Fields.

	   Specify the fields to return.
	*/
	FieldsQueryParameter []string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the shadowcopy get params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ShadowcopyGetParams) WithDefaults() *ShadowcopyGetParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the shadowcopy get params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ShadowcopyGetParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the shadowcopy get params
func (o *ShadowcopyGetParams) WithTimeout(timeout time.Duration) *ShadowcopyGetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the shadowcopy get params
func (o *ShadowcopyGetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the shadowcopy get params
func (o *ShadowcopyGetParams) WithContext(ctx context.Context) *ShadowcopyGetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the shadowcopy get params
func (o *ShadowcopyGetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the shadowcopy get params
func (o *ShadowcopyGetParams) WithHTTPClient(client *http.Client) *ShadowcopyGetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the shadowcopy get params
func (o *ShadowcopyGetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithClientUUIDPathParameter adds the clientUUID to the shadowcopy get params
func (o *ShadowcopyGetParams) WithClientUUIDPathParameter(clientUUID string) *ShadowcopyGetParams {
	o.SetClientUUIDPathParameter(clientUUID)
	return o
}

// SetClientUUIDPathParameter adds the clientUuid to the shadowcopy get params
func (o *ShadowcopyGetParams) SetClientUUIDPathParameter(clientUUID string) {
	o.ClientUUIDPathParameter = clientUUID
}

// WithFieldsQueryParameter adds the fields to the shadowcopy get params
func (o *ShadowcopyGetParams) WithFieldsQueryParameter(fields []string) *ShadowcopyGetParams {
	o.SetFieldsQueryParameter(fields)
	return o
}

// SetFieldsQueryParameter adds the fields to the shadowcopy get params
func (o *ShadowcopyGetParams) SetFieldsQueryParameter(fields []string) {
	o.FieldsQueryParameter = fields
}

// WriteToRequest writes these params to a swagger request
func (o *ShadowcopyGetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param client_uuid
	if err := r.SetPathParam("client_uuid", o.ClientUUIDPathParameter); err != nil {
		return err
	}

	if o.FieldsQueryParameter != nil {

		// binding items for fields
		joinedFields := o.bindParamFields(reg)

		// query array param fields
		if err := r.SetQueryParam("fields", joinedFields...); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindParamShadowcopyGet binds the parameter fields
func (o *ShadowcopyGetParams) bindParamFields(formats strfmt.Registry) []string {
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
