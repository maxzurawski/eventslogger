package dbprovider

import (
	"github.com/jinzhu/gorm"
	"github.com/maxzurawski/eventslogger/dto"
	"github.com/maxzurawski/eventslogger/model"
)

type DBManager interface {
	GetDb() *gorm.DB

	// mapper
	MapToEvent(dto dto.EventDTO) *model.Event
	MapToDto(input *model.Event) dto.EventDTO

	// CRUDs
	SaveEvent(dto dto.EventDTO) (*model.Event, error)

	// search function
	FindBy(dto dto.SearchDTO) ([]model.Event, error)
}

var Mgr DBManager

type manager struct {
	db *gorm.DB
}
