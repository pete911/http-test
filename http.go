package main

import (
	"io"
	"net/http"
	"net/http/httptrace"
	"time"
)

func DefaultTransport() *http.Transport {
	return &http.Transport{}
}

type Client struct {
	*http.Client
	trace *httptrace.ClientTrace
}

func NewClient(tr *http.Transport, trace *httptrace.ClientTrace) *Client {
	return &Client{
		Client: &http.Client{
			Transport: tr,
			Timeout:   10 * time.Second,
		},
		trace: trace,
	}
}

func DiscardResponse(res *http.Response) error {
	defer res.Body.Close()
	_, err := io.Copy(io.Discard, res.Body)
	return err
}

func (c *Client) Do(req *http.Request) (*http.Response, error) {
	req = req.WithContext(httptrace.WithClientTrace(req.Context(), c.trace))
	return c.Client.Do(req)
}
