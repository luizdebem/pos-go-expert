package main

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// A categoria has many products.

type Category struct {
	ID       int `gorm:primaryKey`
	Name     string
	Products []Product `gorm:"many2many:products_categories;"`
}

// A product belongs to a category.

type Product struct {
	ID           int `gorm:primaryKey`
	Name         string
	Price        float64
	Categories   []Category `gorm:"many2many:products_categories;"`
	SerialNumber SerialNumber
	gorm.Model
}

type SerialNumber struct {
	ID        int `gorm:primaryKey`
	Number    string
	ProductID int
}

func main() {
	dsn := "root:root@tcp(localhost:3306)/goexpert?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(&Product{}, &Category{}, &SerialNumber{})

	// category := Category{Name: "Eletronicos"}
	// db.Create(&category)

	// category2 := Category{Name: "Cozinha"}
	// db.Create(&category2)

	// db.Create(&Product{
	// 	Name:       "Mouse",
	// 	Price:      100.0,
	// 	Categories: []Category{category, category2},
	// })

	var categories []Category
	err = db.Model(&Category{}).Preload("Products").Find(&categories).Error
	if err != nil {
		panic(err)
	}
	for _, c := range categories {
		fmt.Println(c.Name)
		for _, p := range c.Products {
			fmt.Println("- ", p.Name)
		}
	}

}
