package dbprovider

import (
	"github.com/xdevices/eventslogger/dto"
	"github.com/xdevices/eventslogger/model"
)

func (mgr *manager) SaveEvent(dto dto.EventDTO) (*model.Event, error) {
	event := mgr.MapToEvent(dto)
	mgr.db.Unscoped().Save(event)
	return event, nil
}
