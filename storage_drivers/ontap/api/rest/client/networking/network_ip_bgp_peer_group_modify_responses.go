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

// NetworkIPBgpPeerGroupModifyReader is a Reader for the NetworkIPBgpPeerGroupModify structure.
type NetworkIPBgpPeerGroupModifyReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *NetworkIPBgpPeerGroupModifyReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewNetworkIPBgpPeerGroupModifyOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewNetworkIPBgpPeerGroupModifyDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewNetworkIPBgpPeerGroupModifyOK creates a NetworkIPBgpPeerGroupModifyOK with default headers values
func NewNetworkIPBgpPeerGroupModifyOK() *NetworkIPBgpPeerGroupModifyOK {
	return &NetworkIPBgpPeerGroupModifyOK{}
}

/* NetworkIPBgpPeerGroupModifyOK describes a response with status code 200, with default header values.

OK
*/
type NetworkIPBgpPeerGroupModifyOK struct {
}

func (o *NetworkIPBgpPeerGroupModifyOK) Error() string {
	return fmt.Sprintf("[PATCH /network/ip/bgp/peer-groups/{uuid}][%d] networkIpBgpPeerGroupModifyOK ", 200)
}

func (o *NetworkIPBgpPeerGroupModifyOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewNetworkIPBgpPeerGroupModifyDefault creates a NetworkIPBgpPeerGroupModifyDefault with default headers values
func NewNetworkIPBgpPeerGroupModifyDefault(code int) *NetworkIPBgpPeerGroupModifyDefault {
	return &NetworkIPBgpPeerGroupModifyDefault{
		_statusCode: code,
	}
}

/* NetworkIPBgpPeerGroupModifyDefault describes a response with status code -1, with default header values.

 ONTAP Error Response Codes
| Error Code | Description |
| ---------- | ----------- |
| 1967171 | Internal error. Fail to access or update BGP peer group. Retry the command, if necessary. |
| 1967188 | Configuring peer address as a next hop requires an effective cluster version of 9.9.1 or later. |
| 53281998 | Failed to rename the BGP peer group because that name is already assigned to a different peer group in the IPspace. |
| 53282006 | BGP peer group could not be updated to use a peer address because the value provided is not a valid peer address. If necessary, try the command again with a routable host address. |
| 53282007 | BGP peer group could not be updated to use a peer address because the address represents a different address family to the address of the associated BGP LIF. If necessary, try the command again with a matching address family. |
| 53282018 | Failed to create BGP peer group because an existing peer group has already established a BGP session between LIF and peer address. If necessary, try the command again with a different BGP LIF or a different peer address. |

*/
type NetworkIPBgpPeerGroupModifyDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the network ip bgp peer group modify default response
func (o *NetworkIPBgpPeerGroupModifyDefault) Code() int {
	return o._statusCode
}

func (o *NetworkIPBgpPeerGroupModifyDefault) Error() string {
	return fmt.Sprintf("[PATCH /network/ip/bgp/peer-groups/{uuid}][%d] network_ip_bgp_peer_group_modify default  %+v", o._statusCode, o.Payload)
}
func (o *NetworkIPBgpPeerGroupModifyDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *NetworkIPBgpPeerGroupModifyDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}