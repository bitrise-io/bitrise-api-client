// Code generated by go-swagger; DO NOT EDIT.

package provisioning_profile

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/bitrise-io/bitrise-api-client/models"
)

// ProvisioningProfileUpdateReader is a Reader for the ProvisioningProfileUpdate structure.
type ProvisioningProfileUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ProvisioningProfileUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewProvisioningProfileUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewProvisioningProfileUpdateBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewProvisioningProfileUpdateUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewProvisioningProfileUpdateNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewProvisioningProfileUpdateInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewProvisioningProfileUpdateOK creates a ProvisioningProfileUpdateOK with default headers values
func NewProvisioningProfileUpdateOK() *ProvisioningProfileUpdateOK {
	return &ProvisioningProfileUpdateOK{}
}

/*ProvisioningProfileUpdateOK handles this case with default header values.

OK
*/
type ProvisioningProfileUpdateOK struct {
	Payload *models.V0ProvisionProfileResponseModel
}

func (o *ProvisioningProfileUpdateOK) Error() string {
	return fmt.Sprintf("[PATCH /apps/{app-slug}/provisioning-profiles/{provisioning-profile-slug}][%d] provisioningProfileUpdateOK  %+v", 200, o.Payload)
}

func (o *ProvisioningProfileUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.V0ProvisionProfileResponseModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewProvisioningProfileUpdateBadRequest creates a ProvisioningProfileUpdateBadRequest with default headers values
func NewProvisioningProfileUpdateBadRequest() *ProvisioningProfileUpdateBadRequest {
	return &ProvisioningProfileUpdateBadRequest{}
}

/*ProvisioningProfileUpdateBadRequest handles this case with default header values.

Bad Request
*/
type ProvisioningProfileUpdateBadRequest struct {
	Payload *models.ServiceStandardErrorRespModel
}

func (o *ProvisioningProfileUpdateBadRequest) Error() string {
	return fmt.Sprintf("[PATCH /apps/{app-slug}/provisioning-profiles/{provisioning-profile-slug}][%d] provisioningProfileUpdateBadRequest  %+v", 400, o.Payload)
}

func (o *ProvisioningProfileUpdateBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ServiceStandardErrorRespModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewProvisioningProfileUpdateUnauthorized creates a ProvisioningProfileUpdateUnauthorized with default headers values
func NewProvisioningProfileUpdateUnauthorized() *ProvisioningProfileUpdateUnauthorized {
	return &ProvisioningProfileUpdateUnauthorized{}
}

/*ProvisioningProfileUpdateUnauthorized handles this case with default header values.

Unauthorized
*/
type ProvisioningProfileUpdateUnauthorized struct {
	Payload *models.ServiceStandardErrorRespModel
}

func (o *ProvisioningProfileUpdateUnauthorized) Error() string {
	return fmt.Sprintf("[PATCH /apps/{app-slug}/provisioning-profiles/{provisioning-profile-slug}][%d] provisioningProfileUpdateUnauthorized  %+v", 401, o.Payload)
}

func (o *ProvisioningProfileUpdateUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ServiceStandardErrorRespModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewProvisioningProfileUpdateNotFound creates a ProvisioningProfileUpdateNotFound with default headers values
func NewProvisioningProfileUpdateNotFound() *ProvisioningProfileUpdateNotFound {
	return &ProvisioningProfileUpdateNotFound{}
}

/*ProvisioningProfileUpdateNotFound handles this case with default header values.

Not Found
*/
type ProvisioningProfileUpdateNotFound struct {
	Payload *models.ServiceStandardErrorRespModel
}

func (o *ProvisioningProfileUpdateNotFound) Error() string {
	return fmt.Sprintf("[PATCH /apps/{app-slug}/provisioning-profiles/{provisioning-profile-slug}][%d] provisioningProfileUpdateNotFound  %+v", 404, o.Payload)
}

func (o *ProvisioningProfileUpdateNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ServiceStandardErrorRespModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewProvisioningProfileUpdateInternalServerError creates a ProvisioningProfileUpdateInternalServerError with default headers values
func NewProvisioningProfileUpdateInternalServerError() *ProvisioningProfileUpdateInternalServerError {
	return &ProvisioningProfileUpdateInternalServerError{}
}

/*ProvisioningProfileUpdateInternalServerError handles this case with default header values.

Internal Server Error
*/
type ProvisioningProfileUpdateInternalServerError struct {
	Payload *models.ServiceStandardErrorRespModel
}

func (o *ProvisioningProfileUpdateInternalServerError) Error() string {
	return fmt.Sprintf("[PATCH /apps/{app-slug}/provisioning-profiles/{provisioning-profile-slug}][%d] provisioningProfileUpdateInternalServerError  %+v", 500, o.Payload)
}

func (o *ProvisioningProfileUpdateInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ServiceStandardErrorRespModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
