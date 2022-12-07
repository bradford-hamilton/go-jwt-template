package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/bradford-hamilton/go-jwt-template/internal/server"
	"github.com/bradford-hamilton/go-jwt-template/internal/storage"
)

func main() {
	db, err := storage.NewDB()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	s := server.New(db)
	port := os.Getenv("GO_JWT_TEMPLATE_SERVER_PORT")
	if port == "" {
		port = "4000"
	}

	fmt.Printf("Serving application on port %s\n", port)
	log.Fatal(http.ListenAndServe(":"+port, s.Mux))
}
