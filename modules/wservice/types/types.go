package types

type Input struct {
	Header Header      `json:"header"`
	Body   interface{} `json:"body"`
}

type Header struct {
	ID          string `json:"id"`
	ReqSequence string `json:"req_sequence"`
}

type Output struct {
	Header Header      `json:"header"`
	Body   interface{} `json:"body"`
}
