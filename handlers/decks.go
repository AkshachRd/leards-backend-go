package handlers

import (
	"github.com/AkshachRd/leards-backend-go/httputils"
	"github.com/AkshachRd/leards-backend-go/models"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"net/http"
)

// GetDeck godoc
// @Summary      Get a single deck by id
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
func (s *Server) GetDeck(c *gin.Context) {
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

// CreateDeck godoc
// @Summary      Create new deck
// @Description  creates new deck in the database
// @Tags         decks
// @Accept       json
// @Produce      json
// @Security     Bearer
// @Param		 createDeckData body httputils.CreateDeckRequest true "Create deck data"
// @Success      200  {object}  httputils.DeckResponse
// @Failure      400  {object}  httputils.HTTPError
// @Failure      500  {object}  httputils.HTTPError
// @Router       /decks [post]
func (s *Server) CreateDeck(c *gin.Context) {
	var input httputils.CreateDeckRequest

	if err := c.ShouldBind(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	deck, err := models.NewDeck(s.db, input.Name, models.Private, input.ParentFolderId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Cannot create new deck"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Deck successfully created",
		"deck":    httputils.Deck{DeckId: deck.ID, Name: deck.Name, Content: make([]httputils.Card, 0)},
	})
}

// UpdateDeck godoc
// @Summary      Updates the deck
// @Description  updates the deck in the database
// @Tags         decks
// @Accept       json
// @Produce      json
// @Security     Bearer
// @Param		 updateDeckData body httputils.UpdateDeckRequest true "Update deck data"
// @Success      200  {object}  httputils.BasicResponse
// @Failure      400  {object}  httputils.HTTPError
// @Failure      500  {object}  httputils.HTTPError
// @Router       /decks/{id} [put]
func (s *Server) UpdateDeck(c *gin.Context) {
	var input httputils.UpdateDeckRequest

	if err := c.ShouldBind(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	var cards []models.Card
	for _, notParsedCard := range input.Content {
		parsedCard := models.Card{}

		if _, err := uuid.Parse(notParsedCard.CardId); err == nil {
			parsedCard.ID = notParsedCard.CardId
		}
		parsedCard.FrontSide = notParsedCard.FrontSide
		parsedCard.BackSide = notParsedCard.BackSide
		parsedCard.DeckID = input.DeckId

		cards = append(cards, parsedCard)
	}

	err := models.UpdateDeck(s.db, input.DeckId, input.Name, models.Private, cards)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Cannot update deck"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Deck successfully updated",
	})
}
