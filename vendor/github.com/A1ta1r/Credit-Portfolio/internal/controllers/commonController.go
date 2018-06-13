package controllers

import (
	"github.com/A1ta1r/Credit-Portfolio/internal/models"
	"github.com/A1ta1r/Credit-Portfolio/internal/utils"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"net/http"
	"strconv"
)

type CommonController struct {
	db gorm.DB
}

func NewCommonController(pg *gorm.DB) CommonController {
	return CommonController{db: *pg}
}

func (cc CommonController) AddRole(c *gin.Context) {
	var role models.Role
	c.BindJSON(&role)
	cc.db.Create(&role)
	c.JSON(http.StatusCreated, gin.H{"role": role})
}

func (cc CommonController) AddCurrency(c *gin.Context) {
	var currency models.Currency
	c.BindJSON(&currency)
	cc.db.Create(&currency)
	c.JSON(http.StatusCreated, gin.H{"currency": currency})
}

func (cc CommonController) AddBank(c *gin.Context) {
	var bank models.Bank
	c.BindJSON(&bank)
	cc.db.Create(&bank)
	c.JSON(http.StatusCreated, gin.H{"bank": bank})
}

func (cc CommonController) AddPaymentType(c *gin.Context) {
	var paymentType models.PaymentType
	c.BindJSON(&paymentType)
	cc.db.Create(&paymentType)
	c.JSON(http.StatusCreated, gin.H{"paymentType": paymentType})
}

func (cc CommonController) GetRole(c *gin.Context) {
	var role models.Role
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnprocessableEntity, gin.H{"message": utils.BadID})
		return
	}
	cc.db.First(&role, id)
	if role.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{"message": utils.ResNotFound})
		return
	}
	cc.db.First(&role, id)
	c.JSON(http.StatusOK, gin.H{"role": role})
}

func (cc CommonController) GetCurrency(c *gin.Context) {
	var currency models.Currency
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnprocessableEntity, gin.H{"message": utils.BadID})
		return
	}
	cc.db.First(&currency, id)
	if currency.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{"message": utils.ResNotFound})
		return
	}
	cc.db.First(&currency, id)
	c.JSON(http.StatusOK, gin.H{"currency": currency})
}

func (cc CommonController) GetBank(c *gin.Context) {
	var bank models.Bank
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnprocessableEntity, gin.H{"message": utils.BadID})
		return
	}
	cc.db.First(&bank, id)
	if bank.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{"message": utils.ResNotFound})
		return
	}
	cc.db.First(&bank, id)
	c.JSON(http.StatusOK, gin.H{"bank": bank})
}

func (cc CommonController) GetPaymentType(c *gin.Context) {
	var paymentType models.PaymentType
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnprocessableEntity, gin.H{"message": utils.BadID})
		return
	}
	cc.db.First(&paymentType, id)
	if paymentType.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{"message": utils.ResNotFound})
		return
	}
	cc.db.First(&paymentType, id)
	c.JSON(http.StatusOK, gin.H{"paymentType": paymentType})
}
