package controllers

import (
	"github.com/a1ta1r/Credit-Portfolio/internal/codes"
	"github.com/a1ta1r/Credit-Portfolio/internal/components/advertisements/entities"
	"github.com/a1ta1r/Credit-Portfolio/internal/components/errors"
	"github.com/a1ta1r/Credit-Portfolio/internal/components/advertisements/storages"
	"github.com/a1ta1r/Credit-Portfolio/internal/components/roles"
	"github.com/a1ta1r/Credit-Portfolio/internal/specification/requests"
	_ "github.com/a1ta1r/Credit-Portfolio/internal/specification/responses"
	"github.com/gin-gonic/gin"
	"gopkg.in/go-playground/validator.v8"
	"net/http"
	"strconv"
)

type AdvertiserController struct {
	advertiserStorage  storages.AdvertiserStorage
	bannerStorage      storages.BannerStorage
	bannerPlaceStorage storages.BannerPlaceStorage
}

func NewAdvertiserController(
	advertiserStorage storages.AdvertiserStorage,
	bannerStorage storages.BannerStorage,
	bannerPlaceStorage storages.BannerPlaceStorage,
) AdvertiserController {
	return AdvertiserController{
		advertiserStorage:  advertiserStorage,
		bannerStorage:      bannerStorage,
		bannerPlaceStorage: bannerPlaceStorage,
	}
}

// @Tags Advertisers
// @Summary Получить список всех рекламодателей
// @Description Метод возвращает список всех имеющихся в системе рекламодателей
// @Produce  json
// @Success 200 {object} responses.AllAdvertisers
// @Router /partners [get]
func (ac AdvertiserController) GetAdvertisers(c *gin.Context) {
	var advertisers []entities.Advertiser
	advertisers, _ = ac.advertiserStorage.GetAdvertisers()
	for i := 0; i < len(advertisers); i++ {
		advertisers[i].Password = ""
	}
	c.JSON(http.StatusOK, gin.H{
		"status":      "OK",
		"count":       len(advertisers),
		"advertisers": advertisers,
	})
}

// @Tags Advertisers
// @Summary Получить рекламодателя по ID
// @Description Метод возвращает рекламодателя по его ID
// @Produce  json
// @Param id path int true "ID рекламодателя"
// @Success 200 {object} responses.OneAdvertiser
// @Failure 404 "{"message": "resource not found"}"
// @Failure 422
// @Router /partners/{id} [get]
func (ac AdvertiserController) GetAdvertiser(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnprocessableEntity, gin.H{"message": codes.BadID})
		return
	}
	advertiser, err := ac.advertiserStorage.GetAdvertiser(uint(id))
	if advertiser.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{"message": codes.ResNotFound})
		return
	}
	advertiser.Password = ""
	c.JSON(http.StatusOK, gin.H{"advertiser": advertiser})
}

// @Tags Advertisers
// @Summary Добавить нового рекламодателя
// @Description Метод добавляет в систему нового рекламодателя с заданными параметрами
// @Accept json
// @Produce  json
// @Param advertiser body requests.NewAdvertiser true "Данные о рекламодателе"
// @Success 201 {object} responses.OneAdvertiser
// @Router /partners [post]
func (ac AdvertiserController) AddAdvertiser(c *gin.Context) {
	var request requests.NewAdvertiser
	var advertiser entities.Advertiser
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
	advertiser = request.ToAdvertiser()
	advertiser.Role = roles.Ads
	advertiser.Password = advertiser.GetHashedPassword()
	err := ac.advertiserStorage.CreateAdvertiser(&advertiser)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": codes.Unhealthy})
		return
	}
	advertiser.Password = ""
	c.JSON(http.StatusCreated, gin.H{"advertiser": advertiser})
}

// @Tags Advertisers
// @Summary Удалить рекламодателя
// @Description Метод удаляет из системы рекламодателя с заданным ID
// @Produce  json
// @Param id path int true "ID рекламодателя"
// @Success 200
// @Failure 422
// @Router /partners/{id} [delete]
func (ac AdvertiserController) DeleteAdvertiser(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnprocessableEntity, gin.H{"error": codes.BadID})
		return
	}
	advertiser := entities.Advertiser{ID: uint(id)}
	ac.advertiserStorage.DeleteAdvertiser(advertiser)
	c.JSON(http.StatusOK, gin.H{"message": codes.ResDeleted})
}

// @Tags Advertisers
// @Summary Обновить рекламодателя
// @Description Метод обновляет рекламодателя заданными параметрами по ID
// @Accept json
// @Produce  json
// @Param id path int true "ID рекламодателя"
// @Param advertiser body requests.UpdateAdvertiser true "Новые данные о рекламодателе"
// @Success 200 {object} responses.OneAdvertiser
// @Failure 404
// @Failure 422
// @Router /partners/{id} [put]
func (ac AdvertiserController) UpdateAdvertiser(c *gin.Context) {
	var request requests.UpdateAdvertiser
	var advertiser entities.Advertiser
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnprocessableEntity, gin.H{"error": codes.BadID})
		return
	}
	advertiser, _ = ac.advertiserStorage.GetAdvertiser(uint(id))
	if advertiser.ID == 0 {
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
	advertiser = request.ToAdvertiser(advertiser)
	_ = ac.advertiserStorage.UpdateAdvertiser(&advertiser)
	advertiser.Password = ""
	c.JSON(http.StatusOK, gin.H{
		"status":     "OK",
		"advertiser": advertiser,
	})
}

func (ac AdvertiserController) GetBannersByAdvertisement(c *gin.Context) {
	var banners []entities.Banner
	id, err := strconv.ParseUint(c.Param("adsid"), 10, 32)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnprocessableEntity, gin.H{"error": codes.BadID})
		return
	}
	banners, _ = ac.bannerStorage.GetBannersByAdvertisement(uint(id))
	c.JSON(http.StatusOK, gin.H{
		"status":  "OK",
		"count":   len(banners),
		"banners": banners,
	})
}
