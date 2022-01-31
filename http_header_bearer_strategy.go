package main

import (
	"net/http"
)

func init() {
	RegisterStrategyFactory("bearer", HttpHeaderBearerStrategyFactory)
}

func HttpHeaderBearerStrategyFactory(config *AuthenticationConfiguration) (HttpStrategy, error) {
	return NewHttpHeaderBearerStrategy(config.Token)
}

type BearerHeaderStrategy struct {
	httpHeaderStrategy *HttpHeaderStrategy
}

func (strategy *BearerHeaderStrategy) Authenticate(request *http.Request) {
	strategy.httpHeaderStrategy.Authenticate(request)
}

func NewHttpHeaderBearerStrategy(token string) (*BearerHeaderStrategy, error) {
	token = "Bearer " + token
	httpHeaderStrategy, err := NewHttpHeaderStrategy("Authorization", token)

	if err != nil {
		return nil, err
	}

	return &BearerHeaderStrategy{
		httpHeaderStrategy: httpHeaderStrategy,
	}, nil
}
