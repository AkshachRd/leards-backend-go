package middlewares

import (
	"encoding/base64"
	_ "github.com/AkshachRd/leards-backend-go/httputils"
	"github.com/AkshachRd/leards-backend-go/models"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

type AuthType int

const (
	BasicAuth AuthType = iota
	BearerAuth
)

func Auth(authType AuthType) gin.HandlerFunc {
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
			user, isCorrectPassword := basicAuth(email, password)
			if !isCorrectPassword {
				respondWithError(http.StatusUnauthorized, "Unauthorized", c)
				return
			}

			c.Set("user_id", user.ID)
			c.Set("email", email)
			c.Set("password", password)
		case BearerAuth:
			token := string(payload)
			user, isValidToken := tokenAuth(token)
			if !isValidToken {
				respondWithError(http.StatusUnauthorized, "Unauthorized", c)
				return
			}

			c.Set("user_id", user.ID)
			c.Set("token", token)
		default:
			respondWithError(http.StatusBadRequest, "Bad Request", c)
			return
		}

		c.Next()
	}
}

func basicAuth(email, password string) (*models.User, bool) {
	user, err := models.FetchUserByEmail(email)
	if err != nil {
		return &models.User{}, false
	}

	isPasswordCorrect := user.IsPasswordCorrect(password)
	if !isPasswordCorrect {
		return &models.User{}, false
	}

	return user, isPasswordCorrect
}

func tokenAuth(token string) (*models.User, bool) {
	user, err := models.FetchUserByToken(token)
	if err != nil {
		return &models.User{}, false
	}

	isValidToken := user.IsTokenValid()
	if !isValidToken {
		return &models.User{}, false
	}

	return user, isValidToken
}

func respondWithError(code int, message string, c *gin.Context) {
	resp := map[string]string{"error": message}

	c.JSON(code, resp)
	c.Abort()
}
