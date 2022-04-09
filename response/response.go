package response

type PayloadType int

const (
	PayloadEmpty PayloadType = iota
	PayloadText
	PayloadJSON
	PayloadStream
)

type HttpResponse struct {
	statusCode  int
	payloadType PayloadType
	payload     interface{}
	contentType string
	logError    error
	errMessage  string
}

func (r *HttpResponse) StatusCode() int {
	return r.statusCode
}

func (r *HttpResponse) Payload() interface{} {
	return r.payload
}

func (r *HttpResponse) PayloadType() PayloadType {
	return r.payloadType
}

func (r *HttpResponse) Error() error {
	return r.logError
}

func (r *HttpResponse) ErrorMessage() string {
	return r.errMessage
}
