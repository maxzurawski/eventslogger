package config

import "github.com/maxzurawski/utilities/discovery"

type EventsloggerEurekaManager struct {
	discovery.Manager
}

func InitEventsloggerEurekaManager() *EventsloggerEurekaManager {
	manager := EventsloggerEurekaManager{
		Manager: discovery.Manager{
			RegistrationTicket: EventsloggerConfigManager().RegistrationTicket(),
			EurekaService:      EventsloggerConfigManager().EurekaService(),
		},
	}
	return &manager
}
