package controllers


import (
	"github.com/a1ta1r/Credit-Portfolio/internal/codes"
	"github.com/a1ta1r/Credit-Portfolio/internal/models"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"net/http"
	"strconv"
)

type ExpenceController struct {
	gormDB gorm.DB
}

func NewExpenceController(db *gorm.DB) ExpenceController {
	return ExpenceController{gormDB: *db}
}

func (expenceController ExpenceController) GetExpenceById(c *gin.Context) {
	var expence models.Income
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnprocessableEntity, gin.H{"message": codes.BadID})
		return
	}
	expenceController.gormDB.Preload("Payments").First(&expence, id)
	if expence.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{"message": codes.ResNotFound})
		return
	}
	c.JSON(http.StatusOK, gin.H{"expence": expence})
}

func (expenceController ExpenceController) AddExpence(c *gin.Context) {
	var expence models.Income
	c.BindJSON(&expence)
	expenceController.gormDB.Create(&expence)
	c.JSON(http.StatusCreated, gin.H{"expence": expence})
}

func (expenceController ExpenceController) UpdateExpenceById(c *gin.Context) {
	var expence models.Income
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnprocessableEntity, gin.H{"message": codes.BadID})
		return
	}
	expenceController.gormDB.Preload("Payments").First(&expence, id)
	if expence.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{"message": codes.ResNotFound})
		return
	}
	c.BindJSON(&expence)
	expenceController.gormDB.Update(&expence)
	c.JSON(http.StatusCreated, gin.H{"expence": expence})
}

func (expenceController ExpenceController) DeleteExpenceById(c *gin.Context) {
	var expence models.Income
	c.BindJSON(&expence)
	expenceController.gormDB.Delete(&expence)
	c.JSON(http.StatusOK, gin.H{"message": codes.ResDeleted})
}
