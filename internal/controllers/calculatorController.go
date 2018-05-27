package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"github.com/a1ta1r/Credit-Portfolio/internal/services"
)

type CalculatorController struct {
	calculatorService services.Calculator
}

func (cc CalculatorController) Calculate(c *gin.Context) {
	paymentPlan := cc.calculatorService.Calculate()
	c.JSON(http.StatusOK, gin.H{
		"payment_plan": paymentPlan,
	})
}

func NewCalculatorController() CalculatorController {
	return CalculatorController{services.Calculator{}}
}
