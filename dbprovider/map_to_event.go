package dbprovider

import (
	"github.com/xdevices/eventslogger/dto"
	"github.com/xdevices/eventslogger/model"
)

func (mgr *manager) MapToEvent(dto dto.EventDTO) *model.Event {
	value := dto.Value.ToString()
	return &model.Event{
		ProcessId:    &dto.ProcessId,
		SensorUuid:   &dto.SensorUuid,
		Topic:        &dto.Topic,
		RoutingKey:   &dto.RoutingKey,
		Cache:        &dto.Cache,
		Max:          &dto.Max,
		Min:          &dto.Min,
		Service:      &dto.Service,
		Value:        &value,
		LogMsg:       &dto.LogMsg,
		ErrorMsg:     &dto.ErrorMsg,
		ErrorDetails: &dto.ErrorDetails,
		Previous:     &dto.Previous,
		Current:      &dto.Current,
		PublishedOn:  &dto.PublishedOn,
	}
}
