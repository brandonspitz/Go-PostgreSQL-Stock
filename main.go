package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/brandonspitz/Go-PostgreSQL-Stock/router"
)

func main() {
	r := router.Router()
	fmt.Println("Starting server on port 8080...")

	log.Fatal(http.ListenAndServe(":8000", r))
}
