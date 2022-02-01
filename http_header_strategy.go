package main

import (
	"fmt"
	"net/http"
)

type HttpHeaderStrategy struct {
	headerName string
	token      string
}

func init() {
	RegisterStrategyFactory("header", HttpHeaderStrategyFactory)
}

func HttpHeaderStrategyFactory(config *AuthenticationConfiguration) (HttpStrategy, error) {
	return NewHttpHeaderStrategy(config.HeaderName, config.Token)
}

func (strategy *HttpHeaderStrategy) Authenticate(request *http.Request) {
	request.Header.Set(strategy.headerName, strategy.token)
}

func NewHttpHeaderStrategy(headerName string, token string) (*HttpHeaderStrategy, error) {
	if err := NotBlank(headerName); err != nil {
		return nil, fmt.Errorf("header name is required")
	}

	if err := NotBlank(token); err != nil {
		return nil, fmt.Errorf("token is required")
	}

	return &HttpHeaderStrategy{
		headerName: headerName,
		token:      token,
	}, nil
}
