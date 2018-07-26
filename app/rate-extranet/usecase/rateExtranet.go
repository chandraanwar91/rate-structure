package usecase

import (
	"time"
	"misteraladin.com/jasmine/rate-structure/helpers"
	"strconv"
	rateExInterfaces "misteraladin.com/jasmine/rate-structure/app/rate-extranet"
	"misteraladin.com/jasmine/rate-structure/models"
	"github.com/gin-gonic/gin"
	"errors"
)

type rateExtranetUsecase struct{
	rateExtranetRepo rateExInterfaces.IRateExtranetRepository
}

func NewrateExtranetUsecase(a rateExInterfaces.IRateExtranetRepository) rateExInterfaces.IRateExtranetUseCase {
	return &rateExtranetUsecase{
		rateExtranetRepo: a,
	}
}

func (a *rateExtranetUsecase) Store(c *gin.Context) (*models.RateExtranet, error) {

	rateParam := mapRate(c)

	msg,exist := a.IsAvailable(rateParam,0)
	if exist {
		return nil, errors.New(msg)
	}

	res, err := a.rateExtranetRepo.Store(rateParam)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (a *rateExtranetUsecase) Update(c *gin.Context) (*models.RateExtranet, error) {
	id, _ := strconv.ParseUint(c.Param("id"), 10, 64)

	m, err := a.rateExtranetRepo.GetById(id)

	if(err != nil){
		return nil,errors.New(err.Error())
	}

	rateParam := mapRate(c)

	msg,exist := a.IsAvailable(rateParam,id)
	if exist {
		return nil, errors.New(msg)
	}

	res,err := a.rateExtranetRepo.Update(rateParam,m)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (a *rateExtranetUsecase) GetByID(c *gin.Context) (*models.RateExtranet, error) {
	id, _ := strconv.ParseUint(c.Param("id"), 10, 64)

	if id == 0 || helpers.Empty(c.Param("id")){
		return nil, errors.New("parameter id not be null")
	}


	res, err := a.rateExtranetRepo.GetById(id)
	if err != nil {
		return nil, err
	}

	if res.ID == 0 {
		return nil,errors.New("Data Not Found")
	}

	return res, nil
}

func (a *rateExtranetUsecase) Fetch(c *gin.Context) ([]*models.RateExtranet, *models.Pagination, error) {
	var (
		pagination *models.Pagination
		total      int
	)

	page, _ := strconv.Atoi(c.Query("page"))
	perPage, _ := strconv.Atoi(c.Query("perpage"))
	applicableTo := c.Query("applicable_to")
	status := c.Query("status")
	bookingStart := c.Query("booking_start")
	bookingEnd := c.Query("booking_end")
	travelStart := c.Query("travel_start")
	travelEnd := c.Query("travel_end")
	hotelID := c.Query("hotel_id")
	roomID := c.Query("room_id")
	discountType := c.Query("discount_type")

	if page < 1 {
		page = 1
	}

	if perPage < 1 {
		perPage = 10
	}

	//Get All
	//txCount.Table("rate_backoffice").Where("applicable_to = ? ",ApplicableTo).Scopes(models.ScopeActive).Count(&total)
	rates,total,err := a.rateExtranetRepo.Fetch(page, perPage,applicableTo,bookingStart,bookingEnd,travelStart,travelEnd,status,hotelID,roomID,discountType)

	if(err != nil){
		return nil,nil,err
	}

	count := len(rates)
	pagination = models.BuildPagination(total, page, perPage, count)

	return rates, pagination, err
}

func  (a *rateExtranetUsecase) IsAvailable(rateParam *models.RateExtranet,id uint64) (string,bool) {
	//check source
	var exist = false

	dataExist,err := a.rateExtranetRepo.GetByDate(rateParam,id)

	if(err != nil){
		return err.Error(),true
	}

	count := len(dataExist)

	if(count > 0){
		for _, rate := range dataExist {
			idStr :=  strconv.Itoa(rate.ID)
			var resultBookingDays = helpers.Explode(rate.BookingDays,",")
			var bookingDays = helpers.Explode(rateParam.BookingDays,",")
			var resultTravelDays = helpers.Explode(rate.TravelDays,",")
			var travelDays = helpers.Explode(rateParam.TravelDays,",")

			for _,bookingDay := range bookingDays {
				if r := helpers.InArray(bookingDay,resultBookingDays); r {
					exist =  true
					return "Duplicate Booking Days "+bookingDay+" Configuration ID "+idStr,exist
					break
				}
			}

			for _,travelDay := range travelDays {
				if r := helpers.InArray(travelDay,resultTravelDays); r {
					exist =  true
					return "Duplicate Travel Days "+travelDay+" Configuration ID "+idStr,exist
					break
				}
			}

			if exist {
				exist =  true
				return "Duplicate Configuration ID "+idStr,exist
			}

		}
	}
	return "",exist

	
}

func mapRate(params *gin.Context) *models.RateExtranet{
	nominal, _ := strconv.ParseFloat(params.PostForm("nominal"), 64)
	commission_added, _ := strconv.ParseFloat(params.PostForm("commission_added"), 64)
	createdBy, _ := strconv.Atoi(params.PostForm("created_by"))
	modifiedBy, _ := strconv.Atoi(params.PostForm("modified_by"))
	bookingStart,_ := time.Parse("2006-01-02", params.PostForm("booking_start"))
	bookingEnd,_ := time.Parse("2006-01-02", params.PostForm("booking_end"))
	travelStart,_ := time.Parse("2006-01-02", params.PostForm("travel_start"))
	travelEnd,_ := time.Parse("2006-01-02", params.PostForm("travel_end"))
	rate := new(models.RateExtranet)
	rate.BookingStart 		= bookingStart
	rate.BookingEnd 	  	= bookingEnd
	rate.DiscountType		= params.PostForm("discount_type")
	rate.Currency			= params.PostForm("currency")
	rate.Nominal			= nominal
	rate.BookingDays		= params.PostForm("booking_days")
	rate.ApplicableTo		= params.PostForm("applicable_to")
	rate.CommissionAdded	= commission_added
	rate.TravelStart		= travelStart
	rate.TravelEnd			= travelEnd
	rate.RoomID  			= params.PostForm("room_id")
	rate.RoomType			= params.PostForm("room_type")
	rate.HotelID			= params.PostForm("hotel_id")
	rate.TravelDays			= params.PostForm("travel_days")
	rate.MinimumStay		= params.PostForm("minimum_stay")
	rate.CancellationPolicy = params.PostForm("cancellation_policy")
	rate.Status				= params.PostForm("status")
	rate.CreatedBy			= createdBy
	rate.ModifiedBy			= modifiedBy
	return rate
}


func (a *rateExtranetUsecase) CheckAvailable(c *gin.Context) error {
	rateParam := mapRate(c)

	id, _ := strconv.ParseUint(c.PostForm("id"), 10, 64)

	msg, err := a.IsAvailable(rateParam,id)

	if err {
		return errors.New(msg)
	}

	return nil
}

func (a *rateExtranetUsecase) Delete(c *gin.Context) error {
	id, _ := strconv.ParseUint(c.Param("id"), 10, 64)

	if id == 0 || helpers.Empty(c.Param("id")){
		return errors.New("parameter id not be null")
	}

	modifiedBy, _ := strconv.Atoi(c.Param("modified_by"))

	m, err := a.rateExtranetRepo.GetById(id)

	if(err != nil){
		return err
	}

	errDel := a.rateExtranetRepo.Delete(m,modifiedBy)
	if errDel != nil {
		return errDel
	}
	return nil
}