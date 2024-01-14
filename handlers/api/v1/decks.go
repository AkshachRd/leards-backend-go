package v1

import (
	"fmt"
	"github.com/AkshachRd/leards-backend-go/services"
	"net/http"

	"github.com/AkshachRd/leards-backend-go/httputils"
	"github.com/AkshachRd/leards-backend-go/models"
	"github.com/gin-gonic/gin"
)

// GetDeck godoc
// @Id			 getDeckById
// @Summary      Get the deck by id
// @Description  fetches the deck from the database
// @Tags         decks
// @Accept       json
// @Produce      json
// @Security     Bearer
// @Param		 deck_id	  path		string	true	"Deck ID"
// @Success      200  {object}  httputils.DeckResponse
// @Failure      400  {object}  httputils.HTTPError
// @Router       /decks/{deck_id} [get]
func GetDeck(c *gin.Context) {
	deckId := c.Param("deck_id")

	deck, err := models.FetchDeckById(deckId, true, true)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input or deck doesn't exist"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Deck successfully fetched",
		"deck":    *httputils.ConvertDeck(deck),
	})
}

// GetDeckSettings godoc
// @Id			 getDeckSettingsByDeckId
// @Summary      Get the deck's settings by deck id
// @Description  fetches the deck from the database and returns settings
// @Tags         decks
// @Accept       json
// @Produce      json
// @Security     Bearer
// @Param		 deck_id	  path		string	true	"Deck ID"
// @Success      200  {object}  httputils.StorageSettingsResponse
// @Failure      400  {object}  httputils.HTTPError
// @Router       /decks/{deck_id}/settings [get]
func GetDeckSettings(c *gin.Context) {
	deckId := c.Param("deck_id")

	deck, err := models.FetchDeckById(deckId, false, true)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input or deck doesn't exist"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message":         "Deck settings successfully fetched",
		"storageSettings": *httputils.ConvertDeckToStorageSettings(deck),
	})
}

// CreateDeck godoc
// @Id			 createNewDeck
// @Summary      Create a new deck
// @Description  creates a new deck in the database
// @Tags         decks
// @Accept       json
// @Produce      json
// @Security     Bearer
// @Param		 createDeckData body httputils.CreateDeckRequest true "Create deck data"
// @Success      200  {object}  httputils.DeckResponse
// @Failure      400  {object}  httputils.HTTPError
// @Failure      500  {object}  httputils.HTTPError
// @Router       /decks [post]
func CreateDeck(c *gin.Context) {
	var input httputils.CreateDeckRequest

	if err := c.ShouldBind(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	deck, err := models.CreateDeck(input.Name, models.AccessTypePrivate, input.ParentFolderId, input.UserId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Cannot create a new deck"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Deck successfully created",
		"deck":    *httputils.ConvertDeck(deck),
	})
}

// UpdateDeck godoc
// @Id			 updateDeckById
// @Summary      Updates the deck by id
// @Description  updates the deck in the database
// @Tags         decks
// @Accept       json
// @Produce      json
// @Security     Bearer
// @Param		 deck_id	  path		string	true	"Deck ID"
// @Param		 updateDeckData body httputils.UpdateDeckRequest true "Update deck data"
// @Success      200  {object}  httputils.DeckResponse
// @Failure      400  {object}  httputils.HTTPError
// @Failure      500  {object}  httputils.HTTPError
// @Router       /decks/{deck_id} [put]
func UpdateDeck(c *gin.Context) {
	var input httputils.UpdateDeckRequest

	if err := c.ShouldBind(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	deckId := c.Param("deck_id")
	deck, err := models.UpdateDeckById(deckId, input.Name)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Cannot update deck"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Deck successfully updated",
		"deck":    *httputils.ConvertDeck(deck),
	})
}

// DeleteDeck godoc
// @Id           deleteDeckById
// @Summary      Delete the deck by id
// @Description  deletes the deck in the database
// @Tags         decks
// @Accept       json
// @Produce      json
// @Security     Bearer
// @Param		 deck_id	  path		string	true	"Deck ID"
// @Success      200  {object}  httputils.BasicResponse
// @Failure      400  {object}  httputils.HTTPError
// @Router       /decks/{deck_id} [delete]
func DeleteDeck(c *gin.Context) {
	deckId := c.Param("deck_id")

	err := models.DeleteDeckById(deckId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input, deck doesn't exist or cannot delete"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Deck successfully deleted",
	})
}

// CloneDeck godoc
// @Id			 cloneDeckById
// @Summary      Clones the deck by id
// @Description  clones the deck in the database
// @Tags         decks
// @Accept       json
// @Produce      json
// @Security     Bearer
// @Param		 cloneDeckData body httputils.CloneDeckRequest true "Clone deck data"
// @Success      200  {object}  httputils.DeckResponse
// @Failure      400  {object}  httputils.HTTPError
// @Failure      500  {object}  httputils.HTTPError
// @Router       /decks/clone [post]
func CloneDeck(c *gin.Context) {
	var input httputils.CloneDeckRequest

	if err := c.ShouldBind(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	deck, err := models.CloneDeck(input.DeckId, input.UserId, input.ParentFolderId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if len(deck.Cards) != 0 {
		repetitionService := services.NewRepetitionService()

		for _, card := range deck.Cards {
			err = repetitionService.CreateRepetition(input.UserId, card.ID)
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{
					"error": fmt.Sprintf("Cannot create repetition: %s", err.Error()),
				})
				return
			}
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Deck successfully cloned",
		"deck":    *httputils.ConvertDeck(deck),
	})
}
