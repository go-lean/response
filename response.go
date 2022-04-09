package response

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

func New(statusCode int) *PartialCustom {
	return &PartialCustom{Partial{&HttpResponse{statusCode: statusCode}}}
}
