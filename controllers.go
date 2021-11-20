package main

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"os"
)

func initializeDb() (*gorm.DB, error) {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=5432 sslmode=disable TimeZone=Europe/Helsinki", os.Getenv("POSTGRES_HOST"), os.Getenv("POSTGRES_USER"), os.Getenv("POSTGRES_PASSWORD"), os.Getenv("POSTGRES_DB"))
	pgConf := postgres.Config{DSN: dsn, PreferSimpleProtocol: true}
	db, err := gorm.Open(postgres.New(pgConf), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	err = db.AutoMigrate(&PingPong{})
	if err != nil {
		return nil, err
	}

	var pp PingPong

	result := db.FirstOrCreate(&pp, PingPong{Counter: 0})
	if result.Error != nil {
		return nil, result.Error
	}
	log.Printf("counter is %d", pp.Counter)

	return db, err
}
