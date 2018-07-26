package rate

import (
	//"strconv"
	//"fmt"
	"misteraladin.com/jasmine/rate-structure/models"
	"github.com/gin-gonic/gin"
	//"strings"
)

type IRateUseCase interface {
	FindAvailableExtranet(c *gin.Context) ([]*models.RateExtranet, error)
	FindAvailableBackoffice(c *gin.Context) ([]*models.RateBackoffice, error)
}