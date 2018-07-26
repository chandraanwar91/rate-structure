package handler

import (
	rateExInterfaces "misteraladin.com/jasmine/rate-structure/app/rate-extranet"
	"github.com/gin-gonic/gin"
	"github.com/asaskevich/govalidator"
	"misteraladin.com/jasmine/rate-structure/helpers"
	"misteraladin.com/jasmine/rate-structure/transformers"
)

//eventRules -- Struct for rate rules
type createRateRules struct {
	BookingStart  string `valid:"required~parameter is empty"`
	BookingEnd string `valid:"required~parameter is empty"`
	TravelStart  string `valid:"required~parameter is empty"`
	TravelEnd string `valid:"required~parameter is empty"`
	BookingDays string `valid:"required~parameter is empty"`
	TravelDays string `valid:"required~parameter is empty"`	
	DiscountType string `valid:"required~parameter is empty"`
	Currency string `valid:"required~parameter is empty"`
	//Nominal string `valid:"required~parameter is empty"`
	MinimumStay string `valid:"required~parameter is empty"`
	ApplicableTo string `valid:"required~parameter is empty"`
	RoomID string `valid:"required~parameter is empty"`
	RoomType string `valid:"required~parameter is empty"`
	HotelId string `valid:"required~parameter is empty"`
	CancellationPolicy string `valid:"required~parameter is empty"`
	Status string `valid:"required~parameter is empty"`
}

type RateExternalHandler struct {
	RUsecase rateExInterfaces.IRateExtranetUseCase 
}

// Create -- function for create event
func (a *RateExternalHandler) CreateRateExtranet(c *gin.Context) {
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
func (a *RateExternalHandler) UpdateRateExtranet(c *gin.Context) {
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
	res.TransformExtranet(rate)
	RespondJSON(c,res)
	return
}

func (a *RateExternalHandler) GetAllRateExtranet(c *gin.Context) {

	rates, pagination, err := a.RUsecase.Fetch(c)
	if err != nil {
		RespondFailValidation(c,err.Error())
	}

	//Transform
	res := new(transformers.CollectionTransformer)
	res.TransformExtranetCollection(rates, pagination)
	RespondJSON(c,res)
	return
}

func (a *RateExternalHandler) ShowRateExtranet(c *gin.Context) {
	rate, err := a.RUsecase.GetByID(c)
	if err != nil {
		RespondFailValidation(c,err.Error())
		return
	}

	//Transform
	res := new(transformers.Transformer)
	res.TransformExtranet(rate)
	RespondJSON(c,res)
	return
}

func (a *RateExternalHandler) CheckRateExtranet(c *gin.Context) {

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
		BookingStart 		: params.PostForm("booking_start"),
		BookingEnd 			: params.PostForm("booking_end"),
		TravelStart 		: params.PostForm("travel_start"),
		TravelEnd			: params.PostForm("travel_end"),
		BookingDays			: params.PostForm("booking_days"),
		TravelDays			: params.PostForm("travel_days"),
		RoomID				: params.PostForm("room_id"),
		RoomType			: params.PostForm("room_type"),
		HotelId				: params.PostForm("hotel_id"),
		DiscountType		: params.PostForm("discount_type"),
		Currency			: params.PostForm("currency"),
		//Nominal			: params.PostForm("nominal"),
		MinimumStay			: params.PostForm("minimum_stay"),
		ApplicableTo		: params.PostForm("applicable_to"),
		CancellationPolicy	: params.PostForm("cancellation_policy"),
		Status				: params.PostForm("status"),
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
// DeleteRateExtranet -- function for delete rate extranet
func (a *RateExternalHandler) DeleteRateExtranet(c *gin.Context) {
	err := a.RUsecase.Delete(c)

	if err != nil {
		RespondFailValidation(c,"Failed to delete Rate Plan Extranet")
		return
	}

	RespondDeleted(c,"Success Delete rate extranet")
	return
}
