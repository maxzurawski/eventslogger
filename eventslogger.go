package main

import (
	"github.com/labstack/echo"
	"github.com/xdevices/eventslogger/config"
	"github.com/xdevices/eventslogger/dbprovider"
	"github.com/xdevices/eventslogger/handlers"
	"github.com/xdevices/eventslogger/observers"
	"github.com/xdevices/eventslogger/publishers"
	"github.com/xdevices/eventslogger/service"
)

func main() {

	go observers.LogsObserver()

	e := echo.New()
	e.GET("/events/", handlers.FindEventsHandler)
	e.Logger.Fatal(e.Start(config.EventsloggerConfigManager().Address()))
}

func init() {
	manager := config.InitEventsloggerEurekaManager()
	manager.SendRegistrationOrFail()
	manager.ScheduleHeartBeat(config.EventsloggerConfigManager().ServiceName(), 10)
	dbprovider.InitDbManager()
	service.Init()
	publishers.InitLogger()
}
