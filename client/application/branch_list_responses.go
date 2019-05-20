// Code generated by go-swagger; DO NOT EDIT.

package application

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/bitrise-io/bitrise-api-client/models"
)

// BranchListReader is a Reader for the BranchList structure.
type BranchListReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *BranchListReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewBranchListOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewBranchListBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 401:
		result := NewBranchListUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewBranchListNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewBranchListInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewBranchListOK creates a BranchListOK with default headers values
func NewBranchListOK() *BranchListOK {
	return &BranchListOK{}
}

/*BranchListOK handles this case with default header values.

OK
*/
type BranchListOK struct {
	Payload *models.V0BranchListResponseModel
}

func (o *BranchListOK) Error() string {
	return fmt.Sprintf("[GET /apps/{app-slug}/branches][%d] branchListOK  %+v", 200, o.Payload)
}

func (o *BranchListOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.V0BranchListResponseModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewBranchListBadRequest creates a BranchListBadRequest with default headers values
func NewBranchListBadRequest() *BranchListBadRequest {
	return &BranchListBadRequest{}
}

/*BranchListBadRequest handles this case with default header values.

Bad Request
*/
type BranchListBadRequest struct {
	Payload *models.ServiceStandardErrorRespModel
}

func (o *BranchListBadRequest) Error() string {
	return fmt.Sprintf("[GET /apps/{app-slug}/branches][%d] branchListBadRequest  %+v", 400, o.Payload)
}

func (o *BranchListBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ServiceStandardErrorRespModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewBranchListUnauthorized creates a BranchListUnauthorized with default headers values
func NewBranchListUnauthorized() *BranchListUnauthorized {
	return &BranchListUnauthorized{}
}

/*BranchListUnauthorized handles this case with default header values.

Unauthorized
*/
type BranchListUnauthorized struct {
	Payload *models.ServiceStandardErrorRespModel
}

func (o *BranchListUnauthorized) Error() string {
	return fmt.Sprintf("[GET /apps/{app-slug}/branches][%d] branchListUnauthorized  %+v", 401, o.Payload)
}

func (o *BranchListUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ServiceStandardErrorRespModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewBranchListNotFound creates a BranchListNotFound with default headers values
func NewBranchListNotFound() *BranchListNotFound {
	return &BranchListNotFound{}
}

/*BranchListNotFound handles this case with default header values.

Not Found
*/
type BranchListNotFound struct {
	Payload *models.ServiceStandardErrorRespModel
}

func (o *BranchListNotFound) Error() string {
	return fmt.Sprintf("[GET /apps/{app-slug}/branches][%d] branchListNotFound  %+v", 404, o.Payload)
}

func (o *BranchListNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ServiceStandardErrorRespModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewBranchListInternalServerError creates a BranchListInternalServerError with default headers values
func NewBranchListInternalServerError() *BranchListInternalServerError {
	return &BranchListInternalServerError{}
}

/*BranchListInternalServerError handles this case with default header values.

Internal Server Error
*/
type BranchListInternalServerError struct {
	Payload *models.ServiceStandardErrorRespModel
}

func (o *BranchListInternalServerError) Error() string {
	return fmt.Sprintf("[GET /apps/{app-slug}/branches][%d] branchListInternalServerError  %+v", 500, o.Payload)
}

func (o *BranchListInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ServiceStandardErrorRespModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
