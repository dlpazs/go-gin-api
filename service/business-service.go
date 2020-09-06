package service

import "github.com/dlpazs/golang-gin-api/entity"

type BusinessService interface {
	Save(entity.Business) entity.Business
	FindAll() []entity.Business
}

type businessService struct {
	businesses []entity.Business
}

func New() BusinessService {
	return &businessService{}
}

func (service *businessService) Save(business entity.Business) entity.Business {
	service.businesses = append(service.businesses, business)
	return business
}

func (service *businessService) FindAll() []entity.Business {
	return service.businesses
}
