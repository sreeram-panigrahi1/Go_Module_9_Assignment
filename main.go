package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Data struct {
	Name string `json:"name"`
	Age   int64  `json:age`
}

func ReqeustHandler(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	getResponse := Data{
		Name: "Sreeram Panigrahi",
		Age:   22,
	}
	res := Data{}
	switch r.Method {
	case "GET":
		{
			err := json.NewEncoder(w).Encode(getResponse)
			if err != nil {
				fmt.Println("ERROR!")
			}
		}
	case "POST": 
		{
			err := json.NewDecoder(r.Body).Decode(&res)
			if err != nil {
				fmt.Println("ERROR!")
			}
			fmt.Fprintf(w, "input: %+v", res)
		}
	default:
		w.WriteHeader(http.StatusNotFound)
	}
}

func main() {

	http.HandleFunc("/", ReqeustHandler)
	fmt.Println("Starting server at Localhost")
	http.ListenAndServe(":8080", nil)

}
