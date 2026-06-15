package main

import (
	"fmt"
	"http/config"
	"http/pkg"
	"http/routes"
	"net/http"
)

func main() {

	connectionkey, err := config.Config()
	if err != nil {
		fmt.Println(err)
		return
	}

	conn, err := pkg.InitDB(connectionkey)
	if err != nil {
		fmt.Println(err)
		return
	}

	routes.Routes(conn)

	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Println("Server error:", err)
	}

}
