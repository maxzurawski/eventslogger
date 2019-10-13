package dbprovider

import (
	"fmt"
	"log"

	"github.com/jinzhu/gorm"
	"github.com/xdevices/eventslogger/model"
	"github.com/xdevices/register/config"
	"github.com/xdevices/utilities/db"

	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

func InitDbManager() {
	dbPath := config.Config().DBPath()

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
