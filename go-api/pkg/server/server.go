package server

import (
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

// Serve create new web server and start to listen
func Serve(addr string) {
	router := mux.NewRouter()

	// Server rest API routes
	serveAPI(router)

	// Create server instance
	var httpServer = &http.Server{
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
		Handler:      router,
		Addr:         addr,
	}

	// Start http server
	log.Printf("Yo Listening on %s\n", httpServer.Addr)
	log.Fatal(httpServer.ListenAndServe())
}
