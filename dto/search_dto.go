package dto

import "time"

type SearchDTO struct {
	ProcessId       string     `json:"processid"`
	Topic           string     `json:"topic"`
	RoutingKey      string     `json:"routing_key"`
	SensorsUuid     string     `json:"sensors_uuid"`
	ErrorMsg        string     `json:"error_msg"`
	Service         string     `json:"service"`
	PublishedOnFrom *time.Time `json:"published_on_from"`
	PublishedOnTo   *time.Time `json:"published_on_to"`
}
