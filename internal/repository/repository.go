package repository

import app "github.com/KudryavtsevIvan/city_information_service"

type CityList interface {
	Create(city app.CityRequest) (string, error)
	Delete(id int) error
	SetPopulation(id, population int) error
	GetFromRegion(region string) ([]string, error)
	GetFromDistrict(district string) ([]string, error)
	GetFromPopulation(start, end int) ([]string, error)
	GetFromFoundation(start, end int) ([]string, error)
	findCities(idx_type int, searchText string) []string
	GetFull(id int) (*app.City, error)
}

type Repository struct {
	CityList
}

func NewRepository(db *DataBase) *Repository {
	return &Repository{
		CityList: NewCityListDB(db),
	}
}