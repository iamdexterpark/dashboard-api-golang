// Code generated by go-swagger; DO NOT EDIT.

package insight

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// GetNetworkInsightApplicationHealthByTimeReader is a Reader for the GetNetworkInsightApplicationHealthByTime structure.
type GetNetworkInsightApplicationHealthByTimeReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetNetworkInsightApplicationHealthByTimeReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetNetworkInsightApplicationHealthByTimeOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetNetworkInsightApplicationHealthByTimeOK creates a GetNetworkInsightApplicationHealthByTimeOK with default headers values
func NewGetNetworkInsightApplicationHealthByTimeOK() *GetNetworkInsightApplicationHealthByTimeOK {
	return &GetNetworkInsightApplicationHealthByTimeOK{}
}

/* GetNetworkInsightApplicationHealthByTimeOK describes a response with status code 200, with default header values.

Successful operation
*/
type GetNetworkInsightApplicationHealthByTimeOK struct {
	Payload []interface{}
}

func (o *GetNetworkInsightApplicationHealthByTimeOK) Error() string {
	return fmt.Sprintf("[GET /networks/{networkId}/insight/applications/{applicationId}/healthByTime][%d] getNetworkInsightApplicationHealthByTimeOK  %+v", 200, o.Payload)
}
func (o *GetNetworkInsightApplicationHealthByTimeOK) GetPayload() []interface{} {
	return o.Payload
}

func (o *GetNetworkInsightApplicationHealthByTimeOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
