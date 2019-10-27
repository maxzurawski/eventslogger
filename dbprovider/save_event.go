package dbprovider

import (
	"github.com/maxzurawski/eventslogger/dto"
	"github.com/maxzurawski/eventslogger/model"
)

func (mgr *manager) SaveEvent(dto dto.EventDTO) (*model.Event, error) {
	event := mgr.MapToEvent(dto)
	mgr.db.Unscoped().Save(event)
	return event, nil
}
