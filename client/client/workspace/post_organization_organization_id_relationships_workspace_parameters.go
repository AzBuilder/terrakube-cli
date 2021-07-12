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

// NewPostOrganizationOrganizationIDRelationshipsWorkspaceParams creates a new PostOrganizationOrganizationIDRelationshipsWorkspaceParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPostOrganizationOrganizationIDRelationshipsWorkspaceParams() *PostOrganizationOrganizationIDRelationshipsWorkspaceParams {
	return &PostOrganizationOrganizationIDRelationshipsWorkspaceParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPostOrganizationOrganizationIDRelationshipsWorkspaceParamsWithTimeout creates a new PostOrganizationOrganizationIDRelationshipsWorkspaceParams object
// with the ability to set a timeout on a request.
func NewPostOrganizationOrganizationIDRelationshipsWorkspaceParamsWithTimeout(timeout time.Duration) *PostOrganizationOrganizationIDRelationshipsWorkspaceParams {
	return &PostOrganizationOrganizationIDRelationshipsWorkspaceParams{
		timeout: timeout,
	}
}

// NewPostOrganizationOrganizationIDRelationshipsWorkspaceParamsWithContext creates a new PostOrganizationOrganizationIDRelationshipsWorkspaceParams object
// with the ability to set a context for a request.
func NewPostOrganizationOrganizationIDRelationshipsWorkspaceParamsWithContext(ctx context.Context) *PostOrganizationOrganizationIDRelationshipsWorkspaceParams {
	return &PostOrganizationOrganizationIDRelationshipsWorkspaceParams{
		Context: ctx,
	}
}

// NewPostOrganizationOrganizationIDRelationshipsWorkspaceParamsWithHTTPClient creates a new PostOrganizationOrganizationIDRelationshipsWorkspaceParams object
// with the ability to set a custom HTTPClient for a request.
func NewPostOrganizationOrganizationIDRelationshipsWorkspaceParamsWithHTTPClient(client *http.Client) *PostOrganizationOrganizationIDRelationshipsWorkspaceParams {
	return &PostOrganizationOrganizationIDRelationshipsWorkspaceParams{
		HTTPClient: client,
	}
}

/* PostOrganizationOrganizationIDRelationshipsWorkspaceParams contains all the parameters to send to the API endpoint
   for the post organization organization ID relationships workspace operation.

   Typically these are written to a http.Request.
*/
type PostOrganizationOrganizationIDRelationshipsWorkspaceParams struct {

	/* OrganizationID.

	   organization Identifier
	*/
	OrganizationID string

	// Relationship.
	Relationship PostOrganizationOrganizationIDRelationshipsWorkspaceBody

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the post organization organization ID relationships workspace params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostOrganizationOrganizationIDRelationshipsWorkspaceParams) WithDefaults() *PostOrganizationOrganizationIDRelationshipsWorkspaceParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the post organization organization ID relationships workspace params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostOrganizationOrganizationIDRelationshipsWorkspaceParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the post organization organization ID relationships workspace params
func (o *PostOrganizationOrganizationIDRelationshipsWorkspaceParams) WithTimeout(timeout time.Duration) *PostOrganizationOrganizationIDRelationshipsWorkspaceParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post organization organization ID relationships workspace params
func (o *PostOrganizationOrganizationIDRelationshipsWorkspaceParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post organization organization ID relationships workspace params
func (o *PostOrganizationOrganizationIDRelationshipsWorkspaceParams) WithContext(ctx context.Context) *PostOrganizationOrganizationIDRelationshipsWorkspaceParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post organization organization ID relationships workspace params
func (o *PostOrganizationOrganizationIDRelationshipsWorkspaceParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post organization organization ID relationships workspace params
func (o *PostOrganizationOrganizationIDRelationshipsWorkspaceParams) WithHTTPClient(client *http.Client) *PostOrganizationOrganizationIDRelationshipsWorkspaceParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post organization organization ID relationships workspace params
func (o *PostOrganizationOrganizationIDRelationshipsWorkspaceParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithOrganizationID adds the organizationID to the post organization organization ID relationships workspace params
func (o *PostOrganizationOrganizationIDRelationshipsWorkspaceParams) WithOrganizationID(organizationID string) *PostOrganizationOrganizationIDRelationshipsWorkspaceParams {
	o.SetOrganizationID(organizationID)
	return o
}

// SetOrganizationID adds the organizationId to the post organization organization ID relationships workspace params
func (o *PostOrganizationOrganizationIDRelationshipsWorkspaceParams) SetOrganizationID(organizationID string) {
	o.OrganizationID = organizationID
}

// WithRelationship adds the relationship to the post organization organization ID relationships workspace params
func (o *PostOrganizationOrganizationIDRelationshipsWorkspaceParams) WithRelationship(relationship PostOrganizationOrganizationIDRelationshipsWorkspaceBody) *PostOrganizationOrganizationIDRelationshipsWorkspaceParams {
	o.SetRelationship(relationship)
	return o
}

// SetRelationship adds the relationship to the post organization organization ID relationships workspace params
func (o *PostOrganizationOrganizationIDRelationshipsWorkspaceParams) SetRelationship(relationship PostOrganizationOrganizationIDRelationshipsWorkspaceBody) {
	o.Relationship = relationship
}

// WriteToRequest writes these params to a swagger request
func (o *PostOrganizationOrganizationIDRelationshipsWorkspaceParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param organizationId
	if err := r.SetPathParam("organizationId", o.OrganizationID); err != nil {
		return err
	}
	if err := r.SetBodyParam(o.Relationship); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
