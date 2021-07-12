// Code generated by go-swagger; DO NOT EDIT.

package module

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

// PostOrganizationOrganizationIDRelationshipsModuleReader is a Reader for the PostOrganizationOrganizationIDRelationshipsModule structure.
type PostOrganizationOrganizationIDRelationshipsModuleReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostOrganizationOrganizationIDRelationshipsModuleReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewPostOrganizationOrganizationIDRelationshipsModuleCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPostOrganizationOrganizationIDRelationshipsModuleCreated creates a PostOrganizationOrganizationIDRelationshipsModuleCreated with default headers values
func NewPostOrganizationOrganizationIDRelationshipsModuleCreated() *PostOrganizationOrganizationIDRelationshipsModuleCreated {
	return &PostOrganizationOrganizationIDRelationshipsModuleCreated{}
}

/* PostOrganizationOrganizationIDRelationshipsModuleCreated describes a response with status code 201, with default header values.

Successful response
*/
type PostOrganizationOrganizationIDRelationshipsModuleCreated struct {
	Payload *PostOrganizationOrganizationIDRelationshipsModuleCreatedBody
}

func (o *PostOrganizationOrganizationIDRelationshipsModuleCreated) Error() string {
	return fmt.Sprintf("[POST /organization/{organizationId}/relationships/module][%d] postOrganizationOrganizationIdRelationshipsModuleCreated  %+v", 201, o.Payload)
}
func (o *PostOrganizationOrganizationIDRelationshipsModuleCreated) GetPayload() *PostOrganizationOrganizationIDRelationshipsModuleCreatedBody {
	return o.Payload
}

func (o *PostOrganizationOrganizationIDRelationshipsModuleCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(PostOrganizationOrganizationIDRelationshipsModuleCreatedBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*PostOrganizationOrganizationIDRelationshipsModuleBody post organization organization ID relationships module body
swagger:model PostOrganizationOrganizationIDRelationshipsModuleBody
*/
type PostOrganizationOrganizationIDRelationshipsModuleBody struct {

	// data
	Data []*PostOrganizationOrganizationIDRelationshipsModuleParamsBodyDataItems0 `json:"data"`
}

// Validate validates this post organization organization ID relationships module body
func (o *PostOrganizationOrganizationIDRelationshipsModuleBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateData(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *PostOrganizationOrganizationIDRelationshipsModuleBody) validateData(formats strfmt.Registry) error {
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

// ContextValidate validate this post organization organization ID relationships module body based on the context it is used
func (o *PostOrganizationOrganizationIDRelationshipsModuleBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateData(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *PostOrganizationOrganizationIDRelationshipsModuleBody) contextValidateData(ctx context.Context, formats strfmt.Registry) error {

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
func (o *PostOrganizationOrganizationIDRelationshipsModuleBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *PostOrganizationOrganizationIDRelationshipsModuleBody) UnmarshalBinary(b []byte) error {
	var res PostOrganizationOrganizationIDRelationshipsModuleBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*PostOrganizationOrganizationIDRelationshipsModuleCreatedBody post organization organization ID relationships module created body
swagger:model PostOrganizationOrganizationIDRelationshipsModuleCreatedBody
*/
type PostOrganizationOrganizationIDRelationshipsModuleCreatedBody struct {

	// data
	Data []*PostOrganizationOrganizationIDRelationshipsModuleCreatedBodyDataItems0 `json:"data"`
}

// Validate validates this post organization organization ID relationships module created body
func (o *PostOrganizationOrganizationIDRelationshipsModuleCreatedBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateData(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *PostOrganizationOrganizationIDRelationshipsModuleCreatedBody) validateData(formats strfmt.Registry) error {
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
					return ve.ValidateName("postOrganizationOrganizationIdRelationshipsModuleCreated" + "." + "data" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this post organization organization ID relationships module created body based on the context it is used
func (o *PostOrganizationOrganizationIDRelationshipsModuleCreatedBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateData(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *PostOrganizationOrganizationIDRelationshipsModuleCreatedBody) contextValidateData(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(o.Data); i++ {

		if o.Data[i] != nil {
			if err := o.Data[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("postOrganizationOrganizationIdRelationshipsModuleCreated" + "." + "data" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (o *PostOrganizationOrganizationIDRelationshipsModuleCreatedBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *PostOrganizationOrganizationIDRelationshipsModuleCreatedBody) UnmarshalBinary(b []byte) error {
	var res PostOrganizationOrganizationIDRelationshipsModuleCreatedBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*PostOrganizationOrganizationIDRelationshipsModuleCreatedBodyDataItems0 post organization organization ID relationships module created body data items0
swagger:model PostOrganizationOrganizationIDRelationshipsModuleCreatedBodyDataItems0
*/
type PostOrganizationOrganizationIDRelationshipsModuleCreatedBodyDataItems0 struct {

	// id
	ID string `json:"id,omitempty"`

	// type
	// Enum: [module]
	Type string `json:"type,omitempty"`
}

// Validate validates this post organization organization ID relationships module created body data items0
func (o *PostOrganizationOrganizationIDRelationshipsModuleCreatedBodyDataItems0) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var postOrganizationOrganizationIdRelationshipsModuleCreatedBodyDataItems0TypeTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["module"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		postOrganizationOrganizationIdRelationshipsModuleCreatedBodyDataItems0TypeTypePropEnum = append(postOrganizationOrganizationIdRelationshipsModuleCreatedBodyDataItems0TypeTypePropEnum, v)
	}
}

const (

	// PostOrganizationOrganizationIDRelationshipsModuleCreatedBodyDataItems0TypeModule captures enum value "module"
	PostOrganizationOrganizationIDRelationshipsModuleCreatedBodyDataItems0TypeModule string = "module"
)

// prop value enum
func (o *PostOrganizationOrganizationIDRelationshipsModuleCreatedBodyDataItems0) validateTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, postOrganizationOrganizationIdRelationshipsModuleCreatedBodyDataItems0TypeTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (o *PostOrganizationOrganizationIDRelationshipsModuleCreatedBodyDataItems0) validateType(formats strfmt.Registry) error {
	if swag.IsZero(o.Type) { // not required
		return nil
	}

	// value enum
	if err := o.validateTypeEnum("type", "body", o.Type); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this post organization organization ID relationships module created body data items0 based on context it is used
func (o *PostOrganizationOrganizationIDRelationshipsModuleCreatedBodyDataItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *PostOrganizationOrganizationIDRelationshipsModuleCreatedBodyDataItems0) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *PostOrganizationOrganizationIDRelationshipsModuleCreatedBodyDataItems0) UnmarshalBinary(b []byte) error {
	var res PostOrganizationOrganizationIDRelationshipsModuleCreatedBodyDataItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*PostOrganizationOrganizationIDRelationshipsModuleParamsBodyDataItems0 post organization organization ID relationships module params body data items0
swagger:model PostOrganizationOrganizationIDRelationshipsModuleParamsBodyDataItems0
*/
type PostOrganizationOrganizationIDRelationshipsModuleParamsBodyDataItems0 struct {

	// id
	ID string `json:"id,omitempty"`

	// type
	// Enum: [module]
	Type string `json:"type,omitempty"`
}

// Validate validates this post organization organization ID relationships module params body data items0
func (o *PostOrganizationOrganizationIDRelationshipsModuleParamsBodyDataItems0) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var postOrganizationOrganizationIdRelationshipsModuleParamsBodyDataItems0TypeTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["module"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		postOrganizationOrganizationIdRelationshipsModuleParamsBodyDataItems0TypeTypePropEnum = append(postOrganizationOrganizationIdRelationshipsModuleParamsBodyDataItems0TypeTypePropEnum, v)
	}
}

const (

	// PostOrganizationOrganizationIDRelationshipsModuleParamsBodyDataItems0TypeModule captures enum value "module"
	PostOrganizationOrganizationIDRelationshipsModuleParamsBodyDataItems0TypeModule string = "module"
)

// prop value enum
func (o *PostOrganizationOrganizationIDRelationshipsModuleParamsBodyDataItems0) validateTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, postOrganizationOrganizationIdRelationshipsModuleParamsBodyDataItems0TypeTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (o *PostOrganizationOrganizationIDRelationshipsModuleParamsBodyDataItems0) validateType(formats strfmt.Registry) error {
	if swag.IsZero(o.Type) { // not required
		return nil
	}

	// value enum
	if err := o.validateTypeEnum("type", "body", o.Type); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this post organization organization ID relationships module params body data items0 based on context it is used
func (o *PostOrganizationOrganizationIDRelationshipsModuleParamsBodyDataItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *PostOrganizationOrganizationIDRelationshipsModuleParamsBodyDataItems0) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *PostOrganizationOrganizationIDRelationshipsModuleParamsBodyDataItems0) UnmarshalBinary(b []byte) error {
	var res PostOrganizationOrganizationIDRelationshipsModuleParamsBodyDataItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}