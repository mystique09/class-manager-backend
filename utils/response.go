package utils

import (
	"encoding/json"
	"io"
)

type Response struct {
	Data  interface{} `json:"data,omitempty"`
	Error string      `json:"error,omitempty"`
}

/*
A function to create a new response.
*/
func NewResponse(data interface{}, error string) Response {
	return Response{
		Data:  data,
		Error: error,
	}
}

func GetJson(b io.Reader, t interface{}) error {
	return json.NewDecoder(b).Decode(t)
}
