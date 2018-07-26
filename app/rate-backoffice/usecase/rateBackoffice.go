package usecase

import (
	//"fmt"
	"time"
	"misteraladin.com/jasmine/rate-structure/helpers"
	//"io/ioutil"
	"strconv"
	rateBoInterfaces "misteraladin.com/jasmine/rate-structure/app/rate-backoffice"
	hotelInterfaces "misteraladin.com/jasmine/rate-structure/app/hotel"
	"misteraladin.com/jasmine/rate-structure/models"
	"github.com/gin-gonic/gin"
	"errors"
)

type rateBackofficeUsecase struct{
	rateBackofficeRepo rateBoInterfaces.IRateBackofficeRepository
	hotelRepo hotelInterfaces.IHotelRepository
}

func NewRateBackofficeUsecase(a rateBoInterfaces.IRateBackofficeRepository,h hotelInterfaces.IHotelRepository) rateBoInterfaces.IRateBackofficeUseCase {
	return &rateBackofficeUsecase{
		rateBackofficeRepo: a,
		hotelRepo: h,
	}
}

func (a *rateBackofficeUsecase) Store(c *gin.Context) (*models.RateBackoffice, error) {

	rateParam := mapRate(c)

	msg,exist := a.IsAvailable(rateParam,0)
	if exist {
		return nil, errors.New(msg)
	}

	res, err := a.rateBackofficeRepo.Store(rateParam)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (a *rateBackofficeUsecase) Update(c *gin.Context) (*models.RateBackoffice, error) {
	id, _ := strconv.ParseUint(c.Param("id"), 10, 64)

	m, err := a.rateBackofficeRepo.GetById(id)

	if(err != nil){
		return nil,errors.New(err.Error())
	}

	rateParam := mapRate(c)

	msg,exist := a.IsAvailable(rateParam,id)
	if exist {
		return nil, errors.New(msg)
	}

	res,err := a.rateBackofficeRepo.Update(rateParam,m)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (a *rateBackofficeUsecase) GetByID(c *gin.Context) (*models.RateBackoffice, error) {
	id, _ := strconv.ParseUint(c.Param("id"), 10, 64)

	if id == 0 || helpers.Empty(c.Param("id")){
		return nil, errors.New("parameter id not be null")
	}

	res, err := a.rateBackofficeRepo.GetById(id)
	if err != nil {
		return nil, err
	}

	if res.ID == 0 {
		return nil,errors.New("Data Not Found")
	}

	return res, nil
}

func (a *rateBackofficeUsecase) Fetch(c *gin.Context) ([]*models.RateBackoffice, *models.Pagination, error) {
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
	stayingStart := c.Query("staying_start")
	stayingEnd := c.Query("staying_end")

	if page < 1 {
		page = 1
	}

	if perPage < 1 {
		perPage = 10
	}

	//Get All
	//txCount.Table("rate_backoffice").Where("applicable_to = ? ",ApplicableTo).Scopes(models.ScopeActive).Count(&total)
	rates,total,err := a.rateBackofficeRepo.Fetch(page,perPage,applicableTo,bookingStart,bookingEnd,stayingStart,stayingEnd,status)

	if(err != nil){
		return nil,nil,err
	}

	count := len(rates)
	pagination = models.BuildPagination(total, page, perPage, count)

	return rates, pagination, err
}

func  (a *rateBackofficeUsecase) IsAvailable(rateParam *models.RateBackoffice,id uint64) (string,bool) {
	//check source
	var exist = false
	var(
		date = "2006-01-02"
		BookingStart  = rateParam.BookingStart.Format(date)
		BookingEnd	  = rateParam.BookingEnd.Format(date)
		StayingStart  = rateParam.StayingStart.Format(date)
		StayingEnd	  = rateParam.StayingEnd.Format(date)
	)

	dataExist,err := a.rateBackofficeRepo.GetByDate(BookingStart,BookingEnd,StayingStart,StayingEnd,id)

	if(err != nil){
		return err.Error(),true
	}

	count := len(dataExist)

	if(count > 0){
		for _, rate := range dataExist {
			idStr :=  strconv.Itoa(rate.ID)
			var hotelsTemp []int
			sourceExist := false
			var resultSources = helpers.Explode(rate.AvailableTo,",")
			var sources = helpers.Explode(rateParam.AvailableTo,",")
			var resultApplicable = helpers.Explode(rate.ApplicableTo,",")
			var applicables = helpers.Explode(rateParam.ApplicableTo,",")
			for _,source := range sources {
				if r := helpers.InArray(source,resultSources); r {
					sourceExist = true
					break
				}
			}

			if(sourceExist){
				if helpers.Empty(rateParam.CountryID) && helpers.Empty(rateParam.CityID) && helpers.Empty(rateParam.HotelGroupID) && helpers.Empty(rateParam.HotelID) && helpers.Empty(rateParam.StarRating) {
					for _,applicable := range applicables {
						if r := helpers.InArray(applicable,resultApplicable); r {
							exist =  true
							return "Duplicate Source "+applicable+" in Configuration ID "+idStr,exist
						}
					}	
				}

				//check hotel exist
				resHotels,_ := a.hotelRepo.GetHotel(rateParam.CountryID,rateParam.CityID,rateParam.HotelGroupID,rateParam.HotelID,rateParam.StarRating,rateParam.HotelTagID)
				hotels,_ :=  a.hotelRepo.GetHotel(rate.CountryID,rate.CityID,rate.HotelGroupID,rate.HotelID,rate.StarRating,rate.HotelTagID) 

				for _,resHotel := range resHotels {
					hotelsTemp = append(hotelsTemp,resHotel.HOTELID)
				}

				for _,hotel := range hotels {
					if r := helpers.InArray(hotel.HOTELID,hotelsTemp); r {
						for _,applicable := range applicables {
							if e := helpers.InArray(applicable,resultApplicable); e {
								exist =  true
								return "Duplicate hotel "+hotel.NAME+" in Configuration ID "+idStr,exist
							}
						}	
						break;
					}
				}

			}

			if exist {
				break
			}

		}
	}
	return "",exist

	
}

func mapRate(params *gin.Context) *models.RateBackoffice{
	nominal, _ := strconv.ParseFloat(params.PostForm("nominal"), 64)
	createdBy, _ := strconv.Atoi(params.PostForm("created_by"))
	bookingStart,_ := time.Parse("2006-01-02", params.PostForm("booking_start"))
	bookingEnd,_ := time.Parse("2006-01-02", params.PostForm("booking_end"))
	stayingStart,_ := time.Parse("2006-01-02", params.PostForm("staying_start"))
	stayingEnd,_ := time.Parse("2006-01-02", params.PostForm("staying_end"))
	rate := new(models.RateBackoffice)
	rate.BookingStart 		= bookingStart
	rate.BookingEnd 	  	= bookingEnd
	rate.DiscountType		= params.PostForm("discount_type")
	rate.Currency			= params.PostForm("currency")
	rate.Nominal			= nominal
	rate.AvailableTo		= params.PostForm("available_to")
	rate.ApplicableTo		= params.PostForm("applicable_to")
	rate.StayingStart		= stayingStart
	rate.StayingEnd			= stayingEnd
	rate.CountryID  		= params.PostForm("country_id")
	rate.CityID				= params.PostForm("city_id")
	rate.HotelGroupID		= params.PostForm("hotel_group_id")
	rate.HotelID			= params.PostForm("hotel_id")
	rate.HotelTagID			= params.PostForm("hotel_tag_id")
	rate.StarRating			= params.PostForm("star_rating")
	rate.Description		= params.PostForm("description")
	rate.Status				= params.PostForm("status")
	rate.CreatedBy			= createdBy
	return rate
}


func (a *rateBackofficeUsecase) CheckAvailable(c *gin.Context) error {
	rateParam := mapRate(c)

	id, _ := strconv.ParseUint(c.PostForm("id"), 10, 64)

	msg, err := a.IsAvailable(rateParam,id)

	if err {
		return errors.New(msg)
	}

	return nil
}