// Code generated by go-swagger; DO NOT EDIT.

package replication

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/vmiko/go-client/pkg/sdk/v2.0/models"
)

// GetReplicationExecutionReader is a Reader for the GetReplicationExecution structure.
type GetReplicationExecutionReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetReplicationExecutionReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetReplicationExecutionOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewGetReplicationExecutionUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetReplicationExecutionForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetReplicationExecutionNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetReplicationExecutionInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetReplicationExecutionOK creates a GetReplicationExecutionOK with default headers values
func NewGetReplicationExecutionOK() *GetReplicationExecutionOK {
	return &GetReplicationExecutionOK{}
}

/*
GetReplicationExecutionOK describes a response with status code 200, with default header values.

Success
*/
type GetReplicationExecutionOK struct {
	Payload *models.ReplicationExecution
}

// IsSuccess returns true when this get replication execution o k response has a 2xx status code
func (o *GetReplicationExecutionOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get replication execution o k response has a 3xx status code
func (o *GetReplicationExecutionOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get replication execution o k response has a 4xx status code
func (o *GetReplicationExecutionOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get replication execution o k response has a 5xx status code
func (o *GetReplicationExecutionOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get replication execution o k response a status code equal to that given
func (o *GetReplicationExecutionOK) IsCode(code int) bool {
	return code == 200
}

func (o *GetReplicationExecutionOK) Error() string {
	return fmt.Sprintf("[GET /replication/executions/{id}][%d] getReplicationExecutionOK  %+v", 200, o.Payload)
}

func (o *GetReplicationExecutionOK) String() string {
	return fmt.Sprintf("[GET /replication/executions/{id}][%d] getReplicationExecutionOK  %+v", 200, o.Payload)
}

func (o *GetReplicationExecutionOK) GetPayload() *models.ReplicationExecution {
	return o.Payload
}

func (o *GetReplicationExecutionOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ReplicationExecution)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetReplicationExecutionUnauthorized creates a GetReplicationExecutionUnauthorized with default headers values
func NewGetReplicationExecutionUnauthorized() *GetReplicationExecutionUnauthorized {
	return &GetReplicationExecutionUnauthorized{}
}

/*
GetReplicationExecutionUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type GetReplicationExecutionUnauthorized struct {

	/* The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *models.Errors
}

// IsSuccess returns true when this get replication execution unauthorized response has a 2xx status code
func (o *GetReplicationExecutionUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get replication execution unauthorized response has a 3xx status code
func (o *GetReplicationExecutionUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get replication execution unauthorized response has a 4xx status code
func (o *GetReplicationExecutionUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this get replication execution unauthorized response has a 5xx status code
func (o *GetReplicationExecutionUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this get replication execution unauthorized response a status code equal to that given
func (o *GetReplicationExecutionUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *GetReplicationExecutionUnauthorized) Error() string {
	return fmt.Sprintf("[GET /replication/executions/{id}][%d] getReplicationExecutionUnauthorized  %+v", 401, o.Payload)
}

func (o *GetReplicationExecutionUnauthorized) String() string {
	return fmt.Sprintf("[GET /replication/executions/{id}][%d] getReplicationExecutionUnauthorized  %+v", 401, o.Payload)
}

func (o *GetReplicationExecutionUnauthorized) GetPayload() *models.Errors {
	return o.Payload
}

func (o *GetReplicationExecutionUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header X-Request-Id
	hdrXRequestID := response.GetHeader("X-Request-Id")

	if hdrXRequestID != "" {
		o.XRequestID = hdrXRequestID
	}

	o.Payload = new(models.Errors)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetReplicationExecutionForbidden creates a GetReplicationExecutionForbidden with default headers values
func NewGetReplicationExecutionForbidden() *GetReplicationExecutionForbidden {
	return &GetReplicationExecutionForbidden{}
}

/*
GetReplicationExecutionForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type GetReplicationExecutionForbidden struct {

	/* The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *models.Errors
}

// IsSuccess returns true when this get replication execution forbidden response has a 2xx status code
func (o *GetReplicationExecutionForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get replication execution forbidden response has a 3xx status code
func (o *GetReplicationExecutionForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get replication execution forbidden response has a 4xx status code
func (o *GetReplicationExecutionForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this get replication execution forbidden response has a 5xx status code
func (o *GetReplicationExecutionForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this get replication execution forbidden response a status code equal to that given
func (o *GetReplicationExecutionForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *GetReplicationExecutionForbidden) Error() string {
	return fmt.Sprintf("[GET /replication/executions/{id}][%d] getReplicationExecutionForbidden  %+v", 403, o.Payload)
}

func (o *GetReplicationExecutionForbidden) String() string {
	return fmt.Sprintf("[GET /replication/executions/{id}][%d] getReplicationExecutionForbidden  %+v", 403, o.Payload)
}

func (o *GetReplicationExecutionForbidden) GetPayload() *models.Errors {
	return o.Payload
}

func (o *GetReplicationExecutionForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header X-Request-Id
	hdrXRequestID := response.GetHeader("X-Request-Id")

	if hdrXRequestID != "" {
		o.XRequestID = hdrXRequestID
	}

	o.Payload = new(models.Errors)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetReplicationExecutionNotFound creates a GetReplicationExecutionNotFound with default headers values
func NewGetReplicationExecutionNotFound() *GetReplicationExecutionNotFound {
	return &GetReplicationExecutionNotFound{}
}

/*
GetReplicationExecutionNotFound describes a response with status code 404, with default header values.

Not found
*/
type GetReplicationExecutionNotFound struct {

	/* The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *models.Errors
}

// IsSuccess returns true when this get replication execution not found response has a 2xx status code
func (o *GetReplicationExecutionNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get replication execution not found response has a 3xx status code
func (o *GetReplicationExecutionNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get replication execution not found response has a 4xx status code
func (o *GetReplicationExecutionNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get replication execution not found response has a 5xx status code
func (o *GetReplicationExecutionNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get replication execution not found response a status code equal to that given
func (o *GetReplicationExecutionNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *GetReplicationExecutionNotFound) Error() string {
	return fmt.Sprintf("[GET /replication/executions/{id}][%d] getReplicationExecutionNotFound  %+v", 404, o.Payload)
}

func (o *GetReplicationExecutionNotFound) String() string {
	return fmt.Sprintf("[GET /replication/executions/{id}][%d] getReplicationExecutionNotFound  %+v", 404, o.Payload)
}

func (o *GetReplicationExecutionNotFound) GetPayload() *models.Errors {
	return o.Payload
}

func (o *GetReplicationExecutionNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header X-Request-Id
	hdrXRequestID := response.GetHeader("X-Request-Id")

	if hdrXRequestID != "" {
		o.XRequestID = hdrXRequestID
	}

	o.Payload = new(models.Errors)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetReplicationExecutionInternalServerError creates a GetReplicationExecutionInternalServerError with default headers values
func NewGetReplicationExecutionInternalServerError() *GetReplicationExecutionInternalServerError {
	return &GetReplicationExecutionInternalServerError{}
}

/*
GetReplicationExecutionInternalServerError describes a response with status code 500, with default header values.

Internal server error
*/
type GetReplicationExecutionInternalServerError struct {

	/* The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *models.Errors
}

// IsSuccess returns true when this get replication execution internal server error response has a 2xx status code
func (o *GetReplicationExecutionInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get replication execution internal server error response has a 3xx status code
func (o *GetReplicationExecutionInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get replication execution internal server error response has a 4xx status code
func (o *GetReplicationExecutionInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get replication execution internal server error response has a 5xx status code
func (o *GetReplicationExecutionInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get replication execution internal server error response a status code equal to that given
func (o *GetReplicationExecutionInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *GetReplicationExecutionInternalServerError) Error() string {
	return fmt.Sprintf("[GET /replication/executions/{id}][%d] getReplicationExecutionInternalServerError  %+v", 500, o.Payload)
}

func (o *GetReplicationExecutionInternalServerError) String() string {
	return fmt.Sprintf("[GET /replication/executions/{id}][%d] getReplicationExecutionInternalServerError  %+v", 500, o.Payload)
}

func (o *GetReplicationExecutionInternalServerError) GetPayload() *models.Errors {
	return o.Payload
}

func (o *GetReplicationExecutionInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header X-Request-Id
	hdrXRequestID := response.GetHeader("X-Request-Id")

	if hdrXRequestID != "" {
		o.XRequestID = hdrXRequestID
	}

	o.Payload = new(models.Errors)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
