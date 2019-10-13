package dto

import (
	"time"

	"github.com/xdevices/utilities/stringutils"
)

type EventDTO struct {
	ID           uint                    `json:"id"`
	ProcessId    string                  `json:"processId"`
	SensorUuid   string                  `json:"sensorUuid"`
	Topic        string                  `json:"topic"`
	RoutingKey   string                  `json:"routingKey"`
	Cache        string                  `json:"cache"`
	Max          string                  `json:"max"`
	Min          string                  `json:"min"`
	Service      string                  `json:"service"`
	Value        stringutils.MultiString `json:"value"`
	LogMsg       string                  `json:"logMsg"`
	ErrorMsg     string                  `json:"errorMsg"`
	ErrorDetails string                  `json:"errorDetails"`
	Previous     string                  `json:"previous"`
	Current      string                  `json:"current"`
	PublishedOn  time.Time               `json:"publishedOn"`
}
