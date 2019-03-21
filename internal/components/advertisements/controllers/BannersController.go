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
	"math/rand"
	"net/http"
	"strconv"
	"time"
)

type BannersController struct {
	bannerStorage storages.BannerStorage
	rnd rand.Rand
}

func NewBannersController(bs storages.BannerStorage) BannersController {
	s := rand.NewSource(time.Now().Unix())
	r := rand.New(s)
	return BannersController{bannerStorage: bs, rnd:*r}
}

// @Tags Banners
// @Summary Получить баннеры по ID рекламного объявления
// @Description Метод возвращает набор баннеров по ID рекламного объявления.
// @Accept json
// @Produce  json
// @Param id path int true "ID рекламного объявления"
// @Success 200 {object} responses.BannersByAds
// @Failure 404
// @Failure 422
// @Router /promotions/{id}/banners [get]
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

// @Tags Banners
// @Summary Получить баннер по ID
// @Description Метод возвращает один баннер c заданным ID
// @Accept json
// @Produce  json
// @Param id path int true "ID баннера"
// @Success 200 {object} responses.OneBanner
// @Failure 404
// @Failure 422
// @Router /banners/{id} [get]
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

// @Tags Banners
// @Summary Удалить баннер по ID
// @Description Метод удаляет один баннер c заданным ID
// @Accept json
// @Produce  json
// @Param id path int true "ID баннера"
// @Success 200
// @Failure 422
// @Router /banners/{id} [delete]
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

// @Tags Banners
// @Summary Добавить новый баннер
// @Description Метод создает один баннер c заданными параметрами
// @Accept json
// @Produce  json
// @Param advertiser body requests.NewBanner true "Данные о баннере"
// @Success 200 {object} responses.OneBanner
// @Failure 422
// @Router /banners [post]
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

// @Tags Banners
// @Summary Добавить новый баннер
// @Description Метод создает один баннер c заданными параметрами
// @Accept json
// @Produce  json
// @Param id path int true "ID баннера"
// @Param advertiser body requests.UpdateBanner true "Данные о баннере"
// @Success 200 {object} responses.OneBanner
// @Failure 404
// @Failure 422
// @Router /banners/{id} [put]
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

func (bc BannersController) IncrViewsForBanner(c *gin.Context) {
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
	banner.Views++
	_ = bc.bannerStorage.UpdateBanner(&banner)
	banner, err = bc.bannerStorage.GetBanner(banner.ID)
	c.JSON(http.StatusOK, gin.H{
		"status": "OK",
		"banner": banner,
	})
}

func (bc BannersController) IncrClicksForBanner(c *gin.Context) {
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
	banner.Clicks++
	_ = bc.bannerStorage.UpdateBanner(&banner)
	banner, err = bc.bannerStorage.GetBanner(banner.ID)
	c.JSON(http.StatusOK, gin.H{
		"status": "OK",
		"banner": banner,
	})
}

func (bc BannersController) GetRandomBanner(c *gin.Context) {

	banners, _ := bc.bannerStorage.GetBanners()
	if len(banners)==0 {
		c.JSON(http.StatusNotFound, gin.H{"error": codes.ResNotFound})
		return
	}
	var activeBanners []entities.Banner
	for _, b := range banners {
		if b.IsVisible {
			activeBanners = append(activeBanners, b)
		}
	}
	banner := activeBanners[bc.rnd.Intn(len(activeBanners))]

	if banner.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": codes.ResNotFound})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status": "OK",
		"banner": banner,
	})
}
