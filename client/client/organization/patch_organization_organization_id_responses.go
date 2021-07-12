// Code generated by go-swagger; DO NOT EDIT.

package organization

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"fmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"

	"azb/client/models"
)

// PatchOrganizationOrganizationIDReader is a Reader for the PatchOrganizationOrganizationID structure.
type PatchOrganizationOrganizationIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PatchOrganizationOrganizationIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewPatchOrganizationOrganizationIDNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPatchOrganizationOrganizationIDNoContent creates a PatchOrganizationOrganizationIDNoContent with default headers values
func NewPatchOrganizationOrganizationIDNoContent() *PatchOrganizationOrganizationIDNoContent {
	return &PatchOrganizationOrganizationIDNoContent{}
}

/* PatchOrganizationOrganizationIDNoContent describes a response with status code 204, with default header values.

Successful response
*/
type PatchOrganizationOrganizationIDNoContent struct {
}

func (o *PatchOrganizationOrganizationIDNoContent) Error() string {
	return fmt.Sprintf("[PATCH /organization/{organizationId}][%d] patchOrganizationOrganizationIdNoContent ", 204)
}

func (o *PatchOrganizationOrganizationIDNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

/*PatchOrganizationOrganizationIDBody patch organization organization ID body
swagger:model PatchOrganizationOrganizationIDBody
*/
type PatchOrganizationOrganizationIDBody struct {

	// data
	Data *models.Organization `json:"data,omitempty"`
}

// Validate validates this patch organization organization ID body
func (o *PatchOrganizationOrganizationIDBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateData(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *PatchOrganizationOrganizationIDBody) validateData(formats strfmt.Registry) error {
	if swag.IsZero(o.Data) { // not required
		return nil
	}

	if o.Data != nil {
		if err := o.Data.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("organization" + "." + "data")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this patch organization organization ID body based on the context it is used
func (o *PatchOrganizationOrganizationIDBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateData(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *PatchOrganizationOrganizationIDBody) contextValidateData(ctx context.Context, formats strfmt.Registry) error {

	if o.Data != nil {
		if err := o.Data.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("organization" + "." + "data")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *PatchOrganizationOrganizationIDBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *PatchOrganizationOrganizationIDBody) UnmarshalBinary(b []byte) error {
	var res PatchOrganizationOrganizationIDBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}