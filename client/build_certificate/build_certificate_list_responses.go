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

// BuildCertificateListReader is a Reader for the BuildCertificateList structure.
type BuildCertificateListReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *BuildCertificateListReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewBuildCertificateListOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewBuildCertificateListBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 401:
		result := NewBuildCertificateListUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewBuildCertificateListNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewBuildCertificateListInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewBuildCertificateListOK creates a BuildCertificateListOK with default headers values
func NewBuildCertificateListOK() *BuildCertificateListOK {
	return &BuildCertificateListOK{}
}

/*BuildCertificateListOK handles this case with default header values.

OK
*/
type BuildCertificateListOK struct {
	Payload *models.V0BuildCertificateListResponseModel
}

func (o *BuildCertificateListOK) Error() string {
	return fmt.Sprintf("[GET /apps/{app-slug}/build-certificates][%d] buildCertificateListOK  %+v", 200, o.Payload)
}

func (o *BuildCertificateListOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.V0BuildCertificateListResponseModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewBuildCertificateListBadRequest creates a BuildCertificateListBadRequest with default headers values
func NewBuildCertificateListBadRequest() *BuildCertificateListBadRequest {
	return &BuildCertificateListBadRequest{}
}

/*BuildCertificateListBadRequest handles this case with default header values.

Bad Request
*/
type BuildCertificateListBadRequest struct {
	Payload *models.ServiceStandardErrorRespModel
}

func (o *BuildCertificateListBadRequest) Error() string {
	return fmt.Sprintf("[GET /apps/{app-slug}/build-certificates][%d] buildCertificateListBadRequest  %+v", 400, o.Payload)
}

func (o *BuildCertificateListBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ServiceStandardErrorRespModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewBuildCertificateListUnauthorized creates a BuildCertificateListUnauthorized with default headers values
func NewBuildCertificateListUnauthorized() *BuildCertificateListUnauthorized {
	return &BuildCertificateListUnauthorized{}
}

/*BuildCertificateListUnauthorized handles this case with default header values.

Unauthorized
*/
type BuildCertificateListUnauthorized struct {
	Payload *models.ServiceStandardErrorRespModel
}

func (o *BuildCertificateListUnauthorized) Error() string {
	return fmt.Sprintf("[GET /apps/{app-slug}/build-certificates][%d] buildCertificateListUnauthorized  %+v", 401, o.Payload)
}

func (o *BuildCertificateListUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ServiceStandardErrorRespModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewBuildCertificateListNotFound creates a BuildCertificateListNotFound with default headers values
func NewBuildCertificateListNotFound() *BuildCertificateListNotFound {
	return &BuildCertificateListNotFound{}
}

/*BuildCertificateListNotFound handles this case with default header values.

Not Found
*/
type BuildCertificateListNotFound struct {
	Payload *models.ServiceStandardErrorRespModel
}

func (o *BuildCertificateListNotFound) Error() string {
	return fmt.Sprintf("[GET /apps/{app-slug}/build-certificates][%d] buildCertificateListNotFound  %+v", 404, o.Payload)
}

func (o *BuildCertificateListNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ServiceStandardErrorRespModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewBuildCertificateListInternalServerError creates a BuildCertificateListInternalServerError with default headers values
func NewBuildCertificateListInternalServerError() *BuildCertificateListInternalServerError {
	return &BuildCertificateListInternalServerError{}
}

/*BuildCertificateListInternalServerError handles this case with default header values.

Internal Server Error
*/
type BuildCertificateListInternalServerError struct {
	Payload *models.ServiceStandardErrorRespModel
}

func (o *BuildCertificateListInternalServerError) Error() string {
	return fmt.Sprintf("[GET /apps/{app-slug}/build-certificates][%d] buildCertificateListInternalServerError  %+v", 500, o.Payload)
}

func (o *BuildCertificateListInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ServiceStandardErrorRespModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
