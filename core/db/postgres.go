package db

import (
  "gorm.io/driver/postgres"
	"gorm.io/gorm"
	"fmt"
	"os"
)

func PostgresConnect() *gorm.DB {

	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable TimeZone=Brazil/East", os.Getenv("POSTGRES_HOST"), os.Getenv("POSTGRES_PORT"), os.Getenv("POSTGRES_USER"), os.Getenv("POSTGRES_PASSWORD"), os.Getenv("POSTGRES_DB"))
	db, err := gorm.Open(postgres.New(postgres.Config{
		DSN: dsn,
		PreferSimpleProtocol: true,
	}), &gorm.Config{}); if err != nil {
			fmt.Errorf("error connecting to database")
	}

	return db
}