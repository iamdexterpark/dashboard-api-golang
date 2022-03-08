// Code generated by go-swagger; DO NOT EDIT.

package appliance

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

// UpdateNetworkApplianceTrafficShapingRulesReader is a Reader for the UpdateNetworkApplianceTrafficShapingRules structure.
type UpdateNetworkApplianceTrafficShapingRulesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateNetworkApplianceTrafficShapingRulesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUpdateNetworkApplianceTrafficShapingRulesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewUpdateNetworkApplianceTrafficShapingRulesOK creates a UpdateNetworkApplianceTrafficShapingRulesOK with default headers values
func NewUpdateNetworkApplianceTrafficShapingRulesOK() *UpdateNetworkApplianceTrafficShapingRulesOK {
	return &UpdateNetworkApplianceTrafficShapingRulesOK{}
}

/* UpdateNetworkApplianceTrafficShapingRulesOK describes a response with status code 200, with default header values.

Successful operation
*/
type UpdateNetworkApplianceTrafficShapingRulesOK struct {
	Payload interface{}
}

func (o *UpdateNetworkApplianceTrafficShapingRulesOK) Error() string {
	return fmt.Sprintf("[PUT /networks/{networkId}/appliance/trafficShaping/rules][%d] updateNetworkApplianceTrafficShapingRulesOK  %+v", 200, o.Payload)
}
func (o *UpdateNetworkApplianceTrafficShapingRulesOK) GetPayload() interface{} {
	return o.Payload
}

func (o *UpdateNetworkApplianceTrafficShapingRulesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*UpdateNetworkApplianceTrafficShapingRulesBody update network appliance traffic shaping rules body
// Example: {"defaultRulesEnabled":true,"rules":[{"definitions":[{"type":"host","value":"google.com"},{"type":"port","value":"9090"},{"type":"ipRange","value":"192.1.0.0"},{"type":"ipRange","value":"192.1.0.0/16"},{"type":"ipRange","value":"10.1.0.0/16:80"},{"type":"localNet","value":"192.168.0.0/16"}],"dscpTagValue":0,"perClientBandwidthLimits":{"bandwidthLimits":{"limitDown":1000000,"limitUp":1000000},"settings":"custom"},"priority":"normal"}]}
swagger:model UpdateNetworkApplianceTrafficShapingRulesBody
*/
type UpdateNetworkApplianceTrafficShapingRulesBody struct {

	// Whether default traffic shaping rules are enabled (true) or disabled (false). There are 4 default rules, which can be seen on your network's traffic shaping page. Note that default rules count against the rule limit of 8.
	DefaultRulesEnabled bool `json:"defaultRulesEnabled,omitempty"`

	//     An array of traffic shaping rules. Rules are applied in the order that
	//     they are specified in. An empty list (or null) means no rules. Note that
	//     you are allowed a maximum of 8 rules.
	//
	Rules []*UpdateNetworkApplianceTrafficShapingRulesParamsBodyRulesItems0 `json:"rules"`
}

// Validate validates this update network appliance traffic shaping rules body
func (o *UpdateNetworkApplianceTrafficShapingRulesBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateRules(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *UpdateNetworkApplianceTrafficShapingRulesBody) validateRules(formats strfmt.Registry) error {
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
					return ve.ValidateName("updateNetworkApplianceTrafficShapingRules" + "." + "rules" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("updateNetworkApplianceTrafficShapingRules" + "." + "rules" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this update network appliance traffic shaping rules body based on the context it is used
func (o *UpdateNetworkApplianceTrafficShapingRulesBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateRules(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *UpdateNetworkApplianceTrafficShapingRulesBody) contextValidateRules(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(o.Rules); i++ {

		if o.Rules[i] != nil {
			if err := o.Rules[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("updateNetworkApplianceTrafficShapingRules" + "." + "rules" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("updateNetworkApplianceTrafficShapingRules" + "." + "rules" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (o *UpdateNetworkApplianceTrafficShapingRulesBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *UpdateNetworkApplianceTrafficShapingRulesBody) UnmarshalBinary(b []byte) error {
	var res UpdateNetworkApplianceTrafficShapingRulesBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*UpdateNetworkApplianceTrafficShapingRulesParamsBodyRulesItems0 update network appliance traffic shaping rules params body rules items0
swagger:model UpdateNetworkApplianceTrafficShapingRulesParamsBodyRulesItems0
*/
type UpdateNetworkApplianceTrafficShapingRulesParamsBodyRulesItems0 struct {

	//     A list of objects describing the definitions of your traffic shaping rule. At least one definition is required.
	//
	// Required: true
	Definitions []*UpdateNetworkApplianceTrafficShapingRulesParamsBodyRulesItems0DefinitionsItems0 `json:"definitions"`

	//     The DSCP tag applied by your rule. null means 'Do not change DSCP tag'.
	//     For a list of possible tag values, use the trafficShaping/dscpTaggingOptions endpoint.
	//
	DscpTagValue int64 `json:"dscpTagValue,omitempty"`

	// per client bandwidth limits
	PerClientBandwidthLimits *UpdateNetworkApplianceTrafficShapingRulesParamsBodyRulesItems0PerClientBandwidthLimits `json:"perClientBandwidthLimits,omitempty"`

	//     A string, indicating the priority level for packets bound to your rule.
	//     Can be 'low', 'normal' or 'high'.
	//
	Priority string `json:"priority,omitempty"`
}

// Validate validates this update network appliance traffic shaping rules params body rules items0
func (o *UpdateNetworkApplianceTrafficShapingRulesParamsBodyRulesItems0) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateDefinitions(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validatePerClientBandwidthLimits(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *UpdateNetworkApplianceTrafficShapingRulesParamsBodyRulesItems0) validateDefinitions(formats strfmt.Registry) error {

	if err := validate.Required("definitions", "body", o.Definitions); err != nil {
		return err
	}

	for i := 0; i < len(o.Definitions); i++ {
		if swag.IsZero(o.Definitions[i]) { // not required
			continue
		}

		if o.Definitions[i] != nil {
			if err := o.Definitions[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("definitions" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("definitions" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (o *UpdateNetworkApplianceTrafficShapingRulesParamsBodyRulesItems0) validatePerClientBandwidthLimits(formats strfmt.Registry) error {
	if swag.IsZero(o.PerClientBandwidthLimits) { // not required
		return nil
	}

	if o.PerClientBandwidthLimits != nil {
		if err := o.PerClientBandwidthLimits.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("perClientBandwidthLimits")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("perClientBandwidthLimits")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this update network appliance traffic shaping rules params body rules items0 based on the context it is used
func (o *UpdateNetworkApplianceTrafficShapingRulesParamsBodyRulesItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateDefinitions(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := o.contextValidatePerClientBandwidthLimits(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *UpdateNetworkApplianceTrafficShapingRulesParamsBodyRulesItems0) contextValidateDefinitions(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(o.Definitions); i++ {

		if o.Definitions[i] != nil {
			if err := o.Definitions[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("definitions" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("definitions" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (o *UpdateNetworkApplianceTrafficShapingRulesParamsBodyRulesItems0) contextValidatePerClientBandwidthLimits(ctx context.Context, formats strfmt.Registry) error {

	if o.PerClientBandwidthLimits != nil {
		if err := o.PerClientBandwidthLimits.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("perClientBandwidthLimits")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("perClientBandwidthLimits")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *UpdateNetworkApplianceTrafficShapingRulesParamsBodyRulesItems0) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *UpdateNetworkApplianceTrafficShapingRulesParamsBodyRulesItems0) UnmarshalBinary(b []byte) error {
	var res UpdateNetworkApplianceTrafficShapingRulesParamsBodyRulesItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*UpdateNetworkApplianceTrafficShapingRulesParamsBodyRulesItems0DefinitionsItems0 update network appliance traffic shaping rules params body rules items0 definitions items0
swagger:model UpdateNetworkApplianceTrafficShapingRulesParamsBodyRulesItems0DefinitionsItems0
*/
type UpdateNetworkApplianceTrafficShapingRulesParamsBodyRulesItems0DefinitionsItems0 struct {

	// The type of definition. Can be one of 'application', 'applicationCategory', 'host', 'port', 'ipRange' or 'localNet'.
	// Required: true
	// Enum: [application applicationCategory host port ipRange localNet]
	Type *string `json:"type"`

	//     If "type" is 'host', 'port', 'ipRange' or 'localNet', then "value" must be a string, matching either
	//     a hostname (e.g. "somesite.com"), a port (e.g. 8080), or an IP range ("192.1.0.0",
	//     "192.1.0.0/16", or "10.1.0.0/16:80"). 'localNet' also supports CIDR notation, excluding
	//     custom ports.
	//      If "type" is 'application' or 'applicationCategory', then "value" must be an object
	//     with the structure { "id": "meraki:layer7/..." }, where "id" is the application category or
	//     application ID (for a list of IDs for your network, use the trafficShaping/applicationCategories
	//     endpoint).
	//
	// Required: true
	Value *string `json:"value"`
}

// Validate validates this update network appliance traffic shaping rules params body rules items0 definitions items0
func (o *UpdateNetworkApplianceTrafficShapingRulesParamsBodyRulesItems0DefinitionsItems0) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateType(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateValue(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var updateNetworkApplianceTrafficShapingRulesParamsBodyRulesItems0DefinitionsItems0TypeTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["application","applicationCategory","host","port","ipRange","localNet"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		updateNetworkApplianceTrafficShapingRulesParamsBodyRulesItems0DefinitionsItems0TypeTypePropEnum = append(updateNetworkApplianceTrafficShapingRulesParamsBodyRulesItems0DefinitionsItems0TypeTypePropEnum, v)
	}
}

const (

	// UpdateNetworkApplianceTrafficShapingRulesParamsBodyRulesItems0DefinitionsItems0TypeApplication captures enum value "application"
	UpdateNetworkApplianceTrafficShapingRulesParamsBodyRulesItems0DefinitionsItems0TypeApplication string = "application"

	// UpdateNetworkApplianceTrafficShapingRulesParamsBodyRulesItems0DefinitionsItems0TypeApplicationCategory captures enum value "applicationCategory"
	UpdateNetworkApplianceTrafficShapingRulesParamsBodyRulesItems0DefinitionsItems0TypeApplicationCategory string = "applicationCategory"

	// UpdateNetworkApplianceTrafficShapingRulesParamsBodyRulesItems0DefinitionsItems0TypeHost captures enum value "host"
	UpdateNetworkApplianceTrafficShapingRulesParamsBodyRulesItems0DefinitionsItems0TypeHost string = "host"

	// UpdateNetworkApplianceTrafficShapingRulesParamsBodyRulesItems0DefinitionsItems0TypePort captures enum value "port"
	UpdateNetworkApplianceTrafficShapingRulesParamsBodyRulesItems0DefinitionsItems0TypePort string = "port"

	// UpdateNetworkApplianceTrafficShapingRulesParamsBodyRulesItems0DefinitionsItems0TypeIPRange captures enum value "ipRange"
	UpdateNetworkApplianceTrafficShapingRulesParamsBodyRulesItems0DefinitionsItems0TypeIPRange string = "ipRange"

	// UpdateNetworkApplianceTrafficShapingRulesParamsBodyRulesItems0DefinitionsItems0TypeLocalNet captures enum value "localNet"
	UpdateNetworkApplianceTrafficShapingRulesParamsBodyRulesItems0DefinitionsItems0TypeLocalNet string = "localNet"
)

// prop value enum
func (o *UpdateNetworkApplianceTrafficShapingRulesParamsBodyRulesItems0DefinitionsItems0) validateTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, updateNetworkApplianceTrafficShapingRulesParamsBodyRulesItems0DefinitionsItems0TypeTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (o *UpdateNetworkApplianceTrafficShapingRulesParamsBodyRulesItems0DefinitionsItems0) validateType(formats strfmt.Registry) error {

	if err := validate.Required("type", "body", o.Type); err != nil {
		return err
	}

	// value enum
	if err := o.validateTypeEnum("type", "body", *o.Type); err != nil {
		return err
	}

	return nil
}

func (o *UpdateNetworkApplianceTrafficShapingRulesParamsBodyRulesItems0DefinitionsItems0) validateValue(formats strfmt.Registry) error {

	if err := validate.Required("value", "body", o.Value); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this update network appliance traffic shaping rules params body rules items0 definitions items0 based on context it is used
func (o *UpdateNetworkApplianceTrafficShapingRulesParamsBodyRulesItems0DefinitionsItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *UpdateNetworkApplianceTrafficShapingRulesParamsBodyRulesItems0DefinitionsItems0) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *UpdateNetworkApplianceTrafficShapingRulesParamsBodyRulesItems0DefinitionsItems0) UnmarshalBinary(b []byte) error {
	var res UpdateNetworkApplianceTrafficShapingRulesParamsBodyRulesItems0DefinitionsItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*UpdateNetworkApplianceTrafficShapingRulesParamsBodyRulesItems0PerClientBandwidthLimits     An object describing the bandwidth settings for your rule.
//
swagger:model UpdateNetworkApplianceTrafficShapingRulesParamsBodyRulesItems0PerClientBandwidthLimits
*/
type UpdateNetworkApplianceTrafficShapingRulesParamsBodyRulesItems0PerClientBandwidthLimits struct {

	// bandwidth limits
	BandwidthLimits *UpdateNetworkApplianceTrafficShapingRulesParamsBodyRulesItems0PerClientBandwidthLimitsBandwidthLimits `json:"bandwidthLimits,omitempty"`

	// How bandwidth limits are applied by your rule. Can be one of 'network default', 'ignore' or 'custom'.
	Settings string `json:"settings,omitempty"`
}

// Validate validates this update network appliance traffic shaping rules params body rules items0 per client bandwidth limits
func (o *UpdateNetworkApplianceTrafficShapingRulesParamsBodyRulesItems0PerClientBandwidthLimits) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateBandwidthLimits(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *UpdateNetworkApplianceTrafficShapingRulesParamsBodyRulesItems0PerClientBandwidthLimits) validateBandwidthLimits(formats strfmt.Registry) error {
	if swag.IsZero(o.BandwidthLimits) { // not required
		return nil
	}

	if o.BandwidthLimits != nil {
		if err := o.BandwidthLimits.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("perClientBandwidthLimits" + "." + "bandwidthLimits")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("perClientBandwidthLimits" + "." + "bandwidthLimits")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this update network appliance traffic shaping rules params body rules items0 per client bandwidth limits based on the context it is used
func (o *UpdateNetworkApplianceTrafficShapingRulesParamsBodyRulesItems0PerClientBandwidthLimits) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateBandwidthLimits(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *UpdateNetworkApplianceTrafficShapingRulesParamsBodyRulesItems0PerClientBandwidthLimits) contextValidateBandwidthLimits(ctx context.Context, formats strfmt.Registry) error {

	if o.BandwidthLimits != nil {
		if err := o.BandwidthLimits.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("perClientBandwidthLimits" + "." + "bandwidthLimits")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("perClientBandwidthLimits" + "." + "bandwidthLimits")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *UpdateNetworkApplianceTrafficShapingRulesParamsBodyRulesItems0PerClientBandwidthLimits) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *UpdateNetworkApplianceTrafficShapingRulesParamsBodyRulesItems0PerClientBandwidthLimits) UnmarshalBinary(b []byte) error {
	var res UpdateNetworkApplianceTrafficShapingRulesParamsBodyRulesItems0PerClientBandwidthLimits
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*UpdateNetworkApplianceTrafficShapingRulesParamsBodyRulesItems0PerClientBandwidthLimitsBandwidthLimits The bandwidth limits object, specifying the upload ('limitUp') and download ('limitDown') speed in Kbps. These are only enforced if 'settings' is set to 'custom'.
swagger:model UpdateNetworkApplianceTrafficShapingRulesParamsBodyRulesItems0PerClientBandwidthLimitsBandwidthLimits
*/
type UpdateNetworkApplianceTrafficShapingRulesParamsBodyRulesItems0PerClientBandwidthLimitsBandwidthLimits struct {

	// The maximum download limit (integer, in Kbps).
	LimitDown int64 `json:"limitDown,omitempty"`

	// The maximum upload limit (integer, in Kbps).
	LimitUp int64 `json:"limitUp,omitempty"`
}

// Validate validates this update network appliance traffic shaping rules params body rules items0 per client bandwidth limits bandwidth limits
func (o *UpdateNetworkApplianceTrafficShapingRulesParamsBodyRulesItems0PerClientBandwidthLimitsBandwidthLimits) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this update network appliance traffic shaping rules params body rules items0 per client bandwidth limits bandwidth limits based on context it is used
func (o *UpdateNetworkApplianceTrafficShapingRulesParamsBodyRulesItems0PerClientBandwidthLimitsBandwidthLimits) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *UpdateNetworkApplianceTrafficShapingRulesParamsBodyRulesItems0PerClientBandwidthLimitsBandwidthLimits) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *UpdateNetworkApplianceTrafficShapingRulesParamsBodyRulesItems0PerClientBandwidthLimitsBandwidthLimits) UnmarshalBinary(b []byte) error {
	var res UpdateNetworkApplianceTrafficShapingRulesParamsBodyRulesItems0PerClientBandwidthLimitsBandwidthLimits
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
