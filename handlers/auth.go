package handlers

import (
	"github.com/AkshachRd/leards-backend-go/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

type RegisterInput struct {
	Username string `json:"username" binding:"required"`
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func (s *Server) Register(c *gin.Context) {
	var input RegisterInput

	if err := c.ShouldBind(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input or username/email already exists"})
		return
	}

	user, err := models.NewUser(s.db, input.Username, input.Email, input.Password)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input or username/email already exists"})
		return
	}

	err = user.GenerateAuthToken(s.db)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Token generation error"})
		return
	}

	if !user.AuthToken.Valid {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Invalid token"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "User successfully created", "token": user.AuthToken.String,
		"token_type": "bearer"})
}

func (s *Server) RefreshToken(c *gin.Context) {
	login, ok := c.Get("login")
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Username was not provided"})
		return
	}

	user, err := models.FetchUserByLogin(s.db, login.(string))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	password, ok := c.Get("password")
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Password was not provided"})
		return
	}

	ok = user.IsPasswordCorrect(password.(string))
	if !ok {
		c.JSON(http.StatusForbidden, gin.H{"error": "Wrong password"})
		return
	}

	err = user.GenerateAuthToken(s.db)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Token generation error"})
		return
	}

	if !user.AuthToken.Valid {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Invalid token"})
	}

	c.JSON(http.StatusOK, gin.H{"message": "Token successfully refreshed", "token": user.AuthToken.String,
		"token_type": "bearer"})
}

func (s *Server) RevokeToken(c *gin.Context) {
	token, ok := c.Get("token")
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{"error": "User token was not provided"})
		return
	}

	user, err := models.FetchUserByToken(s.db, token.(string))
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Token is invalid or expired"})
		return
	}

	err = user.RevokeAuthToken(s.db)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Token successfully revoked"})
}
