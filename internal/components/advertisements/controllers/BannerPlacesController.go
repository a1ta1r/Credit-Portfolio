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

type BannerPlacesController struct {
	bannerPlaceStorage storages.BannerPlaceStorage
}

func NewBannerPlacesController(bs storages.BannerPlaceStorage) BannerPlacesController {
	return BannerPlacesController{bannerPlaceStorage: bs}
}

func (bc BannerPlacesController) GetBannerPlaces(c *gin.Context) {
	var bannerPlaces []entities.BannerPlace
	var err error
	if bannerPlaces, err = bc.bannerPlaceStorage.GetBannerPlaces(); err == gorm.ErrRecordNotFound {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"message": codes.NotFound})
		return
	}
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": codes.InternalError})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status":       "OK",
		"count":        len(bannerPlaces),
		"bannerPlaces": bannerPlaces,
	})
}

func (bc BannerPlacesController) GetBannerPlaceByID(c *gin.Context) {
	var bannerPlace entities.BannerPlace
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnprocessableEntity, gin.H{"message": codes.BadID})
		return
	}
	if bannerPlace, err = bc.bannerPlaceStorage.GetBannerPlace(uint(id)); err == gorm.ErrRecordNotFound {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"message": codes.NotFound})
		return
	}
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": codes.InternalError})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status":      "OK",
		"bannerPlace": bannerPlace,
	})
}

func (bc BannerPlacesController) DeleteBannerPlaceByID(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnprocessableEntity, gin.H{"error": codes.BadID})
		return
	}
	bannerPlace := entities.BannerPlace{ID: uint(id)}
	bc.bannerPlaceStorage.DeleteBannerPlace(bannerPlace)
	c.JSON(http.StatusOK, gin.H{"message": codes.ResDeleted})
}

func (bc BannerPlacesController) AddBannerPlace(c *gin.Context) {
	var request requests.NewBannerPlace
	var bannerPlace entities.BannerPlace
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
	bannerPlace = request.ToBanner()
	err := bc.bannerPlaceStorage.CreateBannerPlace(&bannerPlace)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": codes.Unhealthy})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"bannerPlace": bannerPlace})
}

func (bc BannerPlacesController) UpdateBannerPlace(c *gin.Context) {
	var request requests.UpdateBannerPlace
	var bannerPlace entities.BannerPlace
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnprocessableEntity, gin.H{"error": codes.BadID})
		return
	}
	bannerPlace, _ = bc.bannerPlaceStorage.GetBannerPlace(uint(id))
	if bannerPlace.ID == 0 {
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
	bannerPlace = request.ToBannerPlace(bannerPlace)
	_ = bc.bannerPlaceStorage.UpdateBannerPlace(&bannerPlace)
	c.JSON(http.StatusOK, gin.H{
		"status":      "OK",
		"bannerPlace": bannerPlace,
	})
}
