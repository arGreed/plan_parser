package main

import (
	"fmt"
	"log"
	"os"

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

var filename string = "test_delete_all_comments.txt"
var Line string = "_____________"

func main() {
	data, err := os.ReadFile(filename)
	if err != nil {
		log.Println("Ошибка при чтении файла:", err)
		return
	}

	content := string(data)
	//? Очистка полученного файла от лишних комментариев
	err = clearQueries(&content, "--")
	fmt.Println(content)
	if err != nil {
		log.Println(err)
		return
	}
}
