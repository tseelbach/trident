// Code generated by go-swagger; DO NOT EDIT.

package storage

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

	"github.com/netapp/trident/storage_drivers/ontap/api/rest/models"
)

// NewFileCloneCreateParams creates a new FileCloneCreateParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewFileCloneCreateParams() *FileCloneCreateParams {
	return &FileCloneCreateParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewFileCloneCreateParamsWithTimeout creates a new FileCloneCreateParams object
// with the ability to set a timeout on a request.
func NewFileCloneCreateParamsWithTimeout(timeout time.Duration) *FileCloneCreateParams {
	return &FileCloneCreateParams{
		timeout: timeout,
	}
}

// NewFileCloneCreateParamsWithContext creates a new FileCloneCreateParams object
// with the ability to set a context for a request.
func NewFileCloneCreateParamsWithContext(ctx context.Context) *FileCloneCreateParams {
	return &FileCloneCreateParams{
		Context: ctx,
	}
}

// NewFileCloneCreateParamsWithHTTPClient creates a new FileCloneCreateParams object
// with the ability to set a custom HTTPClient for a request.
func NewFileCloneCreateParamsWithHTTPClient(client *http.Client) *FileCloneCreateParams {
	return &FileCloneCreateParams{
		HTTPClient: client,
	}
}

/* FileCloneCreateParams contains all the parameters to send to the API endpoint
   for the file clone create operation.

   Typically these are written to a http.Request.
*/
type FileCloneCreateParams struct {

	/* Info.

	   Info specification
	*/
	Info *models.FileClone

	/* ReturnRecords.

	   The default is false.  If set to true, the records are returned.
	*/
	ReturnRecordsQueryParameter *bool

	/* ReturnTimeout.

	   The number of seconds to allow the call to execute before returning. When doing a POST, PATCH, or DELETE operation on a single record, the default is 0 seconds.  This means that if an asynchronous operation is started, the server immediately returns HTTP code 202 (Accepted) along with a link to the job.  If a non-zero value is specified for POST, PATCH, or DELETE operations, ONTAP waits that length of time to see if the job completes so it can return something other than 202.
	*/
	ReturnTimeoutQueryParameter *int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the file clone create params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *FileCloneCreateParams) WithDefaults() *FileCloneCreateParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the file clone create params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *FileCloneCreateParams) SetDefaults() {
	var (
		returnRecordsQueryParameterDefault = bool(false)

		returnTimeoutQueryParameterDefault = int64(0)
	)

	val := FileCloneCreateParams{
		ReturnRecordsQueryParameter: &returnRecordsQueryParameterDefault,
		ReturnTimeoutQueryParameter: &returnTimeoutQueryParameterDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the file clone create params
func (o *FileCloneCreateParams) WithTimeout(timeout time.Duration) *FileCloneCreateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the file clone create params
func (o *FileCloneCreateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the file clone create params
func (o *FileCloneCreateParams) WithContext(ctx context.Context) *FileCloneCreateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the file clone create params
func (o *FileCloneCreateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the file clone create params
func (o *FileCloneCreateParams) WithHTTPClient(client *http.Client) *FileCloneCreateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the file clone create params
func (o *FileCloneCreateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithInfo adds the info to the file clone create params
func (o *FileCloneCreateParams) WithInfo(info *models.FileClone) *FileCloneCreateParams {
	o.SetInfo(info)
	return o
}

// SetInfo adds the info to the file clone create params
func (o *FileCloneCreateParams) SetInfo(info *models.FileClone) {
	o.Info = info
}

// WithReturnRecordsQueryParameter adds the returnRecords to the file clone create params
func (o *FileCloneCreateParams) WithReturnRecordsQueryParameter(returnRecords *bool) *FileCloneCreateParams {
	o.SetReturnRecordsQueryParameter(returnRecords)
	return o
}

// SetReturnRecordsQueryParameter adds the returnRecords to the file clone create params
func (o *FileCloneCreateParams) SetReturnRecordsQueryParameter(returnRecords *bool) {
	o.ReturnRecordsQueryParameter = returnRecords
}

// WithReturnTimeoutQueryParameter adds the returnTimeout to the file clone create params
func (o *FileCloneCreateParams) WithReturnTimeoutQueryParameter(returnTimeout *int64) *FileCloneCreateParams {
	o.SetReturnTimeoutQueryParameter(returnTimeout)
	return o
}

// SetReturnTimeoutQueryParameter adds the returnTimeout to the file clone create params
func (o *FileCloneCreateParams) SetReturnTimeoutQueryParameter(returnTimeout *int64) {
	o.ReturnTimeoutQueryParameter = returnTimeout
}

// WriteToRequest writes these params to a swagger request
func (o *FileCloneCreateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Info != nil {
		if err := r.SetBodyParam(o.Info); err != nil {
			return err
		}
	}

	if o.ReturnRecordsQueryParameter != nil {

		// query param return_records
		var qrReturnRecords bool

		if o.ReturnRecordsQueryParameter != nil {
			qrReturnRecords = *o.ReturnRecordsQueryParameter
		}
		qReturnRecords := swag.FormatBool(qrReturnRecords)
		if qReturnRecords != "" {

			if err := r.SetQueryParam("return_records", qReturnRecords); err != nil {
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

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}