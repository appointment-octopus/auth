package migration

import (
	"github.com/appointment-octopus/auth/core/db"
	"github.com/appointment-octopus/auth/services/models"
)

func AutoMigration(){
	db := db.PostgresConnect()
	db.Set("gorm:table_options", "ENGINE=InnoDB")
	if db.Migrator().HasTable(models.User{}) == false {
		db.AutoMigrate(models.User{})
	}
}
