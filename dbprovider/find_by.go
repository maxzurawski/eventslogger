package dbprovider

import (
	"github.com/xdevices/eventslogger/dto"
	"github.com/xdevices/eventslogger/model"
	"github.com/xdevices/utilities/stringutils"
)

func (mgr *manager) FindBy(searchDTO dto.SearchDTO) ([]model.Event, error) {

	var events []model.Event

	db := mgr.db

	if !stringutils.IsZero(searchDTO.ProcessId) {
		db = db.Where("processid LIKE ?", "%"+searchDTO.ProcessId+"%")
	}

	if !stringutils.IsZero(searchDTO.RoutingKey) {
		db = db.Where("routing_key LIKE ?", "%"+searchDTO.RoutingKey+"%")
	}

	if !stringutils.IsZero(searchDTO.Topic) {
		db = db.Where("topic LIKE ?", "%"+searchDTO.Topic+"%")
	}

	if !stringutils.IsZero(searchDTO.SensorsUuid) {
		db = db.Where("sensor_uuid LIKE ?", "%"+searchDTO.SensorsUuid+"%")
	}

	if !stringutils.IsZero(searchDTO.Service) {
		db = db.Where("service LIKE ?", "%"+searchDTO.Service+"%")
	}

	if searchDTO.PublishedOnFrom != nil {
		db = db.Where("published_on >= ?", searchDTO.PublishedOnFrom)
	}

	if searchDTO.PublishedOnTo != nil {
		db = db.Where("published_on <= ?", searchDTO.PublishedOnTo)
	}

	err := db.Order("published_on desc").Find(&events).Error
	return events, err
}
