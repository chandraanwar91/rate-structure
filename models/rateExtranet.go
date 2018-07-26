package models

import(
	"time"
	gorm2 "github.com/jinzhu/gorm"
)
type (
	RateExtranet struct {
		BaseModel
		BookingStart 	time.Time  `json:"booking_start" gorm:"type:date";column:booking_start`
		BookingEnd		time.Time  `json:"booking_end" gorm:"type:date";column:booking_end`
		TravelStart 	time.Time  `json:"travel_start" gorm:"type:date";column:travel_start`
		TravelEnd 		time.Time  `json:"travel_end" gorm:"type:date";column:travel_end`
		DiscountType 	string  `json:"discount_type" gorm:"type:ENUM('amount','percentage')"`
		Currency 		string  `json:"currency" gorm:"type:ENUM('IDR','USD')"`
		Nominal 		float64 `json:"nominal" gorm:"type:decimal(12,2)"`
		CommissionAdded float64 `json:"commission_added" gorm:"type:decimal(12,2)"`
		BookingDays 	string `json:"booking_days" gorm:"type:set('1','2','3','4','5','6','7')";column:booking_days`
		TravelDays 		string `json:"travel_days" gorm:"type:set('1','2','3','4','5','6','7')";column:travel_days`
		RoomID			string `json:"room_id" gorm:"type:text"`
		RoomType		string	`json:"room_type" gorm:"type:varchar(255)"`
		HotelID			string	`json:"hotel_id" gorm:"type:varchar(255)"`
		MinimumStay		string `json:"minimum_stay" gorm:"type:text"`
		CancellationPolicy	string `json:"cancellation_policy" gorm:"type:set('non_refundable','follow_hotel_policy')"`
		ApplicableTo 	string `json:"applicable_to" gorm:"type:ENUM('apps','member')"`
		Status			string `json:"status" gorm:"type:enum('active','disabled')"`
		CreatedBy       int    `json:"created_by" gorm:"column:created_by"`
		ModifiedBy      int    `json:"modified_by" gorm:"column:modified_by"`
	}
)

func (RateExtranet) TableName() string {
	return "rate_plan_extranet"
}

func ScopeRatePlanActive(db *gorm2.DB) *gorm2.DB {
	return db.Where("status = ?", "active")
}

func ScopeRatePlanNotRemoved(db *gorm2.DB) *gorm2.DB {
	return db.Where("status <> ?", "removed")
}

