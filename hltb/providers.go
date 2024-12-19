package hltb

import (
	"io"
	"net/http"
)

type requestProvider interface {
	NewRequest(method string, url string, body io.Reader) (*http.Request, error)
}

type defaultRequestProvider struct{}

func (p *defaultRequestProvider) NewRequest(method string, url string, body io.Reader) (*http.Request, error) {
	return http.NewRequest(method, url, body)
}
