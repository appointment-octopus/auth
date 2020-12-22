package models

import (
	database "github.com/appointment-octopus/auth/core/db"
	"encoding/json"
	"log"
)

type User struct {
 UUID     string `gorm:"primary_key;auto_increment" json:"uuid" form:"-"`
 Username string `gorm:"size:255;not null;unique" json:"username" form:"username"`
 Password string `gorm:"size:255;not null" json:"password" form:"password"`
}

func (user *User) CreateUser() (*User, error) {
	pgConn := database.PostgresConnect()

	var err error
	err = pgConn.Debug().Model(&User{}).Create(&user).Error; if err != nil {
		return &User{}, err
	}

	return user, nil
}

func (user *User) FindUser() (*User, error) {
	pgConn := database.PostgresConnect()

	userFound, err := database.RedisGetValue(user.Username); if err != nil {
		log.Println("searching in postgres")
		var err error
		err = pgConn.Debug().Where(&User{Username: user.Username}).Find(&user).Error; if err != nil {
			return user, err
		}
		userBytes, _ := json.Marshal(user)
		database.RedisSetValue(user.Username, userBytes)
		return user, nil
	}

	log.Println("found in redis")
	json.Unmarshal(userFound, user)
	return user, nil
}
