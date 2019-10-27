package service

import (
	"github.com/maxzurawski/eventslogger/dbprovider"
	"github.com/maxzurawski/eventslogger/dto"
)

type eventsService interface {
	SaveEvent(dto dto.EventDTO) (*dto.EventDTO, error)
	FindBy(dto dto.SearchDTO) (result []dto.EventDTO, err error)
}

var Service eventsService

type service struct {
	mgr dbprovider.DBManager
}
