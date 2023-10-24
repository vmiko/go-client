// Code generated by go-swagger; DO NOT EDIT.

package user

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/vmiko/go-client/pkg/sdk/v2.0/models"
)

// UpdateUserPasswordReader is a Reader for the UpdateUserPassword structure.
type UpdateUserPasswordReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateUserPasswordReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUpdateUserPasswordOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewUpdateUserPasswordBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewUpdateUserPasswordUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewUpdateUserPasswordForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewUpdateUserPasswordInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewUpdateUserPasswordOK creates a UpdateUserPasswordOK with default headers values
func NewUpdateUserPasswordOK() *UpdateUserPasswordOK {
	return &UpdateUserPasswordOK{}
}

/*
UpdateUserPasswordOK describes a response with status code 200, with default header values.

Success
*/
type UpdateUserPasswordOK struct {

	/* The ID of the corresponding request for the response
	 */
	XRequestID string
}

// IsSuccess returns true when this update user password o k response has a 2xx status code
func (o *UpdateUserPasswordOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this update user password o k response has a 3xx status code
func (o *UpdateUserPasswordOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update user password o k response has a 4xx status code
func (o *UpdateUserPasswordOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this update user password o k response has a 5xx status code
func (o *UpdateUserPasswordOK) IsServerError() bool {
	return false
}

// IsCode returns true when this update user password o k response a status code equal to that given
func (o *UpdateUserPasswordOK) IsCode(code int) bool {
	return code == 200
}

func (o *UpdateUserPasswordOK) Error() string {
	return fmt.Sprintf("[PUT /users/{user_id}/password][%d] updateUserPasswordOK ", 200)
}

func (o *UpdateUserPasswordOK) String() string {
	return fmt.Sprintf("[PUT /users/{user_id}/password][%d] updateUserPasswordOK ", 200)
}

func (o *UpdateUserPasswordOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header X-Request-Id
	hdrXRequestID := response.GetHeader("X-Request-Id")

	if hdrXRequestID != "" {
		o.XRequestID = hdrXRequestID
	}

	return nil
}

// NewUpdateUserPasswordBadRequest creates a UpdateUserPasswordBadRequest with default headers values
func NewUpdateUserPasswordBadRequest() *UpdateUserPasswordBadRequest {
	return &UpdateUserPasswordBadRequest{}
}

/*
UpdateUserPasswordBadRequest describes a response with status code 400, with default header values.

Invalid user ID; Password does not meet requirement
*/
type UpdateUserPasswordBadRequest struct {
}

// IsSuccess returns true when this update user password bad request response has a 2xx status code
func (o *UpdateUserPasswordBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update user password bad request response has a 3xx status code
func (o *UpdateUserPasswordBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update user password bad request response has a 4xx status code
func (o *UpdateUserPasswordBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this update user password bad request response has a 5xx status code
func (o *UpdateUserPasswordBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this update user password bad request response a status code equal to that given
func (o *UpdateUserPasswordBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *UpdateUserPasswordBadRequest) Error() string {
	return fmt.Sprintf("[PUT /users/{user_id}/password][%d] updateUserPasswordBadRequest ", 400)
}

func (o *UpdateUserPasswordBadRequest) String() string {
	return fmt.Sprintf("[PUT /users/{user_id}/password][%d] updateUserPasswordBadRequest ", 400)
}

func (o *UpdateUserPasswordBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUpdateUserPasswordUnauthorized creates a UpdateUserPasswordUnauthorized with default headers values
func NewUpdateUserPasswordUnauthorized() *UpdateUserPasswordUnauthorized {
	return &UpdateUserPasswordUnauthorized{}
}

/*
UpdateUserPasswordUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type UpdateUserPasswordUnauthorized struct {

	/* The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *models.Errors
}

// IsSuccess returns true when this update user password unauthorized response has a 2xx status code
func (o *UpdateUserPasswordUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update user password unauthorized response has a 3xx status code
func (o *UpdateUserPasswordUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update user password unauthorized response has a 4xx status code
func (o *UpdateUserPasswordUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this update user password unauthorized response has a 5xx status code
func (o *UpdateUserPasswordUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this update user password unauthorized response a status code equal to that given
func (o *UpdateUserPasswordUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *UpdateUserPasswordUnauthorized) Error() string {
	return fmt.Sprintf("[PUT /users/{user_id}/password][%d] updateUserPasswordUnauthorized  %+v", 401, o.Payload)
}

func (o *UpdateUserPasswordUnauthorized) String() string {
	return fmt.Sprintf("[PUT /users/{user_id}/password][%d] updateUserPasswordUnauthorized  %+v", 401, o.Payload)
}

func (o *UpdateUserPasswordUnauthorized) GetPayload() *models.Errors {
	return o.Payload
}

func (o *UpdateUserPasswordUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewUpdateUserPasswordForbidden creates a UpdateUserPasswordForbidden with default headers values
func NewUpdateUserPasswordForbidden() *UpdateUserPasswordForbidden {
	return &UpdateUserPasswordForbidden{}
}

/*
UpdateUserPasswordForbidden describes a response with status code 403, with default header values.

The caller does not have permission to update the password of the user with given ID, or the old password in request body is not correct.
*/
type UpdateUserPasswordForbidden struct {
}

// IsSuccess returns true when this update user password forbidden response has a 2xx status code
func (o *UpdateUserPasswordForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update user password forbidden response has a 3xx status code
func (o *UpdateUserPasswordForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update user password forbidden response has a 4xx status code
func (o *UpdateUserPasswordForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this update user password forbidden response has a 5xx status code
func (o *UpdateUserPasswordForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this update user password forbidden response a status code equal to that given
func (o *UpdateUserPasswordForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *UpdateUserPasswordForbidden) Error() string {
	return fmt.Sprintf("[PUT /users/{user_id}/password][%d] updateUserPasswordForbidden ", 403)
}

func (o *UpdateUserPasswordForbidden) String() string {
	return fmt.Sprintf("[PUT /users/{user_id}/password][%d] updateUserPasswordForbidden ", 403)
}

func (o *UpdateUserPasswordForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUpdateUserPasswordInternalServerError creates a UpdateUserPasswordInternalServerError with default headers values
func NewUpdateUserPasswordInternalServerError() *UpdateUserPasswordInternalServerError {
	return &UpdateUserPasswordInternalServerError{}
}

/*
UpdateUserPasswordInternalServerError describes a response with status code 500, with default header values.

Internal server error
*/
type UpdateUserPasswordInternalServerError struct {

	/* The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *models.Errors
}

// IsSuccess returns true when this update user password internal server error response has a 2xx status code
func (o *UpdateUserPasswordInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update user password internal server error response has a 3xx status code
func (o *UpdateUserPasswordInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update user password internal server error response has a 4xx status code
func (o *UpdateUserPasswordInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this update user password internal server error response has a 5xx status code
func (o *UpdateUserPasswordInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this update user password internal server error response a status code equal to that given
func (o *UpdateUserPasswordInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *UpdateUserPasswordInternalServerError) Error() string {
	return fmt.Sprintf("[PUT /users/{user_id}/password][%d] updateUserPasswordInternalServerError  %+v", 500, o.Payload)
}

func (o *UpdateUserPasswordInternalServerError) String() string {
	return fmt.Sprintf("[PUT /users/{user_id}/password][%d] updateUserPasswordInternalServerError  %+v", 500, o.Payload)
}

func (o *UpdateUserPasswordInternalServerError) GetPayload() *models.Errors {
	return o.Payload
}

func (o *UpdateUserPasswordInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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
