package controllers

import (
	"github.com/a1ta1r/Credit-Portfolio/internal/codes"
	"github.com/a1ta1r/Credit-Portfolio/internal/components/advertisements/entities"
	"github.com/a1ta1r/Credit-Portfolio/internal/components/advertisements/storages"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"net/http"
	"strconv"
)

type AdvertisementController struct {
	advertisementStorage storages.AdvertisementStorage
}

func NewAdvertisementController(storage storages.AdvertisementStorage) AdvertisementController {
	return AdvertisementController{advertisementStorage: storage}
}

// @Tags Advertisements
// @Summary Получить рекламные объявления рекламодателя
// @Description Метод возвращает данные о рекламных объявлениях рекламодателя с заданным ID
// @Produce  json
// @Param id path int true "ID рекламодателя"
// @Success 200 {object} responses.AllAdvertisements
// @Failure 422
// @Router /advertiser/{id}/ads [get]
func (ac AdvertisementController) GetAdvertisementsByAdvertiser(c *gin.Context) {
	var advertisements []entities.Advertisement
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnprocessableEntity, gin.H{"message": codes.BadID})
		return
	}
	advertisements, _ = ac.advertisementStorage.GetAdvertisementsByAdvertiser(uint(id))
	c.JSON(http.StatusOK, gin.H{
		"status":         "OK",
		"count":          len(advertisements),
		"advertisements": advertisements,
	})
}

// @Tags Advertisements
// @Summary Получить вообще все рекламные объявления
// @Description Метод возвращает данные о всех рекламных объявлениях
// @Produce  json
// @Success 200 {object} responses.AllAdvertisements
// @Failure 422
// @Router /ads [get]
func (ac AdvertisementController) GetAdvertisements(c *gin.Context) {
	var advertisements []entities.Advertisement
	advertisements, _ = ac.advertisementStorage.GetAdvertisements()
	c.JSON(http.StatusOK, gin.H{
		"status":      "OK",
		"count":       len(advertisements),
		"advertisements": advertisements,
	})
}

// @Tags Advertisements
// @Summary Получить рекламное объявление по ID
// @Description Метод возвращает данные рекламном объявлении с данным ID
// @Produce  json
// @Param id path int true "ID рекламного объявления"
// @Success 200 {object} responses.OneAdvertisement
// @Failure 404
// @Failure 422
// @Router /ads/{id} [get]
func (ac AdvertisementController) GetAdvertisement(c *gin.Context) {
	var advertisement entities.Advertisement
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnprocessableEntity, gin.H{"message": codes.BadID})
		return
	}
	advertisement, _ = ac.advertisementStorage.GetAdvertisement(uint(id))
	if advertisement.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{"message": codes.ResNotFound})
		return
	}
	c.JSON(http.StatusOK, gin.H{"advertisement": advertisement})
}

// @Tags Advertisements
// @Summary Добавить рекламное объявление
// @Description Метод добавляет новое рекламне объявление с заданными параметрами
// @Accept json
// @Produce  json
// @Param advertisement body entities.Advertisement true "Данные о рекламном объявлении"
// @Success 201 {object} responses.OneAdvertisement
// @Failure 422
// @Router /ads [post]
func (ac AdvertisementController) AddAdvertisement(c *gin.Context) {
	var advertisement entities.Advertisement
	c.BindJSON(&advertisement)
	err := ac.advertisementStorage.CreateAdvertisement(advertisement)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"message": codes.ResourceExists})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"advertisement": advertisement})
}

// @Tags Advertisements
// @Summary Удалить рекламное объявление
// @Description Метод удаляет рекламне объявление с заданным ID
// @Produce  json
// @Param id path int true "ID рекламного объявления"
// @Success 200
// @Failure 422
// @Router /ads/{id} [delete]
func (ac AdvertisementController) DeleteAdvertisement(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnprocessableEntity, gin.H{"message": codes.BadID})
		return
	}
	advertisement := entities.Advertisement{ID: uint(id)}
	ac.advertisementStorage.DeleteAdvertisement(advertisement)
	c.JSON(http.StatusOK, gin.H{"message": codes.ResDeleted})
}

// @Tags Advertisements
// @Summary Обновить рекламное объявление
// @Description Метод обновляет рекламне объявление с заданным ID
// @Accept json
// @Produce  json
// @Param id path int true "ID рекламного объявления"
// @Param advertisement body entities.Advertisement true "Новые данные о рекламном объявлении"
// @Success 200 {object} responses.OneAdvertisement
// @Failure 404
// @Failure 422
// @Router /ads/{id} [put]
func (ac AdvertisementController) UpdateAdvertisement(c *gin.Context) {
	var advertisement entities.Advertisement
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnprocessableEntity, gin.H{"message": codes.BadID})
		return
	}
	advertisement, _ = ac.advertisementStorage.GetAdvertisement(uint(id))
	if advertisement.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{"message": codes.ResNotFound})
		return
	}

	c.ShouldBindWith(&advertisement, binding.JSON)
	_ = ac.advertisementStorage.UpdateAdvertisement(advertisement)
	c.JSON(http.StatusOK, gin.H{
		"status":     "OK",
		"advertisement": advertisement,
	})
}
