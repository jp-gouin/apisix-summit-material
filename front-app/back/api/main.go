package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

type Config struct {
	PostAuthorUri string `json:"postAuthorUri"`
	PostBookUri   string `json:"postBookUri"`
	AuthorsUri    string `json:"authorsUri"`
	BooksUri      string `json:"booksUri"`
	Token         string `json:"token"`
}

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the HomePage!")
	fmt.Println("Endpoint Hit: homePage")
}

func callapi1(w http.ResponseWriter, r *http.Request) {
	callapi(w, r, os.Getenv("URLAPI1"), true)
}
func callapi2(w http.ResponseWriter, r *http.Request) {
	callapi(w, r, os.Getenv("URLAPI2"), false)
}
func getConfig(w http.ResponseWriter, r *http.Request) {
	for k, v := range r.Header {
		log.Printf("Header field %q, Value %q\n", k, v)
	}
	config := Config{
		PostAuthorUri: os.Getenv("postAuthorUri"),
		PostBookUri:   os.Getenv("postBookUri"),
		AuthorsUri:    os.Getenv("authorsUri"),
		BooksUri:      os.Getenv("booksUri"),
		Token:         string(r.Header.Get("Authorization")),
	}
	json.NewEncoder(w).Encode(config)
}
func getAuthors(w http.ResponseWriter, r *http.Request) {
	req, err := http.NewRequest("GET", os.Getenv("extauthorsUri"), nil)
	req.Header.Add("X-Id-Token", string(r.Header.Get("X-Id-Token")))
	req.Header.Add("X-Userinfo", string(r.Header.Get("X-Userinfo")))
	req.Header.Add("Authorization", string(r.Header.Get("Authorization")))
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Printf("there was an error performing the http request +%+v", err)
	}
	body, _ := ioutil.ReadAll(resp.Body)
	json.NewEncoder(w).Encode(string(body))
}
func postAuthor(w http.ResponseWriter, r *http.Request) {
	reqBody, _ := ioutil.ReadAll(r.Body)
	req, err := http.NewRequest("POST", os.Getenv("extpostAuthorUri"), bytes.NewBuffer(reqBody))
	req.Header.Add("X-Id-Token", string(r.Header.Get("X-Id-Token")))
	req.Header.Add("X-Userinfo", string(r.Header.Get("X-Userinfo")))
	req.Header.Add("Authorization", string(r.Header.Get("Authorization")))
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Printf("there was an error performing the http request +%+v", err)
	}
	body, _ := ioutil.ReadAll(resp.Body)
	json.NewEncoder(w).Encode(string(body))
}

func getBooks(w http.ResponseWriter, r *http.Request) {
	req, err := http.NewRequest("GET", os.Getenv("extbooksUri"), nil)
	req.Header.Add("X-Id-Token", string(r.Header.Get("X-Id-Token")))
	req.Header.Add("X-Userinfo", string(r.Header.Get("X-Userinfo")))
	req.Header.Add("Authorization", string(r.Header.Get("Authorization")))
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Printf("there was an error performing the http request +%+v", err)
	}
	body, _ := ioutil.ReadAll(resp.Body)
	json.NewEncoder(w).Encode(string(body))
}
func postBook(w http.ResponseWriter, r *http.Request) {
	reqBody, _ := ioutil.ReadAll(r.Body)
	req, err := http.NewRequest("POST", os.Getenv("extpostBookUri"), bytes.NewBuffer(reqBody))
	req.Header.Add("X-Id-Token", string(r.Header.Get("X-Id-Token")))
	req.Header.Add("X-Userinfo", string(r.Header.Get("X-Userinfo")))
	req.Header.Add("Authorization", string(r.Header.Get("Authorization")))
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Printf("there was an error performing the http request +%+v", err)
	}
	body, _ := ioutil.ReadAll(resp.Body)
	json.NewEncoder(w).Encode(string(body))
}

func callapi(w http.ResponseWriter, r *http.Request, url string, passthrough bool) {

	for k, v := range r.Header {
		log.Printf("Header field %q, Value %q\n", k, v)
	}
	req, err := http.NewRequest("GET", url, nil)
	if passthrough {
		req.Header.Add("X-Id-Token", string(r.Header.Get("X-Id-Token")))
		req.Header.Add("X-Userinfo", string(r.Header.Get("X-Userinfo")))
		req.Header.Add("Authorization", string(r.Header.Get("Authorization")))
	}

	log.Printf("Security header for the request %s", req.Header.Get("Authorization"))
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Printf("there was an error performing the http request +%+v", err)
	}
	myresp := "Cookie input : "
	for _, cookie := range r.Cookies() {
		fmt.Println("Found a cookie named:", cookie.Name)
		log.Printf("Found a cookie named: %s", cookie.Name)
		myresp = myresp + cookie.Name
	}

	myresp = myresp + " Cookie output : "
	for _, cookie := range resp.Cookies() {
		fmt.Println("Found a cookie named:", cookie.Name)

		myresp = myresp + cookie.Name
	}

	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	myresp = myresp + " API callback : " + string(body)
	json.NewEncoder(w).Encode(myresp)
}
func handleRequests() {
	router := mux.NewRouter()
	sapi := router.PathPrefix("/api/").Subrouter()
	// Where ORIGIN_ALLOWED is like `scheme://dns[:port]`, or `*` (insecure)
	//headersOk := handlers.AllowedHeaders([]string{"X-Requested-With", "Authorization", "SignIn"})
	//originsOk := handlers.AllowedOrigins([]string{"*", "http://localhost:3001"})
	//methodsOk := handlers.AllowedMethods([]string{"GET", "HEAD", "POST", "PUT", "OPTIONS"})

	c := cors.New(cors.Options{
		AllowedOrigins: []string{"*"},                                     // All origins
		AllowedMethods: []string{"GET", "HEAD", "POST", "PUT", "OPTIONS"}, // Allowing only get, just an example
	})

	sapi.HandleFunc("/api1", callapi1)
	sapi.HandleFunc("/api2", callapi2)
	sapi.HandleFunc("/config", getConfig)
	sapi.HandleFunc("/authors", getAuthors)
	sapi.HandleFunc("/author", postAuthor).Methods("POST")
	sapi.HandleFunc("/books", getBooks)
	sapi.HandleFunc("/book", postBook).Methods("POST")
	log.Fatal(http.ListenAndServe(":3000", c.Handler(router)))
}

func main() {
	handleRequests()
}
