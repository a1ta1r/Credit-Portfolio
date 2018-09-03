package services

import (
	"github.com/a1ta1r/Credit-Portfolio/internal/models"
	"github.com/a1ta1r/Credit-Portfolio/internal/storages"
	"time"
)

func NewAgendaService(storageContainer storages.StorageContainer) AgendaService {
	return AgendaService{storageContainer: storageContainer}
}

type AgendaService struct {
	storageContainer storages.StorageContainer
}

func (as AgendaService) GetElementsByTime(from time.Time, to time.Time) []models.AgendaElement {
	return []models.AgendaElement{}
}