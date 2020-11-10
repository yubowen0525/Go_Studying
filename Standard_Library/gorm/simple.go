package main

import (
	"fmt"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// gorm.model
type User struct {
	gorm.Model
	Name  string `gorm:"unique"`
	Price uint
}

func main() {
	Start()
}

func Start() {
	db, err := gorm.Open(sqlite.Open("./test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// Migrate the schema
	result1 := db.AutoMigrate(&User{})
	fmt.Println(result1)
	// Create
	if err := db.Create(&User{Name: "Bowen Yu", Price: 100}).Error; err != nil {
		// 错误处理...
		panic("! user already exists ")
	}
	//test := err.Statement.Model
	//value := reflect.ValueOf(test)
	//fmt.Println(value)
	fmt.Println(err)
	// Read 读进数据Mod user里
	var user User
	db.First(&user, 1) // find product with integer primary key

	//db.First(&user, "name = ?", "bowen") // find product with code D42

	//Update - update product's price to 200
	db.Model(&user).Update("Price", 200)
	// Update - update multiple fields
	//db.Model(&user).Updates(User{Price: 200}) // non-zero fields
	//db.Model(&user).Updates(map[string]interface{}{"Price": 200, "Code": "F42"})

	// Delete - delete product
	//db.Delete(&user, 1)
}
