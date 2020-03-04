package main

import (
	"encoding/json"
	"log"
	"os"

	"github.com/jfoster/go-ves"
)

func main() {
	key := os.Getenv("API_KEY")
	reg := os.Getenv("REG_NUM")

	c := ves.NewClient(ves.SetKey(key))

	v, err := c.Vehicles.GetVehicle(reg)
	if err != nil {
		log.Fatal(err)
	}

	json.NewEncoder(os.Stdout).Encode(v)
}
