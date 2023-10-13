package handlers

import (
	"github.com/AkshachRd/leards-backend-go/httputils"
	"github.com/AkshachRd/leards-backend-go/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

// GetSingleFolder godoc
// @Summary      Get single folder by id
// @Description  fetches the folder from the database
// @Tags         folders
// @Accept       json
// @Produce      json
// @Security     Bearer
// @Param		 folder_id	  path		string	true	"Folder ID"
// @Success      200  {object}  httputils.FolderResponse
// @Failure      400  {object}  httputils.HTTPError
// @Failure      500  {object}  httputils.HTTPError
// @Router       /folders/{folder_id} [get]
func (s *Server) GetSingleFolder(c *gin.Context) {
	folderId := c.Param("folder_id")

	folder, err := models.FetchFolderById(s.db, folderId, true)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input or folder doesn't exist"})
		return
	}

	var path []httputils.Path
	path = append(
		[]httputils.Path{{Name: folder.Name, Id: folder.ID}},
		path...,
	)
	for parentFolder := folder.ParentFolder; parentFolder != nil; {
		path = append(
			[]httputils.Path{{Name: parentFolder.Name, Id: parentFolder.ID}},
			path...,
		)
		parentFolder = parentFolder.ParentFolder
	}

	var content []httputils.Content

	contentDecks, err := models.FetchDecksByParentId(s.db, folder.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Can't fetch decks from the database"})
		return
	}

	for _, contentDeck := range *contentDecks {
		content = append(content, httputils.Content{Id: contentDeck.ID, Name: contentDeck.Name, Type: "deck"})
	}

	contentFolders, err := models.FetchFoldersByParentFolderId(s.db, folder.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Can't fetch folders from the database"})
		return
	}

	for _, contentFolder := range *contentFolders {
		content = append(content, httputils.Content{Id: contentFolder.ID, Name: contentFolder.Name, Type: "folder"})
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Folder successfully fetched",
		"folder":  httputils.Folder{FolderId: folder.ID, Name: folder.Name, Path: path, Content: content},
	})
}
