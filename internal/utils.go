package internal

import "net/http"

type MappedNullable interface {
	ToMap() (map[string]interface{}, error)
}

type RoundTripFunc func(req *http.Request) *http.Response

func (f RoundTripFunc) RoundTrip(req *http.Request) (*http.Response, error) {
	return f(req), nil
}

func NewTestClient(testFn RoundTripFunc) *http.Client {
	return &http.Client{
		Transport: RoundTripFunc(testFn),
	}
}
