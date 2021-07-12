// Code generated by go-swagger; DO NOT EDIT.

package workspace

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

// DeleteOrganizationOrganizationIDRelationshipsWorkspaceReader is a Reader for the DeleteOrganizationOrganizationIDRelationshipsWorkspace structure.
type DeleteOrganizationOrganizationIDRelationshipsWorkspaceReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteOrganizationOrganizationIDRelationshipsWorkspaceReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewDeleteOrganizationOrganizationIDRelationshipsWorkspaceNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDeleteOrganizationOrganizationIDRelationshipsWorkspaceNoContent creates a DeleteOrganizationOrganizationIDRelationshipsWorkspaceNoContent with default headers values
func NewDeleteOrganizationOrganizationIDRelationshipsWorkspaceNoContent() *DeleteOrganizationOrganizationIDRelationshipsWorkspaceNoContent {
	return &DeleteOrganizationOrganizationIDRelationshipsWorkspaceNoContent{}
}

/* DeleteOrganizationOrganizationIDRelationshipsWorkspaceNoContent describes a response with status code 204, with default header values.

Successful response
*/
type DeleteOrganizationOrganizationIDRelationshipsWorkspaceNoContent struct {
}

func (o *DeleteOrganizationOrganizationIDRelationshipsWorkspaceNoContent) Error() string {
	return fmt.Sprintf("[DELETE /organization/{organizationId}/relationships/workspace][%d] deleteOrganizationOrganizationIdRelationshipsWorkspaceNoContent ", 204)
}

func (o *DeleteOrganizationOrganizationIDRelationshipsWorkspaceNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

/*DeleteOrganizationOrganizationIDRelationshipsWorkspaceBody delete organization organization ID relationships workspace body
swagger:model DeleteOrganizationOrganizationIDRelationshipsWorkspaceBody
*/
type DeleteOrganizationOrganizationIDRelationshipsWorkspaceBody struct {

	// data
	Data []*DeleteOrganizationOrganizationIDRelationshipsWorkspaceParamsBodyDataItems0 `json:"data"`
}

// Validate validates this delete organization organization ID relationships workspace body
func (o *DeleteOrganizationOrganizationIDRelationshipsWorkspaceBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateData(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *DeleteOrganizationOrganizationIDRelationshipsWorkspaceBody) validateData(formats strfmt.Registry) error {
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

// ContextValidate validate this delete organization organization ID relationships workspace body based on the context it is used
func (o *DeleteOrganizationOrganizationIDRelationshipsWorkspaceBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateData(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *DeleteOrganizationOrganizationIDRelationshipsWorkspaceBody) contextValidateData(ctx context.Context, formats strfmt.Registry) error {

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
func (o *DeleteOrganizationOrganizationIDRelationshipsWorkspaceBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *DeleteOrganizationOrganizationIDRelationshipsWorkspaceBody) UnmarshalBinary(b []byte) error {
	var res DeleteOrganizationOrganizationIDRelationshipsWorkspaceBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*DeleteOrganizationOrganizationIDRelationshipsWorkspaceParamsBodyDataItems0 delete organization organization ID relationships workspace params body data items0
swagger:model DeleteOrganizationOrganizationIDRelationshipsWorkspaceParamsBodyDataItems0
*/
type DeleteOrganizationOrganizationIDRelationshipsWorkspaceParamsBodyDataItems0 struct {

	// id
	ID string `json:"id,omitempty"`

	// type
	// Enum: [workspace]
	Type string `json:"type,omitempty"`
}

// Validate validates this delete organization organization ID relationships workspace params body data items0
func (o *DeleteOrganizationOrganizationIDRelationshipsWorkspaceParamsBodyDataItems0) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var deleteOrganizationOrganizationIdRelationshipsWorkspaceParamsBodyDataItems0TypeTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["workspace"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		deleteOrganizationOrganizationIdRelationshipsWorkspaceParamsBodyDataItems0TypeTypePropEnum = append(deleteOrganizationOrganizationIdRelationshipsWorkspaceParamsBodyDataItems0TypeTypePropEnum, v)
	}
}

const (

	// DeleteOrganizationOrganizationIDRelationshipsWorkspaceParamsBodyDataItems0TypeWorkspace captures enum value "workspace"
	DeleteOrganizationOrganizationIDRelationshipsWorkspaceParamsBodyDataItems0TypeWorkspace string = "workspace"
)

// prop value enum
func (o *DeleteOrganizationOrganizationIDRelationshipsWorkspaceParamsBodyDataItems0) validateTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, deleteOrganizationOrganizationIdRelationshipsWorkspaceParamsBodyDataItems0TypeTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (o *DeleteOrganizationOrganizationIDRelationshipsWorkspaceParamsBodyDataItems0) validateType(formats strfmt.Registry) error {
	if swag.IsZero(o.Type) { // not required
		return nil
	}

	// value enum
	if err := o.validateTypeEnum("type", "body", o.Type); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this delete organization organization ID relationships workspace params body data items0 based on context it is used
func (o *DeleteOrganizationOrganizationIDRelationshipsWorkspaceParamsBodyDataItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *DeleteOrganizationOrganizationIDRelationshipsWorkspaceParamsBodyDataItems0) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *DeleteOrganizationOrganizationIDRelationshipsWorkspaceParamsBodyDataItems0) UnmarshalBinary(b []byte) error {
	var res DeleteOrganizationOrganizationIDRelationshipsWorkspaceParamsBodyDataItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}