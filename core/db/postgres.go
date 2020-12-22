package db

import (
  "gorm.io/driver/postgres"
	"gorm.io/gorm"
	"fmt"
	"os"
)

func PostgresConnect() *gorm.DB {

	dsn := fmt.Sprintf("user=%s password=%s DB.name=%s port=%s sslmode=disable TimeZone=Brazil/East", os.Getenv("POSTGRES_USER"), os.Getenv("POSTGRES_PASSWORD"), os.Getenv("POSTGRES_HOST"), os.Getenv("POSTGRES_PORT"))

	db, err := gorm.Open(postgres.New(postgres.Config{
		DSN: dsn,
		PreferSimpleProtocol: true,
	}), &gorm.Config{}); if err != nil {
			fmt.Errorf("error connecting to database")
	}

	return db
}