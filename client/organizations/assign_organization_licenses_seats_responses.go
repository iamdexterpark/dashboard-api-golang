// Code generated by go-swagger; DO NOT EDIT.

package organizations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"fmt"
	"io"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// AssignOrganizationLicensesSeatsReader is a Reader for the AssignOrganizationLicensesSeats structure.
type AssignOrganizationLicensesSeatsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AssignOrganizationLicensesSeatsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewAssignOrganizationLicensesSeatsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewAssignOrganizationLicensesSeatsOK creates a AssignOrganizationLicensesSeatsOK with default headers values
func NewAssignOrganizationLicensesSeatsOK() *AssignOrganizationLicensesSeatsOK {
	return &AssignOrganizationLicensesSeatsOK{}
}

/* AssignOrganizationLicensesSeatsOK describes a response with status code 200, with default header values.

Successful operation
*/
type AssignOrganizationLicensesSeatsOK struct {
	Payload interface{}
}

func (o *AssignOrganizationLicensesSeatsOK) Error() string {
	return fmt.Sprintf("[POST /organizations/{organizationId}/licenses/assignSeats][%d] assignOrganizationLicensesSeatsOK  %+v", 200, o.Payload)
}
func (o *AssignOrganizationLicensesSeatsOK) GetPayload() interface{} {
	return o.Payload
}

func (o *AssignOrganizationLicensesSeatsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*AssignOrganizationLicensesSeatsBody assign organization licenses seats body
// Example: {"licenseId":"1234","networkId":"N_24329156","seatCount":20}
swagger:model AssignOrganizationLicensesSeatsBody
*/
type AssignOrganizationLicensesSeatsBody struct {

	// The ID of the SM license to assign seats from
	// Required: true
	LicenseID *string `json:"licenseId"`

	// The ID of the SM network to assign the seats to
	// Required: true
	NetworkID *string `json:"networkId"`

	// The number of seats to assign to the SM network. Must be less than or equal to the total number of seats of the license
	// Required: true
	SeatCount *int64 `json:"seatCount"`
}

// Validate validates this assign organization licenses seats body
func (o *AssignOrganizationLicensesSeatsBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateLicenseID(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateNetworkID(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateSeatCount(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *AssignOrganizationLicensesSeatsBody) validateLicenseID(formats strfmt.Registry) error {

	if err := validate.Required("assignOrganizationLicensesSeats"+"."+"licenseId", "body", o.LicenseID); err != nil {
		return err
	}

	return nil
}

func (o *AssignOrganizationLicensesSeatsBody) validateNetworkID(formats strfmt.Registry) error {

	if err := validate.Required("assignOrganizationLicensesSeats"+"."+"networkId", "body", o.NetworkID); err != nil {
		return err
	}

	return nil
}

func (o *AssignOrganizationLicensesSeatsBody) validateSeatCount(formats strfmt.Registry) error {

	if err := validate.Required("assignOrganizationLicensesSeats"+"."+"seatCount", "body", o.SeatCount); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this assign organization licenses seats body based on context it is used
func (o *AssignOrganizationLicensesSeatsBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *AssignOrganizationLicensesSeatsBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *AssignOrganizationLicensesSeatsBody) UnmarshalBinary(b []byte) error {
	var res AssignOrganizationLicensesSeatsBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}