package main

type AuthenticationConfiguration struct {
	Type       string `json:"type"`
	HeaderName string `json:"headerName;omitempty"`
	Token      string `json:"token"`
}
