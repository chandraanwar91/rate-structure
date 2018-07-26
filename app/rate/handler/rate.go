package handler

import (
	//"fmt"
	rateInterfaces "misteraladin.com/jasmine/rate-structure/app/rate"
	"github.com/gin-gonic/gin"
	"github.com/asaskevich/govalidator"
	"misteraladin.com/jasmine/rate-structure/helpers"
	"misteraladin.com/jasmine/rate-structure/transformers"
)

//eventRules -- Struct for rate rules
type findRateExtranetRules struct {
	CheckIn     string  `valid:"required~parameter is empty"`
	CheckOut	string	`valid:"required~parameter is empty"`
	BookingDate	string	`valid:"required~parameter is empty"`
	RoomIds		string  `valid:"required~parameter is empty"`
	Platform	string	`valid:"required~parameter is empty"`
}

type findRateBackofficeRules struct {
	CheckIn     string  `valid:"required~parameter is empty"`
	CheckOut	string	`valid:"required~parameter is empty"`
	BookingDate	string	`valid:"required~parameter is empty"`
	Platform	string	`valid:"required~parameter is empty"`
}

type RateHandler struct {
	RUsecase rateInterfaces.IRateUseCase
}

// findAvailableExtranet -- function for find available extranet
func (a *RateHandler) FindAvailableExtranet(c *gin.Context) {
	validation := ValidateRate(c)
	if(validation != nil){
		RespondFailValidation(c,validation)
		return
	}
	rates,err := a.RUsecase.FindAvailableExtranet(c)

	res := new(transformers.TransformerInit)
	if err!= nil {
		RespondJSON(c,res)
		return
	}
	res.TransformAvailableExtranet(rates)
	RespondJSON(c,res)
	return
}

// findAvailableBackoffice -- function for find available rate backoffice
func (a *RateHandler) FindAvailableBackoffice(c *gin.Context) {
	validation := ValidateRateBackoffice(c)
	if(validation != nil){
		RespondFailValidation(c,validation)
		return
	}
	rates,err := a.RUsecase.FindAvailableBackoffice(c)

	res := new(transformers.TransformerBOInit)
	if err!= nil {
		RespondJSON(c,res)
		return
	}
	res.TransformAvailableBackoffice(rates)
	RespondJSON(c,res)
	return
}

func ValidateRate(params *gin.Context) interface{}{
	rateRules := &findRateExtranetRules{
		CheckIn  	: params.Query("check_in"),
		CheckOut 	: params.Query("check_out"),
		BookingDate	: params.Query("booking_date"),
		RoomIds		: params.Query("room_ids"),
		Platform	: params.Query("platform"),
	}
	_,err := govalidator.ValidateStruct(rateRules)
	if(err != nil){
		respErr := helpers.ValidationError(rateRules,err)
		
		if(respErr != nil){
			return respErr
		}
	}

	return nil
}

func ValidateRateBackoffice(params *gin.Context) interface{}{
	rateRules := &findRateBackofficeRules{
		CheckIn  	: params.Query("check_in"),
		CheckOut 	: params.Query("check_out"),
		BookingDate	: params.Query("booking_date"),
		Platform	: params.Query("platform"),
	}
	_,err := govalidator.ValidateStruct(rateRules)
	if(err != nil){
		respErr := helpers.ValidationError(rateRules,err)
		
		if(respErr != nil){
			return respErr
		}
	}

	return nil
}
