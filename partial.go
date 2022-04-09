package response

type Partial struct {
	response *HttpResponse
}

type PartialCustom struct {
	Partial
}

func (p *Partial) WithJSON(payload interface{}) *HttpResponse {
	p.response.payload = payload
	p.response.payloadType = PayloadJSON

	return p.response
}

func (p *Partial) WithText(text string) *HttpResponse {
	p.response.payload = text
	p.response.payloadType = PayloadText

	return p.response
}

func (p *PartialCustom) WithoutContent() *HttpResponse {
	p.response.payload = nil
	p.response.payloadType = PayloadEmpty

	return p.response
}

func (p *PartialCustom) WithError(err error, publicMessage string) *HttpResponse {
	p.response.payload = nil
	p.response.payloadType = PayloadEmpty
	p.response.logError = err
	p.response.errMessage = publicMessage

	return p.response
}
