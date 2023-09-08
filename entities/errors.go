package entities

import (
	"errors"
	"fmt"
	"net/http"
)

var (
	ErrNotFound            = errors.New("Not Found")
	ErrServiceNotAvailable = errors.New("Service Is Not Available")
	ErrUnauthorized        = errors.New("Unauthorized")
	ErrForbidden           = errors.New("Forbidden")
)

type ErrRequestWithResponse struct {
	Request  *http.Request
	Response *http.Response
}

func NewErrRequestWithResponse(request *http.Request, response *http.Response) *ErrRequestWithResponse {
	return &ErrRequestWithResponse{
		Request:  request,
		Response: response,
	}
}

func (e ErrRequestWithResponse) Error() string {
	return fmt.Sprintf("error request with response from %s %s", e.Request.Method, e.Request.URL)
}

type ErrRequestWithMessage string

func NewErrRequestWithMessage(message string) ErrRequestWithMessage {
	return ErrRequestWithMessage(message)
}

func (e ErrRequestWithMessage) Error() string {
	return string(e)
}

type ErrRequestWithMessageNotFound string

func (e ErrRequestWithMessageNotFound) Error() string {
	return string(e)
}

func NewErrorRequestWithMessageNotFound(message string) ErrRequestWithMessageNotFound {
	return ErrRequestWithMessageNotFound(message)
}

type ErrRequestUnauthorized string

func (e ErrRequestUnauthorized) Error() string {
	return string(e)
}

func NewErrorRequestUnauthorized(message string) ErrRequestUnauthorized {
	return ErrRequestUnauthorized(message)
}

type ErrRequestWithMessageUnprocessableEntity string

func NewErrRequestWithMessageUnprocessableEntity(message string) ErrRequestWithMessageUnprocessableEntity {
	return ErrRequestWithMessageUnprocessableEntity(message)
}

func (e ErrRequestWithMessageUnprocessableEntity) Error() string {
	return string(e)
}

type ErrRequestWithMessageConflict string

func NewErrRequestWithMessageConflict(message string) ErrRequestWithMessageConflict {
	return ErrRequestWithMessageConflict(message)
}

func (e ErrRequestWithMessageConflict) Error() string {
	return string(e)
}
