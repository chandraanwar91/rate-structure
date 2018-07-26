package transformers

import (
	"misteraladin.com/jasmine/rate-structure/models"
	"time"
)

type (
	RateExtranet struct {
		ID				int    `json:"id" gorm:"type:varchar(255)"`
		BookingStart 	string `json:"booking_start" gorm:"type:varchar(255)"`
		BookingEnd 		string `json:"booking_end" gorm:"type:varchar(255)"`
		TravelStart 	string `json:"travel_start" gorm:"type:varchar(255)"`
		TravelEnd 		string `json:"travel_end" gorm:"type:varchar(255)"`
		DiscountType 	string `json:"discount_type" gorm:"type:int(10)"`
		Currency 		string  `json:"currency" gorm:"type:ENUM('IDR','USD')"`
		Nominal 		float64 `json:"nominal" gorm:"type:decimal(12,2)"`
		CommissionAdded float64 `json:"commission_added" gorm:"type:decimal(12,2)"`
		BookingDays 	string `json:"booking_days"`
		TravelDays 		string `json:"travel_days"`
		ApplicableTo 	string `json:"applicable_to" gorm:"type:ENUM('apps','member')"`
		MinimumStay		string `json:"minimum_stay"`
		CancellationPolicy	string `json:"cancellation_policy"`
		RoomID			string `json:"room_id" gorm:"column:room_id"`
		RoomType		string `json:"room_type" gorm:"column:room_type"`
		Status			string `json:"status"`
		CreatedBy       int    `json:"created_by" gorm:"column:created_by"`
		ModifiedBy      int    `json:"modified_by" gorm:"column:modified_by"`
		CreatedAt string `json:"created_at"`
		UpdatedAt string `json:"updated_at"`
	}
)

func (res *Transformer) TransformExtranet(rate *models.RateExtranet) *Transformer{
	res.Data = assignRateExtranet(rate)
	return res
}

func (res *CollectionTransformer) TransformExtranetCollection(rates []*models.RateExtranet, pagination *models.Pagination) {
	for _, rate := range rates {
		res.Data = append(res.Data, assignRateExtranet(rate))
	}

	res.Meta = models.Meta{Pagination: pagination}
}

func assignRateExtranet(rate *models.RateExtranet) interface{} {
	var date = "2006-01-02"
	result := RateExtranet{}
	result.ID 			= rate.ID
	result.BookingStart = rate.BookingStart.Format(date)
	result.BookingEnd	= rate.BookingEnd.Format(date)
	result.TravelStart 	= rate.TravelStart.Format(date)
	result.TravelEnd	= rate.TravelEnd.Format(date)
	result.RoomID 		= rate.RoomID
	result.RoomType		= rate.RoomType
	result.DiscountType = rate.DiscountType
	result.Currency		= rate.Currency
	result.Nominal		= rate.Nominal
	result.CommissionAdded = rate.CommissionAdded
	result.BookingDays  = rate.BookingDays
	result.TravelDays   = rate.TravelDays
	result.ApplicableTo = rate.ApplicableTo
	result.MinimumStay	= rate.MinimumStay
	result.CancellationPolicy = rate.CancellationPolicy
	result.Status		= rate.Status
	result.CreatedBy	= rate.CreatedBy
	result.ModifiedBy	= rate.ModifiedBy
	result.CreatedAt = rate.CreatedAt.Format(time.RFC3339)
	result.UpdatedAt = rate.UpdatedAt.Format(time.RFC3339)
	return result
}

