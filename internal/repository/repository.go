package repository

import model "github.com/KudryavtsevIvan/city_information_service/internal/model"

type CityList interface {
	Create(city model.CityRequest) (string, error)
	Delete(id int) error
	SetPopulation(id, population int) error
	GetFromRegion(region string) ([]string, error)
	GetFromDistrict(district string) ([]string, error)
	GetFromPopulation(start, end int) ([]string, error)
	GetFromFoundation(start, end int) ([]string, error)
	findCities(idx_type int, searchText string) []string
	GetFull(id int) (*model.City, error)
}

type Repository struct {
	CityList
}

func NewRepository(db *DataBase) *Repository {
	return &Repository{
		CityList: NewCityListDB(db),
	}
}