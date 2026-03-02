package dto

import "strings"

type Error struct {
	Message        string `json:"message"`
	Warn           string `json:"warn"`
	HttpStatusCode int    `json:"http_status_code"`
	Error          any    `json:"error"`
}

type Request struct {
	Method string          `json:"method"`
	Url    string          `json:"url"`
	Body   *strings.Reader `json:"body"`
}
