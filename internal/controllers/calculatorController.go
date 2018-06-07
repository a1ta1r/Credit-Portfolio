package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"net/http"
)

type CalculatorController struct {
	gormDB gorm.DB
}

func (cc CalculatorController) Calculate(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "hello world",
	})
}

func NewCalculatorController(pg *gorm.DB) CalculatorController {
	return CalculatorController{gormDB: *pg}
}
