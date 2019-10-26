package dbprovider

import (
	"fmt"
	"log"

	"github.com/maxzurawski/eventslogger/config"

	"github.com/jinzhu/gorm"
	"github.com/maxzurawski/eventslogger/model"
	"github.com/maxzurawski/utilities/db"

	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

func InitDbManager() {
	dbPath := config.EventsloggerConfigManager().DBPath()

	if path, exists := db.AdjustDBPath(dbPath); !exists {
		dbPath = path
	}

	db, err := gorm.Open("sqlite3", dbPath)
	if err != nil {
		errorMsg := fmt.Sprintf("Failed to init db[%s]:", dbPath)
		log.Fatal(errorMsg, err)
	}

	db.SingularTable(true)
	db.AutoMigrate(&model.Event{})
	Mgr = &manager{db: db}
}
