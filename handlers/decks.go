package handlers

import (
	"github.com/AkshachRd/leards-backend-go/httputils"
	"github.com/AkshachRd/leards-backend-go/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

// GetSingleDeck godoc
// @Summary      Get single deck by id
// @Description  fetches the deck from the database
// @Tags         decks
// @Accept       json
// @Produce      json
// @Security     Bearer
// @Param		 id	  path		string	true	"Deck ID"
// @Success      200  {object}  httputils.DeckResponse
// @Failure      400  {object}  httputils.HTTPError
// @Failure      500  {object}  httputils.HTTPError
// @Router       /decks/{id} [get]
func (s *Server) GetSingleDeck(c *gin.Context) {
	id := c.Param("id")

	deck, err := models.FetchDeckById(s.db, id, true)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input or deck doesn't exist"})
		return
	}

	var content []httputils.Card

	for _, card := range deck.Cards {
		content = append(content, httputils.Card{CardId: card.ID, FrontSide: card.FrontSide, BackSide: card.BackSide})
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Deck successfully fetched",
		"deck":    httputils.Deck{DeckId: deck.ID, Name: deck.Name, Content: content},
	})
}
