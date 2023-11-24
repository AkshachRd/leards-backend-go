package v1

import (
	"github.com/AkshachRd/leards-backend-go/httputils"
	"github.com/AkshachRd/leards-backend-go/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

// GetFavoriteStorages godoc
// @Id			 getFavoriteStorages
// @Summary      get favorite storages by user id
// @Description  returns the user's favorite storages
// @Tags         library
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param        page         query     int     false    "Page" minimum(1)
// @Param        page_size    query     int     false    "Page size" minimum(10)    maximum(100)
// @Param		 user_id	  path		string	true	"User ID"
// @Success      200  {object}  httputils.FavoriteStoragesResponse
// @Failure      500  {object}  httputils.HTTPError
// @Router       /library/{user_id} [get]
func GetFavoriteStorages(c *gin.Context) {
	userId := c.Param("user_id")

	favoriteStoragesContent, err := models.FetchFavoriteStoragesContentByUserId(userId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message":          "Favorite storages successfully fetched",
		"favoriteStorages": *httputils.ConvertFavoriteStoragesContentToContent(favoriteStoragesContent),
	})
}

// AddStorageToFavorite godoc
// @Id			 addStorageToFavorite
// @Summary      add favorite storage by user id, storage id and storage type
// @Description  creates a new favorite storage in the database
// @Tags         library
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param		 user_id	  path		string	true	"User ID" minlength(36)  maxlength(36)
// @Param		 storage_type path		string	true	"Storage type" Enums(deck, folder)
// @Param		 storage_id	  path		string	true	"Storage ID" minlength(36)  maxlength(36)
// @Success      201  {object}  httputils.BasicResponse
// @Failure      500  {object}  httputils.HTTPError
// @Router       /library/{user_id}/{storage_type}/{storage_id} [post]
func AddStorageToFavorite(c *gin.Context) {
	userId := c.Param("user_id")
	storageType := c.Param("storage_type")
	storageId := c.Param("storage_id")

	_, err := models.NewFavoriteStorage(userId, storageId, storageType)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusInternalServerError, gin.H{
		"message": "Successfully added the storage to favorites",
	})
}

// RemoveStorageFromFavorite godoc
// @Id			 removeStorageFromFavorite
// @Summary      remove the storage from favorites by user id, storage id and storage type
// @Description  deletes the favorite storage in the database
// @Tags         library
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param		 user_id	  path		string	true	"User ID" minlength(36)  maxlength(36)
// @Param		 storage_type path		string	true	"Storage type" Enums(deck, folder)
// @Param		 storage_id	  path		string	true	"Storage ID" minlength(36)  maxlength(36)
// @Success      201  {object}  httputils.BasicResponse
// @Failure      500  {object}  httputils.HTTPError
// @Router       /library/{user_id}/{storage_type}/{storage_id} [delete]
func RemoveStorageFromFavorite(c *gin.Context) {
	userId := c.Param("user_id")
	storageType := c.Param("storage_type")
	storageId := c.Param("storage_id")

	err := models.DeleteFavoriteStorage(userId, storageId, storageType)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusInternalServerError, gin.H{
		"message": "Successfully remove the storage from favorites",
	})
}
