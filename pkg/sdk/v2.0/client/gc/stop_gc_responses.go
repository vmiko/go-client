// Code generated by go-swagger; DO NOT EDIT.

package gc

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/vmiko/go-client/pkg/sdk/v2.0/models"
)

// StopGCReader is a Reader for the StopGC structure.
type StopGCReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *StopGCReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewStopGCOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewStopGCUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewStopGCForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewStopGCNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewStopGCInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewStopGCOK creates a StopGCOK with default headers values
func NewStopGCOK() *StopGCOK {
	return &StopGCOK{}
}

/*
StopGCOK describes a response with status code 200, with default header values.

Success
*/
type StopGCOK struct {

	/* The ID of the corresponding request for the response
	 */
	XRequestID string
}

// IsSuccess returns true when this stop Gc o k response has a 2xx status code
func (o *StopGCOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this stop Gc o k response has a 3xx status code
func (o *StopGCOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this stop Gc o k response has a 4xx status code
func (o *StopGCOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this stop Gc o k response has a 5xx status code
func (o *StopGCOK) IsServerError() bool {
	return false
}

// IsCode returns true when this stop Gc o k response a status code equal to that given
func (o *StopGCOK) IsCode(code int) bool {
	return code == 200
}

func (o *StopGCOK) Error() string {
	return fmt.Sprintf("[PUT /system/gc/{gc_id}][%d] stopGcOK ", 200)
}

func (o *StopGCOK) String() string {
	return fmt.Sprintf("[PUT /system/gc/{gc_id}][%d] stopGcOK ", 200)
}

func (o *StopGCOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header X-Request-Id
	hdrXRequestID := response.GetHeader("X-Request-Id")

	if hdrXRequestID != "" {
		o.XRequestID = hdrXRequestID
	}

	return nil
}

// NewStopGCUnauthorized creates a StopGCUnauthorized with default headers values
func NewStopGCUnauthorized() *StopGCUnauthorized {
	return &StopGCUnauthorized{}
}

/*
StopGCUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type StopGCUnauthorized struct {

	/* The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *models.Errors
}

// IsSuccess returns true when this stop Gc unauthorized response has a 2xx status code
func (o *StopGCUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this stop Gc unauthorized response has a 3xx status code
func (o *StopGCUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this stop Gc unauthorized response has a 4xx status code
func (o *StopGCUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this stop Gc unauthorized response has a 5xx status code
func (o *StopGCUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this stop Gc unauthorized response a status code equal to that given
func (o *StopGCUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *StopGCUnauthorized) Error() string {
	return fmt.Sprintf("[PUT /system/gc/{gc_id}][%d] stopGcUnauthorized  %+v", 401, o.Payload)
}

func (o *StopGCUnauthorized) String() string {
	return fmt.Sprintf("[PUT /system/gc/{gc_id}][%d] stopGcUnauthorized  %+v", 401, o.Payload)
}

func (o *StopGCUnauthorized) GetPayload() *models.Errors {
	return o.Payload
}

func (o *StopGCUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewStopGCForbidden creates a StopGCForbidden with default headers values
func NewStopGCForbidden() *StopGCForbidden {
	return &StopGCForbidden{}
}

/*
StopGCForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type StopGCForbidden struct {

	/* The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *models.Errors
}

// IsSuccess returns true when this stop Gc forbidden response has a 2xx status code
func (o *StopGCForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this stop Gc forbidden response has a 3xx status code
func (o *StopGCForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this stop Gc forbidden response has a 4xx status code
func (o *StopGCForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this stop Gc forbidden response has a 5xx status code
func (o *StopGCForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this stop Gc forbidden response a status code equal to that given
func (o *StopGCForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *StopGCForbidden) Error() string {
	return fmt.Sprintf("[PUT /system/gc/{gc_id}][%d] stopGcForbidden  %+v", 403, o.Payload)
}

func (o *StopGCForbidden) String() string {
	return fmt.Sprintf("[PUT /system/gc/{gc_id}][%d] stopGcForbidden  %+v", 403, o.Payload)
}

func (o *StopGCForbidden) GetPayload() *models.Errors {
	return o.Payload
}

func (o *StopGCForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewStopGCNotFound creates a StopGCNotFound with default headers values
func NewStopGCNotFound() *StopGCNotFound {
	return &StopGCNotFound{}
}

/*
StopGCNotFound describes a response with status code 404, with default header values.

Not found
*/
type StopGCNotFound struct {

	/* The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *models.Errors
}

// IsSuccess returns true when this stop Gc not found response has a 2xx status code
func (o *StopGCNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this stop Gc not found response has a 3xx status code
func (o *StopGCNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this stop Gc not found response has a 4xx status code
func (o *StopGCNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this stop Gc not found response has a 5xx status code
func (o *StopGCNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this stop Gc not found response a status code equal to that given
func (o *StopGCNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *StopGCNotFound) Error() string {
	return fmt.Sprintf("[PUT /system/gc/{gc_id}][%d] stopGcNotFound  %+v", 404, o.Payload)
}

func (o *StopGCNotFound) String() string {
	return fmt.Sprintf("[PUT /system/gc/{gc_id}][%d] stopGcNotFound  %+v", 404, o.Payload)
}

func (o *StopGCNotFound) GetPayload() *models.Errors {
	return o.Payload
}

func (o *StopGCNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewStopGCInternalServerError creates a StopGCInternalServerError with default headers values
func NewStopGCInternalServerError() *StopGCInternalServerError {
	return &StopGCInternalServerError{}
}

/*
StopGCInternalServerError describes a response with status code 500, with default header values.

Internal server error
*/
type StopGCInternalServerError struct {

	/* The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *models.Errors
}

// IsSuccess returns true when this stop Gc internal server error response has a 2xx status code
func (o *StopGCInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this stop Gc internal server error response has a 3xx status code
func (o *StopGCInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this stop Gc internal server error response has a 4xx status code
func (o *StopGCInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this stop Gc internal server error response has a 5xx status code
func (o *StopGCInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this stop Gc internal server error response a status code equal to that given
func (o *StopGCInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *StopGCInternalServerError) Error() string {
	return fmt.Sprintf("[PUT /system/gc/{gc_id}][%d] stopGcInternalServerError  %+v", 500, o.Payload)
}

func (o *StopGCInternalServerError) String() string {
	return fmt.Sprintf("[PUT /system/gc/{gc_id}][%d] stopGcInternalServerError  %+v", 500, o.Payload)
}

func (o *StopGCInternalServerError) GetPayload() *models.Errors {
	return o.Payload
}

func (o *StopGCInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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
