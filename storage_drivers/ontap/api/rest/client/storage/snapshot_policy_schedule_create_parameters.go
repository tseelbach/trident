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

// NewSnapshotPolicyScheduleCreateParams creates a new SnapshotPolicyScheduleCreateParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewSnapshotPolicyScheduleCreateParams() *SnapshotPolicyScheduleCreateParams {
	return &SnapshotPolicyScheduleCreateParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewSnapshotPolicyScheduleCreateParamsWithTimeout creates a new SnapshotPolicyScheduleCreateParams object
// with the ability to set a timeout on a request.
func NewSnapshotPolicyScheduleCreateParamsWithTimeout(timeout time.Duration) *SnapshotPolicyScheduleCreateParams {
	return &SnapshotPolicyScheduleCreateParams{
		timeout: timeout,
	}
}

// NewSnapshotPolicyScheduleCreateParamsWithContext creates a new SnapshotPolicyScheduleCreateParams object
// with the ability to set a context for a request.
func NewSnapshotPolicyScheduleCreateParamsWithContext(ctx context.Context) *SnapshotPolicyScheduleCreateParams {
	return &SnapshotPolicyScheduleCreateParams{
		Context: ctx,
	}
}

// NewSnapshotPolicyScheduleCreateParamsWithHTTPClient creates a new SnapshotPolicyScheduleCreateParams object
// with the ability to set a custom HTTPClient for a request.
func NewSnapshotPolicyScheduleCreateParamsWithHTTPClient(client *http.Client) *SnapshotPolicyScheduleCreateParams {
	return &SnapshotPolicyScheduleCreateParams{
		HTTPClient: client,
	}
}

/* SnapshotPolicyScheduleCreateParams contains all the parameters to send to the API endpoint
   for the snapshot policy schedule create operation.

   Typically these are written to a http.Request.
*/
type SnapshotPolicyScheduleCreateParams struct {

	/* Info.

	   Info specification
	*/
	Info *models.SnapshotPolicySchedule

	/* ReturnRecords.

	   The default is false.  If set to true, the records are returned.
	*/
	ReturnRecordsQueryParameter *bool

	/* SnapshotPolicyUUID.

	   Snapshot copy policy UUID
	*/
	SnapshotPolicyUUIDPathParameter string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the snapshot policy schedule create params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SnapshotPolicyScheduleCreateParams) WithDefaults() *SnapshotPolicyScheduleCreateParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the snapshot policy schedule create params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SnapshotPolicyScheduleCreateParams) SetDefaults() {
	var (
		returnRecordsQueryParameterDefault = bool(false)
	)

	val := SnapshotPolicyScheduleCreateParams{
		ReturnRecordsQueryParameter: &returnRecordsQueryParameterDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the snapshot policy schedule create params
func (o *SnapshotPolicyScheduleCreateParams) WithTimeout(timeout time.Duration) *SnapshotPolicyScheduleCreateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the snapshot policy schedule create params
func (o *SnapshotPolicyScheduleCreateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the snapshot policy schedule create params
func (o *SnapshotPolicyScheduleCreateParams) WithContext(ctx context.Context) *SnapshotPolicyScheduleCreateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the snapshot policy schedule create params
func (o *SnapshotPolicyScheduleCreateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the snapshot policy schedule create params
func (o *SnapshotPolicyScheduleCreateParams) WithHTTPClient(client *http.Client) *SnapshotPolicyScheduleCreateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the snapshot policy schedule create params
func (o *SnapshotPolicyScheduleCreateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithInfo adds the info to the snapshot policy schedule create params
func (o *SnapshotPolicyScheduleCreateParams) WithInfo(info *models.SnapshotPolicySchedule) *SnapshotPolicyScheduleCreateParams {
	o.SetInfo(info)
	return o
}

// SetInfo adds the info to the snapshot policy schedule create params
func (o *SnapshotPolicyScheduleCreateParams) SetInfo(info *models.SnapshotPolicySchedule) {
	o.Info = info
}

// WithReturnRecordsQueryParameter adds the returnRecords to the snapshot policy schedule create params
func (o *SnapshotPolicyScheduleCreateParams) WithReturnRecordsQueryParameter(returnRecords *bool) *SnapshotPolicyScheduleCreateParams {
	o.SetReturnRecordsQueryParameter(returnRecords)
	return o
}

// SetReturnRecordsQueryParameter adds the returnRecords to the snapshot policy schedule create params
func (o *SnapshotPolicyScheduleCreateParams) SetReturnRecordsQueryParameter(returnRecords *bool) {
	o.ReturnRecordsQueryParameter = returnRecords
}

// WithSnapshotPolicyUUIDPathParameter adds the snapshotPolicyUUID to the snapshot policy schedule create params
func (o *SnapshotPolicyScheduleCreateParams) WithSnapshotPolicyUUIDPathParameter(snapshotPolicyUUID string) *SnapshotPolicyScheduleCreateParams {
	o.SetSnapshotPolicyUUIDPathParameter(snapshotPolicyUUID)
	return o
}

// SetSnapshotPolicyUUIDPathParameter adds the snapshotPolicyUuid to the snapshot policy schedule create params
func (o *SnapshotPolicyScheduleCreateParams) SetSnapshotPolicyUUIDPathParameter(snapshotPolicyUUID string) {
	o.SnapshotPolicyUUIDPathParameter = snapshotPolicyUUID
}

// WriteToRequest writes these params to a swagger request
func (o *SnapshotPolicyScheduleCreateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// path param snapshot_policy.uuid
	if err := r.SetPathParam("snapshot_policy.uuid", o.SnapshotPolicyUUIDPathParameter); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}