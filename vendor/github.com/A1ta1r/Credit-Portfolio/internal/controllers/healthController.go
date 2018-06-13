package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"net/http"
)

type HealthController struct {
	gormDB gorm.DB
}

//checks if services are still alive and returns 200
//returns 503 otherwise
func (hc HealthController) HealthCheck(c *gin.Context) {
	if hc.gormDB.DB() == nil || hc.gormDB.DB().Ping() != nil {
		c.JSON(http.StatusServiceUnavailable, gin.H{
			"message": "Internal error occurred",
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"message": "I am alive!",
		})
	}
}

func NewHealthController(db *gorm.DB) HealthController {
	return HealthController{
		gormDB: *db,
	}
}
