package main

import (
	"fmt"
	"log"
	"matheuspolitano/context-check/api"
)

func main() {
	users, err := api.GetUsers()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Name: %s %s %s\n", users[0].Name.Title, users[0].Name.First, users[0].Name.Last)

}
