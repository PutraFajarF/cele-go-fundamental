package main

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	dataConn := "host=localhost user=postgres password=postgres db_name=project-go port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	db, err := gorm.Open(postgres.Open(dataConn), &gorm.Config{})
	if err != nil {
		log.Fatal(err.Error())
	}

}
