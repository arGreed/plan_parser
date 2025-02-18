package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var dsn string = "host=localhost user=postgres password=admin dbname=postgres port=5432 sslmode=disable TimeZone=UTC"

func dbInit() (*gorm.DB, error) {
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	return db, nil
}

func main() {
	db, err := dbInit()
	if err != nil {
		log.Println("Ошибка инициализации базы данных")
		return
	}
	router := gin.Default()

	router.GET(getQueryRoute, getQuery(db))
}
