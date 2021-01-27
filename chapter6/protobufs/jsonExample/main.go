package main

import (
	"encoding/json"
	"fmt"

	"github.com/4L3X4NND3RR/chapter6/protobufs/protofiles"
)

func main() {
	p := &protofiles.Person{
		Id: 1234,
		Name: "Roger F",
		Email: "rf@example.com",
		Phones: []*protofiles.Person_PhoneNumber{
			{Number: "555-4321", Type: protofiles.Person_HOME},
		},
	}
	body, _ := json.Marshal(p)
	fmt.Println(string(body))
}
