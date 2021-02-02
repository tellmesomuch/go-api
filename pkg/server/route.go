package server

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

var bookList []Book
var authorList []Author

type Book struct {
	Title  string
	Author Author
}

type Author struct {
	Name     string
	LastName string
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
