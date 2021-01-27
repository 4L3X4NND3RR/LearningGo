package main

import (
	"log"

	"github.com/4L3X4NND3RR/chapter7/basicExample/helper"
)

func main() {
	_, err := helper.InitDB()
	if err != nil {
		log.Println(err)
	}

	log.Println("Database tables are successfully initialized.")
}
