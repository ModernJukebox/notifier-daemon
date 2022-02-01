package main

import (
	"fmt"
	"net/http"
	"net/url"
	"strings"
)

func init() {
	RegisterTransportFactory("http", HttpTransportFactory)
	RegisterTransportFactory("https", HttpTransportFactory)
}

func HttpTransportFactory(u *url.URL, authenticationStrategy *HttpStrategy) (Transport, error) {
	transport, err := NewHttpTransport(u, authenticationStrategy)

	if err != nil {
		return nil, err
	}

	return transport, nil
}

type HttpTransport struct {
	serverUrl *url.URL

	authenticationStrategy *HttpStrategy
}

func NewHttpTransport(u *url.URL, authenticationStrategy *HttpStrategy) (*HttpTransport, error) {
	return &HttpTransport{
		serverUrl:              u,
		authenticationStrategy: authenticationStrategy,
	}, nil
}

func (transport *HttpTransport) Send(data string) error {
	request, err := http.NewRequest("POST", transport.serverUrl.String(), strings.NewReader(data))

	if err != nil {
		return err
	}

	client := &http.Client{}
	response, err := client.Do(request)

	if err != nil {
		return fmt.Errorf("error sending notification: %s", err)
	}

	if 200 > response.StatusCode || response.StatusCode >= 300 {
		return fmt.Errorf("failed to send notification: %s", response.Status)
	}

	return nil
}

func (transport *HttpTransport) beforeSend(request *http.Request) error {
	request.Header.Add("User-Agent", "notifier-daemon/0.1.0-DEV")

	(*transport.authenticationStrategy).Authenticate(request)

	return nil
}
