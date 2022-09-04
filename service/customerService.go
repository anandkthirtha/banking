package service

import (
	"github.com/anandkthirtha/banking/domain"
	"github.com/anandkthirtha/banking/err"
)

type CustomerService interface {
	GetAllCustomer()([]domain.Customer,error)
	GetCustomer(string)(*domain.Customer,*err.AppError)
}

type DefaultCustomerService struct {
	repo domain.CustomerRepository
}

func (s DefaultCustomerService) GetAllCustomer()([]domain.Customer,error)  {
return s.repo.FindAll()
}

func (s DefaultCustomerService) GetCustomer(id string)(*domain.Customer,*err.AppError)  {
	return s.repo.ById(id)
}
func NewCustomerService(repository domain.CustomerRepository) DefaultCustomerService {
	return DefaultCustomerService{repository}
}