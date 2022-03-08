// Code generated by go-swagger; DO NOT EDIT.

package organizations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// GetOrganizationInventoryDeviceReader is a Reader for the GetOrganizationInventoryDevice structure.
type GetOrganizationInventoryDeviceReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetOrganizationInventoryDeviceReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetOrganizationInventoryDeviceOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetOrganizationInventoryDeviceOK creates a GetOrganizationInventoryDeviceOK with default headers values
func NewGetOrganizationInventoryDeviceOK() *GetOrganizationInventoryDeviceOK {
	return &GetOrganizationInventoryDeviceOK{}
}

/* GetOrganizationInventoryDeviceOK describes a response with status code 200, with default header values.

Successful operation
*/
type GetOrganizationInventoryDeviceOK struct {
	Payload interface{}
}

func (o *GetOrganizationInventoryDeviceOK) Error() string {
	return fmt.Sprintf("[GET /organizations/{organizationId}/inventory/devices/{serial}][%d] getOrganizationInventoryDeviceOK  %+v", 200, o.Payload)
}
func (o *GetOrganizationInventoryDeviceOK) GetPayload() interface{} {
	return o.Payload
}

func (o *GetOrganizationInventoryDeviceOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
