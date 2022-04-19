package server

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

var bookList []Book
var authorList []Author

type Book struct {
	Title      string
	AuthorName string
}

type Author struct {
	Name     string
	LastName string
	Number   string
}

// ServeAPI list and serve all rest API route
func serveAPI(r *mux.Router) {

	// health
	r.HandleFunc("/api/health", func(w http.ResponseWriter, r *http.Request) {
		json.NewEncoder(w).Encode(map[string]bool{"ok": true})
	})

	// book routes

	r.HandleFunc("/api/book", func(w http.ResponseWriter, r *http.Request) {
		// Get and parse req object
		var b Book
		err := json.NewDecoder(r.Body).Decode(&b)
		if err != nil {
			returnError(err, w)
			return
		}
		found := false
		for _, a := range authorList {
			if a.LastName == b.AuthorName {
				found = true
			}
		}
		if !found {
			returnError(fmt.Errorf("Author %s is not listed yet", b.AuthorName), w)
			return
		}
		// Add to local store
		bookList = append(bookList, b)

		// Return response
		resp, err := json.Marshal(b)
		if err != nil {
			returnError(err, w)
			return
		}
		json.NewEncoder(w).Encode(map[string]interface{}{"ok": true, "created": string(resp)})
	}).Methods("POST")

	r.HandleFunc("/api/books", func(w http.ResponseWriter, r *http.Request) {
		// Get and parse req object
		var books []string

		// Convert books to json
		for _, b := range bookList {
			resp, err := json.Marshal(b)
			if err != nil {
				returnError(err, w)
			}
			books = append(books, string(resp))
		}

		// Return response
		json.NewEncoder(w).Encode(map[string]interface{}{"ok": true, "books": books})
	}).Methods("get")

	// Author routes

	r.HandleFunc("/api/author", func(w http.ResponseWriter, r *http.Request) {
		// Get and parse req object
		var a Author
		err := json.NewDecoder(r.Body).Decode(&a)
		if err != nil {
			returnError(err, w)
			return
		}
		// Add to local store
		authorList = append(authorList, a)

		// Return response
		resp, err := json.Marshal(a)
		if err != nil {
			returnError(err, w)
			return
		}
		json.NewEncoder(w).Encode(map[string]interface{}{"ok": true, "created": string(resp)})
	}).Methods("POST")

	r.HandleFunc("/api/authors", func(w http.ResponseWriter, r *http.Request) {
		// Get and parse req object
		log.Println("Received a call ! ")
		for _, cookie := range r.Cookies() {
			log.Printf("Found a cookie named: %s", cookie.Name)
		}
		for k, v := range r.Header {
			log.Printf("Header field %q, Value %q\n", k, v)
		}

		var authors []string

		// Convert books to json
		for _, a := range authorList {
			resp, err := json.Marshal(a)
			if err != nil {
				returnError(err, w)
			}
			authors = append(authors, string(resp))
		}

		// Return response
		json.NewEncoder(w).Encode(map[string]interface{}{"ok": true, "authors": authors})
	}).Methods("get")

}

func returnError(err error, w http.ResponseWriter) {
	json.NewEncoder(w).Encode(map[string]interface{}{"ok": false, "error": err})
}
