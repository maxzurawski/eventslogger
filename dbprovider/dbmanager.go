package dbprovider

import (
	"github.com/jinzhu/gorm"
	"github.com/xdevices/eventslogger/dto"
	"github.com/xdevices/eventslogger/model"
)

type DBManager interface {
	GetDb() *gorm.DB

	// mapper
	MapToEvent(dto dto.EventDTO) *model.Event

	// CRUDs
	SaveEvent(dto dto.EventDTO) (*model.Event, error)
}

var Mgr DBManager

type manager struct {
	db *gorm.DB
}
