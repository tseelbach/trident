// Code generated by go-swagger; DO NOT EDIT.

package name_services

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/netapp/trident/storage_drivers/ontap/api/rest/models"
)

// UnixUserModifyReader is a Reader for the UnixUserModify structure.
type UnixUserModifyReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UnixUserModifyReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUnixUserModifyOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewUnixUserModifyDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewUnixUserModifyOK creates a UnixUserModifyOK with default headers values
func NewUnixUserModifyOK() *UnixUserModifyOK {
	return &UnixUserModifyOK{}
}

/* UnixUserModifyOK describes a response with status code 200, with default header values.

OK
*/
type UnixUserModifyOK struct {
}

func (o *UnixUserModifyOK) Error() string {
	return fmt.Sprintf("[PATCH /name-services/unix-users/{svm.uuid}/{name}][%d] unixUserModifyOK ", 200)
}

func (o *UnixUserModifyOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUnixUserModifyDefault creates a UnixUserModifyDefault with default headers values
func NewUnixUserModifyDefault(code int) *UnixUserModifyDefault {
	return &UnixUserModifyDefault{
		_statusCode: code,
	}
}

/* UnixUserModifyDefault describes a response with status code -1, with default header values.

 ONTAP Error Response Codes
| Error Code | Description |
| ---------- | ----------- |
| 23724128   | The specified Unix user full-name contains invalid character ':' |
| 23724089   | The specified UNIX user full-name is too long. Maximum supported length is 256 characters. |
| 23724055   | Internal error. Failed to modify the UNIX user for the SVM. Verify that the cluster is healthy, then try the command again. |
| 23724090   | Configuring individual entries is not supported because file-only configuration is enabled. |

*/
type UnixUserModifyDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the unix user modify default response
func (o *UnixUserModifyDefault) Code() int {
	return o._statusCode
}

func (o *UnixUserModifyDefault) Error() string {
	return fmt.Sprintf("[PATCH /name-services/unix-users/{svm.uuid}/{name}][%d] unix_user_modify default  %+v", o._statusCode, o.Payload)
}
func (o *UnixUserModifyDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *UnixUserModifyDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}