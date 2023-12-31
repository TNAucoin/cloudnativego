package main

import (
	"fmt"
	"github.com/tnaucoin/cloudnativego/api/router"
	"github.com/tnaucoin/cloudnativego/config"
	"github.com/tnaucoin/cloudnativego/util/validator"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	gormlogger "gorm.io/gorm/logger"
	"log"
	"net/http"
)

const fmtDBString = "host=%s user=%s password=%s dbname=%s port=%d sslmode=disable"

// @title CloudNativeGo
// @version 1.0
// @description CloudNative API written in GO
// @contact.name Travis Aucoin
// @contact.url github.com/tnaucoin/cloudnativego
// @host localhost:8080
// @basePath /v1
func main() {
	c := config.New()
	v := validator.New()

	logLevel := gormlogger.Error
	if c.DB.Debug {
		logLevel = gormlogger.Info
	}
	dbString := fmt.Sprintf(fmtDBString, c.DB.Host, c.DB.Username, c.DB.Password, c.DB.DBName, c.DB.Port)
	db, err := gorm.Open(postgres.Open(dbString), &gorm.Config{Logger: gormlogger.Default.LogMode(logLevel)})
	if err != nil {
		log.Fatal("db connection start failure")
		return
	}

	r := router.New(db, v)
	s := &http.Server{
		Addr:         fmt.Sprintf(":%d", c.Server.Port),
		Handler:      r,
		ReadTimeout:  c.Server.TimeoutRead,
		WriteTimeout: c.Server.TimeoutWrite,
		IdleTimeout:  c.Server.TimeoutIdle,
	}
	log.Println("Starting server " + s.Addr)
	if err := s.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatal("server startup failed")
	}
}
