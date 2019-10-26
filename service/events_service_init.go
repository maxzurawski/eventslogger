package service

import "github.com/maxzurawski/eventslogger/dbprovider"

func Init() {
	s := service{}
	s.mgr = dbprovider.Mgr
	Service = &s
}
