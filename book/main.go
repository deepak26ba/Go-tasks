package main

import (
	"books/config"
	"books/pkg"
	"fmt"

	"net/http"

	"github.com/gorilla/mux"
)

func main() {

	connectionkey, err := config.Config()
	if err != nil {
		fmt.Println(err)
		return
	}

	conn, err := pkg.InitDB(connectionkey)
	if err != nil {
		fmt.Println(err, conn)
		return
	}

	r := mux.NewRouter()

	if err := http.ListenAndServe(":8080", r); err != nil {
		fmt.Println("Server error:", err)
	}

}
