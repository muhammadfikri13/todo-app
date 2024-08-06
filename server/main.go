package main

import (
	"fmt"
	"log"
	"net/http"
	"server/router"
)

func main() {
	r := router.Router() // "router" is the package and "Router" is the function
	fmt.Println("Starting server on port 9000...-_-...")

	log.Fatal(http.ListenAndServe(":9000", r))
}
