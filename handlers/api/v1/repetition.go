package v1

import (
	"github.com/AkshachRd/leards-backend-go/httputils"
	"github.com/AkshachRd/leards-backend-go/services"
	"github.com/gin-gonic/gin"
	"net/http"
)

// ReviewCard godoc
// @Id			 reviewCard
// @Summary      Review card
// @Description  reviews card and updates repetition in the database
// @Tags         repetition
// @Accept       json
// @Produce      json
// @Security     Bearer
// @Param		 reviewCardRequest body httputils.ReviewCardRequest true "Review card data"
// @Success      200  {array}   services.SearchResult
// @Failure      400  {object}  httputils.HTTPError
// @Failure      500  {object}  httputils.HTTPError
// @Router       /repetition [put]
func ReviewCard(c *gin.Context) {
	var input httputils.ReviewCardRequest

	if err := c.ShouldBind(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	repetitionService := services.NewRepetitionService()
	err := repetitionService.ReviewCard(input.UserId, input.CardId, input.ReviewAnswer)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Card successfully reviewed",
	})
}

// GetNextCard godoc
// @Id			 getNextCard
// @Summary      Get next card for repetition
// @Description  fetches next card for repetition from the database
// @Tags         repetition
// @Accept       json
// @Produce      json
// @Security     Bearer
// @Param		 user_id	  query		string	true	"User ID"
// @Param		 storage_type path		string	true	"Storage type" Enums(deck, folder)
// @Param		 storage_id	  path		string	true	"Storage ID" minlength(36)  maxlength(36)
// @Success      200  {object}  httputils.Card
// @Failure      500  {object}  httputils.HTTPError
// @Router       /repetition/{storage_type}/{storage_id} [get]
func GetNextCard(c *gin.Context) {
	userId := c.Query("user_id")
	storageType := c.Param("storage_type")
	storageId := c.Param("storage_id")

	repetitionService := services.NewRepetitionService()
	card, err := repetitionService.FetchNextRepetitionCard(userId, storageId, storageType)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if card.ID == "" {
		c.JSON(http.StatusOK, nil)
		return
	}

	c.JSON(http.StatusOK, *httputils.ConvertCard(card))
}

// GetStorageStats godoc
// @Id			 getStorageStats
// @Summary      Get stats of a storage after repetition
// @Description  fetches stats of storage after repetition from the database
// @Tags         repetition
// @Accept       json
// @Produce      json
// @Security     Bearer
// @Param		 user_id	  query		string	true	"User ID"
// @Param		 storage_type path		string	true	"Storage type" Enums(deck, folder)
// @Param		 storage_id	  path		string	true	"Storage ID" minlength(36)  maxlength(36)
// @Success      200  {object}  httputils.RepetitionStatsData
// @Failure      500  {object}  httputils.HTTPError
// @Router       /repetition/{storage_type}/{storage_id}/stats [get]
func GetStorageStats(c *gin.Context) {
	userId := c.Query("user_id")
	storageType := c.Param("storage_type")
	storageId := c.Param("storage_id")

	repetitionService := services.NewRepetitionService()
	repetitionStats, err := repetitionService.GetStorageStats(userId, storageId, storageType)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, *httputils.ConvertRepetitionStats(repetitionStats))
}
