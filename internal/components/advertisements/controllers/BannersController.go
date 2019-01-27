package controllers

import (
	"github.com/a1ta1r/Credit-Portfolio/internal/codes"
	"github.com/a1ta1r/Credit-Portfolio/internal/components/advertisements/entities"
	"github.com/a1ta1r/Credit-Portfolio/internal/components/advertisements/storages"
	"github.com/a1ta1r/Credit-Portfolio/internal/components/errors"
	"github.com/a1ta1r/Credit-Portfolio/internal/specification/requests"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"gopkg.in/go-playground/validator.v8"
	"net/http"
	"strconv"
)

type BannersController struct {
	bannerStorage storages.BannerStorage
}

func NewBannersController(bs storages.BannerStorage) BannersController {
	return BannersController{bannerStorage: bs}
}

func (bc BannersController) GetBannersByAdvertisementID(c *gin.Context) {
	var banners []entities.Banner
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnprocessableEntity, gin.H{"message": codes.BadID})
		return
	}
	if banners, err = bc.bannerStorage.GetBannersByAdvertisement(uint(id)); err == gorm.ErrRecordNotFound {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"message": codes.NotFound})
		return
	}
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": codes.InternalError})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status":  "OK",
		"count":   len(banners),
		"banners": banners,
	})
}

func (bc BannersController) GetBannerByID(c *gin.Context) {
	var banner entities.Banner
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnprocessableEntity, gin.H{"message": codes.BadID})
		return
	}
	if banner, err = bc.bannerStorage.GetBanner(uint(id)); err == gorm.ErrRecordNotFound {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"message": codes.NotFound})
		return
	}
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": codes.InternalError})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status": "OK",
		"banner": banner,
	})
}

func (bc BannersController) DeleteBannerByID(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnprocessableEntity, gin.H{"error": codes.BadID})
		return
	}
	banner := entities.Banner{ID: uint(id)}
	bc.bannerStorage.DeleteBanner(banner)
	c.JSON(http.StatusOK, gin.H{"message": codes.ResDeleted})
}

func (bc BannersController) AddBanner(c *gin.Context) {
	var request requests.NewBanner
	var banner entities.Banner
	if err := c.ShouldBindJSON(&request); err != nil {
		validationErrors, ok := err.(validator.ValidationErrors)
		if !ok {
			c.JSON(http.StatusUnprocessableEntity, gin.H{"error": codes.InvalidJSON})
			return
		}
		errorMsg := errors.GetErrorMessages(validationErrors)
		c.JSON(http.StatusUnprocessableEntity, gin.H{"errors": errorMsg})
		return
	}
	banner = request.ToBanner()
	err := bc.bannerStorage.CreateBanner(&banner)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": codes.Unhealthy})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"banner": banner})
}

func (bc BannersController) UpdateBanner(c *gin.Context) {
	var request requests.UpdateBanner
	var banner entities.Banner
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnprocessableEntity, gin.H{"error": codes.BadID})
		return
	}
	banner, _ = bc.bannerStorage.GetBanner(uint(id))
	if banner.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": codes.ResNotFound})
		return
	}

	if err := c.ShouldBindJSON(&request); err != nil {
		validationErrors, ok := err.(validator.ValidationErrors)
		if !ok {
			c.JSON(http.StatusUnprocessableEntity, gin.H{"error": codes.InvalidJSON})
			return
		}
		errorMsg := errors.GetErrorMessages(validationErrors)
		c.JSON(http.StatusUnprocessableEntity, gin.H{"errors": errorMsg})
		return
	}
	banner = request.ToBanner(banner)
	_ = bc.bannerStorage.UpdateBanner(&banner)
	c.JSON(http.StatusOK, gin.H{
		"status": "OK",
		"banner": banner,
	})
}
