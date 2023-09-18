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
		authHeaderContent := strings.SplitN(c.Request.Header.Get("Authorization"), " ", 2)

		switch {
		case len(authHeaderContent) != 2, authHeaderContent[0] != "Basic", authHeaderContent[0] != "Bearer":
			respondWithError(400, "Bad Request", c)
			return
		}

		payload, err := base64.StdEncoding.DecodeString(authHeaderContent[1])
		if err != nil {
			respondWithError(500, "Internal Server Error", c)
			return
		}

		db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
		if err != nil {
			respondWithError(500, "Internal Server Error", c)
			return
		}

		switch authHeaderContent[0] {
		case "Basic":
			parsedPayload := strings.SplitN(string(payload), ":", 2)
			if len(parsedPayload) != 2 || !basicAuth(db, parsedPayload[0], parsedPayload[1]) {
				respondWithError(401, "Unauthorized", c)
				return
			}
		case "Bearer":
			if !tokenAuth(db, string(payload)) {
				respondWithError(401, "Unauthorized", c)
				return
			}
		default:
			respondWithError(400, "Bad Request", c)
			return
		}

		c.Next()
	}
}

func basicAuth(db *gorm.DB, login, password string) bool {
	user, err := models.FetchUserByLogin(db, login)
	if err != nil {
		return false
	}

	return user.IsPasswordCorrect(password)
}

func tokenAuth(db *gorm.DB, token string) bool {
	_, err := models.FetchUserByToken(db, token)
	if err != nil {
		return false
	}

	return true
}

func respondWithError(code int, message string, c *gin.Context) {
	resp := map[string]string{"error": message}

	c.JSON(code, resp)
	c.Abort()
}
