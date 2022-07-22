package app

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"net/http"
)

type Customer struct {
	Name    string `json:"full_name" xml:"name"`
	City    string `json:"city" xml:"city"`
	Zipcode int    `json:"zipcode" zipcode:"zipcode"`
}

func greet(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hello world")
}

func getCustomers(w http.ResponseWriter, r *http.Request) {

	customers := []Customer{
		{Name: "Anand", City: "blr", Zipcode: 560098},
		{Name: "Pranav", City: "blr", Zipcode: 560098},
	}

	if r.Header.Get("Content-Type") == "application/xml" {
		w.Header().Set("Content-Type", "application/xml")
		xml.NewEncoder(w).Encode(customers)
	} else {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(customers)
	}
}
