package v1

import (
	"fmt"
	"github.com/AkshachRd/leards-backend-go/httputils"
	"github.com/AkshachRd/leards-backend-go/services"
	"github.com/gin-gonic/gin"
	"net/http"
)

// SetStorageAccess godoc
// @Id			 setStorageAccess
// @Summary      Set the storage's access type
// @Description  sets the storage's access type in the database
// @Tags         sharing
// @Accept       json
// @Produce      json
// @Security     Bearer
// @Param		 storage_type path		string	true	"Storage type" Enums(deck, folder)
// @Param		 storage_id	  path		string	true	"Storage ID" minlength(36)  maxlength(36)
// @Param		 SetStorageAccessData body httputils.SetStorageAccessRequest true "Set storage access data"
// @Success      200  {object}  httputils.BasicResponse
// @Failure      400  {object}  httputils.HTTPError
// @Failure      500  {object}  httputils.HTTPError
// @Router       /sharing/{storage_type}/{storage_id}  [put]
func SetStorageAccess(c *gin.Context) {
	var input httputils.SetStorageAccessRequest

	if err := c.ShouldBind(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	sharingService := services.NewSharingService()

	storageType := c.Param("storage_type")
	storageId := c.Param("storage_id")
	if err := sharingService.SetStorageAccess(storageId, storageType, input.Type); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": fmt.Sprintf("Successfully updated %s's access type to %s", storageType, input.Type),
	})
}
