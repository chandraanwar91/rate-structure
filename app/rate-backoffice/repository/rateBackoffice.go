package repository

import (
	"misteraladin.com/jasmine/rate-structure/models"
	//"misteraladin.com/jasmine/rate-structure/helpers"
	"time"
	RateBoInterface "misteraladin.com/jasmine/rate-structure/app/rate-backoffice"
	"errors"
	gorm1 "github.com/jinzhu/gorm"
)
type rateBoRepository struct {
	Conn *gorm1.DB
}

func NewRateBORepository(Conn *gorm1.DB) RateBoInterface.IRateBackofficeRepository {

	return &rateBoRepository{Conn}
}

func (m *rateBoRepository) Store (a *models.RateBackoffice) (*models.RateBackoffice,error) {
	var err error

	a.CreatedAt = time.Now()
	tx := m.Conn.Begin()
	if err = tx.Create(&a).Error; err != nil {
		tx.Rollback()
		return nil,err
	}
	tx.Commit()

	return a,nil
} 

func (m *rateBoRepository) GetByDate(bookingStart string,bookingEnd string,stayingStart string,stayingEnd string,id uint64) ([]models.RateBackoffice,error){
	var err error
	var a []models.RateBackoffice
	tx := m.Conn.Begin()
	if id > 0 {
		if err = tx.Raw("SELECT * FROM rate_backoffice WHERE ( (booking_start <= (?) AND booking_end >= (?)) OR (booking_start BETWEEN (?) AND (?)) OR (booking_end BETWEEN (?) AND (?)) )  AND ( (staying_start <= (?) AND staying_end >= (?)) OR (staying_start BETWEEN (?) AND (?)) OR (staying_end BETWEEN (?) AND (?)) ) AND id <> (?)",bookingStart,bookingEnd,bookingStart,bookingEnd,bookingStart,bookingEnd,stayingStart,stayingEnd,stayingStart,stayingEnd,stayingStart,stayingEnd,id).Scan(&a).Error; err != nil {
			tx.Rollback()
			return a,err
		}
	
		tx.Commit()
	}else{
		if err = tx.Raw("SELECT * FROM rate_backoffice WHERE ( (booking_start <= (?) AND booking_end >= (?)) OR (booking_start BETWEEN (?) AND (?)) OR (booking_end BETWEEN (?) AND (?)) )  AND ( (staying_start <= (?) AND staying_end >= (?)) OR (staying_start BETWEEN (?) AND (?)) OR (staying_end BETWEEN (?) AND (?)) )",bookingStart,bookingEnd,bookingStart,bookingEnd,bookingStart,bookingEnd,stayingStart,stayingEnd,stayingStart,stayingEnd,stayingStart,stayingEnd).Scan(&a).Error; err != nil {
			tx.Rollback()
			return a,err
		}
	
		tx.Commit()
	}

	return a,nil
}

func (m *rateBoRepository) GetById(id uint64) (*models.RateBackoffice, error) {
	var (
		rate   models.RateBackoffice
		err        error
	)
	
	//Initialization
	tx := m.Conn.Begin()
	err = tx.First(&rate,id).Error
	if err != nil {
		tx.Rollback()
		return &rate, nil
	}
	tx.Commit()


	if rate.ID == 0 {
		return nil, errors.New("ID not found")
	}


	return &rate, nil
}

func (m *rateBoRepository) Update(data *models.RateBackoffice,r *models.RateBackoffice) (*models.RateBackoffice,error) {
	var err error

	r.UpdatedAt 	= time.Now()
	r.AvailableTo   = data.AvailableTo
	r.BookingStart 	= data.BookingStart
	r.BookingEnd 	= data.BookingEnd
	r.StayingStart	= data.StayingStart
	r.StayingEnd 	= data.StayingEnd
	r.DiscountType 	= data.DiscountType
	r.Currency		= data.Currency
	r.Nominal		= data.Nominal
	r.CountryID		= data.CountryID
	r.CityID		= data.CityID
	r.HotelGroupID	= data.HotelGroupID
	r.HotelID		= data.HotelID
	r.StarRating	= data.StarRating
	r.HotelTagID	= data.HotelTagID
	r.Description	= data.Description
	r.Status		= data.Status
	r.ModifiedBy	= data.ModifiedBy
	tx := m.Conn.Begin()
	if err = tx.Save(&r).Error; err != nil {
		tx.Rollback()
		return r,err
	}
	tx.Commit()

	return r,err
}

func (m *rateBoRepository) Fetch(page, perPage int,ApplicableTo string,bookingStart string,bookingEnd string,stayingStart string,stayingEnd string,status string) ([]*models.RateBackoffice,int, error) {
	var (
		rates   []*models.RateBackoffice
		total      int
		err        error
	)

	if page < 1 {
		page = 1
	}

	if perPage < 1 {
		perPage = 10
	}

	offset := (page * perPage) - perPage
	//Initialization
	tx := m.Conn.Begin()

	txCount := tx

	//Count All
	//txCount.Table("rate_backoffice").Where("applicable_to = ? ",ApplicableTo).Scopes(models.ScopeActive).Count(&total)
	txCount = txCount.Table("rate_backoffice").Where("applicable_to = ? ",ApplicableTo)
	tx = tx.Where("applicable_to = ? ",ApplicableTo)

	if(bookingStart != "" && bookingEnd != ""){
		txCount = txCount.Where("booking_start >= ?",bookingStart).Where("booking_end <= ?",bookingEnd)
		tx = tx.Where("booking_start >= ?",bookingStart).Where("booking_end <= ?",bookingEnd)
	}

	if(stayingStart != "" && stayingEnd != ""){
		txCount = txCount.Where("staying_start >= ?",stayingStart).Where("staying_end <= ?",stayingEnd)
		tx = tx.Where("staying_start >= ?",stayingStart).Where("staying_end <= ?",stayingEnd)
	}

	if(status != ""){
		txCount = txCount.Where("status = ?",status)
		tx = tx.Where("status = ? ",status)
	}
	txCount.Count(&total)
	err = tx.Limit(perPage).Offset(offset).Find(&rates).Error
	if err != nil {
		tx.Rollback()
		return rates,total, err
	}
	tx.Commit()

	return rates,total, err
}

func (m *rateBoRepository) FindAvailable(params *models.Rate,ApplicableTo string) ([]*models.RateBackoffice,int, error) {
	var (
		date = "2006-01-02"
		checkIn  = params.CheckIn.Format(date)
		checkOut = params.CheckOut.Format(date)
		bookingDate = params.BookingDate.Format(date)
		rates   []*models.RateBackoffice
		err error
	)

	//Initialization
	tx := m.Conn.Begin()

	tx = tx.Table("rate_backoffice").Where("'"+checkIn+"' BETWEEN staying_start and staying_end").Where("'"+checkOut+"' BETWEEN staying_start and staying_end").Where("booking_start <= ?",bookingDate).Where("'"+bookingDate+"' BETWEEN booking_start and booking_end").Where("applicable_to = ?",ApplicableTo).Scopes(models.ScopeActive)

	if err = tx.Find(&rates).Error; err != nil {
		tx.Rollback()
		return nil,0,err
	}

	tx.Commit()

	total := len(rates)

	return rates,total, err
}