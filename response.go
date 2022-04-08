package response

import "net/http"

type payloadType int

const (
	emptyPayload payloadType = iota
	textPayload
	jsonPayload
)

type HttpResponse struct {
	statusCode  int
	payloadType payloadType
	payload     interface{}
	logError    error
	errMessage  string
}

func (r *HttpResponse) StatusCode() int {
	return r.statusCode
}

func (r *HttpResponse) Payload() interface{} {
	return r.payload
}

func (r *HttpResponse) Error() error {
	return r.logError
}

func (r *HttpResponse) ErrorMessage() string {
	return r.errMessage
}

func OKPayload(payload interface{}) *HttpResponse {
	return successfulResponse(http.StatusOK, payload, jsonPayload)
}

func OKText(text string) *HttpResponse {
	return successfulResponse(http.StatusOK, text, textPayload)
}

func Created(payload interface{}) *HttpResponse {
	return successfulResponse(http.StatusCreated, payload, jsonPayload)
}

func Accepted(payload interface{}) *HttpResponse {
	return successfulResponse(http.StatusAccepted, payload, jsonPayload)
}

func NoContent() *HttpResponse {
	return successfulResponse(http.StatusNoContent, nil, emptyPayload)
}

func BadRequest(err error, publicMessage string) *HttpResponse {
	return errorResponse(http.StatusBadRequest, err, publicMessage)
}

func successfulResponse(statusCode int, payload interface{}, payloadType payloadType) *HttpResponse {
	return &HttpResponse{
		statusCode,
		payloadType,
		payload,
		nil,
		"",
	}
}

func errorResponse(statusCode int, err error, publicErrMessage string) *HttpResponse {
	return &HttpResponse{
		statusCode,
		emptyPayload,
		nil,
		err,
		publicErrMessage,
	}
}
