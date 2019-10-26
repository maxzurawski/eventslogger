package service

import "github.com/maxzurawski/eventslogger/dto"

func (s *service) FindBy(dto dto.SearchDTO) (result []dto.EventDTO, err error) {
	events, err := s.mgr.FindBy(dto)
	if err != nil {
		return
	}
	for _, item := range events {
		result = append(result, s.mgr.MapToDto(&item))
	}
	return
}
