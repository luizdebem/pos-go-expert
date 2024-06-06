package main

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Product struct {
	ID    int `gorm:primaryKey`
	Name  string
	Price float64
	gorm.Model
}

// gorm.Model: traz ID, created_at, updated_at e deleted_at de maneira automática.

func main() {
	dsn := "root:root@tcp(localhost:3306)/goexpert?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&Product{})

	// product := &Product{
	// 	Name:  "Notebook",
	// 	Price: 1899.90,
	// }
	// db.Create(product)

	// products := []Product{
	// 	{Name: "Notebook", Price: 1899.90},
	// 	{Name: "Tablet", Price: 999.90},
	// 	{Name: "Mouse", Price: 199.90},
	// }
	// db.Create(&products)

	// var product Product
	// db.First(&product, 1)
	// fmt.Println(product)

	// db.First(&product, "name = ?", "Mouse")
	// fmt.Println(product)

	// var products []Product
	// db.Limit(2).Offset(2).Find(&products)
	// for _, product := range products {
	// 	fmt.Println(product)
	// }

	// db.Where("price > ?", 1).Find(&products)
	// for _, product := range products {
	// 	fmt.Println(product)
	// }

	// db.Where("name LIKE ?", "%o%").Find(&products)
	// for _, product := range products {
	// 	fmt.Println(product)
	// }

	// var p Product
	// db.First(&p, 1)
	// p.Name = "New Notebook"
	// db.Save(&p)

	// var p2 Product
	// db.First(&p2, 1)
	// fmt.Println(p2)

	// db.Delete(&p2)

}
