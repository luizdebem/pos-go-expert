// Golang Standards Project Layout no GitHub.

package main

import (
	"10-apis/configs"
	"10-apis/internal/entity"
	"10-apis/internal/infra/db"
	"10-apis/internal/infra/webserver/handlers"
	"net/http"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {
	config, err := configs.LoadConfig(".")
	if err != nil {
		panic(err)
	}
	println(config)

	database, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	database.AutoMigrate(&entity.Product{}, &entity.User{})

	productDB := db.NewProductDB(database)
	productHandler := handlers.NewProductHandler(productDB)

	http.HandleFunc("/products", productHandler.CreateProduct)
	http.ListenAndServe(":8000", nil)
}

// Ver Go Fiber, Go Echo, Gin
// frameworks web server

// Ver Buffalo
// framework Go completo
