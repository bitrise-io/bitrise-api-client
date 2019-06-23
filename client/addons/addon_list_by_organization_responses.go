// Code generated by go-swagger; DO NOT EDIT.

package addons

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/bitrise-io/bitrise-api-client/models"
)

// AddonListByOrganizationReader is a Reader for the AddonListByOrganization structure.
type AddonListByOrganizationReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AddonListByOrganizationReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewAddonListByOrganizationOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewAddonListByOrganizationBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewAddonListByOrganizationUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewAddonListByOrganizationNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewAddonListByOrganizationInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewAddonListByOrganizationOK creates a AddonListByOrganizationOK with default headers values
func NewAddonListByOrganizationOK() *AddonListByOrganizationOK {
	return &AddonListByOrganizationOK{}
}

/*AddonListByOrganizationOK handles this case with default header values.

OK
*/
type AddonListByOrganizationOK struct {
	Payload *models.V0OwnerAddOnsListResponseModel
}

func (o *AddonListByOrganizationOK) Error() string {
	return fmt.Sprintf("[GET /organizations/{organization-slug}/addons][%d] addonListByOrganizationOK  %+v", 200, o.Payload)
}

func (o *AddonListByOrganizationOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.V0OwnerAddOnsListResponseModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAddonListByOrganizationBadRequest creates a AddonListByOrganizationBadRequest with default headers values
func NewAddonListByOrganizationBadRequest() *AddonListByOrganizationBadRequest {
	return &AddonListByOrganizationBadRequest{}
}

/*AddonListByOrganizationBadRequest handles this case with default header values.

Bad Request
*/
type AddonListByOrganizationBadRequest struct {
	Payload *models.ServiceStandardErrorRespModel
}

func (o *AddonListByOrganizationBadRequest) Error() string {
	return fmt.Sprintf("[GET /organizations/{organization-slug}/addons][%d] addonListByOrganizationBadRequest  %+v", 400, o.Payload)
}

func (o *AddonListByOrganizationBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ServiceStandardErrorRespModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAddonListByOrganizationUnauthorized creates a AddonListByOrganizationUnauthorized with default headers values
func NewAddonListByOrganizationUnauthorized() *AddonListByOrganizationUnauthorized {
	return &AddonListByOrganizationUnauthorized{}
}

/*AddonListByOrganizationUnauthorized handles this case with default header values.

Unauthorized
*/
type AddonListByOrganizationUnauthorized struct {
	Payload *models.ServiceStandardErrorRespModel
}

func (o *AddonListByOrganizationUnauthorized) Error() string {
	return fmt.Sprintf("[GET /organizations/{organization-slug}/addons][%d] addonListByOrganizationUnauthorized  %+v", 401, o.Payload)
}

func (o *AddonListByOrganizationUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ServiceStandardErrorRespModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAddonListByOrganizationNotFound creates a AddonListByOrganizationNotFound with default headers values
func NewAddonListByOrganizationNotFound() *AddonListByOrganizationNotFound {
	return &AddonListByOrganizationNotFound{}
}

/*AddonListByOrganizationNotFound handles this case with default header values.

Not Found
*/
type AddonListByOrganizationNotFound struct {
	Payload *models.ServiceStandardErrorRespModel
}

func (o *AddonListByOrganizationNotFound) Error() string {
	return fmt.Sprintf("[GET /organizations/{organization-slug}/addons][%d] addonListByOrganizationNotFound  %+v", 404, o.Payload)
}

func (o *AddonListByOrganizationNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ServiceStandardErrorRespModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAddonListByOrganizationInternalServerError creates a AddonListByOrganizationInternalServerError with default headers values
func NewAddonListByOrganizationInternalServerError() *AddonListByOrganizationInternalServerError {
	return &AddonListByOrganizationInternalServerError{}
}

/*AddonListByOrganizationInternalServerError handles this case with default header values.

Internal Server Error
*/
type AddonListByOrganizationInternalServerError struct {
	Payload *models.ServiceStandardErrorRespModel
}

func (o *AddonListByOrganizationInternalServerError) Error() string {
	return fmt.Sprintf("[GET /organizations/{organization-slug}/addons][%d] addonListByOrganizationInternalServerError  %+v", 500, o.Payload)
}

func (o *AddonListByOrganizationInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ServiceStandardErrorRespModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
