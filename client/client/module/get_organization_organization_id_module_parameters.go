// Code generated by go-swagger; DO NOT EDIT.

package module

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
	"github.com/go-openapi/swag"
)

// NewGetOrganizationOrganizationIDModuleParams creates a new GetOrganizationOrganizationIDModuleParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetOrganizationOrganizationIDModuleParams() *GetOrganizationOrganizationIDModuleParams {
	return &GetOrganizationOrganizationIDModuleParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetOrganizationOrganizationIDModuleParamsWithTimeout creates a new GetOrganizationOrganizationIDModuleParams object
// with the ability to set a timeout on a request.
func NewGetOrganizationOrganizationIDModuleParamsWithTimeout(timeout time.Duration) *GetOrganizationOrganizationIDModuleParams {
	return &GetOrganizationOrganizationIDModuleParams{
		timeout: timeout,
	}
}

// NewGetOrganizationOrganizationIDModuleParamsWithContext creates a new GetOrganizationOrganizationIDModuleParams object
// with the ability to set a context for a request.
func NewGetOrganizationOrganizationIDModuleParamsWithContext(ctx context.Context) *GetOrganizationOrganizationIDModuleParams {
	return &GetOrganizationOrganizationIDModuleParams{
		Context: ctx,
	}
}

// NewGetOrganizationOrganizationIDModuleParamsWithHTTPClient creates a new GetOrganizationOrganizationIDModuleParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetOrganizationOrganizationIDModuleParamsWithHTTPClient(client *http.Client) *GetOrganizationOrganizationIDModuleParams {
	return &GetOrganizationOrganizationIDModuleParams{
		HTTPClient: client,
	}
}

/* GetOrganizationOrganizationIDModuleParams contains all the parameters to send to the API endpoint
   for the get organization organization ID module operation.

   Typically these are written to a http.Request.
*/
type GetOrganizationOrganizationIDModuleParams struct {

	/* FieldsModule.

	   Selects the set of module fields that should be returned in the result.
	*/
	FieldsModule []string

	/* FilterModule.

	   Filters the collection of module using a 'disjoint' RSQL expression
	*/
	FilterModule *string

	/* Include.

	   Selects the set of relationships that should be expanded as a compound document in the result.
	*/
	Include []string

	/* OrganizationID.

	   organization Identifier
	*/
	OrganizationID string

	/* PageLimit.

	   Maximum number of items to return.  Can be used with page[offset]
	*/
	PageLimit *int64

	/* PageNumber.

	   Number of pages to return.  Can be used with page[size]
	*/
	PageNumber *int64

	/* PageOffset.

	   Offset from 0 to start paginating.  Can be used with page[limit]
	*/
	PageOffset *int64

	/* PageSize.

	   Number of elements per page.  Can be used with page[number]
	*/
	PageSize *int64

	/* PageTotals.

	   For requesting total pages/records be included in the response page meta data
	*/
	PageTotals *string

	/* Sort.

	   Sorts the collection on the selected attributes.  A prefix of '-' sorts descending
	*/
	Sort []string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get organization organization ID module params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetOrganizationOrganizationIDModuleParams) WithDefaults() *GetOrganizationOrganizationIDModuleParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get organization organization ID module params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetOrganizationOrganizationIDModuleParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get organization organization ID module params
func (o *GetOrganizationOrganizationIDModuleParams) WithTimeout(timeout time.Duration) *GetOrganizationOrganizationIDModuleParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get organization organization ID module params
func (o *GetOrganizationOrganizationIDModuleParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get organization organization ID module params
func (o *GetOrganizationOrganizationIDModuleParams) WithContext(ctx context.Context) *GetOrganizationOrganizationIDModuleParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get organization organization ID module params
func (o *GetOrganizationOrganizationIDModuleParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get organization organization ID module params
func (o *GetOrganizationOrganizationIDModuleParams) WithHTTPClient(client *http.Client) *GetOrganizationOrganizationIDModuleParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get organization organization ID module params
func (o *GetOrganizationOrganizationIDModuleParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithFieldsModule adds the fieldsModule to the get organization organization ID module params
func (o *GetOrganizationOrganizationIDModuleParams) WithFieldsModule(fieldsModule []string) *GetOrganizationOrganizationIDModuleParams {
	o.SetFieldsModule(fieldsModule)
	return o
}

// SetFieldsModule adds the fieldsModule to the get organization organization ID module params
func (o *GetOrganizationOrganizationIDModuleParams) SetFieldsModule(fieldsModule []string) {
	o.FieldsModule = fieldsModule
}

// WithFilterModule adds the filterModule to the get organization organization ID module params
func (o *GetOrganizationOrganizationIDModuleParams) WithFilterModule(filterModule *string) *GetOrganizationOrganizationIDModuleParams {
	o.SetFilterModule(filterModule)
	return o
}

// SetFilterModule adds the filterModule to the get organization organization ID module params
func (o *GetOrganizationOrganizationIDModuleParams) SetFilterModule(filterModule *string) {
	o.FilterModule = filterModule
}

// WithInclude adds the include to the get organization organization ID module params
func (o *GetOrganizationOrganizationIDModuleParams) WithInclude(include []string) *GetOrganizationOrganizationIDModuleParams {
	o.SetInclude(include)
	return o
}

// SetInclude adds the include to the get organization organization ID module params
func (o *GetOrganizationOrganizationIDModuleParams) SetInclude(include []string) {
	o.Include = include
}

// WithOrganizationID adds the organizationID to the get organization organization ID module params
func (o *GetOrganizationOrganizationIDModuleParams) WithOrganizationID(organizationID string) *GetOrganizationOrganizationIDModuleParams {
	o.SetOrganizationID(organizationID)
	return o
}

// SetOrganizationID adds the organizationId to the get organization organization ID module params
func (o *GetOrganizationOrganizationIDModuleParams) SetOrganizationID(organizationID string) {
	o.OrganizationID = organizationID
}

// WithPageLimit adds the pageLimit to the get organization organization ID module params
func (o *GetOrganizationOrganizationIDModuleParams) WithPageLimit(pageLimit *int64) *GetOrganizationOrganizationIDModuleParams {
	o.SetPageLimit(pageLimit)
	return o
}

// SetPageLimit adds the pageLimit to the get organization organization ID module params
func (o *GetOrganizationOrganizationIDModuleParams) SetPageLimit(pageLimit *int64) {
	o.PageLimit = pageLimit
}

// WithPageNumber adds the pageNumber to the get organization organization ID module params
func (o *GetOrganizationOrganizationIDModuleParams) WithPageNumber(pageNumber *int64) *GetOrganizationOrganizationIDModuleParams {
	o.SetPageNumber(pageNumber)
	return o
}

// SetPageNumber adds the pageNumber to the get organization organization ID module params
func (o *GetOrganizationOrganizationIDModuleParams) SetPageNumber(pageNumber *int64) {
	o.PageNumber = pageNumber
}

// WithPageOffset adds the pageOffset to the get organization organization ID module params
func (o *GetOrganizationOrganizationIDModuleParams) WithPageOffset(pageOffset *int64) *GetOrganizationOrganizationIDModuleParams {
	o.SetPageOffset(pageOffset)
	return o
}

// SetPageOffset adds the pageOffset to the get organization organization ID module params
func (o *GetOrganizationOrganizationIDModuleParams) SetPageOffset(pageOffset *int64) {
	o.PageOffset = pageOffset
}

// WithPageSize adds the pageSize to the get organization organization ID module params
func (o *GetOrganizationOrganizationIDModuleParams) WithPageSize(pageSize *int64) *GetOrganizationOrganizationIDModuleParams {
	o.SetPageSize(pageSize)
	return o
}

// SetPageSize adds the pageSize to the get organization organization ID module params
func (o *GetOrganizationOrganizationIDModuleParams) SetPageSize(pageSize *int64) {
	o.PageSize = pageSize
}

// WithPageTotals adds the pageTotals to the get organization organization ID module params
func (o *GetOrganizationOrganizationIDModuleParams) WithPageTotals(pageTotals *string) *GetOrganizationOrganizationIDModuleParams {
	o.SetPageTotals(pageTotals)
	return o
}

// SetPageTotals adds the pageTotals to the get organization organization ID module params
func (o *GetOrganizationOrganizationIDModuleParams) SetPageTotals(pageTotals *string) {
	o.PageTotals = pageTotals
}

// WithSort adds the sort to the get organization organization ID module params
func (o *GetOrganizationOrganizationIDModuleParams) WithSort(sort []string) *GetOrganizationOrganizationIDModuleParams {
	o.SetSort(sort)
	return o
}

// SetSort adds the sort to the get organization organization ID module params
func (o *GetOrganizationOrganizationIDModuleParams) SetSort(sort []string) {
	o.Sort = sort
}

// WriteToRequest writes these params to a swagger request
func (o *GetOrganizationOrganizationIDModuleParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.FieldsModule != nil {

		// binding items for fields[module]
		joinedFieldsModule := o.bindParamFieldsModule(reg)

		// query array param fields[module]
		if err := r.SetQueryParam("fields[module]", joinedFieldsModule...); err != nil {
			return err
		}
	}

	if o.FilterModule != nil {

		// query param filter[module]
		var qrFilterModule string

		if o.FilterModule != nil {
			qrFilterModule = *o.FilterModule
		}
		qFilterModule := qrFilterModule
		if qFilterModule != "" {

			if err := r.SetQueryParam("filter[module]", qFilterModule); err != nil {
				return err
			}
		}
	}

	if o.Include != nil {

		// binding items for include
		joinedInclude := o.bindParamInclude(reg)

		// query array param include
		if err := r.SetQueryParam("include", joinedInclude...); err != nil {
			return err
		}
	}

	// path param organizationId
	if err := r.SetPathParam("organizationId", o.OrganizationID); err != nil {
		return err
	}

	if o.PageLimit != nil {

		// query param page[limit]
		var qrPageLimit int64

		if o.PageLimit != nil {
			qrPageLimit = *o.PageLimit
		}
		qPageLimit := swag.FormatInt64(qrPageLimit)
		if qPageLimit != "" {

			if err := r.SetQueryParam("page[limit]", qPageLimit); err != nil {
				return err
			}
		}
	}

	if o.PageNumber != nil {

		// query param page[number]
		var qrPageNumber int64

		if o.PageNumber != nil {
			qrPageNumber = *o.PageNumber
		}
		qPageNumber := swag.FormatInt64(qrPageNumber)
		if qPageNumber != "" {

			if err := r.SetQueryParam("page[number]", qPageNumber); err != nil {
				return err
			}
		}
	}

	if o.PageOffset != nil {

		// query param page[offset]
		var qrPageOffset int64

		if o.PageOffset != nil {
			qrPageOffset = *o.PageOffset
		}
		qPageOffset := swag.FormatInt64(qrPageOffset)
		if qPageOffset != "" {

			if err := r.SetQueryParam("page[offset]", qPageOffset); err != nil {
				return err
			}
		}
	}

	if o.PageSize != nil {

		// query param page[size]
		var qrPageSize int64

		if o.PageSize != nil {
			qrPageSize = *o.PageSize
		}
		qPageSize := swag.FormatInt64(qrPageSize)
		if qPageSize != "" {

			if err := r.SetQueryParam("page[size]", qPageSize); err != nil {
				return err
			}
		}
	}

	if o.PageTotals != nil {

		// query param page[totals]
		var qrPageTotals string

		if o.PageTotals != nil {
			qrPageTotals = *o.PageTotals
		}
		qPageTotals := qrPageTotals
		if qPageTotals != "" {

			if err := r.SetQueryParam("page[totals]", qPageTotals); err != nil {
				return err
			}
		}
	}

	if o.Sort != nil {

		// binding items for sort
		joinedSort := o.bindParamSort(reg)

		// query array param sort
		if err := r.SetQueryParam("sort", joinedSort...); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindParamGetOrganizationOrganizationIDModule binds the parameter fields[module]
func (o *GetOrganizationOrganizationIDModuleParams) bindParamFieldsModule(formats strfmt.Registry) []string {
	fieldsModuleIR := o.FieldsModule

	var fieldsModuleIC []string
	for _, fieldsModuleIIR := range fieldsModuleIR { // explode []string

		fieldsModuleIIV := fieldsModuleIIR // string as string
		fieldsModuleIC = append(fieldsModuleIC, fieldsModuleIIV)
	}

	// items.CollectionFormat: "csv"
	fieldsModuleIS := swag.JoinByFormat(fieldsModuleIC, "csv")

	return fieldsModuleIS
}

// bindParamGetOrganizationOrganizationIDModule binds the parameter include
func (o *GetOrganizationOrganizationIDModuleParams) bindParamInclude(formats strfmt.Registry) []string {
	includeIR := o.Include

	var includeIC []string
	for _, includeIIR := range includeIR { // explode []string

		includeIIV := includeIIR // string as string
		includeIC = append(includeIC, includeIIV)
	}

	// items.CollectionFormat: "csv"
	includeIS := swag.JoinByFormat(includeIC, "csv")

	return includeIS
}

// bindParamGetOrganizationOrganizationIDModule binds the parameter sort
func (o *GetOrganizationOrganizationIDModuleParams) bindParamSort(formats strfmt.Registry) []string {
	sortIR := o.Sort

	var sortIC []string
	for _, sortIIR := range sortIR { // explode []string

		sortIIV := sortIIR // string as string
		sortIC = append(sortIC, sortIIV)
	}

	// items.CollectionFormat: "csv"
	sortIS := swag.JoinByFormat(sortIC, "csv")

	return sortIS
}
