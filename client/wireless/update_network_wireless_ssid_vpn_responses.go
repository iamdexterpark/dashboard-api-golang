// Code generated by go-swagger; DO NOT EDIT.

package wireless

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

// UpdateNetworkWirelessSsidVpnReader is a Reader for the UpdateNetworkWirelessSsidVpn structure.
type UpdateNetworkWirelessSsidVpnReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateNetworkWirelessSsidVpnReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUpdateNetworkWirelessSsidVpnOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewUpdateNetworkWirelessSsidVpnOK creates a UpdateNetworkWirelessSsidVpnOK with default headers values
func NewUpdateNetworkWirelessSsidVpnOK() *UpdateNetworkWirelessSsidVpnOK {
	return &UpdateNetworkWirelessSsidVpnOK{}
}

/* UpdateNetworkWirelessSsidVpnOK describes a response with status code 200, with default header values.

Successful operation
*/
type UpdateNetworkWirelessSsidVpnOK struct {
	Payload interface{}
}

func (o *UpdateNetworkWirelessSsidVpnOK) Error() string {
	return fmt.Sprintf("[PUT /networks/{networkId}/wireless/ssids/{number}/vpn][%d] updateNetworkWirelessSsidVpnOK  %+v", 200, o.Payload)
}
func (o *UpdateNetworkWirelessSsidVpnOK) GetPayload() interface{} {
	return o.Payload
}

func (o *UpdateNetworkWirelessSsidVpnOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*UpdateNetworkWirelessSsidVpnBody update network wireless ssid vpn body
// Example: {"concentrator":{"name":"some concentrator name","networkId":"N_123"},"failover":{"heartbeatInterval":10,"idleTimeout":30,"requestIp":"1.1.1.1"},"splitTunnel":{"enabled":true,"rules":[{"comment":"split tunnel rule 1","destCidr":"1.1.1.1/32","destPort":"any","policy":"allow","protocol":"Any"},{"comment":"split tunnel rule 2","destCidr":"foo.com","destPort":"any","policy":"deny"}]}}
swagger:model UpdateNetworkWirelessSsidVpnBody
*/
type UpdateNetworkWirelessSsidVpnBody struct {

	// concentrator
	Concentrator *UpdateNetworkWirelessSsidVpnParamsBodyConcentrator `json:"concentrator,omitempty"`

	// failover
	Failover *UpdateNetworkWirelessSsidVpnParamsBodyFailover `json:"failover,omitempty"`

	// split tunnel
	SplitTunnel *UpdateNetworkWirelessSsidVpnParamsBodySplitTunnel `json:"splitTunnel,omitempty"`
}

// Validate validates this update network wireless ssid vpn body
func (o *UpdateNetworkWirelessSsidVpnBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateConcentrator(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateFailover(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateSplitTunnel(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *UpdateNetworkWirelessSsidVpnBody) validateConcentrator(formats strfmt.Registry) error {
	if swag.IsZero(o.Concentrator) { // not required
		return nil
	}

	if o.Concentrator != nil {
		if err := o.Concentrator.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("updateNetworkWirelessSsidVpn" + "." + "concentrator")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("updateNetworkWirelessSsidVpn" + "." + "concentrator")
			}
			return err
		}
	}

	return nil
}

func (o *UpdateNetworkWirelessSsidVpnBody) validateFailover(formats strfmt.Registry) error {
	if swag.IsZero(o.Failover) { // not required
		return nil
	}

	if o.Failover != nil {
		if err := o.Failover.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("updateNetworkWirelessSsidVpn" + "." + "failover")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("updateNetworkWirelessSsidVpn" + "." + "failover")
			}
			return err
		}
	}

	return nil
}

func (o *UpdateNetworkWirelessSsidVpnBody) validateSplitTunnel(formats strfmt.Registry) error {
	if swag.IsZero(o.SplitTunnel) { // not required
		return nil
	}

	if o.SplitTunnel != nil {
		if err := o.SplitTunnel.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("updateNetworkWirelessSsidVpn" + "." + "splitTunnel")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("updateNetworkWirelessSsidVpn" + "." + "splitTunnel")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this update network wireless ssid vpn body based on the context it is used
func (o *UpdateNetworkWirelessSsidVpnBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateConcentrator(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := o.contextValidateFailover(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := o.contextValidateSplitTunnel(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *UpdateNetworkWirelessSsidVpnBody) contextValidateConcentrator(ctx context.Context, formats strfmt.Registry) error {

	if o.Concentrator != nil {
		if err := o.Concentrator.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("updateNetworkWirelessSsidVpn" + "." + "concentrator")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("updateNetworkWirelessSsidVpn" + "." + "concentrator")
			}
			return err
		}
	}

	return nil
}

func (o *UpdateNetworkWirelessSsidVpnBody) contextValidateFailover(ctx context.Context, formats strfmt.Registry) error {

	if o.Failover != nil {
		if err := o.Failover.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("updateNetworkWirelessSsidVpn" + "." + "failover")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("updateNetworkWirelessSsidVpn" + "." + "failover")
			}
			return err
		}
	}

	return nil
}

func (o *UpdateNetworkWirelessSsidVpnBody) contextValidateSplitTunnel(ctx context.Context, formats strfmt.Registry) error {

	if o.SplitTunnel != nil {
		if err := o.SplitTunnel.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("updateNetworkWirelessSsidVpn" + "." + "splitTunnel")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("updateNetworkWirelessSsidVpn" + "." + "splitTunnel")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *UpdateNetworkWirelessSsidVpnBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *UpdateNetworkWirelessSsidVpnBody) UnmarshalBinary(b []byte) error {
	var res UpdateNetworkWirelessSsidVpnBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*UpdateNetworkWirelessSsidVpnParamsBodyConcentrator The VPN concentrator settings for this SSID.
swagger:model UpdateNetworkWirelessSsidVpnParamsBodyConcentrator
*/
type UpdateNetworkWirelessSsidVpnParamsBodyConcentrator struct {

	// The NAT ID of the concentrator that should be set.
	NetworkID string `json:"networkId,omitempty"`
}

// Validate validates this update network wireless ssid vpn params body concentrator
func (o *UpdateNetworkWirelessSsidVpnParamsBodyConcentrator) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this update network wireless ssid vpn params body concentrator based on context it is used
func (o *UpdateNetworkWirelessSsidVpnParamsBodyConcentrator) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *UpdateNetworkWirelessSsidVpnParamsBodyConcentrator) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *UpdateNetworkWirelessSsidVpnParamsBodyConcentrator) UnmarshalBinary(b []byte) error {
	var res UpdateNetworkWirelessSsidVpnParamsBodyConcentrator
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*UpdateNetworkWirelessSsidVpnParamsBodyFailover Secondary VPN concentrator settings. This is only used when two VPN concentrators are configured on the SSID.
swagger:model UpdateNetworkWirelessSsidVpnParamsBodyFailover
*/
type UpdateNetworkWirelessSsidVpnParamsBodyFailover struct {

	// Idle timer interval in seconds.
	HeartbeatInterval int64 `json:"heartbeatInterval,omitempty"`

	// Idle timer timeout in seconds.
	IdleTimeout int64 `json:"idleTimeout,omitempty"`

	// IP addressed reserved on DHCP server where SSID will terminate.
	RequestIP string `json:"requestIp,omitempty"`
}

// Validate validates this update network wireless ssid vpn params body failover
func (o *UpdateNetworkWirelessSsidVpnParamsBodyFailover) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this update network wireless ssid vpn params body failover based on context it is used
func (o *UpdateNetworkWirelessSsidVpnParamsBodyFailover) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *UpdateNetworkWirelessSsidVpnParamsBodyFailover) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *UpdateNetworkWirelessSsidVpnParamsBodyFailover) UnmarshalBinary(b []byte) error {
	var res UpdateNetworkWirelessSsidVpnParamsBodyFailover
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*UpdateNetworkWirelessSsidVpnParamsBodySplitTunnel The VPN split tunnel settings for this SSID.
swagger:model UpdateNetworkWirelessSsidVpnParamsBodySplitTunnel
*/
type UpdateNetworkWirelessSsidVpnParamsBodySplitTunnel struct {

	// If true, VPN split tunnel is enabled.
	Enabled bool `json:"enabled,omitempty"`

	// List of VPN split tunnel rules.
	Rules []*UpdateNetworkWirelessSsidVpnParamsBodySplitTunnelRulesItems0 `json:"rules"`
}

// Validate validates this update network wireless ssid vpn params body split tunnel
func (o *UpdateNetworkWirelessSsidVpnParamsBodySplitTunnel) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateRules(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *UpdateNetworkWirelessSsidVpnParamsBodySplitTunnel) validateRules(formats strfmt.Registry) error {
	if swag.IsZero(o.Rules) { // not required
		return nil
	}

	for i := 0; i < len(o.Rules); i++ {
		if swag.IsZero(o.Rules[i]) { // not required
			continue
		}

		if o.Rules[i] != nil {
			if err := o.Rules[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("updateNetworkWirelessSsidVpn" + "." + "splitTunnel" + "." + "rules" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("updateNetworkWirelessSsidVpn" + "." + "splitTunnel" + "." + "rules" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this update network wireless ssid vpn params body split tunnel based on the context it is used
func (o *UpdateNetworkWirelessSsidVpnParamsBodySplitTunnel) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateRules(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *UpdateNetworkWirelessSsidVpnParamsBodySplitTunnel) contextValidateRules(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(o.Rules); i++ {

		if o.Rules[i] != nil {
			if err := o.Rules[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("updateNetworkWirelessSsidVpn" + "." + "splitTunnel" + "." + "rules" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("updateNetworkWirelessSsidVpn" + "." + "splitTunnel" + "." + "rules" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (o *UpdateNetworkWirelessSsidVpnParamsBodySplitTunnel) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *UpdateNetworkWirelessSsidVpnParamsBodySplitTunnel) UnmarshalBinary(b []byte) error {
	var res UpdateNetworkWirelessSsidVpnParamsBodySplitTunnel
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*UpdateNetworkWirelessSsidVpnParamsBodySplitTunnelRulesItems0 update network wireless ssid vpn params body split tunnel rules items0
swagger:model UpdateNetworkWirelessSsidVpnParamsBodySplitTunnelRulesItems0
*/
type UpdateNetworkWirelessSsidVpnParamsBodySplitTunnelRulesItems0 struct {

	// Description for this split tunnel rule (optional).
	Comment string `json:"comment,omitempty"`

	// Destination for this split tunnel rule. IP address, fully-qualified domain names (FQDN) or 'any'.
	// Required: true
	DestCidr *string `json:"destCidr"`

	// Destination port for this split tunnel rule, (integer in the range 1-65535), or 'any'.
	DestPort string `json:"destPort,omitempty"`

	// Traffic policy specified for this split tunnel rule, 'allow' or 'deny'.
	// Required: true
	Policy *string `json:"policy"`

	// Protocol for this split tunnel rule.
	// Enum: [Any TCP UDP]
	Protocol string `json:"protocol,omitempty"`
}

// Validate validates this update network wireless ssid vpn params body split tunnel rules items0
func (o *UpdateNetworkWirelessSsidVpnParamsBodySplitTunnelRulesItems0) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateDestCidr(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validatePolicy(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateProtocol(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *UpdateNetworkWirelessSsidVpnParamsBodySplitTunnelRulesItems0) validateDestCidr(formats strfmt.Registry) error {

	if err := validate.Required("destCidr", "body", o.DestCidr); err != nil {
		return err
	}

	return nil
}

func (o *UpdateNetworkWirelessSsidVpnParamsBodySplitTunnelRulesItems0) validatePolicy(formats strfmt.Registry) error {

	if err := validate.Required("policy", "body", o.Policy); err != nil {
		return err
	}

	return nil
}

var updateNetworkWirelessSsidVpnParamsBodySplitTunnelRulesItems0TypeProtocolPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["Any","TCP","UDP"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		updateNetworkWirelessSsidVpnParamsBodySplitTunnelRulesItems0TypeProtocolPropEnum = append(updateNetworkWirelessSsidVpnParamsBodySplitTunnelRulesItems0TypeProtocolPropEnum, v)
	}
}

const (

	// UpdateNetworkWirelessSsidVpnParamsBodySplitTunnelRulesItems0ProtocolAny captures enum value "Any"
	UpdateNetworkWirelessSsidVpnParamsBodySplitTunnelRulesItems0ProtocolAny string = "Any"

	// UpdateNetworkWirelessSsidVpnParamsBodySplitTunnelRulesItems0ProtocolTCP captures enum value "TCP"
	UpdateNetworkWirelessSsidVpnParamsBodySplitTunnelRulesItems0ProtocolTCP string = "TCP"

	// UpdateNetworkWirelessSsidVpnParamsBodySplitTunnelRulesItems0ProtocolUDP captures enum value "UDP"
	UpdateNetworkWirelessSsidVpnParamsBodySplitTunnelRulesItems0ProtocolUDP string = "UDP"
)

// prop value enum
func (o *UpdateNetworkWirelessSsidVpnParamsBodySplitTunnelRulesItems0) validateProtocolEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, updateNetworkWirelessSsidVpnParamsBodySplitTunnelRulesItems0TypeProtocolPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (o *UpdateNetworkWirelessSsidVpnParamsBodySplitTunnelRulesItems0) validateProtocol(formats strfmt.Registry) error {
	if swag.IsZero(o.Protocol) { // not required
		return nil
	}

	// value enum
	if err := o.validateProtocolEnum("protocol", "body", o.Protocol); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this update network wireless ssid vpn params body split tunnel rules items0 based on context it is used
func (o *UpdateNetworkWirelessSsidVpnParamsBodySplitTunnelRulesItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *UpdateNetworkWirelessSsidVpnParamsBodySplitTunnelRulesItems0) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *UpdateNetworkWirelessSsidVpnParamsBodySplitTunnelRulesItems0) UnmarshalBinary(b []byte) error {
	var res UpdateNetworkWirelessSsidVpnParamsBodySplitTunnelRulesItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
