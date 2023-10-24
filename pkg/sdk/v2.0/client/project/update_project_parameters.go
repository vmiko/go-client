// Code generated by go-swagger; DO NOT EDIT.

package project

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

	"github.com/vmiko/go-client/pkg/sdk/v2.0/models"
)

// NewUpdateProjectParams creates a new UpdateProjectParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewUpdateProjectParams() *UpdateProjectParams {
	return &UpdateProjectParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateProjectParamsWithTimeout creates a new UpdateProjectParams object
// with the ability to set a timeout on a request.
func NewUpdateProjectParamsWithTimeout(timeout time.Duration) *UpdateProjectParams {
	return &UpdateProjectParams{
		timeout: timeout,
	}
}

// NewUpdateProjectParamsWithContext creates a new UpdateProjectParams object
// with the ability to set a context for a request.
func NewUpdateProjectParamsWithContext(ctx context.Context) *UpdateProjectParams {
	return &UpdateProjectParams{
		Context: ctx,
	}
}

// NewUpdateProjectParamsWithHTTPClient creates a new UpdateProjectParams object
// with the ability to set a custom HTTPClient for a request.
func NewUpdateProjectParamsWithHTTPClient(client *http.Client) *UpdateProjectParams {
	return &UpdateProjectParams{
		HTTPClient: client,
	}
}

/*
UpdateProjectParams contains all the parameters to send to the API endpoint

	for the update project operation.

	Typically these are written to a http.Request.
*/
type UpdateProjectParams struct {

	/* XIsResourceName.

	   The flag to indicate whether the parameter which supports both name and id in the path is the name of the resource. When the X-Is-Resource-Name is false and the parameter can be converted to an integer, the parameter will be as an id, otherwise, it will be as a name.
	*/
	XIsResourceName *bool

	/* XRequestID.

	   An unique ID for the request
	*/
	XRequestID *string

	/* Project.

	   Updates of project.
	*/
	Project *models.ProjectReq

	/* ProjectNameOrID.

	   The name or id of the project
	*/
	ProjectNameOrID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the update project params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateProjectParams) WithDefaults() *UpdateProjectParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the update project params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateProjectParams) SetDefaults() {
	var (
		xIsResourceNameDefault = bool(false)
	)

	val := UpdateProjectParams{
		XIsResourceName: &xIsResourceNameDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the update project params
func (o *UpdateProjectParams) WithTimeout(timeout time.Duration) *UpdateProjectParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update project params
func (o *UpdateProjectParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update project params
func (o *UpdateProjectParams) WithContext(ctx context.Context) *UpdateProjectParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update project params
func (o *UpdateProjectParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update project params
func (o *UpdateProjectParams) WithHTTPClient(client *http.Client) *UpdateProjectParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update project params
func (o *UpdateProjectParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithXIsResourceName adds the xIsResourceName to the update project params
func (o *UpdateProjectParams) WithXIsResourceName(xIsResourceName *bool) *UpdateProjectParams {
	o.SetXIsResourceName(xIsResourceName)
	return o
}

// SetXIsResourceName adds the xIsResourceName to the update project params
func (o *UpdateProjectParams) SetXIsResourceName(xIsResourceName *bool) {
	o.XIsResourceName = xIsResourceName
}

// WithXRequestID adds the xRequestID to the update project params
func (o *UpdateProjectParams) WithXRequestID(xRequestID *string) *UpdateProjectParams {
	o.SetXRequestID(xRequestID)
	return o
}

// SetXRequestID adds the xRequestId to the update project params
func (o *UpdateProjectParams) SetXRequestID(xRequestID *string) {
	o.XRequestID = xRequestID
}

// WithProject adds the project to the update project params
func (o *UpdateProjectParams) WithProject(project *models.ProjectReq) *UpdateProjectParams {
	o.SetProject(project)
	return o
}

// SetProject adds the project to the update project params
func (o *UpdateProjectParams) SetProject(project *models.ProjectReq) {
	o.Project = project
}

// WithProjectNameOrID adds the projectNameOrID to the update project params
func (o *UpdateProjectParams) WithProjectNameOrID(projectNameOrID string) *UpdateProjectParams {
	o.SetProjectNameOrID(projectNameOrID)
	return o
}

// SetProjectNameOrID adds the projectNameOrId to the update project params
func (o *UpdateProjectParams) SetProjectNameOrID(projectNameOrID string) {
	o.ProjectNameOrID = projectNameOrID
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateProjectParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.XIsResourceName != nil {

		// header param X-Is-Resource-Name
		if err := r.SetHeaderParam("X-Is-Resource-Name", swag.FormatBool(*o.XIsResourceName)); err != nil {
			return err
		}
	}

	if o.XRequestID != nil {

		// header param X-Request-Id
		if err := r.SetHeaderParam("X-Request-Id", *o.XRequestID); err != nil {
			return err
		}
	}
	if o.Project != nil {
		if err := r.SetBodyParam(o.Project); err != nil {
			return err
		}
	}

	// path param project_name_or_id
	if err := r.SetPathParam("project_name_or_id", o.ProjectNameOrID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
