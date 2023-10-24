// Code generated by go-swagger; DO NOT EDIT.

package robot

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"

	"github.com/vmiko/go-client/pkg/sdk/v2.0/models"
)

// ListRobotReader is a Reader for the ListRobot structure.
type ListRobotReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListRobotReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewListRobotOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewListRobotBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewListRobotNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewListRobotInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewListRobotOK creates a ListRobotOK with default headers values
func NewListRobotOK() *ListRobotOK {
	return &ListRobotOK{}
}

/*
ListRobotOK describes a response with status code 200, with default header values.

Success
*/
type ListRobotOK struct {

	/* Link refers to the previous page and next page
	 */
	Link string

	/* The total count of robot accounts
	 */
	XTotalCount int64

	Payload []*models.Robot
}

// IsSuccess returns true when this list robot o k response has a 2xx status code
func (o *ListRobotOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this list robot o k response has a 3xx status code
func (o *ListRobotOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this list robot o k response has a 4xx status code
func (o *ListRobotOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this list robot o k response has a 5xx status code
func (o *ListRobotOK) IsServerError() bool {
	return false
}

// IsCode returns true when this list robot o k response a status code equal to that given
func (o *ListRobotOK) IsCode(code int) bool {
	return code == 200
}

func (o *ListRobotOK) Error() string {
	return fmt.Sprintf("[GET /robots][%d] listRobotOK  %+v", 200, o.Payload)
}

func (o *ListRobotOK) String() string {
	return fmt.Sprintf("[GET /robots][%d] listRobotOK  %+v", 200, o.Payload)
}

func (o *ListRobotOK) GetPayload() []*models.Robot {
	return o.Payload
}

func (o *ListRobotOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header Link
	hdrLink := response.GetHeader("Link")

	if hdrLink != "" {
		o.Link = hdrLink
	}

	// hydrates response header X-Total-Count
	hdrXTotalCount := response.GetHeader("X-Total-Count")

	if hdrXTotalCount != "" {
		valxTotalCount, err := swag.ConvertInt64(hdrXTotalCount)
		if err != nil {
			return errors.InvalidType("X-Total-Count", "header", "int64", hdrXTotalCount)
		}
		o.XTotalCount = valxTotalCount
	}

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListRobotBadRequest creates a ListRobotBadRequest with default headers values
func NewListRobotBadRequest() *ListRobotBadRequest {
	return &ListRobotBadRequest{}
}

/*
ListRobotBadRequest describes a response with status code 400, with default header values.

Bad request
*/
type ListRobotBadRequest struct {

	/* The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *models.Errors
}

// IsSuccess returns true when this list robot bad request response has a 2xx status code
func (o *ListRobotBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this list robot bad request response has a 3xx status code
func (o *ListRobotBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this list robot bad request response has a 4xx status code
func (o *ListRobotBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this list robot bad request response has a 5xx status code
func (o *ListRobotBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this list robot bad request response a status code equal to that given
func (o *ListRobotBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *ListRobotBadRequest) Error() string {
	return fmt.Sprintf("[GET /robots][%d] listRobotBadRequest  %+v", 400, o.Payload)
}

func (o *ListRobotBadRequest) String() string {
	return fmt.Sprintf("[GET /robots][%d] listRobotBadRequest  %+v", 400, o.Payload)
}

func (o *ListRobotBadRequest) GetPayload() *models.Errors {
	return o.Payload
}

func (o *ListRobotBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewListRobotNotFound creates a ListRobotNotFound with default headers values
func NewListRobotNotFound() *ListRobotNotFound {
	return &ListRobotNotFound{}
}

/*
ListRobotNotFound describes a response with status code 404, with default header values.

Not found
*/
type ListRobotNotFound struct {

	/* The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *models.Errors
}

// IsSuccess returns true when this list robot not found response has a 2xx status code
func (o *ListRobotNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this list robot not found response has a 3xx status code
func (o *ListRobotNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this list robot not found response has a 4xx status code
func (o *ListRobotNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this list robot not found response has a 5xx status code
func (o *ListRobotNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this list robot not found response a status code equal to that given
func (o *ListRobotNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *ListRobotNotFound) Error() string {
	return fmt.Sprintf("[GET /robots][%d] listRobotNotFound  %+v", 404, o.Payload)
}

func (o *ListRobotNotFound) String() string {
	return fmt.Sprintf("[GET /robots][%d] listRobotNotFound  %+v", 404, o.Payload)
}

func (o *ListRobotNotFound) GetPayload() *models.Errors {
	return o.Payload
}

func (o *ListRobotNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewListRobotInternalServerError creates a ListRobotInternalServerError with default headers values
func NewListRobotInternalServerError() *ListRobotInternalServerError {
	return &ListRobotInternalServerError{}
}

/*
ListRobotInternalServerError describes a response with status code 500, with default header values.

Internal server error
*/
type ListRobotInternalServerError struct {

	/* The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *models.Errors
}

// IsSuccess returns true when this list robot internal server error response has a 2xx status code
func (o *ListRobotInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this list robot internal server error response has a 3xx status code
func (o *ListRobotInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this list robot internal server error response has a 4xx status code
func (o *ListRobotInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this list robot internal server error response has a 5xx status code
func (o *ListRobotInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this list robot internal server error response a status code equal to that given
func (o *ListRobotInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *ListRobotInternalServerError) Error() string {
	return fmt.Sprintf("[GET /robots][%d] listRobotInternalServerError  %+v", 500, o.Payload)
}

func (o *ListRobotInternalServerError) String() string {
	return fmt.Sprintf("[GET /robots][%d] listRobotInternalServerError  %+v", 500, o.Payload)
}

func (o *ListRobotInternalServerError) GetPayload() *models.Errors {
	return o.Payload
}

func (o *ListRobotInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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
