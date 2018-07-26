package RateExtranet

import (
	"misteraladin.com/jasmine/rate-structure/models"
)

type IRateExtranetRepository interface {
	Store(m *models.RateExtranet) (*models.RateExtranet,error)
	Update(data *models.RateExtranet,m *models.RateExtranet) (*models.RateExtranet,error)
	GetByDate(rateParam *models.RateExtranet,id uint64) ([]models.RateExtranet,error)
	GetById(id uint64) (*models.RateExtranet, error)
	Fetch(page, perPage int,ApplicableTo string,bookingStart string,bookingEnd string,travelStart string,travelEnd string,status string,hotelID string,roomID string,discountType string) ([]*models.RateExtranet,int, error)
	Delete(data *models.RateExtranet,ModifiedBy int) error
	FindAvailable(m *models.Rate,ApplicableTo string) ([]*models.RateExtranet,int,error)
}