package main

import (
	"log"

	"github.com/4L3X4NND3RR/chapter7/base62Example/base62"
)


func main() {
	x := 100
       base62String := base62.ToBase62(x)
       log.Println(base62String)
       normalNumber := base62.ToBase10(base62String)
       log.Println(normalNumber)
}
