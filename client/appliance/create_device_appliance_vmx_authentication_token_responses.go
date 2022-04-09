// Code generated by go-swagger; DO NOT EDIT.

package appliance

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// CreateDeviceApplianceVmxAuthenticationTokenReader is a Reader for the CreateDeviceApplianceVmxAuthenticationToken structure.
type CreateDeviceApplianceVmxAuthenticationTokenReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateDeviceApplianceVmxAuthenticationTokenReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewCreateDeviceApplianceVmxAuthenticationTokenCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewCreateDeviceApplianceVmxAuthenticationTokenCreated creates a CreateDeviceApplianceVmxAuthenticationTokenCreated with default headers values
func NewCreateDeviceApplianceVmxAuthenticationTokenCreated() *CreateDeviceApplianceVmxAuthenticationTokenCreated {
	return &CreateDeviceApplianceVmxAuthenticationTokenCreated{}
}

/* CreateDeviceApplianceVmxAuthenticationTokenCreated describes a response with status code 201, with default header values.

Successful operation
*/
type CreateDeviceApplianceVmxAuthenticationTokenCreated struct {
	Payload interface{}
}

func (o *CreateDeviceApplianceVmxAuthenticationTokenCreated) Error() string {
	return fmt.Sprintf("[POST /devices/{serial}/appliance/vmx/authenticationToken][%d] createDeviceApplianceVmxAuthenticationTokenCreated  %+v", 201, o.Payload)
}
func (o *CreateDeviceApplianceVmxAuthenticationTokenCreated) GetPayload() interface{} {
	return o.Payload
}

func (o *CreateDeviceApplianceVmxAuthenticationTokenCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
