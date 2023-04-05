package main

import (
	"fmt"
	"log"
	"net/http"

	"golang-to-do-backend/router"

	"github.com/gorilla/handlers"
)

func main() {
	r := router.Router()
	fmt.Println("Starting server on the port 9000...")
	corsObj := handlers.AllowedOrigins([]string{"*"})

	log.Fatal(http.ListenAndServe(":9000", handlers.CORS(corsObj)(r)))

}
