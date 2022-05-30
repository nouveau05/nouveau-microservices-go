package main

import (
	"log"
	"net/http"
	"os"
	"github.com/nouveau05/nouveau-microservices-go/router"
)

func main() {

	r := router.Router()

	// Determine port for HTTP service.
        port := os.Getenv("PORT")
        if port == "" {
                port = "8080"
                log.Printf("defaulting to port %s", port)
        }

        // Start HTTP server.
        log.Printf("listening on port %s", port)
        if err := http.ListenAndServe(":"+port, r); err != nil {
                log.Fatal(err)
        }
}

