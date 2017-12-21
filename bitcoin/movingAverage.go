package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
)

type Person struct {
	test int
}

func main() {
	data, err := ioutil.ReadFile(`test.json`)

	var persons Person
	if err := json.Unmarshal(data, &persons); err != nil {
		log.Fatal(err)
	}

	if err != nil {
		log.Fatal(err)
	}

	fmt.Print(string(data))
}
