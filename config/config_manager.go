package config

import (
	"fmt"
	"os"

	"github.com/maxzurawski/utilities/config"
	"github.com/maxzurawski/utilities/rabbit"
)

type eventsloggerConfigManager struct {
	config.Manager
	proxyService string
	dbPath       string
	rabbit.RabbitMQManager
}

var instance *eventsloggerConfigManager

func EventsloggerConfigManager() *eventsloggerConfigManager {
	if instance == nil {
		instance = new(eventsloggerConfigManager)
		instance.Init()
	}
	return instance
}

func (e *eventsloggerConfigManager) Init() {
	e.Manager.Init()

	if proxyService, err := os.LookupEnv("PROXY_SERVICE"); !err {
		e.proxyService = "http://localhost:8000/api"
	} else {
		e.proxyService = proxyService
	}

	if dbPath, err := os.LookupEnv("DB_PATH"); !err {
		panic(fmt.Sprintf("set DB_PATH and try again"))
	} else {
		e.dbPath = dbPath
	}

	if e.ConnectToRabbit() {
		e.RabbitMQManager.InitConnection(e.RabbitURL())
	}
}

func (e *eventsloggerConfigManager) ProxyService() string {
	return e.proxyService
}

func (e *eventsloggerConfigManager) DBPath() string {
	return e.dbPath
}
