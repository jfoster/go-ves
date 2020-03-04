package ves

import (
	"net/http"
	"reflect"
	"testing"
)

func TestSetClient(t *testing.T) {
	type args struct {
		client *http.Client
	}
	tests := []struct {
		name string
		args args
		want OptFunc
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SetClient(tt.args.client); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SetClient() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSetKey(t *testing.T) {
	type args struct {
		key string
	}
	tests := []struct {
		name string
		args args
		want OptFunc
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SetKey(tt.args.key); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SetKey() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSetBaseURL(t *testing.T) {
	type args struct {
		url string
	}
	tests := []struct {
		name string
		args args
		want OptFunc
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SetBaseURL(tt.args.url); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SetBaseURL() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSetVersion(t *testing.T) {
	type args struct {
		v int
	}
	tests := []struct {
		name string
		args args
		want OptFunc
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SetVersion(tt.args.v); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SetVersion() = %v, want %v", got, tt.want)
			}
		})
	}
}
