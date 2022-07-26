package domain

import (
	"database/sql"
	"err"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"time"
)


type CustomerRepositoryDb struct {
	client *sql.DB
}

func (d CustomerRepositoryDb) FindAll()([]Customer, error)  {

	findAllSql:="select customer_id, name, city, zipcode, date_of_birth, status from customers"

	rows, err:=d.client.Query(findAllSql)

	if err!= nil {
		log.Println("error while querying cu table" + err.Error())
		return nil, err
	}
customers:=make([]Customer, 0)
	for rows.Next(){
		var c Customer
		err:=rows.Scan(&c.Id, &c.Name, &c.City, &c.Zipcode, &c.DateofBirth, &c.Status)
		if err!= nil {
			log.Println("error while scanning cu table" + err.Error())
			return nil, err
		}
		customers=append(customers,c)
	}
	return customers, nil
}

func (d CustomerRepositoryDb) ById(id string)(*Customer, *err.AppError)  {
	customerSql:="select customer_id, name, city, zipcode, date_of_birth, status from customers where customer_id=?"
	row:=d.client.QueryRow(customerSql, id)
	var c Customer
	err:=row.Scan(&c.Id, &c.Name, &c.City, &c.Zipcode, &c.DateofBirth, &c.Status)
	if err != nil{
		if err == sql.ErrNoRows {
			return nil, err.NewNotFoundError("customer not found")
		}else{
		log.Println("error while scancning customer"+err.Error())
		return nil, err.NewUnexpectedError("unexpected database error")
		}
	}
return &c,nil
}

func NewCustomerRepositoryDb() CustomerRepositoryDb {
	client, err := sql.Open("mysql", "root:Hello#123@tcp(localhost:3306)/banking")
	if err != nil {
		panic(err)
	}
	// See "Important settings" section.
	client.SetConnMaxLifetime(time.Minute * 3)
	client.SetMaxOpenConns(10)
	client.SetMaxIdleConns(10)
	return CustomerRepositoryDb{client}

}
