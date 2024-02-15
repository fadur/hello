package main

import (
	"fmt"

	"github.com/fadur/hello/cmd"
	"github.com/fadur/hello/internal/models"
)

func main() {

	err := models.InitDB()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("starting server on port 8080")
	if err := cmd.Server(); err != nil {
		fmt.Println(err)
	}

}
