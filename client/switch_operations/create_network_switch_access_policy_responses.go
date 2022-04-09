// Code generated by go-swagger; DO NOT EDIT.

package switch_operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// CreateNetworkSwitchAccessPolicyReader is a Reader for the CreateNetworkSwitchAccessPolicy structure.
type CreateNetworkSwitchAccessPolicyReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateNetworkSwitchAccessPolicyReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewCreateNetworkSwitchAccessPolicyCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewCreateNetworkSwitchAccessPolicyCreated creates a CreateNetworkSwitchAccessPolicyCreated with default headers values
func NewCreateNetworkSwitchAccessPolicyCreated() *CreateNetworkSwitchAccessPolicyCreated {
	return &CreateNetworkSwitchAccessPolicyCreated{}
}

/* CreateNetworkSwitchAccessPolicyCreated describes a response with status code 201, with default header values.

Successful operation
*/
type CreateNetworkSwitchAccessPolicyCreated struct {
	Payload interface{}
}

func (o *CreateNetworkSwitchAccessPolicyCreated) Error() string {
	return fmt.Sprintf("[POST /networks/{networkId}/switch/accessPolicies][%d] createNetworkSwitchAccessPolicyCreated  %+v", 201, o.Payload)
}
func (o *CreateNetworkSwitchAccessPolicyCreated) GetPayload() interface{} {
	return o.Payload
}

func (o *CreateNetworkSwitchAccessPolicyCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*CreateNetworkSwitchAccessPolicyBody create network switch access policy body
// Example: {"accessPolicyType":"Hybrid authentication","dot1x":{"controlDirection":"inbound"},"guestVlanId":100,"hostMode":"Single-Host","increaseAccessSpeed":false,"name":"Access policy #1","radius":{"criticalAuth":{"dataVlanId":100,"suspendPortBounce":true,"voiceVlanId":100},"failedAuthVlanId":100,"reAuthenticationInterval":120,"suspendReAuthentication":true},"radiusAccountingEnabled":true,"radiusAccountingServers":[{"host":"1.2.3.4","port":22,"secret":"secret"}],"radiusCoaSupportEnabled":false,"radiusGroupAttribute":"11","radiusServers":[{"host":"1.2.3.4","port":22,"secret":"secret"}],"radiusTestingEnabled":false,"urlRedirectWalledGardenEnabled":true,"urlRedirectWalledGardenRanges":["192.168.1.0/24"],"voiceVlanClients":true}
swagger:model CreateNetworkSwitchAccessPolicyBody
*/
type CreateNetworkSwitchAccessPolicyBody struct {

	// Access Type of the policy. Automatically 'Hybrid authentication' when hostMode is 'Multi-Domain'.
	// Enum: [802.1x MAC authentication bypass Hybrid authentication]
	AccessPolicyType string `json:"accessPolicyType,omitempty"`

	// ID for the guest VLAN allow unauthorized devices access to limited network resources
	GuestVlanID int64 `json:"guestVlanId,omitempty"`

	// Choose the Host Mode for the access policy.
	// Required: true
	// Enum: [Single-Host Multi-Domain Multi-Host Multi-Auth]
	HostMode *string `json:"hostMode"`

	// Enabling this option will make switches execute 802.1X and MAC-bypass authentication simultaneously so that clients authenticate faster. Only required when accessPolicyType is 'Hybrid Authentication.
	IncreaseAccessSpeed bool `json:"increaseAccessSpeed,omitempty"`

	// Name of the access policy
	// Required: true
	Name *string `json:"name"`

	// radius
	Radius *CreateNetworkSwitchAccessPolicyParamsBodyRadius `json:"radius,omitempty"`

	// Enable to send start, interim-update and stop messages to a configured RADIUS accounting server for tracking connected clients
	// Required: true
	RadiusAccountingEnabled *bool `json:"radiusAccountingEnabled"`

	// List of RADIUS accounting servers to require connecting devices to authenticate against before granting network access
	RadiusAccountingServers []*CreateNetworkSwitchAccessPolicyParamsBodyRadiusAccountingServersItems0 `json:"radiusAccountingServers"`

	// Change of authentication for RADIUS re-authentication and disconnection
	// Required: true
	RadiusCoaSupportEnabled *bool `json:"radiusCoaSupportEnabled"`

	// Acceptable values are `""` for None, or `"11"` for Group Policies ACL
	RadiusGroupAttribute string `json:"radiusGroupAttribute,omitempty"`

	// List of RADIUS servers to require connecting devices to authenticate against before granting network access
	// Required: true
	RadiusServers []*CreateNetworkSwitchAccessPolicyParamsBodyRadiusServersItems0 `json:"radiusServers"`

	// If enabled, Meraki devices will periodically send access-request messages to these RADIUS servers
	// Required: true
	RadiusTestingEnabled *bool `json:"radiusTestingEnabled"`

	// Enable to restrict access for clients to a specific set of IP addresses or hostnames prior to authentication
	// Required: true
	URLRedirectWalledGardenEnabled *bool `json:"urlRedirectWalledGardenEnabled"`

	// IP address ranges, in CIDR notation, to restrict access for clients to a specific set of IP addresses or hostnames prior to authentication
	URLRedirectWalledGardenRanges []string `json:"urlRedirectWalledGardenRanges"`

	// CDP/LLDP capable voice clients will be able to use this VLAN. Automatically true when hostMode is 'Multi-Domain'.
	VoiceVlanClients bool `json:"voiceVlanClients,omitempty"`
}

// Validate validates this create network switch access policy body
func (o *CreateNetworkSwitchAccessPolicyBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateAccessPolicyType(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateHostMode(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateRadius(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateRadiusAccountingEnabled(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateRadiusAccountingServers(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateRadiusCoaSupportEnabled(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateRadiusServers(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateRadiusTestingEnabled(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateURLRedirectWalledGardenEnabled(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var createNetworkSwitchAccessPolicyBodyTypeAccessPolicyTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["802.1x","MAC authentication bypass","Hybrid authentication"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		createNetworkSwitchAccessPolicyBodyTypeAccessPolicyTypePropEnum = append(createNetworkSwitchAccessPolicyBodyTypeAccessPolicyTypePropEnum, v)
	}
}

const (

	// CreateNetworkSwitchAccessPolicyBodyAccessPolicyTypeNr802Dot1x captures enum value "802.1x"
	CreateNetworkSwitchAccessPolicyBodyAccessPolicyTypeNr802Dot1x string = "802.1x"

	// CreateNetworkSwitchAccessPolicyBodyAccessPolicyTypeMACAuthenticationBypass captures enum value "MAC authentication bypass"
	CreateNetworkSwitchAccessPolicyBodyAccessPolicyTypeMACAuthenticationBypass string = "MAC authentication bypass"

	// CreateNetworkSwitchAccessPolicyBodyAccessPolicyTypeHybridAuthentication captures enum value "Hybrid authentication"
	CreateNetworkSwitchAccessPolicyBodyAccessPolicyTypeHybridAuthentication string = "Hybrid authentication"
)

// prop value enum
func (o *CreateNetworkSwitchAccessPolicyBody) validateAccessPolicyTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, createNetworkSwitchAccessPolicyBodyTypeAccessPolicyTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (o *CreateNetworkSwitchAccessPolicyBody) validateAccessPolicyType(formats strfmt.Registry) error {
	if swag.IsZero(o.AccessPolicyType) { // not required
		return nil
	}

	// value enum
	if err := o.validateAccessPolicyTypeEnum("createNetworkSwitchAccessPolicy"+"."+"accessPolicyType", "body", o.AccessPolicyType); err != nil {
		return err
	}

	return nil
}

var createNetworkSwitchAccessPolicyBodyTypeHostModePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["Single-Host","Multi-Domain","Multi-Host","Multi-Auth"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		createNetworkSwitchAccessPolicyBodyTypeHostModePropEnum = append(createNetworkSwitchAccessPolicyBodyTypeHostModePropEnum, v)
	}
}

const (

	// CreateNetworkSwitchAccessPolicyBodyHostModeSingleDashHost captures enum value "Single-Host"
	CreateNetworkSwitchAccessPolicyBodyHostModeSingleDashHost string = "Single-Host"

	// CreateNetworkSwitchAccessPolicyBodyHostModeMultiDashDomain captures enum value "Multi-Domain"
	CreateNetworkSwitchAccessPolicyBodyHostModeMultiDashDomain string = "Multi-Domain"

	// CreateNetworkSwitchAccessPolicyBodyHostModeMultiDashHost captures enum value "Multi-Host"
	CreateNetworkSwitchAccessPolicyBodyHostModeMultiDashHost string = "Multi-Host"

	// CreateNetworkSwitchAccessPolicyBodyHostModeMultiDashAuth captures enum value "Multi-Auth"
	CreateNetworkSwitchAccessPolicyBodyHostModeMultiDashAuth string = "Multi-Auth"
)

// prop value enum
func (o *CreateNetworkSwitchAccessPolicyBody) validateHostModeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, createNetworkSwitchAccessPolicyBodyTypeHostModePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (o *CreateNetworkSwitchAccessPolicyBody) validateHostMode(formats strfmt.Registry) error {

	if err := validate.Required("createNetworkSwitchAccessPolicy"+"."+"hostMode", "body", o.HostMode); err != nil {
		return err
	}

	// value enum
	if err := o.validateHostModeEnum("createNetworkSwitchAccessPolicy"+"."+"hostMode", "body", *o.HostMode); err != nil {
		return err
	}

	return nil
}

func (o *CreateNetworkSwitchAccessPolicyBody) validateName(formats strfmt.Registry) error {

	if err := validate.Required("createNetworkSwitchAccessPolicy"+"."+"name", "body", o.Name); err != nil {
		return err
	}

	return nil
}

func (o *CreateNetworkSwitchAccessPolicyBody) validateRadius(formats strfmt.Registry) error {
	if swag.IsZero(o.Radius) { // not required
		return nil
	}

	if o.Radius != nil {
		if err := o.Radius.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("createNetworkSwitchAccessPolicy" + "." + "radius")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("createNetworkSwitchAccessPolicy" + "." + "radius")
			}
			return err
		}
	}

	return nil
}

func (o *CreateNetworkSwitchAccessPolicyBody) validateRadiusAccountingEnabled(formats strfmt.Registry) error {

	if err := validate.Required("createNetworkSwitchAccessPolicy"+"."+"radiusAccountingEnabled", "body", o.RadiusAccountingEnabled); err != nil {
		return err
	}

	return nil
}

func (o *CreateNetworkSwitchAccessPolicyBody) validateRadiusAccountingServers(formats strfmt.Registry) error {
	if swag.IsZero(o.RadiusAccountingServers) { // not required
		return nil
	}

	for i := 0; i < len(o.RadiusAccountingServers); i++ {
		if swag.IsZero(o.RadiusAccountingServers[i]) { // not required
			continue
		}

		if o.RadiusAccountingServers[i] != nil {
			if err := o.RadiusAccountingServers[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("createNetworkSwitchAccessPolicy" + "." + "radiusAccountingServers" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("createNetworkSwitchAccessPolicy" + "." + "radiusAccountingServers" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (o *CreateNetworkSwitchAccessPolicyBody) validateRadiusCoaSupportEnabled(formats strfmt.Registry) error {

	if err := validate.Required("createNetworkSwitchAccessPolicy"+"."+"radiusCoaSupportEnabled", "body", o.RadiusCoaSupportEnabled); err != nil {
		return err
	}

	return nil
}

func (o *CreateNetworkSwitchAccessPolicyBody) validateRadiusServers(formats strfmt.Registry) error {

	if err := validate.Required("createNetworkSwitchAccessPolicy"+"."+"radiusServers", "body", o.RadiusServers); err != nil {
		return err
	}

	for i := 0; i < len(o.RadiusServers); i++ {
		if swag.IsZero(o.RadiusServers[i]) { // not required
			continue
		}

		if o.RadiusServers[i] != nil {
			if err := o.RadiusServers[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("createNetworkSwitchAccessPolicy" + "." + "radiusServers" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("createNetworkSwitchAccessPolicy" + "." + "radiusServers" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (o *CreateNetworkSwitchAccessPolicyBody) validateRadiusTestingEnabled(formats strfmt.Registry) error {

	if err := validate.Required("createNetworkSwitchAccessPolicy"+"."+"radiusTestingEnabled", "body", o.RadiusTestingEnabled); err != nil {
		return err
	}

	return nil
}

func (o *CreateNetworkSwitchAccessPolicyBody) validateURLRedirectWalledGardenEnabled(formats strfmt.Registry) error {

	if err := validate.Required("createNetworkSwitchAccessPolicy"+"."+"urlRedirectWalledGardenEnabled", "body", o.URLRedirectWalledGardenEnabled); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this create network switch access policy body based on the context it is used
func (o *CreateNetworkSwitchAccessPolicyBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateRadius(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := o.contextValidateRadiusAccountingServers(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := o.contextValidateRadiusServers(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *CreateNetworkSwitchAccessPolicyBody) contextValidateRadius(ctx context.Context, formats strfmt.Registry) error {

	if o.Radius != nil {
		if err := o.Radius.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("createNetworkSwitchAccessPolicy" + "." + "radius")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("createNetworkSwitchAccessPolicy" + "." + "radius")
			}
			return err
		}
	}

	return nil
}

func (o *CreateNetworkSwitchAccessPolicyBody) contextValidateRadiusAccountingServers(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(o.RadiusAccountingServers); i++ {

		if o.RadiusAccountingServers[i] != nil {
			if err := o.RadiusAccountingServers[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("createNetworkSwitchAccessPolicy" + "." + "radiusAccountingServers" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("createNetworkSwitchAccessPolicy" + "." + "radiusAccountingServers" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (o *CreateNetworkSwitchAccessPolicyBody) contextValidateRadiusServers(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(o.RadiusServers); i++ {

		if o.RadiusServers[i] != nil {
			if err := o.RadiusServers[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("createNetworkSwitchAccessPolicy" + "." + "radiusServers" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("createNetworkSwitchAccessPolicy" + "." + "radiusServers" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (o *CreateNetworkSwitchAccessPolicyBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *CreateNetworkSwitchAccessPolicyBody) UnmarshalBinary(b []byte) error {
	var res CreateNetworkSwitchAccessPolicyBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*CreateNetworkSwitchAccessPolicyParamsBodyRadius Object for RADIUS Settings
swagger:model CreateNetworkSwitchAccessPolicyParamsBodyRadius
*/
type CreateNetworkSwitchAccessPolicyParamsBodyRadius struct {

	// critical auth
	CriticalAuth *CreateNetworkSwitchAccessPolicyParamsBodyRadiusCriticalAuth `json:"criticalAuth,omitempty"`
}

// Validate validates this create network switch access policy params body radius
func (o *CreateNetworkSwitchAccessPolicyParamsBodyRadius) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateCriticalAuth(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *CreateNetworkSwitchAccessPolicyParamsBodyRadius) validateCriticalAuth(formats strfmt.Registry) error {
	if swag.IsZero(o.CriticalAuth) { // not required
		return nil
	}

	if o.CriticalAuth != nil {
		if err := o.CriticalAuth.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("createNetworkSwitchAccessPolicy" + "." + "radius" + "." + "criticalAuth")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("createNetworkSwitchAccessPolicy" + "." + "radius" + "." + "criticalAuth")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this create network switch access policy params body radius based on the context it is used
func (o *CreateNetworkSwitchAccessPolicyParamsBodyRadius) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateCriticalAuth(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *CreateNetworkSwitchAccessPolicyParamsBodyRadius) contextValidateCriticalAuth(ctx context.Context, formats strfmt.Registry) error {

	if o.CriticalAuth != nil {
		if err := o.CriticalAuth.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("createNetworkSwitchAccessPolicy" + "." + "radius" + "." + "criticalAuth")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("createNetworkSwitchAccessPolicy" + "." + "radius" + "." + "criticalAuth")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *CreateNetworkSwitchAccessPolicyParamsBodyRadius) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *CreateNetworkSwitchAccessPolicyParamsBodyRadius) UnmarshalBinary(b []byte) error {
	var res CreateNetworkSwitchAccessPolicyParamsBodyRadius
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*CreateNetworkSwitchAccessPolicyParamsBodyRadiusAccountingServersItems0 create network switch access policy params body radius accounting servers items0
swagger:model CreateNetworkSwitchAccessPolicyParamsBodyRadiusAccountingServersItems0
*/
type CreateNetworkSwitchAccessPolicyParamsBodyRadiusAccountingServersItems0 struct {

	// Public IP address of the RADIUS accounting server
	// Required: true
	Host *string `json:"host"`

	// UDP port that the RADIUS Accounting server listens on for access requests
	// Required: true
	Port *int64 `json:"port"`

	// RADIUS client shared secret
	// Required: true
	Secret *string `json:"secret"`
}

// Validate validates this create network switch access policy params body radius accounting servers items0
func (o *CreateNetworkSwitchAccessPolicyParamsBodyRadiusAccountingServersItems0) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateHost(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validatePort(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateSecret(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *CreateNetworkSwitchAccessPolicyParamsBodyRadiusAccountingServersItems0) validateHost(formats strfmt.Registry) error {

	if err := validate.Required("host", "body", o.Host); err != nil {
		return err
	}

	return nil
}

func (o *CreateNetworkSwitchAccessPolicyParamsBodyRadiusAccountingServersItems0) validatePort(formats strfmt.Registry) error {

	if err := validate.Required("port", "body", o.Port); err != nil {
		return err
	}

	return nil
}

func (o *CreateNetworkSwitchAccessPolicyParamsBodyRadiusAccountingServersItems0) validateSecret(formats strfmt.Registry) error {

	if err := validate.Required("secret", "body", o.Secret); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this create network switch access policy params body radius accounting servers items0 based on context it is used
func (o *CreateNetworkSwitchAccessPolicyParamsBodyRadiusAccountingServersItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *CreateNetworkSwitchAccessPolicyParamsBodyRadiusAccountingServersItems0) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *CreateNetworkSwitchAccessPolicyParamsBodyRadiusAccountingServersItems0) UnmarshalBinary(b []byte) error {
	var res CreateNetworkSwitchAccessPolicyParamsBodyRadiusAccountingServersItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*CreateNetworkSwitchAccessPolicyParamsBodyRadiusCriticalAuth Critical auth settings for when authentication is rejected by the RADIUS server
swagger:model CreateNetworkSwitchAccessPolicyParamsBodyRadiusCriticalAuth
*/
type CreateNetworkSwitchAccessPolicyParamsBodyRadiusCriticalAuth struct {

	// VLAN that clients who use data will be placed on when RADIUS authentication fails. Will be null if hostMode is Multi-Auth
	DataVlanID int64 `json:"dataVlanId,omitempty"`

	// Enable to suspend port bounce when RADIUS servers are unreachable
	SuspendPortBounce bool `json:"suspendPortBounce,omitempty"`

	// VLAN that clients who use voice will be placed on when RADIUS authentication fails. Will be null if hostMode is Multi-Auth
	VoiceVlanID int64 `json:"voiceVlanId,omitempty"`
}

// Validate validates this create network switch access policy params body radius critical auth
func (o *CreateNetworkSwitchAccessPolicyParamsBodyRadiusCriticalAuth) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this create network switch access policy params body radius critical auth based on context it is used
func (o *CreateNetworkSwitchAccessPolicyParamsBodyRadiusCriticalAuth) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *CreateNetworkSwitchAccessPolicyParamsBodyRadiusCriticalAuth) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *CreateNetworkSwitchAccessPolicyParamsBodyRadiusCriticalAuth) UnmarshalBinary(b []byte) error {
	var res CreateNetworkSwitchAccessPolicyParamsBodyRadiusCriticalAuth
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*CreateNetworkSwitchAccessPolicyParamsBodyRadiusServersItems0 create network switch access policy params body radius servers items0
swagger:model CreateNetworkSwitchAccessPolicyParamsBodyRadiusServersItems0
*/
type CreateNetworkSwitchAccessPolicyParamsBodyRadiusServersItems0 struct {

	// Public IP address of the RADIUS server
	// Required: true
	Host *string `json:"host"`

	// UDP port that the RADIUS server listens on for access requests
	// Required: true
	Port *int64 `json:"port"`

	// RADIUS client shared secret
	// Required: true
	Secret *string `json:"secret"`
}

// Validate validates this create network switch access policy params body radius servers items0
func (o *CreateNetworkSwitchAccessPolicyParamsBodyRadiusServersItems0) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateHost(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validatePort(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateSecret(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *CreateNetworkSwitchAccessPolicyParamsBodyRadiusServersItems0) validateHost(formats strfmt.Registry) error {

	if err := validate.Required("host", "body", o.Host); err != nil {
		return err
	}

	return nil
}

func (o *CreateNetworkSwitchAccessPolicyParamsBodyRadiusServersItems0) validatePort(formats strfmt.Registry) error {

	if err := validate.Required("port", "body", o.Port); err != nil {
		return err
	}

	return nil
}

func (o *CreateNetworkSwitchAccessPolicyParamsBodyRadiusServersItems0) validateSecret(formats strfmt.Registry) error {

	if err := validate.Required("secret", "body", o.Secret); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this create network switch access policy params body radius servers items0 based on context it is used
func (o *CreateNetworkSwitchAccessPolicyParamsBodyRadiusServersItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *CreateNetworkSwitchAccessPolicyParamsBodyRadiusServersItems0) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *CreateNetworkSwitchAccessPolicyParamsBodyRadiusServersItems0) UnmarshalBinary(b []byte) error {
	var res CreateNetworkSwitchAccessPolicyParamsBodyRadiusServersItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
