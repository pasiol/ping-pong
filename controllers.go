package main

import (
	"database/sql"
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"os"
	"time"
)

func initializeDb() (*gorm.DB, *sql.DB, error) {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=5432 sslmode=disable TimeZone=Europe/Helsinki", os.Getenv("POSTGRES_HOST"), os.Getenv("POSTGRES_USER"), os.Getenv("POSTGRES_PASSWORD"), os.Getenv("POSTGRES_DB"))
	pgConf := postgres.Config{DSN: dsn, PreferSimpleProtocol: true}
	db, err := gorm.Open(postgres.New(pgConf), &gorm.Config{})
	if err != nil {
		for n := 1; n <= 5; n++ {
			db, err = gorm.Open(postgres.New(pgConf), &gorm.Config{})
			if err != nil && n == 5 {
				return nil, nil, err
			} else if err == nil {
				break
			}
			time.Sleep(time.Duration(10 * n))
		}
	}
	sqlDB, err := db.DB()
	sqlDB.SetMaxIdleConns(2)
	sqlDB.SetMaxOpenConns(10)
	sqlDB.SetConnMaxLifetime(1)

	err = db.AutoMigrate(&PingPong{})
	if err != nil {
		return nil, nil, err
	}

	var pp PingPong

	result := db.FirstOrCreate(&pp, PingPong{Counter: 0})
	if result.Error != nil {
		return nil, nil, result.Error
	}
	log.Printf("counter is %d", pp.Counter)

	return db, sqlDB, err
}
