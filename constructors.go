package response

import "net/http"

// region 2xx

func OK() *Partial {
	return &Partial{&HttpResponse{statusCode: http.StatusOK}}
}

func Created() *Partial {
	return &Partial{&HttpResponse{statusCode: http.StatusCreated}}
}

func Accepted() *Partial {
	return &Partial{&HttpResponse{statusCode: http.StatusAccepted}}
}

func NoContent() *HttpResponse {
	return &HttpResponse{statusCode: http.StatusNoContent, payloadType: Empty}
}

// endregion 2xx

// region 4xx

func BadRequest(err error, publicMessage string) *HttpResponse {
	return &HttpResponse{statusCode: http.StatusBadRequest, logError: err, errMessage: publicMessage}
}

func Unauthorized(err error, publicMessage string) *HttpResponse {
	return &HttpResponse{statusCode: http.StatusUnauthorized, logError: err, errMessage: publicMessage}
}

func PaymentRequired(err error, publicMessage string) *HttpResponse {
	return &HttpResponse{statusCode: http.StatusPaymentRequired, logError: err, errMessage: publicMessage}
}

func Forbidden(err error, publicMessage string) *HttpResponse {
	return &HttpResponse{statusCode: http.StatusForbidden, logError: err, errMessage: publicMessage}
}

func NotFound(err error, publicMessage string) *HttpResponse {
	return &HttpResponse{statusCode: http.StatusNotFound, logError: err, errMessage: publicMessage}
}

func NotAcceptable(err error, publicMessage string) *HttpResponse {
	return &HttpResponse{statusCode: http.StatusNotAcceptable, logError: err, errMessage: publicMessage}
}

func Conflict(err error, publicMessage string) *HttpResponse {
	return &HttpResponse{statusCode: http.StatusConflict, logError: err, errMessage: publicMessage}
}

// endregion 4xx

// region 5xx

func InternalServerError(err error, publicMessage string) *HttpResponse {
	return &HttpResponse{statusCode: http.StatusInternalServerError, logError: err, errMessage: publicMessage}
}

func NotImplemented(err error, publicMessage string) *HttpResponse {
	return &HttpResponse{statusCode: http.StatusNotImplemented, logError: err, errMessage: publicMessage}
}

func ServiceUnavailable(err error, publicMessage string) *HttpResponse {
	return &HttpResponse{statusCode: http.StatusServiceUnavailable, logError: err, errMessage: publicMessage}
}

func InsufficientStorage(err error, publicMessage string) *HttpResponse {
	return &HttpResponse{statusCode: http.StatusInsufficientStorage, logError: err, errMessage: publicMessage}
}

// endregion 5xx
