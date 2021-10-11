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

// CifsSearchPathCreateReader is a Reader for the CifsSearchPathCreate structure.
type CifsSearchPathCreateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CifsSearchPathCreateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewCifsSearchPathCreateCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewCifsSearchPathCreateDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewCifsSearchPathCreateCreated creates a CifsSearchPathCreateCreated with default headers values
func NewCifsSearchPathCreateCreated() *CifsSearchPathCreateCreated {
	return &CifsSearchPathCreateCreated{}
}

/* CifsSearchPathCreateCreated describes a response with status code 201, with default header values.

Created
*/
type CifsSearchPathCreateCreated struct {
	Payload *models.CifsSearchPathResponse
}

func (o *CifsSearchPathCreateCreated) Error() string {
	return fmt.Sprintf("[POST /protocols/cifs/home-directory/search-paths][%d] cifsSearchPathCreateCreated  %+v", 201, o.Payload)
}
func (o *CifsSearchPathCreateCreated) GetPayload() *models.CifsSearchPathResponse {
	return o.Payload
}

func (o *CifsSearchPathCreateCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.CifsSearchPathResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCifsSearchPathCreateDefault creates a CifsSearchPathCreateDefault with default headers values
func NewCifsSearchPathCreateDefault(code int) *CifsSearchPathCreateDefault {
	return &CifsSearchPathCreateDefault{
		_statusCode: code,
	}
}

/* CifsSearchPathCreateDefault describes a response with status code -1, with default header values.

 ONTAP Error Response Codes
| Error Code | Description |
| ---------- | ----------- |
| 655551     | Invalid home-directory search-path path    |
| 655462     | The specified path is an invalid file-type |

*/
type CifsSearchPathCreateDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the cifs search path create default response
func (o *CifsSearchPathCreateDefault) Code() int {
	return o._statusCode
}

func (o *CifsSearchPathCreateDefault) Error() string {
	return fmt.Sprintf("[POST /protocols/cifs/home-directory/search-paths][%d] cifs_search_path_create default  %+v", o._statusCode, o.Payload)
}
func (o *CifsSearchPathCreateDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *CifsSearchPathCreateDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}