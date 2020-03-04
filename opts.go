package ves

import "net/http"

func SetClient(client *http.Client) OptFunc {
	return func(c *Client) {
		c.httpClient = client
	}
}

func SetKey(key string) OptFunc {
	return func(c *Client) {
		c.Key = key
	}
}
