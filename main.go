package main

import (
	//"misteraladin.com/jasmine/rate-structure-new/app/rate-backoffice/interfaces"
	"fmt"
	"log"
	"github.com/gin-gonic/gin"
	"github.com/gin-contrib/cors"
	"misteraladin.com/jasmine/rate-structure/config"
	rateBoUcase "misteraladin.com/jasmine/rate-structure/app/rate-backoffice/usecase"
	rateExUcase "misteraladin.com/jasmine/rate-structure/app/rate-extranet/usecase"
	rateUcase "misteraladin.com/jasmine/rate-structure/app/rate/usecase"
	routes "misteraladin.com/jasmine/rate-structure/app"
	rateBoRepo "misteraladin.com/jasmine/rate-structure/app/rate-backoffice/repository"
	hotelRepo "misteraladin.com/jasmine/rate-structure/app/hotel/repository"
	rateExRepo "misteraladin.com/jasmine/rate-structure/app/rate-extranet/repository"
	"misteraladin.com/jasmine/rate-structure/db"
	"misteraladin.com/jasmine/rate-structure/middleware"
	"github.com/getsentry/raven-go"
)

var appConfig = config.Config.App

func init() {
	raven.SetDSN(config.Config.SENTRY.SentryDSN)
}

func main() {
	r := gin.New()

	// setting Middleware
	middL := middleware.InitMiddleware()
	r.Use(gin.Logger())
	r.Use(middL.SentryLogHandler())
	r.Use(gin.Recovery())

	//CORS
	r.Use(cors.New(cors.Config{
		//AllowOrigins: []string{"*"},
		AllowAllOrigins: true,
		AllowMethods: []string{"GET", "HEAD", "PUT", "PATCH", "POST", "DELETE"},
	}))

	//setting database connection
	db := gorm.MysqlConn()
	dbHotel := gorm.MysqlConn1()

	//setting up Backoffice routes
	rb := rateBoRepo.NewRateBORepository(db)
	htl := hotelRepo.NewMysqlHotelRepository(dbHotel)
	au := rateBoUcase.NewRateBackofficeUsecase(rb,htl)
	routes.NewRateBoHttpHandler(r,au)

	//setting up extranet routes
	re := rateExRepo.NewRateEXRepository(db)
	ae := rateExUcase.NewrateExtranetUsecase(re)
	routes.NewRateExHttpHandler(r,ae)

	//setting up rate routes
	ar := rateUcase.NewRateUsecase(rb,re)
	routes.NewRateHttpHandler(r,ar)

	// Server
	if err := r.Run(fmt.Sprintf("%s:%s", appConfig.HttpAddr, appConfig.HttpPort)); err != nil {
		log.Fatal(err)
	}
}
