package main

import (
	"fmt"
	"net/http"
)

type StrategyFactory = func(config *AuthenticationConfiguration) (HttpStrategy, error)

var strategyFactories = make(map[string]StrategyFactory)

func RegisterStrategyFactory(t string, factory StrategyFactory) {
	strategyFactories[t] = factory
}

func NewStrategy(config *AuthenticationConfiguration) (HttpStrategy, error) {
	factory, ok := strategyFactories[config.Type]

	if !ok {
		return nil, fmt.Errorf("no authentication strategy factory registered for type %s", config.Type)
	}

	return factory(config)
}

type HttpStrategy interface {
	Authenticate(request *http.Request)
}
