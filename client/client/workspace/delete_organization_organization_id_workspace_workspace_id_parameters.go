// Code generated by go-swagger; DO NOT EDIT.

package workspace

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
)

// NewDeleteOrganizationOrganizationIDWorkspaceWorkspaceIDParams creates a new DeleteOrganizationOrganizationIDWorkspaceWorkspaceIDParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDeleteOrganizationOrganizationIDWorkspaceWorkspaceIDParams() *DeleteOrganizationOrganizationIDWorkspaceWorkspaceIDParams {
	return &DeleteOrganizationOrganizationIDWorkspaceWorkspaceIDParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteOrganizationOrganizationIDWorkspaceWorkspaceIDParamsWithTimeout creates a new DeleteOrganizationOrganizationIDWorkspaceWorkspaceIDParams object
// with the ability to set a timeout on a request.
func NewDeleteOrganizationOrganizationIDWorkspaceWorkspaceIDParamsWithTimeout(timeout time.Duration) *DeleteOrganizationOrganizationIDWorkspaceWorkspaceIDParams {
	return &DeleteOrganizationOrganizationIDWorkspaceWorkspaceIDParams{
		timeout: timeout,
	}
}

// NewDeleteOrganizationOrganizationIDWorkspaceWorkspaceIDParamsWithContext creates a new DeleteOrganizationOrganizationIDWorkspaceWorkspaceIDParams object
// with the ability to set a context for a request.
func NewDeleteOrganizationOrganizationIDWorkspaceWorkspaceIDParamsWithContext(ctx context.Context) *DeleteOrganizationOrganizationIDWorkspaceWorkspaceIDParams {
	return &DeleteOrganizationOrganizationIDWorkspaceWorkspaceIDParams{
		Context: ctx,
	}
}

// NewDeleteOrganizationOrganizationIDWorkspaceWorkspaceIDParamsWithHTTPClient creates a new DeleteOrganizationOrganizationIDWorkspaceWorkspaceIDParams object
// with the ability to set a custom HTTPClient for a request.
func NewDeleteOrganizationOrganizationIDWorkspaceWorkspaceIDParamsWithHTTPClient(client *http.Client) *DeleteOrganizationOrganizationIDWorkspaceWorkspaceIDParams {
	return &DeleteOrganizationOrganizationIDWorkspaceWorkspaceIDParams{
		HTTPClient: client,
	}
}

/* DeleteOrganizationOrganizationIDWorkspaceWorkspaceIDParams contains all the parameters to send to the API endpoint
   for the delete organization organization ID workspace workspace ID operation.

   Typically these are written to a http.Request.
*/
type DeleteOrganizationOrganizationIDWorkspaceWorkspaceIDParams struct {

	/* OrganizationID.

	   organization Identifier
	*/
	OrganizationID string

	/* WorkspaceID.

	   workspace Identifier
	*/
	WorkspaceID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the delete organization organization ID workspace workspace ID params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteOrganizationOrganizationIDWorkspaceWorkspaceIDParams) WithDefaults() *DeleteOrganizationOrganizationIDWorkspaceWorkspaceIDParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the delete organization organization ID workspace workspace ID params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteOrganizationOrganizationIDWorkspaceWorkspaceIDParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the delete organization organization ID workspace workspace ID params
func (o *DeleteOrganizationOrganizationIDWorkspaceWorkspaceIDParams) WithTimeout(timeout time.Duration) *DeleteOrganizationOrganizationIDWorkspaceWorkspaceIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete organization organization ID workspace workspace ID params
func (o *DeleteOrganizationOrganizationIDWorkspaceWorkspaceIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete organization organization ID workspace workspace ID params
func (o *DeleteOrganizationOrganizationIDWorkspaceWorkspaceIDParams) WithContext(ctx context.Context) *DeleteOrganizationOrganizationIDWorkspaceWorkspaceIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete organization organization ID workspace workspace ID params
func (o *DeleteOrganizationOrganizationIDWorkspaceWorkspaceIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete organization organization ID workspace workspace ID params
func (o *DeleteOrganizationOrganizationIDWorkspaceWorkspaceIDParams) WithHTTPClient(client *http.Client) *DeleteOrganizationOrganizationIDWorkspaceWorkspaceIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete organization organization ID workspace workspace ID params
func (o *DeleteOrganizationOrganizationIDWorkspaceWorkspaceIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithOrganizationID adds the organizationID to the delete organization organization ID workspace workspace ID params
func (o *DeleteOrganizationOrganizationIDWorkspaceWorkspaceIDParams) WithOrganizationID(organizationID string) *DeleteOrganizationOrganizationIDWorkspaceWorkspaceIDParams {
	o.SetOrganizationID(organizationID)
	return o
}

// SetOrganizationID adds the organizationId to the delete organization organization ID workspace workspace ID params
func (o *DeleteOrganizationOrganizationIDWorkspaceWorkspaceIDParams) SetOrganizationID(organizationID string) {
	o.OrganizationID = organizationID
}

// WithWorkspaceID adds the workspaceID to the delete organization organization ID workspace workspace ID params
func (o *DeleteOrganizationOrganizationIDWorkspaceWorkspaceIDParams) WithWorkspaceID(workspaceID string) *DeleteOrganizationOrganizationIDWorkspaceWorkspaceIDParams {
	o.SetWorkspaceID(workspaceID)
	return o
}

// SetWorkspaceID adds the workspaceId to the delete organization organization ID workspace workspace ID params
func (o *DeleteOrganizationOrganizationIDWorkspaceWorkspaceIDParams) SetWorkspaceID(workspaceID string) {
	o.WorkspaceID = workspaceID
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteOrganizationOrganizationIDWorkspaceWorkspaceIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param organizationId
	if err := r.SetPathParam("organizationId", o.OrganizationID); err != nil {
		return err
	}

	// path param workspaceId
	if err := r.SetPathParam("workspaceId", o.WorkspaceID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
