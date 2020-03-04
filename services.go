package ves

import (
	"bytes"
	"encoding/json"
)

func newVehiclesService(c *Client) *VehiclesService {
	return &VehiclesService{
		BasePath: "vehicles",
		c:        c,
	}
}

type VehiclesService struct {
	BasePath string

	c *Client
}

func (vs *VehiclesService) GetVehicle(reg string) (v *Vehicle, err error) {
	v = new(Vehicle)

	b, _ := json.Marshal(VehicleRequest{
		RegistrationNumber: reg,
	})

	req := vs.c.NewRequest("POST", vs.BasePath, bytes.NewBuffer(b))

	_, err = req.Do(v)
	if err != nil {
		return nil, err
	}
	return v, nil
}
