package repository

import (
	"strconv"
	"misteraladin.com/jasmine/rate-structure/models"
	"time"
	RateExInterface "misteraladin.com/jasmine/rate-structure/app/rate-extranet"
	"errors"
	gorm1 "github.com/jinzhu/gorm"
	"misteraladin.com/jasmine/rate-structure/helpers"

)

type rateExRepository struct {
	Conn *gorm1.DB
}

func NewRateEXRepository(Conn *gorm1.DB) RateExInterface.IRateExtranetRepository {

	return &rateExRepository{Conn}
}

func (m *rateExRepository) Store (a *models.RateExtranet) (*models.RateExtranet,error) {
	var err error

	tx := m.Conn.Begin()
	a.CreatedAt = time.Now()
	if err = tx.Create(&a).Error; err != nil {
		tx.Rollback()
		return nil,err
	}
	tx.Commit()

	return a,nil
} 

func (m *rateExRepository) GetByDate(rateParam *models.RateExtranet,id uint64) ([]models.RateExtranet,error){
	var(
		date = "2006-01-02"
		BookingStart = rateParam.BookingStart.Format(date)
		BookingEnd	 = rateParam.BookingEnd.Format(date)
		TravelStart  = rateParam.TravelStart.Format(date)
		TravelEnd	 = rateParam.TravelEnd.Format(date)
		RoomID		 = rateParam.RoomID
		ApplicableTo = rateParam.ApplicableTo
		err error
	)
	var a []models.RateExtranet
	tx := m.Conn.Begin()
	if id > 0 {
		if err = tx.Raw("SELECT * FROM rate_plan_extranet WHERE ( (booking_start <= (?) AND booking_end >= (?)) OR (booking_start BETWEEN (?) AND (?)) OR (booking_end BETWEEN (?) AND (?)) )  AND ( (travel_start <= (?) AND travel_end >= (?)) OR (travel_start BETWEEN (?) AND (?)) OR (travel_end BETWEEN (?) AND (?)) ) and room_id = ? AND applicable_to = ? AND status = ? AND id <> ?",BookingStart,BookingEnd,BookingStart,BookingEnd,BookingStart,BookingEnd,TravelStart,TravelEnd,TravelStart,TravelEnd,TravelStart,TravelEnd,RoomID,ApplicableTo,"active",id).Scan(&a).Error; err != nil {
			tx.Rollback()
			return a,err
		}
	
		tx.Commit()
	}else{
		if err = tx.Raw("SELECT * FROM rate_plan_extranet WHERE ( (booking_start <= (?) AND booking_end >= (?)) OR (booking_start BETWEEN (?) AND (?)) OR (booking_end BETWEEN (?) AND (?)) )  AND ( (travel_start <= (?) AND travel_end >= (?)) OR (travel_start BETWEEN (?) AND (?)) OR (travel_end BETWEEN (?) AND (?)) ) and room_id = ?  AND applicable_to = ? AND status = ?",BookingStart,BookingEnd,BookingStart,BookingEnd,BookingStart,BookingEnd,TravelStart,TravelEnd,TravelStart,TravelEnd,TravelStart,TravelEnd,RoomID,ApplicableTo,"active").Scan(&a).Error; err != nil {
			tx.Rollback()
			return a,err
		}
	
		tx.Commit()
	}

	return a,nil
}

func (m *rateExRepository) GetById(id uint64) (*models.RateExtranet, error) {
	var (
		rate   models.RateExtranet
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

func (m *rateExRepository) Update(data *models.RateExtranet,r *models.RateExtranet) (*models.RateExtranet,error) {
	var err error

	r.UpdatedAt 	= time.Now()
	r.BookingStart 	= data.BookingStart
	r.BookingEnd 	= data.BookingEnd
	r.TravelStart	= data.TravelStart
	r.TravelEnd 	= data.TravelEnd
	r.RoomID		= data.RoomID
	r.RoomType		= data.RoomType
	r.HotelID		= data.HotelID
	r.DiscountType 	= data.DiscountType
	r.Currency		= data.Currency
	r.Nominal		= data.Nominal
	r.CommissionAdded = data.CommissionAdded
	r.BookingDays	= data.BookingDays
	r.TravelDays	= data.TravelDays
	r.MinimumStay	= data.MinimumStay
	r.CancellationPolicy	= data.CancellationPolicy
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

func (m *rateExRepository) Fetch(page, perPage int,ApplicableTo string,bookingStart string,bookingEnd string,travelStart string,travelEnd string,status string,hotelID string,roomID string,discountType string) ([]*models.RateExtranet,int, error) {
	var (
		rates   []*models.RateExtranet
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
	txCount = txCount.Table("rate_plan_extranet").Where("applicable_to = ? ",ApplicableTo).Scopes(models.ScopeRatePlanNotRemoved)
	tx = tx.Where("applicable_to = ? ",ApplicableTo).Scopes(models.ScopeRatePlanNotRemoved)

	if(hotelID != ""){
		txCount = txCount.Where("hotel_id = ?",hotelID)
		tx = tx.Where("hotel_id = ?",hotelID)
	}

	if(roomID != ""){
		txCount = txCount.Where("room_id = ?",roomID)
		tx = tx.Where("room_id >= ?",roomID)
	}

	if(discountType != "" && discountType != "all"){
		txCount = txCount.Where("discount_type = ?",discountType)
		tx = tx.Where("discount_type >= ?",discountType)
	}

	if(bookingStart != "" && bookingEnd != ""){
		txCount = txCount.Where("booking_start >= ?",bookingStart).Where("booking_end <= ?",bookingEnd)
		tx = tx.Where("booking_start >= ?",bookingStart).Where("booking_end <= ?",bookingEnd)
	}

	if(travelStart != "" && travelEnd != ""){
		txCount = txCount.Where("travel_start >= ?",travelStart).Where("travel_end <= ?",travelEnd)
		tx = tx.Where("travel_start >= ?",travelStart).Where("travel_end <= ?",travelEnd)
	}

	if(status != "" && status != "all"){
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

func (m *rateExRepository) Delete (a *models.RateExtranet,ModifiedBy int) error {
	var err error

	a.CreatedAt = time.Now()
	a.UpdatedAt 	= time.Now()
	a.Status		= "removed"
	a.ModifiedBy	= ModifiedBy
	tx := m.Conn.Begin()
	if err = tx.Create(&a).Error; err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()

	return nil
} 

func (m *rateExRepository) FindAvailable(params *models.Rate,ApplicableTo string) ([]*models.RateExtranet,int, error) {
	var (
		date = "2006-01-02"
		checkIn  = params.CheckIn.Format(date)
		checkOut = params.CheckOut.Format(date)
		bookingDate = params.BookingDate.Format(date)
		day = helpers.GetDayOffWeek(bookingDate)
		dayStr = strconv.Itoa(day)
		checkInDay = helpers.GetDayOffWeek(checkIn)
		checkInDayStr = strconv.Itoa(checkInDay)
		minStay = helpers.CalNight(checkIn,checkOut)
		roomIds = helpers.Explode(params.RoomIds,",")
		rates   []*models.RateExtranet
		err error
	)
	//Initialization
	tx := m.Conn.Begin()

	tx = tx.Table("rate_plan_extranet").Where("travel_days like '%"+checkInDayStr+"%'").Where("'"+checkIn+"' BETWEEN travel_start and travel_end").Where("'"+checkOut+"' BETWEEN travel_start and travel_end").Where("'"+bookingDate+"' BETWEEN booking_start and booking_end").Where("booking_days like '%"+dayStr+"%'").Where("minimum_stay <= ?",minStay).Where("room_id IN (?)",roomIds).Where("applicable_to = ?",ApplicableTo).Scopes(models.ScopeRatePlanActive).Group("room_id")

	if err = tx.Find(&rates).Error; err != nil {
		tx.Rollback()
		return nil,0,err
	}

	tx.Commit()

	total := len(rates)

	return rates,total, err
}