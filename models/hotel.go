package models

import(
	gorm2 "github.com/jinzhu/gorm"
)
type (
	Hotel struct {
		HOTELID        int       `json:"hotel_id" gorm:"column:hotel_id"`
		NAME		   string	`json:"hotel_name" gorm:"column:name"`
	}
)

func (Hotel) TableName() string {
	return "hotel"
}

func scopeHotelActive(db *gorm2.DB) *gorm2.DB {
	return db.Where("status = ?", "enable")
}

