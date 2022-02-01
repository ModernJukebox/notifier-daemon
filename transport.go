package main

import (
	"fmt"
	"net/url"
)

type TransportFactory = func(u *url.URL, authenticationStrategy *HttpStrategy) (Transport, error)

var transportFactories = make(map[string]TransportFactory)

func RegisterTransportFactory(scheme string, factory TransportFactory) {
	transportFactories[scheme] = factory
}

func NewTransport(u *url.URL, authenticationStrategy *HttpStrategy) (Transport, error) {
	factory, ok := transportFactories[u.Scheme]

	if !ok {
		return nil, fmt.Errorf("no transport factory registered for dsn %s", u.Redacted())
	}

	return factory(u, authenticationStrategy)
}

type Transport interface {
	Send(data string) error
}
