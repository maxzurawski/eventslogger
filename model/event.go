package model

import (
	"time"
)

type Event struct {
	ID           *uint      `gorm:"primary_key"`
	ProcessId    *string    `gorm:"column:processid"`
	SensorUuid   *string    `gorm:"column:sensor_uuid"`
	Topic        *string    `gorm:"column:topic"`
	RoutingKey   *string    `gorm:"column:routing_key"`
	Cache        *string    `gorm:"column:cache"`
	Max          *string    `gorm:"column:max"`
	Min          *string    `gorm:"column:min"`
	Service      *string    `gorm:"column:service"`
	Value        *string    `gorm:"column:value"`
	LogMsg       *string    `gorm:"column:log_msg"`
	ErrorMsg     *string    `gorm:"column:error_msg"`
	ErrorDetails *string    `gorm:"column:error_details"`
	Previous     *string    `gorm:"column:previous"`
	Current      *string    `gorm:"column:current"`
	PublishedOn  *time.Time `gorm:"column:published_on"`
}
