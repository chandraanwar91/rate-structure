package hotel

import (
	"misteraladin.com/jasmine/rate-structure/models"
)

type IHotelRepository interface {
	GetHotel(countryID string,cityID string,hotelGroupID string,hotelID string,starRating string,hotelTagID string) ([]models.Hotel,error)
}