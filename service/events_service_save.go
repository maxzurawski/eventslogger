package service

import (
	"github.com/xdevices/eventslogger/dbprovider"
	"github.com/xdevices/eventslogger/dto"
)

func (s *service) SaveEvent(dto dto.EventDTO) (*dto.EventDTO, error) {
	event, err := dbprovider.Mgr.SaveEvent(dto)
	if err != nil {
		return nil, err
	}
	result := s.mgr.MapToDto(event)
	return &result, nil
}
