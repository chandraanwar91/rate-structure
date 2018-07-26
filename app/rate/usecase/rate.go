package usecase

import (
	//"fmt"
	"time"
	//"misteraladin.com/jasmine/rate-structure/helpers"
	//"io/ioutil"
	"strconv"
	rateBoInterfaces "misteraladin.com/jasmine/rate-structure/app/rate-backoffice"
	rateExInterfaces "misteraladin.com/jasmine/rate-structure/app/rate-extranet"
	rateInterfaces "misteraladin.com/jasmine/rate-structure/app/rate"
	"misteraladin.com/jasmine/rate-structure/models"
	"github.com/gin-gonic/gin"
)

type rateUsecase struct{
	rateBackofficeRepo rateBoInterfaces.IRateBackofficeRepository
	rateExtranetRepo rateExInterfaces.IRateExtranetRepository
}

func NewRateUsecase(a rateBoInterfaces.IRateBackofficeRepository,e rateExInterfaces.IRateExtranetRepository) rateInterfaces.IRateUseCase  {
	return &rateUsecase{
		rateBackofficeRepo: a,
		rateExtranetRepo: e,
	}
}

func (a *rateUsecase) FindAvailableExtranet(c *gin.Context) ([]*models.RateExtranet, error) {
	var error error

	rateParam := mapRate(c)

	var applicableTo = ""

	if rateParam.Platform == "ios" || rateParam.Platform == "android" {
		applicableTo = "apps"
	}
	
	if rateParam.IsUserLogin > 0 {
		applicableTo = "member"
	}

	if applicableTo == ""{
		return nil,error
	}

	rates,total,err := a.rateExtranetRepo.FindAvailable(rateParam,applicableTo)
	if total <= 0 {
		return nil,err
	}

	return rates,err
}

func (a *rateUsecase) FindAvailableBackoffice(c *gin.Context) ([]*models.RateBackoffice, error) {
	var error error

	rateParam := mapRate(c)

	var applicableTo = ""

	if rateParam.Platform == "ios" || rateParam.Platform == "android" {
		applicableTo = "apps"
	}
	
	if rateParam.IsUserLogin > 0 {
		applicableTo = "member"
	}

	if applicableTo == ""{
		return nil,error
	}

	rates,total,err := a.rateBackofficeRepo.FindAvailable(rateParam,applicableTo)
	if total <= 0 {
		return nil,err
	}

	return rates,err
}

func mapRate(params *gin.Context) *models.Rate{
	isUserLogin,_ := strconv.Atoi(params.Query("is_user_login"))
	checkIn,_ := time.Parse("2006-01-02", params.Query("check_in"))
	checkOut,_ := time.Parse("2006-01-02", params.Query("check_out"))
	bookingDate,_ := time.Parse("2006-01-02", params.Query("booking_date"))
	rate := new(models.Rate)
	rate.CheckIn 		= checkIn
	rate.CheckOut 	  	= checkOut
	rate.BookingDate	= bookingDate
	rate.RoomIds		= params.Query("room_ids")
	rate.Platform		= params.Query("platform")
	rate.IsUserLogin	= isUserLogin
	return rate
}