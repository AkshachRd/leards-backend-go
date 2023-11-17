package v1

import (
	"github.com/AkshachRd/leards-backend-go/httputils"
	"github.com/AkshachRd/leards-backend-go/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

// CreateFolder godoc
// @Id			 createNewFolder
// @Summary      Create a new folder
// @Description  creates a new folder in the database
// @Tags         folders
// @Accept       json
// @Produce      json
// @Security     Bearer
// @Param		 createFolderData body httputils.CreateFolderRequest true "Create folder data"
// @Success      201  {object}  httputils.FolderResponse
// @Failure      400  {object}  httputils.HTTPError
// @Failure      500  {object}  httputils.HTTPError
// @Router       /folders [post]
func CreateFolder(c *gin.Context) {
	var input httputils.CreateFolderRequest

	if err := c.ShouldBind(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	folder, err := models.NewFolder(input.Name, models.Private, &input.ParentFolderId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Cannot create a new folder"})
		return
	}

	folder, err = models.FetchFolderById(folder.ID, false, true, true, false, true)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Cannot fetch the created folder"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "Folder successfully created",
		"folder":  *httputils.ConvertFolder(folder),
	})
}

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
func GetFolder(c *gin.Context) {
	folderId := c.Param("folder_id")
	folder, err := models.FetchFolderById(folderId, true, true, true, false, true)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input or folder doesn't exist"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Folder successfully fetched",
		"folder":  *httputils.ConvertFolder(folder),
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
func UpdateFolder(c *gin.Context) {
	var input httputils.UpdateFolderRequest

	if err := c.ShouldBind(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	folderId := c.Param("folder_id")
	folder, err := models.UpdateFolderById(folderId, input.Name, models.Access(input.AccessType))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Cannot update folder"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Folder successfully updated",
		"folder":  *httputils.ConvertFolder(folder),
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
func DeleteFolder(c *gin.Context) {
	folderId := c.Param("folder_id")

	err := models.DeleteFolderById(folderId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input, folder doesn't exist or cannot delete"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Folder successfully deleted",
	})
}
