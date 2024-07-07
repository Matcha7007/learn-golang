package main

import (
	"log"
	"net/http"
)

// define a home handler function which writes a bytes slice containing
// "Hello from Snippetbox" as the response body.
func home(w http.ResponseWriter, r *http.Request) {
	// Check if the current request URL path exactly matches "/". If it doesn't
	// the http.NotFound() function to send a 404 response to the client.
	// Importantly, we then return from the handler. If we don't return the hand
	// would keep executing and also write the "Hello from SnippetBox" message.
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	w.Write([]byte("Hello from Snippetbox"))
}

// Add a showSnippet handler function.
func showSnippet(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Display a specific snippet..."))
}

// Add a createSnippet handler function.
func createSnippet(w http.ResponseWriter, r *http.Request) {
	// Use r.Method to check whether the request is using POST or not.
	// If it's not, use the w.WriteHeader() method to send a 405 status code and
	// the w.Write() method to write a "Method Not Allowed" response body. We
	// then return from the function so that the subsequent code is not executed
	// if r.Method != "POST" {
	// 	w.WriteHeader(405)
	// 	w.Write([]byte("Method Now Allowed"))
	// 	return
	// }

	// Use the Header().Set() method to add an 'Allow: POST' header to the
	// response header map. The first parameter is the header name, and
	// the second parameter is the header value.
	// if r.Method != "POST" {
	//  w.Header().Set("Allow", "POST")
	// 	w.WriteHeader(405)
	// 	w.Write([]byte("Method Now Allowed"))
	// 	return
	// }

	if r.Method != "POST" {
		w.Header().Set("Allow", "POST")
		// Use the http.Error() function to send a 405 status code and "Method N
		// Allowed" string as the response body.
		// http.StatusMethodNotAllowed return status code 405
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	// Set a new cache-control header. If an existing "Cache-Control" header exists
	// it will be overwritten.
	w.Header().Set("Cache-Control", "public, max-age=31536000")
	// In contrast, the Add() method appends a new "Cache-Control" header and can
	// be called multiple times.
	w.Header().Add("Cache-Control", "public")
	w.Header().Add("Cache-Control", "max-age=31536000")
	// Delete all values for the "Cache-Control" header.
	w.Header().Del("Cache-Control")
	// Retrieve the first value for the "Cache-Control" header.
	w.Header().Get("Cache-Control")

	w.Header()["X-XSS-Protection"] = []string{"1; mode=block"}

	w.Write([]byte("Create a new snippet..."))
}

func main() {
	// Use the http.NewServeMux() function to initialize a new servemux, then
	// register the home function as the handler for the "/" URL pattern.
	mux := http.NewServeMux()
	// "/" handler for home function
	mux.HandleFunc("/", home)
	// "/snippet" handler for showSnippet function
	mux.HandleFunc("/snippet", showSnippet)
	// "/snippet/create" handler for createSnippet function
	mux.HandleFunc("/snippet/create", createSnippet)

	// Use the http.ListenAndServe() function to start a new web server. We pas
	// two parameters: the TCP network address to listen on (in this case ":4000
	// and the servemux we just created. If http.ListenAndServe() returns an er
	// we use the log.Fatal() function to log the error message and exit.
	log.Println("Starting server on :4000")
	err := http.ListenAndServe(":4000", mux)
	log.Fatal(err)
}
