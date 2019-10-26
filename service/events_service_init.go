package service

import "github.com/xdevices/eventslogger/dbprovider"

func Init() {
	s := service{}
	s.mgr = dbprovider.Mgr
	Service = &s
}
