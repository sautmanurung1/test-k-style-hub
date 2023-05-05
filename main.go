package main

import (
	"fmt"
	"skincare/infrsatructure/http/server"
)

func main() {
	app := server.Server()
	err := app.Start(":8080")
	if err != nil {
		fmt.Println("Disini error nya : ", err)
	}
}
