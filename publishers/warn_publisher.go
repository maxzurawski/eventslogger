package publishers

import (
	"github.com/labstack/gommon/log"
	"github.com/maxzurawski/eventslogger/config"
	"github.com/maxzurawski/utilities/rabbit/crosscutting"
	"github.com/maxzurawski/utilities/rabbit/publishing"
)

type logger struct {
	*publishing.Publisher
}

var publisher *publishing.Publisher
var loggerInstance *logger

func InitLogger() {
	if publisher == nil && config.EventsloggerConfigManager().ConnectToRabbit() {
		publisher = config.EventsloggerConfigManager().InitPublisher()
		publisher.DeclareTopicExchange(crosscutting.TopicLogs.String())
	}
}

func Logger() *logger {
	if loggerInstance == nil {
		loggerInstance = new(logger)
		loggerInstance.Publisher = publisher
	}
	return loggerInstance
}

func (l *logger) Warn(processId, sensorUuid, msg string) {
	if !config.EventsloggerConfigManager().ConnectToRabbit() {
		log.Info("connection to rabbit disabled")
		return
	}

	l.Reset()
	l.PublishWarn(processId,
		sensorUuid,
		config.EventsloggerConfigManager().ServiceName(),
		msg,
	)
}
