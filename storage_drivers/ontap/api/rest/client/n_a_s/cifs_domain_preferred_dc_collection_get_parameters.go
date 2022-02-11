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

// NewCifsDomainPreferredDcCollectionGetParams creates a new CifsDomainPreferredDcCollectionGetParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewCifsDomainPreferredDcCollectionGetParams() *CifsDomainPreferredDcCollectionGetParams {
	return &CifsDomainPreferredDcCollectionGetParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewCifsDomainPreferredDcCollectionGetParamsWithTimeout creates a new CifsDomainPreferredDcCollectionGetParams object
// with the ability to set a timeout on a request.
func NewCifsDomainPreferredDcCollectionGetParamsWithTimeout(timeout time.Duration) *CifsDomainPreferredDcCollectionGetParams {
	return &CifsDomainPreferredDcCollectionGetParams{
		timeout: timeout,
	}
}

// NewCifsDomainPreferredDcCollectionGetParamsWithContext creates a new CifsDomainPreferredDcCollectionGetParams object
// with the ability to set a context for a request.
func NewCifsDomainPreferredDcCollectionGetParamsWithContext(ctx context.Context) *CifsDomainPreferredDcCollectionGetParams {
	return &CifsDomainPreferredDcCollectionGetParams{
		Context: ctx,
	}
}

// NewCifsDomainPreferredDcCollectionGetParamsWithHTTPClient creates a new CifsDomainPreferredDcCollectionGetParams object
// with the ability to set a custom HTTPClient for a request.
func NewCifsDomainPreferredDcCollectionGetParamsWithHTTPClient(client *http.Client) *CifsDomainPreferredDcCollectionGetParams {
	return &CifsDomainPreferredDcCollectionGetParams{
		HTTPClient: client,
	}
}

/* CifsDomainPreferredDcCollectionGetParams contains all the parameters to send to the API endpoint
   for the cifs domain preferred dc collection get operation.

   Typically these are written to a http.Request.
*/
type CifsDomainPreferredDcCollectionGetParams struct {

	/* Fields.

	   Specify the fields to return.
	*/
	FieldsQueryParameter []string

	/* Fqdn.

	   Filter by fqdn
	*/
	FqdnQueryParameter *string

	/* MaxRecords.

	   Limit the number of records returned.
	*/
	MaxRecordsQueryParameter *int64

	/* OrderBy.

	   Order results by specified fields and optional [asc|desc] direction. Default direction is 'asc' for ascending.
	*/
	OrderByQueryParameter []string

	/* ReturnRecords.

	   The default is true for GET calls.  When set to false, only the number of records is returned.

	   Default: true
	*/
	ReturnRecordsQueryParameter *bool

	/* ReturnTimeout.

	   The number of seconds to allow the call to execute before returning.  When iterating over a collection, the default is 15 seconds.  ONTAP returns earlier if either max records or the end of the collection is reached.

	   Default: 15
	*/
	ReturnTimeoutQueryParameter *int64

	/* ServerIP.

	   Filter by server_ip
	*/
	ServerIPQueryParameter *string

	/* SvmUUID.

	   UUID of the SVM to which this object belongs.
	*/
	SVMUUIDPathParameter string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the cifs domain preferred dc collection get params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CifsDomainPreferredDcCollectionGetParams) WithDefaults() *CifsDomainPreferredDcCollectionGetParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the cifs domain preferred dc collection get params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CifsDomainPreferredDcCollectionGetParams) SetDefaults() {
	var (
		returnRecordsQueryParameterDefault = bool(true)

		returnTimeoutQueryParameterDefault = int64(15)
	)

	val := CifsDomainPreferredDcCollectionGetParams{
		ReturnRecordsQueryParameter: &returnRecordsQueryParameterDefault,
		ReturnTimeoutQueryParameter: &returnTimeoutQueryParameterDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the cifs domain preferred dc collection get params
func (o *CifsDomainPreferredDcCollectionGetParams) WithTimeout(timeout time.Duration) *CifsDomainPreferredDcCollectionGetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the cifs domain preferred dc collection get params
func (o *CifsDomainPreferredDcCollectionGetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the cifs domain preferred dc collection get params
func (o *CifsDomainPreferredDcCollectionGetParams) WithContext(ctx context.Context) *CifsDomainPreferredDcCollectionGetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the cifs domain preferred dc collection get params
func (o *CifsDomainPreferredDcCollectionGetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the cifs domain preferred dc collection get params
func (o *CifsDomainPreferredDcCollectionGetParams) WithHTTPClient(client *http.Client) *CifsDomainPreferredDcCollectionGetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the cifs domain preferred dc collection get params
func (o *CifsDomainPreferredDcCollectionGetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithFieldsQueryParameter adds the fields to the cifs domain preferred dc collection get params
func (o *CifsDomainPreferredDcCollectionGetParams) WithFieldsQueryParameter(fields []string) *CifsDomainPreferredDcCollectionGetParams {
	o.SetFieldsQueryParameter(fields)
	return o
}

// SetFieldsQueryParameter adds the fields to the cifs domain preferred dc collection get params
func (o *CifsDomainPreferredDcCollectionGetParams) SetFieldsQueryParameter(fields []string) {
	o.FieldsQueryParameter = fields
}

// WithFqdnQueryParameter adds the fqdn to the cifs domain preferred dc collection get params
func (o *CifsDomainPreferredDcCollectionGetParams) WithFqdnQueryParameter(fqdn *string) *CifsDomainPreferredDcCollectionGetParams {
	o.SetFqdnQueryParameter(fqdn)
	return o
}

// SetFqdnQueryParameter adds the fqdn to the cifs domain preferred dc collection get params
func (o *CifsDomainPreferredDcCollectionGetParams) SetFqdnQueryParameter(fqdn *string) {
	o.FqdnQueryParameter = fqdn
}

// WithMaxRecordsQueryParameter adds the maxRecords to the cifs domain preferred dc collection get params
func (o *CifsDomainPreferredDcCollectionGetParams) WithMaxRecordsQueryParameter(maxRecords *int64) *CifsDomainPreferredDcCollectionGetParams {
	o.SetMaxRecordsQueryParameter(maxRecords)
	return o
}

// SetMaxRecordsQueryParameter adds the maxRecords to the cifs domain preferred dc collection get params
func (o *CifsDomainPreferredDcCollectionGetParams) SetMaxRecordsQueryParameter(maxRecords *int64) {
	o.MaxRecordsQueryParameter = maxRecords
}

// WithOrderByQueryParameter adds the orderBy to the cifs domain preferred dc collection get params
func (o *CifsDomainPreferredDcCollectionGetParams) WithOrderByQueryParameter(orderBy []string) *CifsDomainPreferredDcCollectionGetParams {
	o.SetOrderByQueryParameter(orderBy)
	return o
}

// SetOrderByQueryParameter adds the orderBy to the cifs domain preferred dc collection get params
func (o *CifsDomainPreferredDcCollectionGetParams) SetOrderByQueryParameter(orderBy []string) {
	o.OrderByQueryParameter = orderBy
}

// WithReturnRecordsQueryParameter adds the returnRecords to the cifs domain preferred dc collection get params
func (o *CifsDomainPreferredDcCollectionGetParams) WithReturnRecordsQueryParameter(returnRecords *bool) *CifsDomainPreferredDcCollectionGetParams {
	o.SetReturnRecordsQueryParameter(returnRecords)
	return o
}

// SetReturnRecordsQueryParameter adds the returnRecords to the cifs domain preferred dc collection get params
func (o *CifsDomainPreferredDcCollectionGetParams) SetReturnRecordsQueryParameter(returnRecords *bool) {
	o.ReturnRecordsQueryParameter = returnRecords
}

// WithReturnTimeoutQueryParameter adds the returnTimeout to the cifs domain preferred dc collection get params
func (o *CifsDomainPreferredDcCollectionGetParams) WithReturnTimeoutQueryParameter(returnTimeout *int64) *CifsDomainPreferredDcCollectionGetParams {
	o.SetReturnTimeoutQueryParameter(returnTimeout)
	return o
}

// SetReturnTimeoutQueryParameter adds the returnTimeout to the cifs domain preferred dc collection get params
func (o *CifsDomainPreferredDcCollectionGetParams) SetReturnTimeoutQueryParameter(returnTimeout *int64) {
	o.ReturnTimeoutQueryParameter = returnTimeout
}

// WithServerIPQueryParameter adds the serverIP to the cifs domain preferred dc collection get params
func (o *CifsDomainPreferredDcCollectionGetParams) WithServerIPQueryParameter(serverIP *string) *CifsDomainPreferredDcCollectionGetParams {
	o.SetServerIPQueryParameter(serverIP)
	return o
}

// SetServerIPQueryParameter adds the serverIp to the cifs domain preferred dc collection get params
func (o *CifsDomainPreferredDcCollectionGetParams) SetServerIPQueryParameter(serverIP *string) {
	o.ServerIPQueryParameter = serverIP
}

// WithSVMUUIDPathParameter adds the svmUUID to the cifs domain preferred dc collection get params
func (o *CifsDomainPreferredDcCollectionGetParams) WithSVMUUIDPathParameter(svmUUID string) *CifsDomainPreferredDcCollectionGetParams {
	o.SetSVMUUIDPathParameter(svmUUID)
	return o
}

// SetSVMUUIDPathParameter adds the svmUuid to the cifs domain preferred dc collection get params
func (o *CifsDomainPreferredDcCollectionGetParams) SetSVMUUIDPathParameter(svmUUID string) {
	o.SVMUUIDPathParameter = svmUUID
}

// WriteToRequest writes these params to a swagger request
func (o *CifsDomainPreferredDcCollectionGetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if o.FqdnQueryParameter != nil {

		// query param fqdn
		var qrFqdn string

		if o.FqdnQueryParameter != nil {
			qrFqdn = *o.FqdnQueryParameter
		}
		qFqdn := qrFqdn
		if qFqdn != "" {

			if err := r.SetQueryParam("fqdn", qFqdn); err != nil {
				return err
			}
		}
	}

	if o.MaxRecordsQueryParameter != nil {

		// query param max_records
		var qrMaxRecords int64

		if o.MaxRecordsQueryParameter != nil {
			qrMaxRecords = *o.MaxRecordsQueryParameter
		}
		qMaxRecords := swag.FormatInt64(qrMaxRecords)
		if qMaxRecords != "" {

			if err := r.SetQueryParam("max_records", qMaxRecords); err != nil {
				return err
			}
		}
	}

	if o.OrderByQueryParameter != nil {

		// binding items for order_by
		joinedOrderBy := o.bindParamOrderBy(reg)

		// query array param order_by
		if err := r.SetQueryParam("order_by", joinedOrderBy...); err != nil {
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

	if o.ServerIPQueryParameter != nil {

		// query param server_ip
		var qrServerIP string

		if o.ServerIPQueryParameter != nil {
			qrServerIP = *o.ServerIPQueryParameter
		}
		qServerIP := qrServerIP
		if qServerIP != "" {

			if err := r.SetQueryParam("server_ip", qServerIP); err != nil {
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

// bindParamCifsDomainPreferredDcCollectionGet binds the parameter fields
func (o *CifsDomainPreferredDcCollectionGetParams) bindParamFields(formats strfmt.Registry) []string {
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

// bindParamCifsDomainPreferredDcCollectionGet binds the parameter order_by
func (o *CifsDomainPreferredDcCollectionGetParams) bindParamOrderBy(formats strfmt.Registry) []string {
	orderByIR := o.OrderByQueryParameter

	var orderByIC []string
	for _, orderByIIR := range orderByIR { // explode []string

		orderByIIV := orderByIIR // string as string
		orderByIC = append(orderByIC, orderByIIV)
	}

	// items.CollectionFormat: "csv"
	orderByIS := swag.JoinByFormat(orderByIC, "csv")

	return orderByIS
}