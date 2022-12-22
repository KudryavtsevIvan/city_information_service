package service

import (
	model "github.com/KudryavtsevIvan/city_information_service/internal/model"
	"github.com/KudryavtsevIvan/city_information_service/internal/repository"
)

type City interface {
	Create(city model.CityRequest) (string, error)
	Delete(id int) error
	SetPopulation(id, population int) error
	GetFromRegion(region string) ([]string, error)
	GetFromDistrict(district string) ([]string, error)
	GetFromPopulation(start, end int) ([]string, error)
	GetFromFoundation(start, end int) ([]string, error)
	GetFull(id int) (*model.City, error)
}

type Service struct {
	City
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		City: NewCityService(repos.CityList),
	}
}