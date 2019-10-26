package dbprovider

import (
	"github.com/xdevices/eventslogger/dto"
	"github.com/xdevices/eventslogger/model"
	"github.com/xdevices/utilities/stringutils"
)

func (mgr *manager) MapToDto(input *model.Event) dto.EventDTO {
	if input == nil {
		return dto.EventDTO{}
	}

	dto := dto.EventDTO{
		ID:           *input.ID,
		ProcessId:    *input.ProcessId,
		SensorUuid:   *input.SensorUuid,
		Topic:        *input.Topic,
		RoutingKey:   *input.RoutingKey,
		Cache:        *input.Cache,
		Max:          *input.Max,
		Min:          *input.Min,
		Service:      *input.Service,
		Value:        stringutils.ToMultiString(*input.Value),
		LogMsg:       *input.LogMsg,
		ErrorMsg:     *input.ErrorMsg,
		ErrorDetails: *input.ErrorDetails,
		Previous:     *input.Previous,
		Current:      *input.Current,
		PublishedOn:  *input.PublishedOn,
	}

	return dto
}
