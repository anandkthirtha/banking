package app

import (
	"log"
	"net/http"

	"github.com/anandkthirtha/banking/domain"
	"github.com/anandkthirtha/banking/service"

	"github.com/gorilla/mux"
)

func Start() {

	router := mux.NewRouter()

	//ch:= CustomerHandlers{service.NewCustomerService(domain.NewCustomerRepositoryStub())}
	ch := CustomerHandlers{service.NewCustomerService(domain.NewCustomerRepositoryDb())}
	//	router.HandleFunc("/greet", greet).Methods(http.MethodGet)
	router.HandleFunc("/customers", ch.getCustomers).Methods(http.MethodGet)
	router.HandleFunc("/customers/{customer_id[0-9]+}", ch.getCustomer).Methods(http.MethodGet)
	//	router.HandleFunc("/customers", createCustomer).Methods(http.MethodPost)

	//	router.HandleFunc("/customers/{customer_id:[0-9]+}", getCustomer).Methods(http.MethodGet)

	log.Fatal(http.ListenAndServe("localhost:8001", router))
}

//func createCustomer(w http.ResponseWriter, r *http.Request) {
//	fmt.Fprint(w, "post request received")
//}

//func getCustomer(w http.ResponseWriter, r *http.Request) {
//	vars := mux.Vars(r)
//	fmt.Fprint(w, vars["customer_id"])
//}
