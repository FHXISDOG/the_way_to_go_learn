package main

import (
	"encoding/json"
	"errors"
	"fmt"
)

type Address struct {
	Type    string
	City    string
	Country string
}

type VCard struct {
	FirstName string
	LastName  string
	Address
	Remark string
}

var errNotFound error = errors.New("path not found")

func main() {
	pa := Address{"private", "Aartselaar", "Belgium"}
	// wa := &Address{"work", "Boom", "Belgium"}
	vc := VCard{"Jan", "Kersschot", pa, "none"}
	// fmt.Printf("%v: \n", vc) // {Jan Kersschot [0x126d2b80 0x126d2be0] none}:
	// JSON format:
	js, _ := json.Marshal(vc)
	fmt.Printf("JSON format: %s", js)
}
