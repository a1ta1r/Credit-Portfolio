package controllers

import (
	"github.com/a1ta1r/Credit-Portfolio/internal/codes"
	"github.com/a1ta1r/Credit-Portfolio/internal/models"
	"github.com/a1ta1r/Credit-Portfolio/internal/storages"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"net/http"
	"strconv"
)

type AdvertisementsController struct {
	advertiserStorage    storages.AdvertiserStorage
	advertisementStorage storages.AdvertisementStorage
	bannerStorage        storages.BannerStorage
	bannerPlaceStorage   storages.BannerPlaceStorage
}

func NewAdvertisementController(
	advertiserStorage storages.AdvertiserStorage,
	advertisementStorage storages.AdvertisementStorage,
	bannerStorage storages.BannerStorage,
	bannerPlaceStorage storages.BannerPlaceStorage,
) AdvertisementsController {
	return AdvertisementsController{
		advertiserStorage:    advertiserStorage,
		advertisementStorage: advertisementStorage,
		bannerStorage:        bannerStorage,
		bannerPlaceStorage:   bannerPlaceStorage,
	}
}

func (ac AdvertisementsController) GetAdvertisers(c *gin.Context) {
	var advertisers []models.Advertiser
	advertisers, _ = ac.advertiserStorage.GetAdvertisers()
	c.JSON(http.StatusOK, gin.H{
		"status":      "OK",
		"count":       len(advertisers),
		"advertisers": advertisers,
	})
}

func (ac AdvertisementsController) GetAdvertiser(c *gin.Context) {
	var advertiser models.Advertiser
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnprocessableEntity, gin.H{"message": codes.BadID})
		return
	}
	advertiser, _ = ac.advertiserStorage.GetAdvertiser(uint(id))
	if advertiser.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{"message": codes.ResNotFound})
		return
	}
	c.JSON(http.StatusOK, gin.H{"advertiser": advertiser})
}

func (ac AdvertisementsController) AddAdvertiser(c *gin.Context) {
	var advertiser models.Advertiser
	c.BindJSON(&advertiser)
	err := ac.advertiserStorage.CreateAdvertiser(advertiser)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": codes.InternalError})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"advertiser": advertiser})
}

func (ac AdvertisementsController) DeleteAdvertiser(c *gin.Context) {
	var advertiser models.Advertiser
	c.BindJSON(&advertiser)
	ac.advertiserStorage.DeleteAdvertiser(advertiser)
	c.JSON(http.StatusOK, gin.H{"message": codes.ResDeleted})
}

func (ac AdvertisementsController) UpdateAdvertiser(c *gin.Context) {
	var advertiser models.Advertiser
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnprocessableEntity, gin.H{"message": codes.BadID})
		return
	}
	advertiser, _ = ac.advertiserStorage.GetAdvertiser(uint(id))
	if advertiser.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{"message": codes.ResNotFound})
		return
	}

	c.ShouldBindWith(&advertiser, binding.JSON)
	_ = ac.advertiserStorage.UpdateAdvertiser(advertiser)
	c.JSON(http.StatusOK, gin.H{
		"status":     "OK",
		"advertiser": advertiser,
	})
}

func (ac AdvertisementsController) GetAdvertisementsByAdvertiser(c *gin.Context) {
	var advertisements []models.Advertisement
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

func (ac AdvertisementsController) GetAdvertisements(c *gin.Context) {
	var advertisement []models.Advertisement
	advertisement, _ = ac.advertisementStorage.GetAdvertisements()
	c.JSON(http.StatusOK, gin.H{
		"status":      "OK",
		"count":       len(advertisement),
		"advertisers": advertisement,
	})
}

func (ac AdvertisementsController) GetAdvertisement(c *gin.Context) {
	var advertisement models.Advertiser
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnprocessableEntity, gin.H{"message": codes.BadID})
		return
	}
	advertisement, _ = ac.advertiserStorage.GetAdvertiser(uint(id))
	if advertisement.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{"message": codes.ResNotFound})
		return
	}
	c.JSON(http.StatusOK, gin.H{"advertiser": advertisement})
}

func (ac AdvertisementsController) AddAdvertisement(c *gin.Context) {
	var advertiser models.Advertiser
	c.BindJSON(&advertiser)
	err := ac.advertiserStorage.CreateAdvertiser(advertiser)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": codes.InternalError})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"advertiser": advertiser})
}

func (ac AdvertisementsController) DeleteAdvertisement(c *gin.Context) {
	var advertiser models.Advertiser
	c.BindJSON(&advertiser)
	ac.advertiserStorage.DeleteAdvertiser(advertiser)
	c.JSON(http.StatusOK, gin.H{"message": codes.ResDeleted})
}

func (ac AdvertisementsController) UpdateAdvertisement(c *gin.Context) {
	var advertiser models.Advertiser
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnprocessableEntity, gin.H{"message": codes.BadID})
		return
	}
	advertiser, _ = ac.advertiserStorage.GetAdvertiser(uint(id))
	if advertiser.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{"message": codes.ResNotFound})
		return
	}

	c.ShouldBindWith(&advertiser, binding.JSON)
	_ = ac.advertiserStorage.UpdateAdvertiser(advertiser)
	c.JSON(http.StatusOK, gin.H{
		"status":     "OK",
		"advertiser": advertiser,
	})
}

func (ac AdvertisementsController) GetBannersByAdvertisement(c *gin.Context) {
	var banners []models.Banner
	id, err := strconv.ParseUint(c.Param("adsid"), 10, 32)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnprocessableEntity, gin.H{"message": codes.BadID})
		return
	}
	banners, _ = ac.bannerStorage.GetBannersByAdvertisement(uint(id))
	c.JSON(http.StatusOK, gin.H{
		"status":  "OK",
		"count":   len(banners),
		"banners": banners,
	})
}
