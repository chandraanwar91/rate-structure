package transformers

import (
	"misteraladin.com/jasmine/rate-structure/models"
	"time"
)

type (
	RateBackoffice struct {
		ID				int    `json:"id" gorm:"type:varchar(255)"`
		BookingStart 	string `json:"booking_start" gorm:"type:varchar(255)"`
		BookingEnd 		string `json:"booking_end" gorm:"type:varchar(255)"`
		StayingStart 	string `json:"staying_start" gorm:"type:varchar(255)"`
		StayingEnd 		string `json:"staying_end" gorm:"type:varchar(255)"`
		DiscountType 	string `json:"discount_type" gorm:"type:int(10)"`
		Currency 		string  `json:"currency" gorm:"type:ENUM('IDR','USD')"`
		Nominal 		float64 `json:"nominal" gorm:"type:decimal(12,2)"`
		AvailableTo 	string `json:"available_to" gorm:"type:set('direct_contract','expedia','hotelbeds')"`
		ApplicableTo 	string `json:"applicable_to" gorm:"type:ENUM('apps','member')"`
		CountryID		string `json:"country_id" gorm:"type:text"`
		CityID			string `json:"city_id" gorm:"type:text"`
		HotelGroupID	string `json:"hotel_group_id" gorm:"type:text"`
		HotelID			string `json:"hotel_id" gorm:"type:text"`
		StarRating		string `json:"star_rating" gorm:"type:text"`
		HotelTagID		string `json:"hotel_tag_id" gorm:"type:text"`
		Description		string `json:"description" gorm:"type:text"`
		Status			string `json:"status" gorm:"type:enum('enable','disable','removed')"`
		CreatedBy       int    `json:"created_by" gorm:"column:created_by"`
		ModifiedBy      int    `json:"modified_by" gorm:"column:modified_by"`
		CreatedAt string `json:"created_at"`
		UpdatedAt string `json:"updated_at"`
	}
)

func (res *Transformer) Transform(rate *models.RateBackoffice) *Transformer{
	res.Data = assignRate(rate)
	return res
}

func (res *CollectionTransformer) TransformCollection(rates []*models.RateBackoffice, pagination *models.Pagination) {
	for _, rate := range rates {
		res.Data = append(res.Data, assignRate(rate))
	}

	res.Meta = models.Meta{Pagination: pagination}
}

func assignRate(rate *models.RateBackoffice) interface{} {
	var date = "2006-01-02"
	result := RateBackoffice{}
	result.ID 			= rate.ID
	result.BookingStart = rate.BookingStart.Format(date)
	result.BookingEnd	= rate.BookingEnd.Format(date)
	result.StayingStart = rate.StayingStart.Format(date)
	result.StayingEnd	= rate.StayingEnd.Format(date)
	result.DiscountType = rate.DiscountType
	result.Currency		= rate.Currency
	result.Nominal		= rate.Nominal
	result.AvailableTo  = rate.AvailableTo
	result.ApplicableTo = rate.ApplicableTo
	result.CountryID    = rate.CountryID
	result.CityID		= rate.CityID
	result.HotelGroupID = rate.HotelGroupID
	result.HotelID		= rate.HotelID
	result.StarRating	= rate.StarRating
	result.HotelTagID   = rate.HotelTagID
	result.Description	= rate.Description
	result.Status		= rate.Status
	result.CreatedBy	= rate.CreatedBy
	result.ModifiedBy	= rate.ModifiedBy
	result.CreatedAt = rate.CreatedAt.Format(time.RFC3339)
	result.UpdatedAt = rate.UpdatedAt.Format(time.RFC3339)
	return result
}

