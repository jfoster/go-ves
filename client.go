package ves

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

type Client struct {
	BaseURL string
	Key     string
	Version int

	httpClient *http.Client

	Vehicles *VehiclesService
}

type OptFunc func(c *Client)

func NewClient(opts ...OptFunc) *Client {
	c := &Client{
		BaseURL: VESAPIUrl,
		Version: VESAPIVersion,
		httpClient: &http.Client{
			Timeout: time.Second * 5,
		},
	}

	for _, f := range opts {
		f(c)
	}

	c.Vehicles = newVehiclesService(c)

	return c
}

func (c *Client) do(req *http.Request, v interface{}) (*http.Response, error) {
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode < http.StatusOK || resp.StatusCode >= http.StatusBadRequest {
		return nil, fmt.Errorf("%s", resp.Body)
	}

	return resp, json.NewDecoder(resp.Body).Decode(v)
}

func (c *Client) NewRequest(method string, path string, body io.Reader) *Request {
	url := fmt.Sprintf("%s/v%d/%s", c.BaseURL, c.Version, path)

	req, err := http.NewRequest(method, url, body)

	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json")

	req.Header.Set("x-api-key", c.Key)

	return &Request{req, c, err}
}

type Request struct {
	*http.Request

	c *Client
	e error
}

func (r *Request) Do(v interface{}) (*http.Response, error) {
	if r.e != nil {
		return nil, r.e
	}

	resp, err := r.c.httpClient.Do(r.Request)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if v != nil {
		return resp, json.NewDecoder(resp.Body).Decode(v)
	}

	return resp, nil
}
