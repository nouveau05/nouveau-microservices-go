package main

import (
	"fmt"
	"go-nouveau-postgres-api/router"
	"log"
	"net/http"
)

func main() {

	r := router.Router()

	fmt.Println("Starting server on the port 8086.....")
	log.Fatal(http.ListenAndServe(":8086", r))

}
