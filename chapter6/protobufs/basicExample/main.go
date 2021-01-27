package main

import (
	"fmt"

	"github.com/4L3X4NND3RR/chapter6/protobufs/protofiles"
	"google.golang.org/protobuf/proto"
)

func main() {
	p := &protofiles.Person{
		Id:    1234,
		Name:  "Roger F",
		Email: "rf@example.com",
		Phones: []*protofiles.Person_PhoneNumber{
			{Number: "555-4321", Type: protofiles.Person_HOME},
		},
	}

	p1 := &protofiles.Person{}
	body, _ := proto.Marshal(p)
	_ = proto.Unmarshal(body, p1)
	fmt.Println("Original struct loaded from proto file: ", p, "\n")
	fmt.Println("Marshalled proto data: ", body, "\n")
	fmt.Println("Unmarshalled struct: ", p1)
}
