package service

import (
	"github.com/xdevices/eventslogger/dbprovider"
	"github.com/xdevices/eventslogger/dto"
)

type eventsService interface {
	SaveEvent(dto dto.EventDTO) (*dto.EventDTO, error)
	FindBy(dto dto.SearchDTO) (result []dto.EventDTO, err error)
}

var Service eventsService

type service struct {
	mgr dbprovider.DBManager
}
