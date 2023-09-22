package handlers

import (
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
// @Success      201  {object}  httputils.TokenResponse
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

	if !user.AuthToken.Valid {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Invalid new token"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "User successfully created", "token": user.AuthToken.String,
		"token_type": "bearer"})
}
