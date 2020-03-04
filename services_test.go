package ves

import (
	"reflect"
	"testing"
)

func Test_newVehiclesService(t *testing.T) {
	type args struct {
		c *Client
	}
	tests := []struct {
		name string
		args args
		want *VehiclesService
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := newVehiclesService(tt.args.c); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("newVehiclesService() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestVehiclesService_GetVehicle(t *testing.T) {
	type args struct {
		reg string
	}
	tests := []struct {
		name    string
		vs      *VehiclesService
		args    args
		wantV   *Vehicle
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotV, err := tt.vs.GetVehicle(tt.args.reg)
			if (err != nil) != tt.wantErr {
				t.Errorf("VehiclesService.GetVehicle() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotV, tt.wantV) {
				t.Errorf("VehiclesService.GetVehicle() = %v, want %v", gotV, tt.wantV)
			}
		})
	}
}
