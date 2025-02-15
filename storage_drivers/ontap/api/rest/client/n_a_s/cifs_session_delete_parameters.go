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

// NewCifsSessionDeleteParams creates a new CifsSessionDeleteParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewCifsSessionDeleteParams() *CifsSessionDeleteParams {
	return &CifsSessionDeleteParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewCifsSessionDeleteParamsWithTimeout creates a new CifsSessionDeleteParams object
// with the ability to set a timeout on a request.
func NewCifsSessionDeleteParamsWithTimeout(timeout time.Duration) *CifsSessionDeleteParams {
	return &CifsSessionDeleteParams{
		timeout: timeout,
	}
}

// NewCifsSessionDeleteParamsWithContext creates a new CifsSessionDeleteParams object
// with the ability to set a context for a request.
func NewCifsSessionDeleteParamsWithContext(ctx context.Context) *CifsSessionDeleteParams {
	return &CifsSessionDeleteParams{
		Context: ctx,
	}
}

// NewCifsSessionDeleteParamsWithHTTPClient creates a new CifsSessionDeleteParams object
// with the ability to set a custom HTTPClient for a request.
func NewCifsSessionDeleteParamsWithHTTPClient(client *http.Client) *CifsSessionDeleteParams {
	return &CifsSessionDeleteParams{
		HTTPClient: client,
	}
}

/*
CifsSessionDeleteParams contains all the parameters to send to the API endpoint

	for the cifs session delete operation.

	Typically these are written to a http.Request.
*/
type CifsSessionDeleteParams struct {

	/* ConnectionID.

	   Unique identifier for the SMB connection.
	*/
	ConnectionIDPathParameter int64

	/* Identifier.

	   Unique identifier for the SMB session.
	*/
	IdentifierPathParameter int64

	/* NodeUUID.

	   Node UUID.
	*/
	NodeUUIDPathParameter string

	/* SvmUUID.

	   UUID of the SVM to which this object belongs.
	*/
	SVMUUIDPathParameter string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the cifs session delete params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CifsSessionDeleteParams) WithDefaults() *CifsSessionDeleteParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the cifs session delete params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CifsSessionDeleteParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the cifs session delete params
func (o *CifsSessionDeleteParams) WithTimeout(timeout time.Duration) *CifsSessionDeleteParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the cifs session delete params
func (o *CifsSessionDeleteParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the cifs session delete params
func (o *CifsSessionDeleteParams) WithContext(ctx context.Context) *CifsSessionDeleteParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the cifs session delete params
func (o *CifsSessionDeleteParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the cifs session delete params
func (o *CifsSessionDeleteParams) WithHTTPClient(client *http.Client) *CifsSessionDeleteParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the cifs session delete params
func (o *CifsSessionDeleteParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithConnectionIDPathParameter adds the connectionID to the cifs session delete params
func (o *CifsSessionDeleteParams) WithConnectionIDPathParameter(connectionID int64) *CifsSessionDeleteParams {
	o.SetConnectionIDPathParameter(connectionID)
	return o
}

// SetConnectionIDPathParameter adds the connectionId to the cifs session delete params
func (o *CifsSessionDeleteParams) SetConnectionIDPathParameter(connectionID int64) {
	o.ConnectionIDPathParameter = connectionID
}

// WithIdentifierPathParameter adds the identifier to the cifs session delete params
func (o *CifsSessionDeleteParams) WithIdentifierPathParameter(identifier int64) *CifsSessionDeleteParams {
	o.SetIdentifierPathParameter(identifier)
	return o
}

// SetIdentifierPathParameter adds the identifier to the cifs session delete params
func (o *CifsSessionDeleteParams) SetIdentifierPathParameter(identifier int64) {
	o.IdentifierPathParameter = identifier
}

// WithNodeUUIDPathParameter adds the nodeUUID to the cifs session delete params
func (o *CifsSessionDeleteParams) WithNodeUUIDPathParameter(nodeUUID string) *CifsSessionDeleteParams {
	o.SetNodeUUIDPathParameter(nodeUUID)
	return o
}

// SetNodeUUIDPathParameter adds the nodeUuid to the cifs session delete params
func (o *CifsSessionDeleteParams) SetNodeUUIDPathParameter(nodeUUID string) {
	o.NodeUUIDPathParameter = nodeUUID
}

// WithSVMUUIDPathParameter adds the svmUUID to the cifs session delete params
func (o *CifsSessionDeleteParams) WithSVMUUIDPathParameter(svmUUID string) *CifsSessionDeleteParams {
	o.SetSVMUUIDPathParameter(svmUUID)
	return o
}

// SetSVMUUIDPathParameter adds the svmUuid to the cifs session delete params
func (o *CifsSessionDeleteParams) SetSVMUUIDPathParameter(svmUUID string) {
	o.SVMUUIDPathParameter = svmUUID
}

// WriteToRequest writes these params to a swagger request
func (o *CifsSessionDeleteParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param connection_id
	if err := r.SetPathParam("connection_id", swag.FormatInt64(o.ConnectionIDPathParameter)); err != nil {
		return err
	}

	// path param identifier
	if err := r.SetPathParam("identifier", swag.FormatInt64(o.IdentifierPathParameter)); err != nil {
		return err
	}

	// path param node.uuid
	if err := r.SetPathParam("node.uuid", o.NodeUUIDPathParameter); err != nil {
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
