package models

import (
	"encoding/json"
	"log"

	database "github.com/appointment-octopus/auth/core/db"
)

type User struct {
	IdUser      int    `gorm:"primary_key;auto_increment;Column:iduser" json:"idUser" form:"idUser"`
	CPF         string `gorm:"size:11;not null;unique;Column:cpf" json:"CPF" form:"CPF"`
	Gender      string `gorm:"size:15;not null;Column:gender" json:"gender" form:"gender"`
	DateOfBirth string `gorm:"not null;Column:dateofbirth" json:"dateOfBirth" form:"dateOfBirth"`
	Email       string `gorm:"size:150;not null;unique;Column:email" json:"email" form:"email"`
	FullName    string `gorm:"size:150;not null;Column:fullname" json:"fullName" form:"fullName"`
	Password    string `gorm:"size:255;not null;Column:password" json:"password" form:"password"`
}

func (User) TableName() string {
	return "users"
}

func (user *User) CreateUser() (*User, error) {
	pgConn := database.PostgresConnect()

	var err error
	err = pgConn.Debug().Model(&User{}).Create(&user).Error
	if err != nil {
		return &User{}, err
	}

	return user, nil
}

func (user *User) FindUser() (*User, error) {
	pgConn := database.PostgresConnect()

	userFound, err := database.RedisGetValue(user.Email)
	if err != nil {
		log.Println("searching in postgres")
		var err error
		err = pgConn.Debug().Where(&User{Email: user.Email}).Find(&user).Error
		if err != nil {
			return user, err
		}
		userBytes, _ := json.Marshal(user)
		database.RedisSetValue(user.Email, userBytes)
		return user, nil
	}

	log.Println("found in redis")
	json.Unmarshal(userFound, user)
	return user, nil
}
