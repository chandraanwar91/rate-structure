package repository

import (
	//"fmt"
	"misteraladin.com/jasmine/rate-structure/models"
	"misteraladin.com/jasmine/rate-structure/helpers"
	hotelInterface "misteraladin.com/jasmine/rate-structure/app/hotel"
	gorm1 "github.com/jinzhu/gorm"
)
type hotelRepository struct {
	Conn *gorm1.DB
}

func NewMysqlHotelRepository(Conn *gorm1.DB) hotelInterface.IHotelRepository {
	return &hotelRepository{Conn}
}

func (m *hotelRepository) GetHotel(countryID string,cityID string,hotelGroupID string,hotelID string,starRating string,hotelTagID string) ([]models.Hotel,error) {
	var	hotels   []models.Hotel
	var err error
	tx := m.Conn.Begin()

	tx = tx.Table("hotel").Select("hotel.hotel_id,hotel.name").Joins("left join hotel_group on hotel_group.group_id = hotel.group_id").Joins("left join join_hotel_tag on join_hotel_tag.hotel_id = hotel.hotel_id").Joins("left join hotel_tag on hotel_tag.tag_id = join_hotel_tag.tag_id").Joins("join location_area on location_area.area_id = hotel.area_id").Joins("join location_city on location_city.city_id = location_area.city_id").Joins("join location_state on location_state.state_id = location_city.state_id").Joins("join location_country on location_country.country_id = location_state.country_id").Where("hotel.approval_status = ?","approved").Where("hotel.status = ?","enable")

	//checkif country not empty
	if r := helpers.Empty(countryID); !r {
		countryExp := helpers.Explode(countryID,",")
		tx = tx.Where("location_country.country_id IN (?)",countryExp)
	}

	//check if city not empty
	if r := helpers.Empty(cityID); !r {
		cityExp := helpers.Explode(cityID,",")
		tx = tx.Where("location_city.city_id IN (?)",cityExp)
	}

	//check if hotel group not empty
	if r := helpers.Empty(hotelGroupID); !r {
		groupExp := helpers.Explode(hotelGroupID,",")
		tx = tx.Where("hotel.group_id IN (?)",groupExp)
	}

	//check if hotel id not empty
	if r := helpers.Empty(hotelID); !r {
		hotelExp := helpers.Explode(hotelID,",")
		tx = tx.Where("hotel.hotel_id IN (?)",hotelExp)
	}

	//check if star not empty
	if r := helpers.Empty(starRating); !r {
		starExp := helpers.Explode(starRating,",")
		tx = tx.Where("hotel.star_rating IN (?)",starExp)
	}

	//check if hotel tag not empty
	if r := helpers.Empty(hotelTagID); !r {
		tagExp := helpers.Explode(hotelTagID,",")
		tx = tx.Where("join_hotel_tag.tag_id IN (?)",tagExp)
	}

	if err = tx.Find(&hotels).Error; err != nil {
		tx.Rollback()
		return nil,err
	}

	tx.Commit()

	return hotels,err

}