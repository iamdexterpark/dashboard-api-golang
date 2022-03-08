// Code generated by go-swagger; DO NOT EDIT.

package networks

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// GetNetworkNetworkHealthChannelUtilizationReader is a Reader for the GetNetworkNetworkHealthChannelUtilization structure.
type GetNetworkNetworkHealthChannelUtilizationReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetNetworkNetworkHealthChannelUtilizationReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetNetworkNetworkHealthChannelUtilizationOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetNetworkNetworkHealthChannelUtilizationOK creates a GetNetworkNetworkHealthChannelUtilizationOK with default headers values
func NewGetNetworkNetworkHealthChannelUtilizationOK() *GetNetworkNetworkHealthChannelUtilizationOK {
	return &GetNetworkNetworkHealthChannelUtilizationOK{}
}

/* GetNetworkNetworkHealthChannelUtilizationOK describes a response with status code 200, with default header values.

Successful operation
*/
type GetNetworkNetworkHealthChannelUtilizationOK struct {

	/* A comma-separated list of first, last, prev, and next relative links used for subsequent paginated requests.
	 */
	Link string

	Payload []interface{}
}

func (o *GetNetworkNetworkHealthChannelUtilizationOK) Error() string {
	return fmt.Sprintf("[GET /networks/{networkId}/networkHealth/channelUtilization][%d] getNetworkNetworkHealthChannelUtilizationOK  %+v", 200, o.Payload)
}
func (o *GetNetworkNetworkHealthChannelUtilizationOK) GetPayload() []interface{} {
	return o.Payload
}

func (o *GetNetworkNetworkHealthChannelUtilizationOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header Link
	hdrLink := response.GetHeader("Link")

	if hdrLink != "" {
		o.Link = hdrLink
	}

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
