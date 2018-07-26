package app

import (
	rateBoInterfaces "misteraladin.com/jasmine/rate-structure/app/rate-backoffice"
	rateExInterfaces "misteraladin.com/jasmine/rate-structure/app/rate-extranet"
	rateInterfaces "misteraladin.com/jasmine/rate-structure/app/rate"
	handlerBO "misteraladin.com/jasmine/rate-structure/app/rate-backoffice/handler"
	handlerEX "misteraladin.com/jasmine/rate-structure/app/rate-extranet/handler"
	handlerRate "misteraladin.com/jasmine/rate-structure/app/rate/handler"
	"github.com/gin-gonic/gin"
)

func NewRateBoHttpHandler(r *gin.Engine,us rateBoInterfaces.IRateBackofficeUseCase) {
	handler := &handlerBO.RateBackofficeHandler{
		RUsecase: us,
	}
	rate := r.Group("/rate-structure")
	rate.POST("/backoffice",handler.CreateRateBackoffice)
	rate.PUT("/backoffice/:id",handler.UpdateRateBackoffice)
	rate.GET("/backoffice",handler.GetAllRateBackoffice)
	rate.GET("/backoffice/:id",handler.ShowRateBackoffice)
	rate.POST("/backoffice/check",handler.CheckRateBackoffice)
}

func NewRateExHttpHandler(r *gin.Engine,us rateExInterfaces.IRateExtranetUseCase) {
	handler := &handlerEX.RateExternalHandler{
		RUsecase: us,
	}
	rate := r.Group("/rate-structure")
	rate.POST("/extranet",handler.CreateRateExtranet)
	rate.GET("/extranet",handler.GetAllRateExtranet)
	rate.GET("/extranet/:id", handler.ShowRateExtranet)
	rate.PUT("/extranet/:id",handler.UpdateRateExtranet)
	rate.POST("/extranet/check",handler.CheckRateExtranet)
	rate.DELETE("/extranet/:id",handler.DeleteRateExtranet)
}

func NewRateHttpHandler(r *gin.Engine,us rateInterfaces.IRateUseCase) {
	handler := &handlerRate.RateHandler{
		RUsecase: us,
	}
	rate := r.Group("/rate-structure")
	rate.GET("/search/extranet",handler.FindAvailableExtranet)
	rate.GET("/search/backoffice",handler.FindAvailableBackoffice)
}