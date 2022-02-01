package main

import (
	"net/http"
	"net/url"
	"strconv"
)

func init() {
	RegisterTransportFactory("mercure", MercureTransportFactory)
}

func MercureTransportFactory(u *url.URL, authenticationStrategy *HttpStrategy) (Transport, error) {
	q := u.Query()
	isPrivate := false

	if q.Has("private") {
		private, err := strconv.ParseBool(q.Get("private"))

		if err != nil {
			return nil, err
		}

		isPrivate = private
	}

	topics := make([]string, len(q["topic"]))
	for i, v := range q["topic"] {
		topics[i] = v
	}

	u.RawQuery = ""
	u.Scheme = "https"

	transport, err := NewMercureTransport(isPrivate, topics, u, authenticationStrategy)

	if err != nil {
		return nil, err
	}

	return transport, nil
}

type MercureTransport struct {
	private bool
	topics  []string

	httpTransport *HttpTransport
}

func NewMercureTransport(private bool, topics []string, u *url.URL, authenticationStrategy *HttpStrategy) (*MercureTransport, error) {
	if err := All(topics, NotBlank); err != nil {
		return nil, err
	}

	u.Path = "/.well-known/mercure"

	httpTransport, err := NewHttpTransport(u, authenticationStrategy)

	if err != nil {
		return nil, err
	}

	httpTransport.beforeSend = func(request *http.Request) error {
		request.Header.Set("Content-Type", "application/x-www-form-urlencoded")

		return nil
	}

	return &MercureTransport{
		private:       private,
		topics:        topics,
		httpTransport: httpTransport,
	}, nil
}

func (transport *MercureTransport) Send(data string) error {
	form := url.Values{}

	for _, topic := range transport.topics {
		form.Add("topic", url.QueryEscape(topic))
	}

	form.Add("data", url.QueryEscape(data))

	if transport.private {
		form.Add("private", "on")
	}

	return transport.httpTransport.Send(form.Encode())
}
