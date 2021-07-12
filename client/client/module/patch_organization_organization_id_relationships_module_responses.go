// Code generated by go-swagger; DO NOT EDIT.

package module

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// PatchOrganizationOrganizationIDRelationshipsModuleReader is a Reader for the PatchOrganizationOrganizationIDRelationshipsModule structure.
type PatchOrganizationOrganizationIDRelationshipsModuleReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PatchOrganizationOrganizationIDRelationshipsModuleReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewPatchOrganizationOrganizationIDRelationshipsModuleNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPatchOrganizationOrganizationIDRelationshipsModuleNoContent creates a PatchOrganizationOrganizationIDRelationshipsModuleNoContent with default headers values
func NewPatchOrganizationOrganizationIDRelationshipsModuleNoContent() *PatchOrganizationOrganizationIDRelationshipsModuleNoContent {
	return &PatchOrganizationOrganizationIDRelationshipsModuleNoContent{}
}

/* PatchOrganizationOrganizationIDRelationshipsModuleNoContent describes a response with status code 204, with default header values.

Successful response
*/
type PatchOrganizationOrganizationIDRelationshipsModuleNoContent struct {
}

func (o *PatchOrganizationOrganizationIDRelationshipsModuleNoContent) Error() string {
	return fmt.Sprintf("[PATCH /organization/{organizationId}/relationships/module][%d] patchOrganizationOrganizationIdRelationshipsModuleNoContent ", 204)
}

func (o *PatchOrganizationOrganizationIDRelationshipsModuleNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

/*PatchOrganizationOrganizationIDRelationshipsModuleBody patch organization organization ID relationships module body
swagger:model PatchOrganizationOrganizationIDRelationshipsModuleBody
*/
type PatchOrganizationOrganizationIDRelationshipsModuleBody struct {

	// data
	Data []*PatchOrganizationOrganizationIDRelationshipsModuleParamsBodyDataItems0 `json:"data"`
}

// Validate validates this patch organization organization ID relationships module body
func (o *PatchOrganizationOrganizationIDRelationshipsModuleBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateData(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *PatchOrganizationOrganizationIDRelationshipsModuleBody) validateData(formats strfmt.Registry) error {
	if swag.IsZero(o.Data) { // not required
		return nil
	}

	for i := 0; i < len(o.Data); i++ {
		if swag.IsZero(o.Data[i]) { // not required
			continue
		}

		if o.Data[i] != nil {
			if err := o.Data[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("relationship" + "." + "data" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this patch organization organization ID relationships module body based on the context it is used
func (o *PatchOrganizationOrganizationIDRelationshipsModuleBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateData(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *PatchOrganizationOrganizationIDRelationshipsModuleBody) contextValidateData(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(o.Data); i++ {

		if o.Data[i] != nil {
			if err := o.Data[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("relationship" + "." + "data" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (o *PatchOrganizationOrganizationIDRelationshipsModuleBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *PatchOrganizationOrganizationIDRelationshipsModuleBody) UnmarshalBinary(b []byte) error {
	var res PatchOrganizationOrganizationIDRelationshipsModuleBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*PatchOrganizationOrganizationIDRelationshipsModuleParamsBodyDataItems0 patch organization organization ID relationships module params body data items0
swagger:model PatchOrganizationOrganizationIDRelationshipsModuleParamsBodyDataItems0
*/
type PatchOrganizationOrganizationIDRelationshipsModuleParamsBodyDataItems0 struct {

	// id
	ID string `json:"id,omitempty"`

	// type
	// Enum: [module]
	Type string `json:"type,omitempty"`
}

// Validate validates this patch organization organization ID relationships module params body data items0
func (o *PatchOrganizationOrganizationIDRelationshipsModuleParamsBodyDataItems0) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var patchOrganizationOrganizationIdRelationshipsModuleParamsBodyDataItems0TypeTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["module"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		patchOrganizationOrganizationIdRelationshipsModuleParamsBodyDataItems0TypeTypePropEnum = append(patchOrganizationOrganizationIdRelationshipsModuleParamsBodyDataItems0TypeTypePropEnum, v)
	}
}

const (

	// PatchOrganizationOrganizationIDRelationshipsModuleParamsBodyDataItems0TypeModule captures enum value "module"
	PatchOrganizationOrganizationIDRelationshipsModuleParamsBodyDataItems0TypeModule string = "module"
)

// prop value enum
func (o *PatchOrganizationOrganizationIDRelationshipsModuleParamsBodyDataItems0) validateTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, patchOrganizationOrganizationIdRelationshipsModuleParamsBodyDataItems0TypeTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (o *PatchOrganizationOrganizationIDRelationshipsModuleParamsBodyDataItems0) validateType(formats strfmt.Registry) error {
	if swag.IsZero(o.Type) { // not required
		return nil
	}

	// value enum
	if err := o.validateTypeEnum("type", "body", o.Type); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this patch organization organization ID relationships module params body data items0 based on context it is used
func (o *PatchOrganizationOrganizationIDRelationshipsModuleParamsBodyDataItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *PatchOrganizationOrganizationIDRelationshipsModuleParamsBodyDataItems0) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *PatchOrganizationOrganizationIDRelationshipsModuleParamsBodyDataItems0) UnmarshalBinary(b []byte) error {
	var res PatchOrganizationOrganizationIDRelationshipsModuleParamsBodyDataItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
