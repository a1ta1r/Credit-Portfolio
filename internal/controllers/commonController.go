package controllers

import (
	"github.com/a1ta1r/Credit-Portfolio/internal/codes"
	"github.com/a1ta1r/Credit-Portfolio/internal/models"
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

func (cc CommonController) AddCurrency(c *gin.Context) {
	var currency models.Currency
	c.BindJSON(&currency)
	cc.db.Create(&currency)
	c.JSON(http.StatusCreated, currency)
}

func (cc CommonController) DeleteCurrency(c *gin.Context) {
	var currency models.Currency
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnprocessableEntity, gin.H{"message": codes.BadID})
		return
	}
	cc.db.First(&currency, id)
	if currency.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{"message": codes.ResNotFound})
		return
	}
	cc.db.First(&currency, id)
	cc.db.Delete(&currency)
	c.JSON(http.StatusOK, gin.H{"currency deleted": currency.Name})
}

func (cc CommonController) UpdateCurrency(c *gin.Context) {
	var currency models.Currency
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnprocessableEntity, gin.H{"message": codes.BadID})
		return
	}
	cc.db.First(&currency, id)
	if currency.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{"message": codes.ResNotFound})
		return
	}
	c.BindJSON(&currency)
	cc.db.Save(&currency)
	c.JSON(http.StatusOK, gin.H{"bank": currency})
}

func (cc CommonController) AddBank(c *gin.Context) {
	var bank models.Bank
	c.BindJSON(&bank)
	cc.db.Create(&bank)
	c.JSON(http.StatusCreated, bank)
}

func (cc CommonController) GetCurrency(c *gin.Context) {
	var currency models.Currency
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnprocessableEntity, gin.H{"message": codes.BadID})
		return
	}
	cc.db.First(&currency, id)
	if currency.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{"message": codes.ResNotFound})
		return
	}
	cc.db.First(&currency, id)
	c.JSON(http.StatusOK, gin.H{"currency": currency})
}

func (cc CommonController) GetBank(c *gin.Context) {
	var bank models.Bank
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnprocessableEntity, gin.H{"message": codes.BadID})
		return
	}
	cc.db.First(&bank, id)
	if bank.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{"message": codes.ResNotFound})
		return
	}
	cc.db.First(&bank, id)
	c.JSON(http.StatusOK, gin.H{"bank": bank})
}

func (cc CommonController) DeleteBank(c *gin.Context) {
	var bank models.Bank
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnprocessableEntity, gin.H{"message": codes.BadID})
		return
	}
	cc.db.First(&bank, id)
	if bank.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{"message": codes.ResNotFound})
		return
	}
	cc.db.First(&bank, id)
	cc.db.Delete(&bank)
	c.JSON(http.StatusOK, gin.H{"bank deleted": bank.Name})
}

func (cc CommonController) UpdateBank(c *gin.Context) {
	var bank models.Bank
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnprocessableEntity, gin.H{"message": codes.BadID})
		return
	}
	cc.db.First(&bank, id)
	if bank.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{"message": codes.ResNotFound})
		return
	}
	c.BindJSON(&bank)
	cc.db.Save(&bank)
	c.JSON(http.StatusOK, gin.H{"bank updated": bank.Name})
}
