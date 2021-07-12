// Code generated by go-swagger; DO NOT EDIT.

package module

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// DeleteOrganizationOrganizationIDModuleModuleIDReader is a Reader for the DeleteOrganizationOrganizationIDModuleModuleID structure.
type DeleteOrganizationOrganizationIDModuleModuleIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteOrganizationOrganizationIDModuleModuleIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewDeleteOrganizationOrganizationIDModuleModuleIDNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDeleteOrganizationOrganizationIDModuleModuleIDNoContent creates a DeleteOrganizationOrganizationIDModuleModuleIDNoContent with default headers values
func NewDeleteOrganizationOrganizationIDModuleModuleIDNoContent() *DeleteOrganizationOrganizationIDModuleModuleIDNoContent {
	return &DeleteOrganizationOrganizationIDModuleModuleIDNoContent{}
}

/* DeleteOrganizationOrganizationIDModuleModuleIDNoContent describes a response with status code 204, with default header values.

Successful response
*/
type DeleteOrganizationOrganizationIDModuleModuleIDNoContent struct {
}

func (o *DeleteOrganizationOrganizationIDModuleModuleIDNoContent) Error() string {
	return fmt.Sprintf("[DELETE /organization/{organizationId}/module/{moduleId}][%d] deleteOrganizationOrganizationIdModuleModuleIdNoContent ", 204)
}

func (o *DeleteOrganizationOrganizationIDModuleModuleIDNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
