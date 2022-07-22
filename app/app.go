package app

import (
	"net/http"
)

func Start() {
	http.HandleFunc("/greet", greet)
	http.HandleFunc("/customers", getCustomers)
	http.ListenAndServe("localhost:8001", nil)
}
