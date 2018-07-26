package RateBackoffice

import (
	"misteraladin.com/jasmine/rate-structure/models"
)

type IRateBackofficeRepository interface {
	Store(m *models.RateBackoffice) (*models.RateBackoffice,error)
	Update(data *models.RateBackoffice,m *models.RateBackoffice) (*models.RateBackoffice,error)
	GetByDate(bookingStart string,bookingEnd string,stayingStart string,stayingEnd string,id uint64) ([]models.RateBackoffice,error)
	GetById(id uint64) (*models.RateBackoffice, error)
	Fetch(page, perPage int,ApplicableTo string,bookingStart string,bookingEnd string,stayingStart string,stayingEnd string,status string) ([]*models.RateBackoffice,int, error)
	FindAvailable(params *models.Rate,ApplicableTo string) ([]*models.RateBackoffice,int, error)
}