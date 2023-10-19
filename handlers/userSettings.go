package handlers

import (
	"github.com/AkshachRd/leards-backend-go/httputils"
	"github.com/AkshachRd/leards-backend-go/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

// GetUserSettings godoc
// @Id           getUserSettings
// @Summary      Get the user settings
// @Description  fetches the user settings from the database
// @Tags         userSettings
// @Accept       json
// @Produce      json
// @Security     Bearer
// @Param		 user_id	  path		string	true	"User ID"
// @Success      200  {object}  httputils.UserSettingsResponse
// @Failure      400  {object}  httputils.HTTPError
// @Router       /accounts/{user_id}/settings [get]
func (s *Server) GetUserSettings(c *gin.Context) {
	userId := c.Param("user_id")
	userSettings, err := models.FetchUserSettingsByUserId(s.db, userId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input or user doesn't exist"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message":  "Settings successfully fetched",
		"settings": *httputils.ConvertUserSettings(userSettings),
	})
}

// UpdateUserSettings godoc
// @Id           updateUserSettings
// @Summary      Update the user settings
// @Description  updates the user settings in the database
// @Tags         userSettings
// @Accept       json
// @Produce      json
// @Security     Bearer
// @Param		 user_id	  path		string	true	"User ID"
// @Param		 updateUserSettingsData body httputils.UpdateUserSettingsRequest true "Update user settings data"
// @Success      200  {object}  httputils.UserSettingsResponse
// @Failure      400  {object}  httputils.HTTPError
// @Router       /accounts/{user_id}/settings [put]
func (s *Server) UpdateUserSettings(c *gin.Context) {
	var input httputils.UpdateUserSettingsRequest

	if err := c.ShouldBind(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	userId := c.Param("user_id")
	userSettings, err := models.UpdateUserSettings(s.db, userId, input.Settings)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Cannot update user settings"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message":  "Settings successfully updated",
		"settings": *httputils.ConvertUserSettings(userSettings),
	})
}
