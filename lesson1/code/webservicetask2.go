package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

// START OMIT
type Reponse struct {
	Code   int    `json:"code"`
	Result string `json:"result"`
}

func hello(w http.ResponseWriter, r *http.Request) {
	res := Reponse{
		Code:   200,
		Result: fmt.Sprintf("%s %s", "hello", r.FormValue("name")),
	}

	json, err := json.Marshal(res)
	if err != nil {
		log.Fatal(err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(json)
}

// END OMIT
func main() {
	// Add ads the function thats going to handle that response
	http.HandleFunc("/", hello)
	// Starts the web server
	// The first argument in this method is the port you want your server to run on
	// The second is a handler. However we have already added this in the line above so we just pass in nil
	http.ListenAndServe(":8000", nil)
}
