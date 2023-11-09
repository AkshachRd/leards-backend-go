package handlers

import (
	"github.com/AkshachRd/leards-backend-go/httputils"
	"github.com/AkshachRd/leards-backend-go/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

// GetCards godoc
// @Id			 getCardsByDeckId
// @Summary      Get all deck's cards
// @Description  fetches cards of the deck from the database
// @Tags         cards
// @Accept       json
// @Produce      json
// @Security     Bearer
// @Param		 folder_id	  path		string	true	"Folder ID"
// @Param		 deck_id	  path		string	true	"Deck ID"
// @Success      200  {object}  httputils.CardsResponse
// @Failure      400  {object}  httputils.HTTPError
// @Failure      500  {object}  httputils.HTTPError
// @Router       /folders/{folder_id}/decks/{deck_id}/cards [get]
func (s *Server) GetCards(c *gin.Context) {
	deckId := c.Param("deck_id")

	cards, err := models.FetchCardsByDeckId(s.DB, deckId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input or deck doesn't exist"})
		return
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
// @Router       /folders/{folder_id}/decks/{deck_id}/cards [put]
func (s *Server) SyncCards(c *gin.Context) {
	var input httputils.SyncCardsRequest

	if err := c.ShouldBind(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	deckId := c.Param("deck_id")

	cards, err := models.FetchCardsByDeckId(s.DB, deckId)
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

	err = models.CreateCards(s.DB, createCards)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Cannot create cards"})
		return
	}

	err = models.UpdateCards(s.DB, updateCards)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Cannot update cards"})
		return
	}

	err = models.DeleteCards(s.DB, deleteCards)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Cannot delete cards"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Cards successfully synced",
	})
}
