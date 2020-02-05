package types

import (
	"bytes"
	"encoding/json"
	"net/http"
	"regexp"
)

type Response struct {
	Code int
	Data interface{}
	Msg string
}

func NewResponse(data interface{}) *Response{
	response := Response{}
	response.Code = 0
	response.Msg = "ok"
	response.Data = data
	return &response
}

func (resp *Response) Render(w http.ResponseWriter, r *http.Request) error {
	// Pre-processing before a response is marshalled and sent across the wire
	return nil
}

var keyMatchRegex = regexp.MustCompile(`"\w+":`)
var wordBarrierRegex = regexp.MustCompile(`([a-z])([A-Z])`)

func (resp Response) MarshalJSON() ([]byte, error) {
	// 防止递归
	type Response_ Response
	marshalled, err := json.Marshal(Response_(resp))
	converted := keyMatchRegex.ReplaceAllFunc(
		marshalled,
		func(match []byte) []byte {
			return bytes.ToLower(wordBarrierRegex.ReplaceAll(
				match,
				[]byte(`${1}_${2}`),
			))
		},
	)
	return converted, err
}
