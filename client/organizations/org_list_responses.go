// Code generated by go-swagger; DO NOT EDIT.

package organizations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/bitrise-io/bitrise-api-client/models"
)

// OrgListReader is a Reader for the OrgList structure.
type OrgListReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *OrgListReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewOrgListOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewOrgListBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewOrgListUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewOrgListNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewOrgListInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewOrgListOK creates a OrgListOK with default headers values
func NewOrgListOK() *OrgListOK {
	return &OrgListOK{}
}

/*OrgListOK handles this case with default header values.

OK
*/
type OrgListOK struct {
	Payload *models.V0OrganizationListRespModel
}

func (o *OrgListOK) Error() string {
	return fmt.Sprintf("[GET /organizations][%d] orgListOK  %+v", 200, o.Payload)
}

func (o *OrgListOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.V0OrganizationListRespModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewOrgListBadRequest creates a OrgListBadRequest with default headers values
func NewOrgListBadRequest() *OrgListBadRequest {
	return &OrgListBadRequest{}
}

/*OrgListBadRequest handles this case with default header values.

Bad Request
*/
type OrgListBadRequest struct {
	Payload *models.ServiceStandardErrorRespModel
}

func (o *OrgListBadRequest) Error() string {
	return fmt.Sprintf("[GET /organizations][%d] orgListBadRequest  %+v", 400, o.Payload)
}

func (o *OrgListBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ServiceStandardErrorRespModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewOrgListUnauthorized creates a OrgListUnauthorized with default headers values
func NewOrgListUnauthorized() *OrgListUnauthorized {
	return &OrgListUnauthorized{}
}

/*OrgListUnauthorized handles this case with default header values.

Unauthorized
*/
type OrgListUnauthorized struct {
	Payload *models.ServiceStandardErrorRespModel
}

func (o *OrgListUnauthorized) Error() string {
	return fmt.Sprintf("[GET /organizations][%d] orgListUnauthorized  %+v", 401, o.Payload)
}

func (o *OrgListUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ServiceStandardErrorRespModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewOrgListNotFound creates a OrgListNotFound with default headers values
func NewOrgListNotFound() *OrgListNotFound {
	return &OrgListNotFound{}
}

/*OrgListNotFound handles this case with default header values.

Not Found
*/
type OrgListNotFound struct {
	Payload *models.ServiceStandardErrorRespModel
}

func (o *OrgListNotFound) Error() string {
	return fmt.Sprintf("[GET /organizations][%d] orgListNotFound  %+v", 404, o.Payload)
}

func (o *OrgListNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ServiceStandardErrorRespModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewOrgListInternalServerError creates a OrgListInternalServerError with default headers values
func NewOrgListInternalServerError() *OrgListInternalServerError {
	return &OrgListInternalServerError{}
}

/*OrgListInternalServerError handles this case with default header values.

Internal Server Error
*/
type OrgListInternalServerError struct {
	Payload *models.ServiceStandardErrorRespModel
}

func (o *OrgListInternalServerError) Error() string {
	return fmt.Sprintf("[GET /organizations][%d] orgListInternalServerError  %+v", 500, o.Payload)
}

func (o *OrgListInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ServiceStandardErrorRespModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
