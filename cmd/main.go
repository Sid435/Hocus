package main

import (
	"fmt"
	"log"

	"github.com/sid/Hocus/hocus"
)

func main() {
	user := map[string]string{
		"name": "siddharth",
		"age":  "36",
	}
	_ = user
	db, err := hocus.New()
	if err != nil {
		log.Fatal(err)
	}
	coll, err := db.CreateCollection("users")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%+v\n", coll)

}
