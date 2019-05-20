// Code generated by go-swagger; DO NOT EDIT.

package avatar_candidate

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/bitrise-io/bitrise-api-client/models"
)

// AvatarCandidateCreateReader is a Reader for the AvatarCandidateCreate structure.
type AvatarCandidateCreateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AvatarCandidateCreateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 201:
		result := NewAvatarCandidateCreateCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewAvatarCandidateCreateBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 401:
		result := NewAvatarCandidateCreateUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewAvatarCandidateCreateNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewAvatarCandidateCreateInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewAvatarCandidateCreateCreated creates a AvatarCandidateCreateCreated with default headers values
func NewAvatarCandidateCreateCreated() *AvatarCandidateCreateCreated {
	return &AvatarCandidateCreateCreated{}
}

/*AvatarCandidateCreateCreated handles this case with default header values.

Created
*/
type AvatarCandidateCreateCreated struct {
	Payload models.V0AvatarCandidateCreateResponseItems
}

func (o *AvatarCandidateCreateCreated) Error() string {
	return fmt.Sprintf("[POST /apps/{app-slug}/avatar-candidates][%d] avatarCandidateCreateCreated  %+v", 201, o.Payload)
}

func (o *AvatarCandidateCreateCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAvatarCandidateCreateBadRequest creates a AvatarCandidateCreateBadRequest with default headers values
func NewAvatarCandidateCreateBadRequest() *AvatarCandidateCreateBadRequest {
	return &AvatarCandidateCreateBadRequest{}
}

/*AvatarCandidateCreateBadRequest handles this case with default header values.

Bad Request
*/
type AvatarCandidateCreateBadRequest struct {
	Payload *models.ServiceStandardErrorRespModel
}

func (o *AvatarCandidateCreateBadRequest) Error() string {
	return fmt.Sprintf("[POST /apps/{app-slug}/avatar-candidates][%d] avatarCandidateCreateBadRequest  %+v", 400, o.Payload)
}

func (o *AvatarCandidateCreateBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ServiceStandardErrorRespModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAvatarCandidateCreateUnauthorized creates a AvatarCandidateCreateUnauthorized with default headers values
func NewAvatarCandidateCreateUnauthorized() *AvatarCandidateCreateUnauthorized {
	return &AvatarCandidateCreateUnauthorized{}
}

/*AvatarCandidateCreateUnauthorized handles this case with default header values.

Unauthorized
*/
type AvatarCandidateCreateUnauthorized struct {
	Payload *models.ServiceStandardErrorRespModel
}

func (o *AvatarCandidateCreateUnauthorized) Error() string {
	return fmt.Sprintf("[POST /apps/{app-slug}/avatar-candidates][%d] avatarCandidateCreateUnauthorized  %+v", 401, o.Payload)
}

func (o *AvatarCandidateCreateUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ServiceStandardErrorRespModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAvatarCandidateCreateNotFound creates a AvatarCandidateCreateNotFound with default headers values
func NewAvatarCandidateCreateNotFound() *AvatarCandidateCreateNotFound {
	return &AvatarCandidateCreateNotFound{}
}

/*AvatarCandidateCreateNotFound handles this case with default header values.

Not Found
*/
type AvatarCandidateCreateNotFound struct {
	Payload *models.ServiceStandardErrorRespModel
}

func (o *AvatarCandidateCreateNotFound) Error() string {
	return fmt.Sprintf("[POST /apps/{app-slug}/avatar-candidates][%d] avatarCandidateCreateNotFound  %+v", 404, o.Payload)
}

func (o *AvatarCandidateCreateNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ServiceStandardErrorRespModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAvatarCandidateCreateInternalServerError creates a AvatarCandidateCreateInternalServerError with default headers values
func NewAvatarCandidateCreateInternalServerError() *AvatarCandidateCreateInternalServerError {
	return &AvatarCandidateCreateInternalServerError{}
}

/*AvatarCandidateCreateInternalServerError handles this case with default header values.

Internal Server Error
*/
type AvatarCandidateCreateInternalServerError struct {
	Payload *models.ServiceStandardErrorRespModel
}

func (o *AvatarCandidateCreateInternalServerError) Error() string {
	return fmt.Sprintf("[POST /apps/{app-slug}/avatar-candidates][%d] avatarCandidateCreateInternalServerError  %+v", 500, o.Payload)
}

func (o *AvatarCandidateCreateInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ServiceStandardErrorRespModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
