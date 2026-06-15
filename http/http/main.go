package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {

	http.HandleFunc("/post", postHandler)

	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Println("Server error:", err)

	}

}

func postHandler(w http.ResponseWriter, r *http.Request) {

	if r.Method != "POST" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	bodyBytes, err := io.ReadAll(r.Body)
	if err != nil {
		fmt.Printf("Failed receiving request : %v", err)
		return
	}

	fmt.Fprintln(w, string(bodyBytes))
}
