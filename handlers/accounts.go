package handlers

import (
	"fmt"
	"github.com/AkshachRd/leards-backend-go/httputils"
	"github.com/AkshachRd/leards-backend-go/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

// CreateUser godoc
// @Summary      Register new user
// @Description  creates new user and returns a token
// @Tags         accounts
// @Accept       json
// @Produce      json
// @Param		 createUserData body httputils.CreateUserRequest true "User register data"
// @Success      201  {object}  httputils.UserResponse
// @Failure      400  {object}  httputils.HTTPError
// @Failure      404  {object}  httputils.HTTPError
// @Failure      500  {object}  httputils.HTTPError
// @Router       /accounts [post]
func (s *Server) CreateUser(c *gin.Context) {
	var input httputils.CreateUserRequest

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
	if err != nil || !user.AuthToken.Valid {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Invalid new token"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "User successfully created",
		"userId":  user.ID.String(),
		"token":   user.AuthToken.String,
	})
}

// LoginUser godoc
// @Summary      Login an existing user
// @Description  returns user id of an existing user
// @Tags         accounts
// @Accept       json
// @Produce      json
// @Security     BasicAuth
// @Success      201  {object}  httputils.UserResponse
// @Failure      400  {object}  httputils.HTTPError
// @Failure      401  {object}  httputils.HTTPError
// @Failure      500  {object}  httputils.HTTPError
// @Router       /accounts [get]
func (s *Server) LoginUser(c *gin.Context) {
	email, ok := c.Get("email")
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User is not authorized"})
		return
	}

	switch email.(type) {
	case string:
		break
	default:
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid login"})
		return
	}

	user, err := models.FetchUserByEmail(s.db, fmt.Sprintf("%v", email))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid login or user do not exist"})
		return
	}

	err = user.GenerateAuthToken(s.db)
	if err != nil || !user.AuthToken.Valid {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Invalid new token"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "User successfully signed in",
		"userId":  user.ID.String(),
		"token":   user.AuthToken.String,
	})
}
