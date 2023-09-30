package handlers

import (
	"github.com/AkshachRd/leards-backend-go/models"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"net/http"
)

// RefreshToken godoc
// @Summary      Refresh user's token
// @Description  when token is expired you need to refresh it
// @Tags         auth
// @Accept       json
// @Produce      json
// @Security     Bearer
// @Param		 id	  path		string	true	"User ID"
// @Success      200  {object}  httputils.TokenResponse
// @Failure      400  {object}  httputils.HTTPError
// @Failure      403  {object}  httputils.HTTPError
// @Failure      404  {object}  httputils.HTTPError
// @Failure      500  {object}  httputils.HTTPError
// @Router       /auth/{id} [get]
func (s *Server) RefreshToken(c *gin.Context) {
	id := c.Param("id")
	userUuid, err := uuid.Parse(id)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "User id parsing error"})
		return
	}

	user, err := models.FetchUserById(s.db, userUuid)
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

	c.JSON(http.StatusOK, gin.H{"message": "Token successfully refreshed", "token": user.AuthToken.String})
}

// RevokeToken godoc
// @Summary      Revokes user's token
// @Description  when user signs out token needs to be revoked
// @Tags         auth
// @Accept       json
// @Produce      json
// @Security     Bearer
// @Param		 id	  path		string	true	"User ID"
// @Success      200  {object}  httputils.BasicResponse
// @Failure      400  {object}  httputils.HTTPError
// @Failure      401  {object}  httputils.HTTPError
// @Failure      500  {object}  httputils.HTTPError
// @Router       /auth/{id} [delete]
func (s *Server) RevokeToken(c *gin.Context) {
	id := c.Param("id")
	userUuid, err := uuid.Parse(id)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "User id parsing error"})
		return
	}

	user, err := models.FetchUserById(s.db, userUuid)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	err = user.RevokeAuthToken(s.db)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Token successfully revoked"})
}
