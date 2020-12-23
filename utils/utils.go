package utils

import (
	"os"
	"log"
)

func RootDir() string {
	dir, err := os.Getwd(); if err != nil {
		log.Println("error getting root dir")
		panic(err)
	}

	return dir
}
