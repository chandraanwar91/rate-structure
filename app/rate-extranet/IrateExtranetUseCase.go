package RateExtranet

import (
	"misteraladin.com/jasmine/rate-structure/models"
	"github.com/gin-gonic/gin"
)

type IRateExtranetUseCase interface {
	Store(c *gin.Context) (*models.RateExtranet, error)
	Update(c *gin.Context) (*models.RateExtranet, error)
	GetByID(c *gin.Context) (*models.RateExtranet,error)
	Fetch(c *gin.Context) ([]*models.RateExtranet,*models.Pagination,error)
	CheckAvailable(c *gin.Context) error
	Delete(c *gin.Context) error
}