package main

import (
	"bufio"
	"encoding/json"
	"flag"
	"log"
	"os"
	"path/filepath"

	"github.com/jfoster/go-ves"
)

func main() {
	flag.Parse()
	args := flag.Args()
	var regs []string

	if len(args) == 0 {
		log.Fatal("No args specified")
	}

	if filepath.Ext(args[0]) == ".txt" {
		file, _ := os.Open(args[0])

		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			regs = append(regs, scanner.Text())
		}
	} else {
		regs = args
	}

	key := os.Getenv("API_KEY")
	c := ves.NewClient(ves.SetKey(key))

	for _, reg := range regs {
		v, err := c.Vehicles.GetVehicle(reg)
		if err != nil {
			log.Print(err)
			continue
		}

		json.NewEncoder(os.Stdout).Encode(v)
	}
}
