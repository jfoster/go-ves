package main

import (
	"encoding/json"
	"flag"
	"log"
	"os"

	"github.com/jfoster/go-ves"
)

func main() {
	flag.Parse()
	args := flag.Args()

	if len(args) != 1 {
		log.Fatal("incorrect number or args")
	}

	key := os.Getenv("API_KEY")

	c := ves.NewClient(ves.SetKey(key))

	v, err := c.Vehicles.GetVehicle(args[0])
	if err != nil {
		log.Panic(err)
	}

	json.NewEncoder(os.Stdout).Encode(v)
}
