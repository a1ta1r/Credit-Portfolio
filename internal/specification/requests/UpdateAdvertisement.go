package requests

import "github.com/a1ta1r/Credit-Portfolio/internal/components/advertisements/entities"

type UpdateAdvertisement struct {
	IsActive *bool   `json:"isActive"`
	Title    *string `json:"title"`
}

func (ua UpdateAdvertisement) ToAdvertisement(advertisement entities.Advertisement) entities.Advertisement {
	if ua.IsActive != nil {
		advertisement.IsActive = *ua.IsActive
	}
	if ua.Title != nil {
		advertisement.Title = *ua.Title
	}
	return advertisement
}
