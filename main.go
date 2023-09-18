package main

import (
	"github.com/AkshachRd/leards-backend-go/handlers"
	"github.com/AkshachRd/leards-backend-go/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"log"
)

func main() {
	r := SetupRouter()

	r.Run(":8080")
}

func DbInit() *gorm.DB {
	db, err := models.Setup()
	if err != nil {
		log.Println("Problem setting up database")
	}
	return db
}

func SetupRouter() *gin.Engine {
	r := gin.Default()

	db := DbInit()

	server := handlers.NewServer(db)

	router := r.Group("/api")
	authorizedRouter := r.Group("/api", server.AuthService())

	router.POST("/register", server.Register)
	authorizedRouter.POST("/refresh-token", server.RefreshToken)
	authorizedRouter.POST("/revoke-token", server.RevokeToken)

	return r
}
