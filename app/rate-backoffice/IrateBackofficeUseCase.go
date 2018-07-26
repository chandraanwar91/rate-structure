package RateBackoffice

import (
	"misteraladin.com/jasmine/rate-structure/models"
	"github.com/gin-gonic/gin"
)

type IRateBackofficeUseCase interface {
	Store(c *gin.Context) (*models.RateBackoffice, error)
	Update(c *gin.Context) (*models.RateBackoffice, error)
	GetByID(c *gin.Context) (*models.RateBackoffice,error)
	Fetch(c *gin.Context) ([]*models.RateBackoffice,*models.Pagination,error)
	CheckAvailable(c *gin.Context) error
}