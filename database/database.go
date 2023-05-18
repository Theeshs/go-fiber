package database

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func openDB(dsn string) (*gorm.DB, error) {
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect with db")
	}
	log.Println("Connected successfully to DB")

	db.Logger = logger.Default.LogMode(logger.Info)

	return db, nil

}

func ConnectDB(dsn string) (*gorm.DB, error) {
	conncetion, err := openDB(dsn)

	if err != nil {
		return nil, err
	}

	log.Println("Connected to Postgress")
	return conncetion, nil
}
