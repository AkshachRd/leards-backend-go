package main

import (
	"encoding/base64"
	"github.com/AkshachRd/leards-backend-go/models"
	"github.com/gin-gonic/gin"
	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
	"strings"
)

func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		auth := strings.SplitN(c.Request.Header.Get("Authorization"), " ", 2)

		if len(auth) != 2 || auth[0] != "Basic" {
			respondWithError(401, "Unauthorized", c)
			return
		}
		payload, _ := base64.StdEncoding.DecodeString(auth[1])
		pair := strings.SplitN(string(payload), ":", 2)

		if len(pair) != 2 || !authenticateUser(pair[0], pair[1]) {
			respondWithError(401, "Unauthorized", c)
			return
		}

		c.Next()
	}
}

func authenticateUser(login, password string) bool {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	var user *models.User
	db.Where("name = ?", login).Or("email = ?", login).First(user)

	tx := db.First(&user)
	if tx.Error != nil {
		return false
	}
	return user.IsPasswordCorrect(password)
}

func respondWithError(code int, message string, c *gin.Context) {
	resp := map[string]string{"error": message}

	c.JSON(code, resp)
	c.Abort()
}
