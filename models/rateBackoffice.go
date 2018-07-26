package models

import(
	"time"
	gorm2 "github.com/jinzhu/gorm"
	"misteraladin.com/jasmine/rate-structure/helpers"
)
type (
	RateBackoffice struct {
		BaseModel
		BookingStart 	time.Time  `json:"booking_start" gorm:"type:date";column:booking_start`
		BookingEnd 		time.Time  `json:"booking_end" gorm:"type:date";column:booking_end`
		StayingStart 	time.Time  `json:"staying_start" gorm:"type:date";column:staying_start`
		StayingEnd 		time.Time  `json:"staying_end" gorm:"type:date";column:staying_end`
		DiscountType 	string  `json:"discount_type" gorm:"type:ENUM('amount','percentage')"`
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
	}
)

func (RateBackoffice) TableName() string {
	return "rate_backoffice"
}

func ScopeActive(db *gorm2.DB) *gorm2.DB {
	return db.Where("status = ?", "enable")
}

func scopeGetEvent(identifier string) func(db *gorm2.DB) *gorm2.DB{
	return func (db *gorm2.DB) *gorm2.DB {
		if helpers.IsNumeric(identifier){
			return db.Where("id = ?",identifier)
		}

		return db.Where("slug = ?",identifier)
	}
}

