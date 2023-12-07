package v1

import (
	"github.com/AkshachRd/leards-backend-go/httputils"
	"github.com/AkshachRd/leards-backend-go/services"
	"github.com/gin-gonic/gin"
	"net/http"
)

// AddTagsToStorage godoc
// @Id			 addTagsToStorage
// @Summary      Add tags to the storage
// @Description  creates new tags or takes ones that already exist and links them in the database
// @Tags         tags
// @Accept       json
// @Produce      json
// @Security     Bearer
// @Param		 user_id	  path		string	true	"User ID" minlength(36)  maxlength(36)
// @Param		 storage_type path		string	true	"Storage type" Enums(deck, folder)
// @Param		 storage_id	  path		string	true	"Storage ID" minlength(36)  maxlength(36)
// @Param		 addTagsToStorageData body httputils.TagsRequest true "Add tags to storage data"
// @Success      200  {object}  httputils.BasicResponse
// @Failure      400  {object}  httputils.HTTPError
// @Failure      500  {object}  httputils.HTTPError
// @Router       /tags/{user_id}/{storage_type}/{storage_id} [post]
func AddTagsToStorage(c *gin.Context) {
	//userId := c.Param("user_id")
	storageType := c.Param("storage_type")
	storageId := c.Param("storage_id")
	var input httputils.TagsRequest

	if err := c.ShouldBind(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	tagService, err := services.NewTagService(storageId, storageType)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err = tagService.AddTags(input.Tags); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Tags successfully added",
	})
}

// RemoveTagsFromStorage godoc
// @Id			 removeTagsFromStorage
// @Summary      removes tags from the storage
// @Description  deletes tags and unlink them from storage in the database
// @Tags         tags
// @Accept       json
// @Produce      json
// @Security     Bearer
// @Param		 user_id	  path		string	true	"User ID" minlength(36)  maxlength(36)
// @Param		 storage_type path		string	true	"Storage type" Enums(deck, folder)
// @Param		 storage_id	  path		string	true	"Storage ID" minlength(36)  maxlength(36)
// @Param		 removeTagsToStorageData body httputils.TagsRequest true "Remove tags to storage data"
// @Success      200  {object}  httputils.BasicResponse
// @Failure      400  {object}  httputils.HTTPError
// @Failure      500  {object}  httputils.HTTPError
// @Router       /tags/{user_id}/{storage_type}/{storage_id} [delete]
func RemoveTagsFromStorage(c *gin.Context) {
	//userId := c.Param("user_id")
	storageType := c.Param("storage_type")
	storageId := c.Param("storage_id")
	var input httputils.TagsRequest

	if err := c.ShouldBind(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	tagService, err := services.NewTagService(storageId, storageType)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err = tagService.RemoveTags(input.Tags); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Tags successfully removed",
	})
}
