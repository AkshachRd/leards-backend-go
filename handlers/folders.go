package handlers

import (
	"github.com/AkshachRd/leards-backend-go/httputils"
	"github.com/AkshachRd/leards-backend-go/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

// GetFolder godoc
// @Id           getFolderById
// @Summary      Get the folder by id
// @Description  fetches the folder from the database
// @Tags         folders
// @Accept       json
// @Produce      json
// @Security     Bearer
// @Param		 folder_id	  path		string	true	"Folder ID"
// @Success      200  {object}  httputils.FolderResponse
// @Failure      400  {object}  httputils.HTTPError
// @Router       /folders/{folder_id} [get]
func (s *Server) GetFolder(c *gin.Context) {
	folderId := c.Param("folder_id")
	folder, err := models.FetchFolderById(s.db, folderId, true, true, true)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input or folder doesn't exist"})
		return
	}

	responseFolder := httputils.ConvertFolder(folder)

	c.JSON(http.StatusOK, gin.H{
		"message": "Folder successfully fetched",
		"folder":  responseFolder,
	})
}

// UpdateFolder godoc
// @Id           updateFolderById
// @Summary      Update the folder by id
// @Description  updates the folder in the database
// @Tags         folders
// @Accept       json
// @Produce      json
// @Security     Bearer
// @Param		 folder_id	  path		string	true	"Folder ID"
// @Param		 updateFolderData body httputils.UpdateFolderRequest true "Update folder data"
// @Success      200  {object}  httputils.FolderResponse
// @Failure      400  {object}  httputils.HTTPError
// @Router       /folders/{folder_id} [put]
func (s *Server) UpdateFolder(c *gin.Context) {
	var input httputils.UpdateFolderRequest

	if err := c.ShouldBind(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	folderId := c.Param("folder_id")
	folder, err := models.UpdateFolderById(s.db, folderId, input.Name, models.Access(input.AccessType))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Cannot update deck"})
		return
	}

	responseFolder := httputils.ConvertFolder(folder)

	c.JSON(http.StatusOK, gin.H{
		"message": "Folder successfully updated",
		"folder":  responseFolder,
	})
}

// DeleteFolder godoc
// @Id           deleteFolderById
// @Summary      Delete the folder by id
// @Description  deletes the folder in the database
// @Tags         folders
// @Accept       json
// @Produce      json
// @Security     Bearer
// @Param		 folder_id	  path		string	true	"Folder ID"
// @Success      200  {object}  httputils.BasicResponse
// @Failure      400  {object}  httputils.HTTPError
// @Router       /folders/{folder_id} [delete]
func (s *Server) DeleteFolder(c *gin.Context) {
	folderId := c.Param("folder_id")

	err := models.DeleteFolderById(s.db, folderId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input, folder doesn't exist or cannot delete"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Folder successfully deleted",
	})
}
