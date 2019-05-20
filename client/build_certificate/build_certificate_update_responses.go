// Code generated by go-swagger; DO NOT EDIT.

package build_certificate

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/bitrise-io/bitrise-api-client/models"
)

// BuildCertificateUpdateReader is a Reader for the BuildCertificateUpdate structure.
type BuildCertificateUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *BuildCertificateUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewBuildCertificateUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewBuildCertificateUpdateBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 401:
		result := NewBuildCertificateUpdateUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewBuildCertificateUpdateNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewBuildCertificateUpdateInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewBuildCertificateUpdateOK creates a BuildCertificateUpdateOK with default headers values
func NewBuildCertificateUpdateOK() *BuildCertificateUpdateOK {
	return &BuildCertificateUpdateOK{}
}

/*BuildCertificateUpdateOK handles this case with default header values.

OK
*/
type BuildCertificateUpdateOK struct {
	Payload *models.V0BuildCertificateResponseModel
}

func (o *BuildCertificateUpdateOK) Error() string {
	return fmt.Sprintf("[PATCH /apps/{app-slug}/build-certificates/{build-certificate-slug}][%d] buildCertificateUpdateOK  %+v", 200, o.Payload)
}

func (o *BuildCertificateUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.V0BuildCertificateResponseModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewBuildCertificateUpdateBadRequest creates a BuildCertificateUpdateBadRequest with default headers values
func NewBuildCertificateUpdateBadRequest() *BuildCertificateUpdateBadRequest {
	return &BuildCertificateUpdateBadRequest{}
}

/*BuildCertificateUpdateBadRequest handles this case with default header values.

Bad Request
*/
type BuildCertificateUpdateBadRequest struct {
	Payload *models.ServiceStandardErrorRespModel
}

func (o *BuildCertificateUpdateBadRequest) Error() string {
	return fmt.Sprintf("[PATCH /apps/{app-slug}/build-certificates/{build-certificate-slug}][%d] buildCertificateUpdateBadRequest  %+v", 400, o.Payload)
}

func (o *BuildCertificateUpdateBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ServiceStandardErrorRespModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewBuildCertificateUpdateUnauthorized creates a BuildCertificateUpdateUnauthorized with default headers values
func NewBuildCertificateUpdateUnauthorized() *BuildCertificateUpdateUnauthorized {
	return &BuildCertificateUpdateUnauthorized{}
}

/*BuildCertificateUpdateUnauthorized handles this case with default header values.

Unauthorized
*/
type BuildCertificateUpdateUnauthorized struct {
	Payload *models.ServiceStandardErrorRespModel
}

func (o *BuildCertificateUpdateUnauthorized) Error() string {
	return fmt.Sprintf("[PATCH /apps/{app-slug}/build-certificates/{build-certificate-slug}][%d] buildCertificateUpdateUnauthorized  %+v", 401, o.Payload)
}

func (o *BuildCertificateUpdateUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ServiceStandardErrorRespModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewBuildCertificateUpdateNotFound creates a BuildCertificateUpdateNotFound with default headers values
func NewBuildCertificateUpdateNotFound() *BuildCertificateUpdateNotFound {
	return &BuildCertificateUpdateNotFound{}
}

/*BuildCertificateUpdateNotFound handles this case with default header values.

Not Found
*/
type BuildCertificateUpdateNotFound struct {
	Payload *models.ServiceStandardErrorRespModel
}

func (o *BuildCertificateUpdateNotFound) Error() string {
	return fmt.Sprintf("[PATCH /apps/{app-slug}/build-certificates/{build-certificate-slug}][%d] buildCertificateUpdateNotFound  %+v", 404, o.Payload)
}

func (o *BuildCertificateUpdateNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ServiceStandardErrorRespModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewBuildCertificateUpdateInternalServerError creates a BuildCertificateUpdateInternalServerError with default headers values
func NewBuildCertificateUpdateInternalServerError() *BuildCertificateUpdateInternalServerError {
	return &BuildCertificateUpdateInternalServerError{}
}

/*BuildCertificateUpdateInternalServerError handles this case with default header values.

Internal Server Error
*/
type BuildCertificateUpdateInternalServerError struct {
	Payload *models.ServiceStandardErrorRespModel
}

func (o *BuildCertificateUpdateInternalServerError) Error() string {
	return fmt.Sprintf("[PATCH /apps/{app-slug}/build-certificates/{build-certificate-slug}][%d] buildCertificateUpdateInternalServerError  %+v", 500, o.Payload)
}

func (o *BuildCertificateUpdateInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ServiceStandardErrorRespModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
