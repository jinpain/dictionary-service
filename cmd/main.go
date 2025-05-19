package main

import (
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/his-vita/dictionary-service/internal/api/router"
	v1 "github.com/his-vita/dictionary-service/internal/api/v1"
	"github.com/his-vita/dictionary-service/internal/config"
	"github.com/his-vita/dictionary-service/internal/insurance"
	"github.com/rs/zerolog"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	logger := zerolog.New(os.Stderr).With().Timestamp().Logger()

	path := config.LoadPath()

	logger.Info().Interface("path", path).Msg("loaded config")

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

	logger.Info().Interface("dbCfg", dbCfg).Msg("db loaded")

	serverCfg := config.LoadServer(path.Server)

	insuranceService := insurance.NewService(db)

	insuranceHandler := v1.NewInsuranceHandler(insuranceService)

	server := gin.Default()

	rg := server.Group("/api/v1")

	router.Insurance(rg, insuranceHandler)

	if err := server.Run(fmt.Sprintf("%s:%v", serverCfg.Host, serverCfg.Port)); err != nil {
		panic(err)
	}
}
