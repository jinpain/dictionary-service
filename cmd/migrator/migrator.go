package main

import (
	"fmt"
	"log"

	"github.com/his-vita/dictionary-service/internal/config"
	"github.com/his-vita/dictionary-service/internal/insurance"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	path := config.LoadPath()

	dbCfg := config.LoadDB(path.DB)

	connStr := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s",
		dbCfg.Host, dbCfg.Port, dbCfg.User, dbCfg.Password, dbCfg.DbName)

	db, err := gorm.Open(postgres.New(postgres.Config{
		DSN:                  connStr,
		PreferSimpleProtocol: true,
	}), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	modelsToMigrate := []interface{}{
		&insurance.Model{},
	}

	for _, model := range modelsToMigrate {
		if err := db.AutoMigrate(model); err != nil {
			log.Fatalf("failed to migrate model %T: %v", model, err)
		}
	}

	log.Println("Database migrated successfully!")
}
