package main

import (
	"github.com/AkshachRd/leards-backend-go/models"
	"github.com/gin-gonic/gin"
	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
	"net/http"
)

func main() {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// Migrate the schema
	err = db.AutoMigrate(&models.AccessType{}, &models.Card{}, &models.Deck{}, &models.Folder{}, &models.Permission{}, &models.PermissionType{}, &models.User{})
	if err != nil {
		return
	}
	user, err := models.NewUser("Admin", "admin@leards.space", "123")
	if err != nil {
		return
	}
	// Create
	db.Create(user)
	//
	//// Read
	//var product Product
	//db.First(&product, 1)                 // find product with integer primary key
	//db.First(&product, "code = ?", "D42") // find product with code D42
	//
	//// Update - update product's price to 200
	//db.Model(&product).Update("Price", 200)
	//// Update - update multiple fields
	//db.Model(&product).Updates(Product{Price: 200, Code: "F42"}) // non-zero fields
	//db.Model(&product).Updates(map[string]interface{}{"Price": 200, "Code": "F42"})
	//
	//// Delete - delete product
	//db.Delete(&product, 1)

	r := gin.Default()

	// Group using gin.BasicAuth() middleware
	// gin.Accounts is a shortcut for map[string]string
	authorized := r.Group("/", Auth())

	// hit "localhost:8080/admin/dashboard
	authorized.GET("/hello", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"success": true})
	})

	// Listen and serve on 0.0.0.0:8080
	r.Run(":8080")
}
