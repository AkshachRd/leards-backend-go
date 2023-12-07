package v1

import (
	_ "github.com/AkshachRd/leards-backend-go/httputils"
	"github.com/AkshachRd/leards-backend-go/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

// RefreshToken godoc
// @Id			 refreshToken
// @Summary      Refresh user's token
// @Description  when token is expired you need to refresh it
// @Tags         auth
// @Accept       json
// @Produce      json
// @Security     Bearer
// @Param		 user_id	  path		string	true	"User ID"
// @Success      200  {object}  httputils.TokenResponse
// @Failure      400  {object}  httputils.HTTPError
// @Failure      403  {object}  httputils.HTTPError
// @Failure      404  {object}  httputils.HTTPError
// @Failure      500  {object}  httputils.HTTPError
// @Router       /auth/{user_id} [get]
func RefreshToken(c *gin.Context) {
	userId := c.Param("user_id")

	user, err := models.FetchUserById(userId)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	if !user.AuthToken.Valid {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "User has no token"})
	}

	err = user.RefreshAuthToken()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Token generation error"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Token successfully refreshed", "token": user.AuthToken.String})
}

// RevokeToken godoc
// @Id			 revokeToken
// @Summary      Revokes user's token
// @Description  when user signs out token needs to be revoked
// @Tags         auth
// @Accept       json
// @Produce      json
// @Security     Bearer
// @Param		 user_id	  path		string	true	"User ID"
// @Success      200  {object}  httputils.BasicResponse
// @Failure      401  {object}  httputils.HTTPError
// @Failure      500  {object}  httputils.HTTPError
// @Router       /auth/{user_id} [delete]
func RevokeToken(c *gin.Context) {
	userId := c.Param("user_id")

	user, err := models.FetchUserById(userId)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	err = user.RevokeAuthToken()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Token successfully revoked"})
}
