package v1

import (
	"fmt"
	"github.com/AkshachRd/leards-backend-go/httputils"
	"github.com/AkshachRd/leards-backend-go/models"
	"github.com/AkshachRd/leards-backend-go/settings"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
	"path/filepath"
)

// CreateUser godoc
// @Id			 registerNewUser
// @Summary      Register a new user
// @Description  creates a new user and returns a token
// @Tags         accounts
// @Accept       json
// @Produce      json
// @Param		 createUserData body httputils.CreateUserRequest true "User register data"
// @Success      201  {object}  httputils.UserResponse
// @Failure      400  {object}  httputils.HTTPError
// @Failure      404  {object}  httputils.HTTPError
// @Failure      500  {object}  httputils.HTTPError
// @Router       /accounts [post]
func CreateUser(c *gin.Context) {
	var input httputils.CreateUserRequest

	if err := c.ShouldBind(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	user, err := models.CreateUser(input.Username, input.Email, input.Password)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input or username/email already exists"})
		return
	}

	err = user.GenerateAuthToken()
	if err != nil || !user.AuthToken.Valid {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Invalid new token"})
		return
	}

	userSettings, err := models.FetchUserSettingsByUserId(user.ID)
	if err != nil || !user.AuthToken.Valid {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Can't fetch user settings"})
		return
	}

	user.Settings = *userSettings

	c.JSON(http.StatusCreated, gin.H{
		"message": "User successfully created",
		"user":    *httputils.ConvertUser(user, c.Request.Host),
	})
}

// LoginUser godoc
// @Id			 loginUser
// @Summary      Login the user
// @Description  returns the user
// @Tags         accounts
// @Accept       json
// @Produce      json
// @Security     BasicAuth
// @Success      201  {object}  httputils.UserResponse
// @Failure      400  {object}  httputils.HTTPError
// @Failure      401  {object}  httputils.HTTPError
// @Failure      500  {object}  httputils.HTTPError
// @Router       /accounts [get]
func LoginUser(c *gin.Context) {
	email, ok := c.Get("email")
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User is not authorized"})
		return
	}

	switch email.(type) {
	case string:
		break
	default:
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	user, err := models.FetchUserByEmail(fmt.Sprintf("%v", email))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid login or user do not exist"})
		return
	}

	err = user.GenerateAuthToken()
	if err != nil || !user.AuthToken.Valid {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Invalid new token"})
		return
	}

	userSettings, err := models.FetchUserSettingsByUserId(user.ID)
	if err != nil || !user.AuthToken.Valid {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Can't fetch user settings"})
		return
	}

	user.Settings = *userSettings

	c.JSON(http.StatusOK, gin.H{
		"message": "User successfully signed in",
		"user":    *httputils.ConvertUser(user, c.Request.Host),
	})
}

// UpdateUser godoc
// @Id			 UpdateUser
// @Summary      Updates the user by id
// @Description  updates the user in the database
// @Tags         accounts
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param		 user_id	  path		string	true	"User ID"
// @Param		 updateUserData body httputils.UpdateUserRequest true "Update user data"
// @Success      200  {object}  httputils.UserResponse
// @Failure      400  {object}  httputils.HTTPError
// @Failure      500  {object}  httputils.HTTPError
// @Router       /accounts/{user_id} [put]
func UpdateUser(c *gin.Context) {
	userId := c.Param("user_id")

	var input httputils.UpdateUserRequest

	if err := c.ShouldBind(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	user, err := models.UpdateUserById(userId, input.Name, input.Email, input.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Cannot update user"})
		return
	}

	userSettings, err := models.FetchUserSettingsByUserId(user.ID)
	if err != nil || !user.AuthToken.Valid {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Can't fetch user settings"})
		return
	}

	user.Settings = *userSettings

	c.JSON(http.StatusOK, gin.H{
		"message": "User successfully updated",
		"user":    *httputils.ConvertUser(user, c.Request.Host),
	})
}

// GetAvatar godoc
// @Id			 getAvatarByUserId
// @Summary      get avatar by user id
// @Description  returns the user's avatar
// @Tags         accounts
// @Accept       json
// @Produce      png
// @Produce      images/jpg
// @Produce      jpeg
// @Produce      gif
// @Param		 user_id	  path		string	true	"User ID"
// @Success      200  {file}  file
// @Failure      400  {object}  httputils.HTTPError
// @Failure      404  {object}  httputils.HTTPError
// @Router       /accounts/{user_id}/avatar [get]
func GetAvatar(c *gin.Context) {
	userId := c.Param("user_id")

	user, err := models.FetchUserById(userId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "User not found"})
		return
	}

	if !user.ProfileIconPath.Valid {
		c.JSON(http.StatusNotFound, gin.H{"error": "Avatar not found"})
		return
	}

	avatarPath := settings.AppSettings.EnvVars.AvatarBasePath + "/" + user.ProfileIconPath.String
	c.Header("Content-Disposition", fmt.Sprintf("attachment; filename=%s", user.ProfileIconPath.String))
	c.File(avatarPath)
}

// UploadAvatar godoc
// @Id			 uploadAvatarByUserId
// @Summary      upload avatar by user id
// @Description  updates the user's avatar
// @Tags         accounts
// @Accept       multipart/form-data
// @Produce      json
// @Security     BearerAuth
// @Param		 user_id	  path		string	true	"User ID"
// @Param avatar formData file true "Avatar image file (JPG, JPEG, PNG, or GIF)"
// @Success      201  {object}  httputils.BasicResponse
// @Failure      400  {object}  httputils.HTTPError
// @Failure      500  {object}  httputils.HTTPError
// @Router       /accounts/{user_id}/avatar [put]
func UploadAvatar(c *gin.Context) {
	userId := c.Param("user_id")

	user, err := models.FetchUserById(userId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "User not found"})
		return
	}

	file, err := c.FormFile("avatar")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "No file uploaded"})
		return
	}

	if !httputils.IsAvatarInSizeRange(file.Size) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Avatar file size is too large"})
		return
	}

	ext := filepath.Ext(file.Filename)
	if !httputils.IsAvatarHasAllowedExtension(ext) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid avatar file extension"})
		return
	}

	avatarFilename := fmt.Sprintf("avatar_%s%s", user.ID, ext)
	avatarSavePath := settings.AppSettings.EnvVars.AvatarBasePath + "/" + avatarFilename

	if user.ProfileIconPath.Valid {
		if err = os.Remove(settings.AppSettings.EnvVars.AvatarBasePath + "/" + user.ProfileIconPath.String); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete previous avatar"})
			return
		}
	} else {
		if err = user.SetProfileIconPath(avatarFilename); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update avatar in DB"})
			return
		}
	}

	if err = c.SaveUploadedFile(file, avatarSavePath); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save file"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Avatar uploaded successfully"})
}
