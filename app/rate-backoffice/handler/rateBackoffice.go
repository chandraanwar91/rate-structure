package handler

import (
	rateBoInterfaces "misteraladin.com/jasmine/rate-structure/app/rate-backoffice"
	"github.com/gin-gonic/gin"
	"github.com/asaskevich/govalidator"
	"misteraladin.com/jasmine/rate-structure/helpers"
	"misteraladin.com/jasmine/rate-structure/transformers"
)

//eventRules -- Struct for rate rules
type createRateRules struct {
	BookingStart  string `valid:"required~parameter is empty"`
	BookingEnd string `valid:"required~parameter is empty"`
	StayingStart  string `valid:"required~parameter is empty"`
	StayingEnd string `valid:"required~parameter is empty"`	
	DiscountType string `valid:"required~parameter is empty"`
	Currency string `valid:"required~parameter is empty"`
	Nominal string `valid:"required~parameter is empty"`
	AvailableTo string `valid:"required~parameter is empty"`
	ApplicableTo string `valid:"required~parameter is empty"`
	Description string `valid:"required~parameter is empty"`
	Status string `valid:"required~parameter is empty"`
}

type RateBackofficeHandler struct {
	RUsecase rateBoInterfaces.IRateBackofficeUseCase 
}

// Create -- function for create event
func (a *RateBackofficeHandler) CreateRateBackoffice(c *gin.Context) {
	validation := ValidateRate(c)
	if(validation != nil){
		RespondFailValidation(c,validation)
		return
	}
	_,err := a.RUsecase.Store(c)
	if err!= nil {
		RespondFailValidation(c,"Failed create rate Detail : "+err.Error())
		return
	}
	RespondCreated(c,"Resource Created")
	return
}

// UpdateRateBackoffice -- function for update rate backoffice
func (a *RateBackofficeHandler) UpdateRateBackoffice(c *gin.Context) {
	validation := ValidateRate(c)
	if(validation != nil){
		RespondFailValidation(c,validation)
		return
	}

	_,err := a.RUsecase.Update(c)

	if err!= nil {
		RespondFailValidation(c,"Failed update rate Detail : "+err.Error())
		return
	}

	rate, err := a.RUsecase.GetByID(c)

	if err != nil {
		RespondFailValidation(c,"Failed update rate Detail : "+err.Error())
	}

	res := new(transformers.Transformer)
	res.Transform(rate)
	RespondJSON(c,res)
	return
}

func (a *RateBackofficeHandler) GetAllRateBackoffice(c *gin.Context) {

	rates, pagination, err := a.RUsecase.Fetch(c)
	if err != nil {
		RespondFailValidation(c,err.Error())
	}

	//Transform
	res := new(transformers.CollectionTransformer)
	res.TransformCollection(rates, pagination)
	RespondJSON(c,res)
	return
}

func (a *RateBackofficeHandler) ShowRateBackoffice(c *gin.Context) {
	rate, err := a.RUsecase.GetByID(c)
	if err != nil {
		RespondFailValidation(c,err.Error())
		return
	}

	//Transform
	res := new(transformers.Transformer)
	res.Transform(rate)
	RespondJSON(c,res)
	return
}

func (a *RateBackofficeHandler) CheckRateBackoffice(c *gin.Context) {

	var res = make(map[string]string)

	res["status"] = "1"
	res["message"] = ""

	err := a.RUsecase.CheckAvailable(c)
	if err != nil {
		res["status"] = "0"
		res["message"] = err.Error()
	}

	RespondJSON(c,res)
	return 
}


func ValidateRate(params *gin.Context) interface{}{
	rateRules := &createRateRules{
		BookingStart 	: params.PostForm("booking_start"),
		BookingEnd 		: params.PostForm("booking_end"),
		StayingStart 	: params.PostForm("staying_start"),
		StayingEnd		: params.PostForm("staying_end"),
		DiscountType	: params.PostForm("discount_type"),
		Currency		: params.PostForm("currency"),
		Nominal			: params.PostForm("nominal"),
		AvailableTo		: params.PostForm("available_to"),
		ApplicableTo	: params.PostForm("applicable_to"),
		Description		: params.PostForm("description"),
		Status			: params.PostForm("status"),
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
