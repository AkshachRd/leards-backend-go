package v1

import (
	"github.com/AkshachRd/leards-backend-go/httputils"
	"github.com/AkshachRd/leards-backend-go/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

// GetStorageCards godoc
// @Id			 getStorageCards
// @Summary      Get all storage's cards
// @Description  fetches cards of the storage from the database
// @Tags         cards
// @Accept       json
// @Produce      json
// @Security     Bearer
// @Param		 storage_type path		string	true	"Storage type" Enums(deck, folder)
// @Param		 storage_id	  path		string	true	"Storage ID" minlength(36)  maxlength(36)
// @Success      200  {object}  httputils.CardsResponse
// @Failure      400  {object}  httputils.HTTPError
// @Failure      500  {object}  httputils.HTTPError
// @Router       /cards/{storage_type}/{storage_id} [get]
func GetStorageCards(c *gin.Context) {
	storageType := c.Param("storage_type")
	storageId := c.Param("storage_id")

	var cards *[]models.Card
	var err error
	switch storageType {
	case "deck":
		cards, err = models.FetchCardsByDeckId(storageId)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input or deck doesn't exist"})
			return
		}
	case "folder":
		cards, err = models.FetchCardsByFolderId(storageId)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input or deck doesn't exist"})
			return
		}
	}

	var content []httputils.Card

	for _, card := range *cards {
		content = append(content, httputils.Card{CardId: card.ID, FrontSide: card.FrontSide, BackSide: card.BackSide})
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Cards successfully fetched",
		"cards":   content,
	})
}

// SyncCards godoc
// @Id			 syncCardsByDeckId
// @Summary      Synchronizes cards
// @Description  adds card without id, updates card with id, deletes card if it's not presented inside the request
// @Tags         cards
// @Accept       json
// @Produce      json
// @Security     Bearer
// @Param		 folder_id	  path		string	true	"Folder ID"
// @Param		 deck_id	  path		string	true	"Deck ID"
// @Param		 syncCardsRequest body httputils.SyncCardsRequest true "Sync cards data"
// @Success      200  {object}  httputils.BasicResponse
// @Failure      400  {object}  httputils.HTTPError
// @Failure      500  {object}  httputils.HTTPError
// @Router       /cards/deck/{deck_id} [put]
func SyncCards(c *gin.Context) {
	var input httputils.SyncCardsRequest

	if err := c.ShouldBind(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	deckId := c.Param("deck_id")

	cards, err := models.FetchCardsByDeckId(deckId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input or deck doesn't exist"})
		return
	}

	cardsMap := make(map[string]models.Card)
	for _, card := range *cards {
		cardsMap[card.ID] = card
	}

	var createCards, updateCards, deleteCards []models.Card

	for _, syncCard := range input.Cards {
		if card, found := cardsMap[syncCard.CardId]; found {
			card.FrontSide = syncCard.FrontSide
			card.BackSide = syncCard.BackSide
			updateCards = append(updateCards, card)

			delete(cardsMap, syncCard.CardId)
		} else {
			createCards = append(createCards, models.Card{
				FrontSide: syncCard.FrontSide,
				BackSide:  syncCard.BackSide,
				DeckID:    deckId,
			})
		}
	}

	for _, deleteCard := range cardsMap {
		deleteCards = append(deleteCards, deleteCard)
	}

	err = models.CreateCards(createCards)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Cannot create cards"})
		return
	}

	err = models.UpdateCards(updateCards)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Cannot update cards"})
		return
	}

	err = models.DeleteCards(deleteCards)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Cannot delete cards"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Cards successfully synced",
	})
}
