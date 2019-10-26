package main

import (
	"github.com/labstack/echo"
	"github.com/maxzurawski/eventslogger/config"
	"github.com/maxzurawski/eventslogger/dbprovider"
	"github.com/maxzurawski/eventslogger/handlers"
	"github.com/maxzurawski/eventslogger/observers"
	"github.com/maxzurawski/eventslogger/publishers"
	"github.com/maxzurawski/eventslogger/service"
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
