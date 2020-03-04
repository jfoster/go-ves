package ves

import (
	"bytes"
	"encoding/json"
)

func newVehiclesService(c *Client) *VehiclesService {
	return &VehiclesService{
		BasePath: "vehicles",
		client:   c,
	}
}

type VehiclesService struct {
	BasePath string

	client *Client
}

func (vs *VehiclesService) GetVehicle(reg string) (v *Vehicle, err error) {
	v = new(Vehicle)

	b, err := json.Marshal(VehicleRequest{
		RegistrationNumber: reg,
	})
	if err != nil {
		return nil, err
	}

	req := vs.client.NewRequest("POST", vs.BasePath, bytes.NewBuffer(b))
	_, err = req.Do(v)
	if err != nil {
		return nil, err
	}
	return v, nil
}
