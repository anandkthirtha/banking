package domain

import "github.com/anandkthirtha/banking/err"

type Customer struct {
	Id string
	Name string
	City string
	Zipcode string
	DateofBirth string
	Status string
}

type CustomerRepository interface {
	FindAll()([]Customer,error)
	ById(string)(*Customer,*err.AppError)
}
