package handlers

import (
	"encoding/base64"
	_ "github.com/AkshachRd/leards-backend-go/httputils"
	"github.com/AkshachRd/leards-backend-go/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	"strings"
)

type AuthType int

const (
	BasicAuth AuthType = iota
	BearerAuth
)

func (s *Server) AuthService(authType AuthType) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeaderContent := strings.SplitN(c.Request.Header.Get("Authorization"), " ", 2)

		if len(authHeaderContent) != 2 {
			respondWithError(http.StatusBadRequest, "Wrong authorization arguments count", c)
			return
		}

		payload, err := base64.StdEncoding.DecodeString(authHeaderContent[1])
		if err != nil {
			respondWithError(http.StatusInternalServerError, "Decode error", c)
			return
		}

		switch authType {
		case BasicAuth:
			parsedPayload := strings.SplitN(string(payload), ":", 2)
			if len(parsedPayload) != 2 {
				respondWithError(http.StatusUnauthorized, "Unauthorized", c)
				return
			}

			email, password := parsedPayload[0], parsedPayload[1]
			if !basicAuth(s.db, email, password) {
				respondWithError(http.StatusUnauthorized, "Unauthorized", c)
				return
			}

			c.Set("email", email)
			c.Set("password", password)
		case BearerAuth:
			token := string(payload)
			if !tokenAuth(s.db, token) {
				respondWithError(http.StatusUnauthorized, "Unauthorized", c)
				return
			}

			c.Set("token", token)
		default:
			respondWithError(http.StatusBadRequest, "Bad Request", c)
			return
		}

		c.Next()
	}
}

func basicAuth(db *gorm.DB, email, password string) bool {
	user, err := models.FetchUserByEmail(db, email)
	if err != nil {
		return false
	}

	return user.IsPasswordCorrect(password)
}

func tokenAuth(db *gorm.DB, token string) bool {
	user, err := models.FetchUserByToken(db, token)
	if err != nil {
		return false
	}

	return user.IsTokenValid()
}

func respondWithError(code int, message string, c *gin.Context) {
	resp := map[string]string{"error": message}

	c.JSON(code, resp)
	c.Abort()
}
