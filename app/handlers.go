package app

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"github.com/anandkthirtha/banking/service"
	"github.com/gorilla/mux"
	"net/http"
)

type Customer struct {
	Name    string `json:"full_name" xml:"name"`
	City    string `json:"city" xml:"city"`
	Zipcode int    `json:"zipcode" zipcode:"zipcode"`
}

//func greet(w http.ResponseWriter, r *http.Request) {
//	fmt.Fprintf(w, "hello world")
//}
type CustomerHandlers struct {
	service service.CustomerService
}
func (ch *CustomerHandlers) getCustomers(w http.ResponseWriter, r *http.Request) {

	//customers := []Customer{
//		{Name: "Anand", City: "blr", Zipcode: 560098},
//		{Name: "Pranav", City: "blr", Zipcode: 560098},
//	}

	customers, _ :=ch.service.GetAllCustomer()

	if r.Header.Get("Content-Type") == "application/xml" {
		w.Header().Set("Content-Type", "application/xml")
		xml.NewEncoder(w).Encode(customers)
	} else {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(customers)
	}
}

func (ch *CustomerHandlers) getCustomer(w http.ResponseWriter, r *http.Request) {
	vars:=mux.Vars(r)
	id:=vars["customer_id"]

	customer,err:=ch.service.GetCustomer(id)
	if err!=nil{
		w.WriteHeader(err.Code)
		fmt.Fprintf(w,err.Message)
	} else {
		w.Header().Add("Content-Type", "application/json")
		json.NewEncoder(w).Encode(customer)
	}
}
