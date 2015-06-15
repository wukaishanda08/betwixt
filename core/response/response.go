// Various responses to a LWM2M Server request
// Typically responses represent a CoAP Request code
package response

import (
	. "github.com/zubairhamed/betwixt"
	"github.com/zubairhamed/canopus"
	"github.com/zubairhamed/go-commons/typeval"
)

// 201 Created
func Created() Lwm2mResponse {
	return &CreatedResponse{}
}

type CreatedResponse struct {
}

func (r *CreatedResponse) GetResponseCode() canopus.CoapCode {
	return canopus.COAPCODE_201_CREATED
}

func (r *CreatedResponse) GetResponseValue() typeval.Value {
	return typeval.Empty()
}

// 202 Deleted
func Deleted() Lwm2mResponse {
	return &DeletedResponse{}
}

type DeletedResponse struct {
}

func (r *DeletedResponse) GetResponseCode() canopus.CoapCode {
	return canopus.COAPCODE_202_DELETED
}

func (r *DeletedResponse) GetResponseValue() typeval.Value {
	return typeval.Empty()
}

// 204 Changed
func Changed() Lwm2mResponse {
	return &ChangedResponse{}
}

type ChangedResponse struct {
}

func (r *ChangedResponse) GetResponseCode() canopus.CoapCode {
	return canopus.COAPCODE_204_CHANGED
}

func (r *ChangedResponse) GetResponseValue() typeval.Value {
	return typeval.Empty()
}

// 205 Content
func Content(val typeval.Value) Lwm2mResponse {
	return &ContentResponse{
		val: val,
	}
}

type ContentResponse struct {
	val typeval.Value
}

func (r *ContentResponse) GetResponseCode() canopus.CoapCode {
	return canopus.COAPCODE_205_CONTENT
}

func (r *ContentResponse) GetResponseValue() typeval.Value {
	return r.val
}

// 400 Bad Request
func BadRequest() Lwm2mResponse {
	return &BadRequestResponse{}
}

type BadRequestResponse struct {
}

func (r *BadRequestResponse) GetResponseCode() canopus.CoapCode {
	return canopus.COAPCODE_400_BAD_REQUEST
}

func (r *BadRequestResponse) GetResponseValue() typeval.Value {
	return typeval.Empty()
}

// 401 Unauthorized
func Unauthorized() Lwm2mResponse {
	return &UnauthorizedResponse{}
}

type UnauthorizedResponse struct {
}

func (r *UnauthorizedResponse) GetResponseCode() canopus.CoapCode {
	return canopus.COAPCODE_401_UNAUTHORIZED
}

func (r *UnauthorizedResponse) GetResponseValue() typeval.Value {
	return typeval.Empty()
}

// 404 Not Found
func NotFound() Lwm2mResponse {
	return &NotFoundResponse{}
}

type NotFoundResponse struct {
}

func (r *NotFoundResponse) GetResponseCode() canopus.CoapCode {
	return canopus.COAPCODE_404_NOT_FOUND
}

func (r *NotFoundResponse) GetResponseValue() typeval.Value {
	return typeval.Empty()
}

// 405 Method Not Allowed
func MethodNotAllowed() Lwm2mResponse {
	return &MethodNotAllowedResponse{}
}

type MethodNotAllowedResponse struct {
}

func (r *MethodNotAllowedResponse) GetResponseCode() canopus.CoapCode {
	return canopus.COAPCODE_405_METHOD_NOT_ALLOWED
}

func (r *MethodNotAllowedResponse) GetResponseValue() typeval.Value {
	return typeval.Empty()
}

// 409 Conflict
func Conflict() Lwm2mResponse {
	return &ConflictResponse{}
}

type ConflictResponse struct {
}

func (r *ConflictResponse) GetResponseCode() canopus.CoapCode {
	return canopus.COAPCODE_409_CONFLICT
}

func (r *ConflictResponse) GetResponseValue() typeval.Value {
	return typeval.Empty()
}
