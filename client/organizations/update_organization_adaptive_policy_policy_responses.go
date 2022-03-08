// Code generated by go-swagger; DO NOT EDIT.

package organizations

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

// UpdateOrganizationAdaptivePolicyPolicyReader is a Reader for the UpdateOrganizationAdaptivePolicyPolicy structure.
type UpdateOrganizationAdaptivePolicyPolicyReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateOrganizationAdaptivePolicyPolicyReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUpdateOrganizationAdaptivePolicyPolicyOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewUpdateOrganizationAdaptivePolicyPolicyOK creates a UpdateOrganizationAdaptivePolicyPolicyOK with default headers values
func NewUpdateOrganizationAdaptivePolicyPolicyOK() *UpdateOrganizationAdaptivePolicyPolicyOK {
	return &UpdateOrganizationAdaptivePolicyPolicyOK{}
}

/* UpdateOrganizationAdaptivePolicyPolicyOK describes a response with status code 200, with default header values.

Successful operation
*/
type UpdateOrganizationAdaptivePolicyPolicyOK struct {
	Payload interface{}
}

func (o *UpdateOrganizationAdaptivePolicyPolicyOK) Error() string {
	return fmt.Sprintf("[PUT /organizations/{organizationId}/adaptivePolicy/policies/{adaptivePolicyId}][%d] updateOrganizationAdaptivePolicyPolicyOK  %+v", 200, o.Payload)
}
func (o *UpdateOrganizationAdaptivePolicyPolicyOK) GetPayload() interface{} {
	return o.Payload
}

func (o *UpdateOrganizationAdaptivePolicyPolicyOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*UpdateOrganizationAdaptivePolicyPolicyBody update organization adaptive policy policy body
// Example: {"acls":[{"id":"444","name":"Block web"}],"destinationGroup":{"id":"333","name":"IoT Servers","sgt":51},"lastEntryRule":"allow","sourceGroup":{"id":"222","name":"IoT Devices","sgt":50}}
swagger:model UpdateOrganizationAdaptivePolicyPolicyBody
*/
type UpdateOrganizationAdaptivePolicyPolicyBody struct {

	// An ordered array of adaptive policy ACLs (each requires one unique attribute) that apply to this policy
	//
	Acls []*UpdateOrganizationAdaptivePolicyPolicyParamsBodyAclsItems0 `json:"acls"`

	// destination group
	DestinationGroup *UpdateOrganizationAdaptivePolicyPolicyParamsBodyDestinationGroup `json:"destinationGroup,omitempty"`

	// The rule to apply if there is no matching ACL
	//
	// Enum: [default allow deny]
	LastEntryRule string `json:"lastEntryRule,omitempty"`

	// source group
	SourceGroup *UpdateOrganizationAdaptivePolicyPolicyParamsBodySourceGroup `json:"sourceGroup,omitempty"`
}

// Validate validates this update organization adaptive policy policy body
func (o *UpdateOrganizationAdaptivePolicyPolicyBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateAcls(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateDestinationGroup(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateLastEntryRule(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateSourceGroup(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *UpdateOrganizationAdaptivePolicyPolicyBody) validateAcls(formats strfmt.Registry) error {
	if swag.IsZero(o.Acls) { // not required
		return nil
	}

	for i := 0; i < len(o.Acls); i++ {
		if swag.IsZero(o.Acls[i]) { // not required
			continue
		}

		if o.Acls[i] != nil {
			if err := o.Acls[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("updateOrganizationAdaptivePolicyPolicy" + "." + "acls" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("updateOrganizationAdaptivePolicyPolicy" + "." + "acls" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (o *UpdateOrganizationAdaptivePolicyPolicyBody) validateDestinationGroup(formats strfmt.Registry) error {
	if swag.IsZero(o.DestinationGroup) { // not required
		return nil
	}

	if o.DestinationGroup != nil {
		if err := o.DestinationGroup.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("updateOrganizationAdaptivePolicyPolicy" + "." + "destinationGroup")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("updateOrganizationAdaptivePolicyPolicy" + "." + "destinationGroup")
			}
			return err
		}
	}

	return nil
}

var updateOrganizationAdaptivePolicyPolicyBodyTypeLastEntryRulePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["default","allow","deny"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		updateOrganizationAdaptivePolicyPolicyBodyTypeLastEntryRulePropEnum = append(updateOrganizationAdaptivePolicyPolicyBodyTypeLastEntryRulePropEnum, v)
	}
}

const (

	// UpdateOrganizationAdaptivePolicyPolicyBodyLastEntryRuleDefault captures enum value "default"
	UpdateOrganizationAdaptivePolicyPolicyBodyLastEntryRuleDefault string = "default"

	// UpdateOrganizationAdaptivePolicyPolicyBodyLastEntryRuleAllow captures enum value "allow"
	UpdateOrganizationAdaptivePolicyPolicyBodyLastEntryRuleAllow string = "allow"

	// UpdateOrganizationAdaptivePolicyPolicyBodyLastEntryRuleDeny captures enum value "deny"
	UpdateOrganizationAdaptivePolicyPolicyBodyLastEntryRuleDeny string = "deny"
)

// prop value enum
func (o *UpdateOrganizationAdaptivePolicyPolicyBody) validateLastEntryRuleEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, updateOrganizationAdaptivePolicyPolicyBodyTypeLastEntryRulePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (o *UpdateOrganizationAdaptivePolicyPolicyBody) validateLastEntryRule(formats strfmt.Registry) error {
	if swag.IsZero(o.LastEntryRule) { // not required
		return nil
	}

	// value enum
	if err := o.validateLastEntryRuleEnum("updateOrganizationAdaptivePolicyPolicy"+"."+"lastEntryRule", "body", o.LastEntryRule); err != nil {
		return err
	}

	return nil
}

func (o *UpdateOrganizationAdaptivePolicyPolicyBody) validateSourceGroup(formats strfmt.Registry) error {
	if swag.IsZero(o.SourceGroup) { // not required
		return nil
	}

	if o.SourceGroup != nil {
		if err := o.SourceGroup.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("updateOrganizationAdaptivePolicyPolicy" + "." + "sourceGroup")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("updateOrganizationAdaptivePolicyPolicy" + "." + "sourceGroup")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this update organization adaptive policy policy body based on the context it is used
func (o *UpdateOrganizationAdaptivePolicyPolicyBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateAcls(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := o.contextValidateDestinationGroup(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := o.contextValidateSourceGroup(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *UpdateOrganizationAdaptivePolicyPolicyBody) contextValidateAcls(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(o.Acls); i++ {

		if o.Acls[i] != nil {
			if err := o.Acls[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("updateOrganizationAdaptivePolicyPolicy" + "." + "acls" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("updateOrganizationAdaptivePolicyPolicy" + "." + "acls" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (o *UpdateOrganizationAdaptivePolicyPolicyBody) contextValidateDestinationGroup(ctx context.Context, formats strfmt.Registry) error {

	if o.DestinationGroup != nil {
		if err := o.DestinationGroup.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("updateOrganizationAdaptivePolicyPolicy" + "." + "destinationGroup")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("updateOrganizationAdaptivePolicyPolicy" + "." + "destinationGroup")
			}
			return err
		}
	}

	return nil
}

func (o *UpdateOrganizationAdaptivePolicyPolicyBody) contextValidateSourceGroup(ctx context.Context, formats strfmt.Registry) error {

	if o.SourceGroup != nil {
		if err := o.SourceGroup.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("updateOrganizationAdaptivePolicyPolicy" + "." + "sourceGroup")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("updateOrganizationAdaptivePolicyPolicy" + "." + "sourceGroup")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *UpdateOrganizationAdaptivePolicyPolicyBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *UpdateOrganizationAdaptivePolicyPolicyBody) UnmarshalBinary(b []byte) error {
	var res UpdateOrganizationAdaptivePolicyPolicyBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*UpdateOrganizationAdaptivePolicyPolicyParamsBodyAclsItems0 update organization adaptive policy policy params body acls items0
swagger:model UpdateOrganizationAdaptivePolicyPolicyParamsBodyAclsItems0
*/
type UpdateOrganizationAdaptivePolicyPolicyParamsBodyAclsItems0 struct {

	// The ID of the adaptive policy ACL
	ID string `json:"id,omitempty"`

	// The name of the adaptive policy ACL
	Name string `json:"name,omitempty"`
}

// Validate validates this update organization adaptive policy policy params body acls items0
func (o *UpdateOrganizationAdaptivePolicyPolicyParamsBodyAclsItems0) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this update organization adaptive policy policy params body acls items0 based on context it is used
func (o *UpdateOrganizationAdaptivePolicyPolicyParamsBodyAclsItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *UpdateOrganizationAdaptivePolicyPolicyParamsBodyAclsItems0) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *UpdateOrganizationAdaptivePolicyPolicyParamsBodyAclsItems0) UnmarshalBinary(b []byte) error {
	var res UpdateOrganizationAdaptivePolicyPolicyParamsBodyAclsItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*UpdateOrganizationAdaptivePolicyPolicyParamsBodyDestinationGroup The destination adaptive policy group (requires one unique attribute)
//
swagger:model UpdateOrganizationAdaptivePolicyPolicyParamsBodyDestinationGroup
*/
type UpdateOrganizationAdaptivePolicyPolicyParamsBodyDestinationGroup struct {

	// The ID of the destination adaptive policy group
	ID string `json:"id,omitempty"`

	// The name of the destination adaptive policy group
	Name string `json:"name,omitempty"`

	// The SGT of the destination adaptive policy group
	Sgt int64 `json:"sgt,omitempty"`
}

// Validate validates this update organization adaptive policy policy params body destination group
func (o *UpdateOrganizationAdaptivePolicyPolicyParamsBodyDestinationGroup) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this update organization adaptive policy policy params body destination group based on context it is used
func (o *UpdateOrganizationAdaptivePolicyPolicyParamsBodyDestinationGroup) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *UpdateOrganizationAdaptivePolicyPolicyParamsBodyDestinationGroup) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *UpdateOrganizationAdaptivePolicyPolicyParamsBodyDestinationGroup) UnmarshalBinary(b []byte) error {
	var res UpdateOrganizationAdaptivePolicyPolicyParamsBodyDestinationGroup
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*UpdateOrganizationAdaptivePolicyPolicyParamsBodySourceGroup The source adaptive policy group (requires one unique attribute)
//
swagger:model UpdateOrganizationAdaptivePolicyPolicyParamsBodySourceGroup
*/
type UpdateOrganizationAdaptivePolicyPolicyParamsBodySourceGroup struct {

	// The ID of the source adaptive policy group
	ID string `json:"id,omitempty"`

	// The name of the source adaptive policy group
	Name string `json:"name,omitempty"`

	// The SGT of the source adaptive policy group
	Sgt int64 `json:"sgt,omitempty"`
}

// Validate validates this update organization adaptive policy policy params body source group
func (o *UpdateOrganizationAdaptivePolicyPolicyParamsBodySourceGroup) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this update organization adaptive policy policy params body source group based on context it is used
func (o *UpdateOrganizationAdaptivePolicyPolicyParamsBodySourceGroup) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *UpdateOrganizationAdaptivePolicyPolicyParamsBodySourceGroup) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *UpdateOrganizationAdaptivePolicyPolicyParamsBodySourceGroup) UnmarshalBinary(b []byte) error {
	var res UpdateOrganizationAdaptivePolicyPolicyParamsBodySourceGroup
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
