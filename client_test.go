package ves

import (
	"io"
	"net/http"
	"reflect"
	"testing"
)

func TestNewClient(t *testing.T) {
	type args struct {
		opts []OptFunc
	}
	tests := []struct {
		name string
		args args
		want *Client
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewClient(tt.args.opts...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewClient() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestClient_do(t *testing.T) {
	type args struct {
		req *http.Request
		v   interface{}
	}
	tests := []struct {
		name    string
		c       *Client
		args    args
		want    *http.Response
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.c.do(tt.args.req, tt.args.v)
			if (err != nil) != tt.wantErr {
				t.Errorf("Client.do() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Client.do() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestClient_NewRequest(t *testing.T) {
	type args struct {
		method string
		path   string
		body   io.Reader
	}
	tests := []struct {
		name string
		c    *Client
		args args
		want *Request
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.c.NewRequest(tt.args.method, tt.args.path, tt.args.body); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Client.NewRequest() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRequest_Do(t *testing.T) {
	type args struct {
		v interface{}
	}
	tests := []struct {
		name    string
		r       *Request
		args    args
		want    *http.Response
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.r.Do(tt.args.v)
			if (err != nil) != tt.wantErr {
				t.Errorf("Request.Do() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Request.Do() = %v, want %v", got, tt.want)
			}
		})
	}
}
