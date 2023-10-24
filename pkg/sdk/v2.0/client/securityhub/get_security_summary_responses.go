// Code generated by go-swagger; DO NOT EDIT.

package securityhub

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/vmiko/go-client/pkg/sdk/v2.0/models"
)

// GetSecuritySummaryReader is a Reader for the GetSecuritySummary structure.
type GetSecuritySummaryReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetSecuritySummaryReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetSecuritySummaryOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewGetSecuritySummaryUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetSecuritySummaryForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetSecuritySummaryNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetSecuritySummaryInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetSecuritySummaryOK creates a GetSecuritySummaryOK with default headers values
func NewGetSecuritySummaryOK() *GetSecuritySummaryOK {
	return &GetSecuritySummaryOK{}
}

/*
GetSecuritySummaryOK describes a response with status code 200, with default header values.

Success
*/
type GetSecuritySummaryOK struct {
	Payload *models.SecuritySummary
}

// IsSuccess returns true when this get security summary o k response has a 2xx status code
func (o *GetSecuritySummaryOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get security summary o k response has a 3xx status code
func (o *GetSecuritySummaryOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get security summary o k response has a 4xx status code
func (o *GetSecuritySummaryOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get security summary o k response has a 5xx status code
func (o *GetSecuritySummaryOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get security summary o k response a status code equal to that given
func (o *GetSecuritySummaryOK) IsCode(code int) bool {
	return code == 200
}

func (o *GetSecuritySummaryOK) Error() string {
	return fmt.Sprintf("[GET /security/summary][%d] getSecuritySummaryOK  %+v", 200, o.Payload)
}

func (o *GetSecuritySummaryOK) String() string {
	return fmt.Sprintf("[GET /security/summary][%d] getSecuritySummaryOK  %+v", 200, o.Payload)
}

func (o *GetSecuritySummaryOK) GetPayload() *models.SecuritySummary {
	return o.Payload
}

func (o *GetSecuritySummaryOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SecuritySummary)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetSecuritySummaryUnauthorized creates a GetSecuritySummaryUnauthorized with default headers values
func NewGetSecuritySummaryUnauthorized() *GetSecuritySummaryUnauthorized {
	return &GetSecuritySummaryUnauthorized{}
}

/*
GetSecuritySummaryUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type GetSecuritySummaryUnauthorized struct {

	/* The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *models.Errors
}

// IsSuccess returns true when this get security summary unauthorized response has a 2xx status code
func (o *GetSecuritySummaryUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get security summary unauthorized response has a 3xx status code
func (o *GetSecuritySummaryUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get security summary unauthorized response has a 4xx status code
func (o *GetSecuritySummaryUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this get security summary unauthorized response has a 5xx status code
func (o *GetSecuritySummaryUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this get security summary unauthorized response a status code equal to that given
func (o *GetSecuritySummaryUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *GetSecuritySummaryUnauthorized) Error() string {
	return fmt.Sprintf("[GET /security/summary][%d] getSecuritySummaryUnauthorized  %+v", 401, o.Payload)
}

func (o *GetSecuritySummaryUnauthorized) String() string {
	return fmt.Sprintf("[GET /security/summary][%d] getSecuritySummaryUnauthorized  %+v", 401, o.Payload)
}

func (o *GetSecuritySummaryUnauthorized) GetPayload() *models.Errors {
	return o.Payload
}

func (o *GetSecuritySummaryUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetSecuritySummaryForbidden creates a GetSecuritySummaryForbidden with default headers values
func NewGetSecuritySummaryForbidden() *GetSecuritySummaryForbidden {
	return &GetSecuritySummaryForbidden{}
}

/*
GetSecuritySummaryForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type GetSecuritySummaryForbidden struct {

	/* The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *models.Errors
}

// IsSuccess returns true when this get security summary forbidden response has a 2xx status code
func (o *GetSecuritySummaryForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get security summary forbidden response has a 3xx status code
func (o *GetSecuritySummaryForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get security summary forbidden response has a 4xx status code
func (o *GetSecuritySummaryForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this get security summary forbidden response has a 5xx status code
func (o *GetSecuritySummaryForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this get security summary forbidden response a status code equal to that given
func (o *GetSecuritySummaryForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *GetSecuritySummaryForbidden) Error() string {
	return fmt.Sprintf("[GET /security/summary][%d] getSecuritySummaryForbidden  %+v", 403, o.Payload)
}

func (o *GetSecuritySummaryForbidden) String() string {
	return fmt.Sprintf("[GET /security/summary][%d] getSecuritySummaryForbidden  %+v", 403, o.Payload)
}

func (o *GetSecuritySummaryForbidden) GetPayload() *models.Errors {
	return o.Payload
}

func (o *GetSecuritySummaryForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetSecuritySummaryNotFound creates a GetSecuritySummaryNotFound with default headers values
func NewGetSecuritySummaryNotFound() *GetSecuritySummaryNotFound {
	return &GetSecuritySummaryNotFound{}
}

/*
GetSecuritySummaryNotFound describes a response with status code 404, with default header values.

Not found
*/
type GetSecuritySummaryNotFound struct {

	/* The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *models.Errors
}

// IsSuccess returns true when this get security summary not found response has a 2xx status code
func (o *GetSecuritySummaryNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get security summary not found response has a 3xx status code
func (o *GetSecuritySummaryNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get security summary not found response has a 4xx status code
func (o *GetSecuritySummaryNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get security summary not found response has a 5xx status code
func (o *GetSecuritySummaryNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get security summary not found response a status code equal to that given
func (o *GetSecuritySummaryNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *GetSecuritySummaryNotFound) Error() string {
	return fmt.Sprintf("[GET /security/summary][%d] getSecuritySummaryNotFound  %+v", 404, o.Payload)
}

func (o *GetSecuritySummaryNotFound) String() string {
	return fmt.Sprintf("[GET /security/summary][%d] getSecuritySummaryNotFound  %+v", 404, o.Payload)
}

func (o *GetSecuritySummaryNotFound) GetPayload() *models.Errors {
	return o.Payload
}

func (o *GetSecuritySummaryNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetSecuritySummaryInternalServerError creates a GetSecuritySummaryInternalServerError with default headers values
func NewGetSecuritySummaryInternalServerError() *GetSecuritySummaryInternalServerError {
	return &GetSecuritySummaryInternalServerError{}
}

/*
GetSecuritySummaryInternalServerError describes a response with status code 500, with default header values.

Internal server error
*/
type GetSecuritySummaryInternalServerError struct {

	/* The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *models.Errors
}

// IsSuccess returns true when this get security summary internal server error response has a 2xx status code
func (o *GetSecuritySummaryInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get security summary internal server error response has a 3xx status code
func (o *GetSecuritySummaryInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get security summary internal server error response has a 4xx status code
func (o *GetSecuritySummaryInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get security summary internal server error response has a 5xx status code
func (o *GetSecuritySummaryInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get security summary internal server error response a status code equal to that given
func (o *GetSecuritySummaryInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *GetSecuritySummaryInternalServerError) Error() string {
	return fmt.Sprintf("[GET /security/summary][%d] getSecuritySummaryInternalServerError  %+v", 500, o.Payload)
}

func (o *GetSecuritySummaryInternalServerError) String() string {
	return fmt.Sprintf("[GET /security/summary][%d] getSecuritySummaryInternalServerError  %+v", 500, o.Payload)
}

func (o *GetSecuritySummaryInternalServerError) GetPayload() *models.Errors {
	return o.Payload
}

func (o *GetSecuritySummaryInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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
