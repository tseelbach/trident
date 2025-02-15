// Code generated by go-swagger; DO NOT EDIT.

package networking

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/netapp/trident/storage_drivers/ontap/api/rest/models"
)

// FcInterfaceCreateReader is a Reader for the FcInterfaceCreate structure.
type FcInterfaceCreateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *FcInterfaceCreateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewFcInterfaceCreateCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewFcInterfaceCreateDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewFcInterfaceCreateCreated creates a FcInterfaceCreateCreated with default headers values
func NewFcInterfaceCreateCreated() *FcInterfaceCreateCreated {
	return &FcInterfaceCreateCreated{}
}

/*
FcInterfaceCreateCreated describes a response with status code 201, with default header values.

Created
*/
type FcInterfaceCreateCreated struct {
	Payload *models.FcInterfaceResponse
}

// IsSuccess returns true when this fc interface create created response has a 2xx status code
func (o *FcInterfaceCreateCreated) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this fc interface create created response has a 3xx status code
func (o *FcInterfaceCreateCreated) IsRedirect() bool {
	return false
}

// IsClientError returns true when this fc interface create created response has a 4xx status code
func (o *FcInterfaceCreateCreated) IsClientError() bool {
	return false
}

// IsServerError returns true when this fc interface create created response has a 5xx status code
func (o *FcInterfaceCreateCreated) IsServerError() bool {
	return false
}

// IsCode returns true when this fc interface create created response a status code equal to that given
func (o *FcInterfaceCreateCreated) IsCode(code int) bool {
	return code == 201
}

func (o *FcInterfaceCreateCreated) Error() string {
	return fmt.Sprintf("[POST /network/fc/interfaces][%d] fcInterfaceCreateCreated  %+v", 201, o.Payload)
}

func (o *FcInterfaceCreateCreated) String() string {
	return fmt.Sprintf("[POST /network/fc/interfaces][%d] fcInterfaceCreateCreated  %+v", 201, o.Payload)
}

func (o *FcInterfaceCreateCreated) GetPayload() *models.FcInterfaceResponse {
	return o.Payload
}

func (o *FcInterfaceCreateCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.FcInterfaceResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewFcInterfaceCreateDefault creates a FcInterfaceCreateDefault with default headers values
func NewFcInterfaceCreateDefault(code int) *FcInterfaceCreateDefault {
	return &FcInterfaceCreateDefault{
		_statusCode: code,
	}
}

/*
	FcInterfaceCreateDefault describes a response with status code -1, with default header values.

	ONTAP Error Response Codes

| Error Code | Description |
| ---------- | ----------- |
| 1966140 | An interface with the same name already exists. |
| 1966217 | The specified port is not valid on the node provided. |
| 2621462 | The supplied SVM does not exist. |
| 2621706 | The specified `svm.uuid` and `svm.name` do not refer to the same SVM. |
| 2621707 | No SVM was specified. Either `svm.name` or `svm.uuid` must be supplied. |
| 5373966 | A Fibre Channel interface with the _fcp_ protocol cannot be created in an SVM that is configured for NVMe. |
| 5374102 | The specified Fibre Channel interface cannot be created because the Fibre Channel adapter is down. Bring the adapter up and try again. |
| 5374871 | The Fibre Channel port identified by the specified UUID does not refer to the same port as that identified by the specified node name and/or port name. |
| 5374872 | If either `location.port.node.name` or `location.port.name` is supplied, both properties must be supplied. |
| 5374873 | The Fibre Channel port must be specified using either `location.port.uuid` or `location.port.node.name` and `location.port.name`. |
| 72089652 | An NVMe service must be created before creating a Fibre Channel interface using the NVMe over FC data protocol. |
| 72089672 | The specified Fibre Channel port does not support the NVMe over FC data protocol. |
| 72089900 | A Fibre Channel interface with the _fc\_nvme_ protocol cannot be created in an SVM that is configured for a SAN protocol. |
*/
type FcInterfaceCreateDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the fc interface create default response
func (o *FcInterfaceCreateDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this fc interface create default response has a 2xx status code
func (o *FcInterfaceCreateDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this fc interface create default response has a 3xx status code
func (o *FcInterfaceCreateDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this fc interface create default response has a 4xx status code
func (o *FcInterfaceCreateDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this fc interface create default response has a 5xx status code
func (o *FcInterfaceCreateDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this fc interface create default response a status code equal to that given
func (o *FcInterfaceCreateDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *FcInterfaceCreateDefault) Error() string {
	return fmt.Sprintf("[POST /network/fc/interfaces][%d] fc_interface_create default  %+v", o._statusCode, o.Payload)
}

func (o *FcInterfaceCreateDefault) String() string {
	return fmt.Sprintf("[POST /network/fc/interfaces][%d] fc_interface_create default  %+v", o._statusCode, o.Payload)
}

func (o *FcInterfaceCreateDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *FcInterfaceCreateDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
