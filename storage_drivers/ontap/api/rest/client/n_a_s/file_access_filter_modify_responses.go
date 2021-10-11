// Code generated by go-swagger; DO NOT EDIT.

package n_a_s

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/netapp/trident/storage_drivers/ontap/api/rest/models"
)

// FileAccessFilterModifyReader is a Reader for the FileAccessFilterModify structure.
type FileAccessFilterModifyReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *FileAccessFilterModifyReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewFileAccessFilterModifyOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewFileAccessFilterModifyDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewFileAccessFilterModifyOK creates a FileAccessFilterModifyOK with default headers values
func NewFileAccessFilterModifyOK() *FileAccessFilterModifyOK {
	return &FileAccessFilterModifyOK{}
}

/* FileAccessFilterModifyOK describes a response with status code 200, with default header values.

OK
*/
type FileAccessFilterModifyOK struct {
}

func (o *FileAccessFilterModifyOK) Error() string {
	return fmt.Sprintf("[PATCH /protocols/file-access-tracing/filters/{svm.uuid}/{index}][%d] fileAccessFilterModifyOK ", 200)
}

func (o *FileAccessFilterModifyOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewFileAccessFilterModifyDefault creates a FileAccessFilterModifyDefault with default headers values
func NewFileAccessFilterModifyDefault(code int) *FileAccessFilterModifyDefault {
	return &FileAccessFilterModifyDefault{
		_statusCode: code,
	}
}

/* FileAccessFilterModifyDefault describes a response with status code -1, with default header values.

Error
*/
type FileAccessFilterModifyDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the file access filter modify default response
func (o *FileAccessFilterModifyDefault) Code() int {
	return o._statusCode
}

func (o *FileAccessFilterModifyDefault) Error() string {
	return fmt.Sprintf("[PATCH /protocols/file-access-tracing/filters/{svm.uuid}/{index}][%d] file_access_filter_modify default  %+v", o._statusCode, o.Payload)
}
func (o *FileAccessFilterModifyDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *FileAccessFilterModifyDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}