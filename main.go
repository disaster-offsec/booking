package main 

import (
	"log"
	"flag"
	"net/http"

	"booking/internal/handlers"
	"booking/internal/config"
	"booking/internal/storage"
)

func main() {
	port := flag.String("port", "8080", "port for service")
	flag.Parse()

	cfg := config.LoadDBConfig()
	if err := storage.InitDB(cfg); err != nil {
    	log.Fatal(err)
	}
	defer storage.DB.Close()
	
	mux := http.NewServeMux()
	mux.HandleFunc("GET /ping", handlers.Ping)
	mux.HandleFunc("POST /book", handlers.Book)
	mux.HandleFunc("GET /booklist", handlers.Booklist)


	log.Printf("Server has started on port %s", *port)
	if err := http.ListenAndServe(":" + *port, mux); err != nil {
		log.Fatal(err)
	}
}
