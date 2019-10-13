package observers

import (
	"encoding/json"

	"github.com/xdevices/eventslogger/config"
	"github.com/xdevices/eventslogger/dto"
	"github.com/xdevices/eventslogger/service"
	"github.com/xdevices/utilities/rabbit/crosscutting"
)

func LogsObserver() {
	observer := config.EventsloggerConfigManager().InitObserver()
	defer config.EventsloggerConfigManager().RabbitMQManager.CloseConnection()
	defer observer.Channel.Close()

	observer.DeclareTopicExchange(crosscutting.TopicLogs.String())
	observer.BindQueue(observer.Queue, "#", crosscutting.TopicLogs.String())
	deliveries := observer.Observe()

	for msg := range deliveries {
		eventDTO := dto.EventDTO{}
		_ = json.Unmarshal(msg.Body, &eventDTO)
		eventDTO.Topic = msg.Exchange
		eventDTO.RoutingKey = msg.RoutingKey
		_, _ = service.Service.SaveEvent(eventDTO)
	}
}
