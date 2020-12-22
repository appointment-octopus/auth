package migration

import (
	"github.com/appointment-octopus/auth/core/db"
	"github.com/appointment-octopus/auth/services/models"
)

func AutoMigration(){
	db := db.PostgresConnect()
	// defer db.Close()
	db.AutoMigrate(models.User{})
}
