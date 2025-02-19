package main

import (
	"fmt"
	"log"
	"os"
	"strings"

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

var filename string = "coupleQueries.txt"
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
	if err != nil {
		log.Println(err)
		return
	}

	queries := strings.Split(content, ";")
	fmt.Println(querySorter(&queries))
}

////  Сильно упрощённый перечень служебных слов, для определения назначения запроса
//func getStartQueryWords() map[string]bool {
//	words := make(map[string]bool)
//	words["select"] = true
//	words["create"] = true
//	words["insert"] = true
//	words["update"] = true
//	words["delete"] = true
//	words["with"] = true
//	words["drop"] = true
//	words["alter"] = true
//	words["truncate"] = true

//	return words
//}

// ? Сильно упрощённый сортировщик запросов
func querySorter(queries *[]string) ([]string, []string, []string, []string) {
	var create []string
	var delete []string
	var get []string
	var others []string

	//! Нереально упрощённая версия
	for _, query := range *queries {
		words := strings.Fields(strings.ToLower(query))
		if len(words) == 0 {
			others = append(others, query)
			continue
		}
		firstWord := words[0]
		switch firstWord {
		case "select":
			fmt.Println(1)
			get = append(get, query)
		case "create":
			fmt.Println(2)
			create = append(create, query)
		case "delete":
			fmt.Println(3)
			delete = append(delete, query)
		default:
			fmt.Println(4)
			others = append(others, query)
		}
	}
	return create, delete, get, others
}
