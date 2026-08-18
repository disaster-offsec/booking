package main 

import (
	"log"
	"flag"
	"net/http"

	"booking/internal/handlers"
)

func main() {
	port := flag.String("port", "8080", "port for service")
	flag.Parse()
	
	mux := http.NewServeMux()
	mux.HandleFunc("GET /ping", handlers.Ping)

	log.Printf("Server has started on port %s", port)
	if err := http.ListenAndServe(":" + *port, mux); err != nil {
		log.Fatal(err)
	}
}
