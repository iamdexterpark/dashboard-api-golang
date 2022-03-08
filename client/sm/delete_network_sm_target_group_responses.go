// Code generated by go-swagger; DO NOT EDIT.

package sm

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// DeleteNetworkSmTargetGroupReader is a Reader for the DeleteNetworkSmTargetGroup structure.
type DeleteNetworkSmTargetGroupReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteNetworkSmTargetGroupReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewDeleteNetworkSmTargetGroupNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDeleteNetworkSmTargetGroupNoContent creates a DeleteNetworkSmTargetGroupNoContent with default headers values
func NewDeleteNetworkSmTargetGroupNoContent() *DeleteNetworkSmTargetGroupNoContent {
	return &DeleteNetworkSmTargetGroupNoContent{}
}

/* DeleteNetworkSmTargetGroupNoContent describes a response with status code 204, with default header values.

Successful operation
*/
type DeleteNetworkSmTargetGroupNoContent struct {
}

func (o *DeleteNetworkSmTargetGroupNoContent) Error() string {
	return fmt.Sprintf("[DELETE /networks/{networkId}/sm/targetGroups/{targetGroupId}][%d] deleteNetworkSmTargetGroupNoContent ", 204)
}

func (o *DeleteNetworkSmTargetGroupNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
