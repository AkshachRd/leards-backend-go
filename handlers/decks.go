package handlers

import (
	"github.com/AkshachRd/leards-backend-go/httputils"
	"github.com/AkshachRd/leards-backend-go/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

// GetDeck godoc
// @Summary      Get a single deck by id
// @Description  fetches the deck from the database
// @Tags         decks
// @Accept       json
// @Produce      json
// @Security     Bearer
// @Param		 folder_id	  path		string	true	"Folder ID"
// @Param		 deck_id	  path		string	true	"Deck ID"
// @Success      200  {object}  httputils.DeckResponse
// @Failure      400  {object}  httputils.HTTPError
// @Router       /folders/{folder_id}/decks/{deck_id} [get]
func (s *Server) GetDeck(c *gin.Context) {
	deckId := c.Param("deckId")

	deck, err := models.FetchDeckById(s.db, deckId, true)
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
// @Param		 folder_id	  path		string	true	"Folder ID"
// @Param		 createDeckData body httputils.CreateDeckRequest true "Create deck data"
// @Success      200  {object}  httputils.DeckResponse
// @Failure      400  {object}  httputils.HTTPError
// @Failure      500  {object}  httputils.HTTPError
// @Router       /folders/{folder_id}/decks [post]
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
// @Param		 folder_id	  path		string	true	"Folder ID"
// @Param		 deck_id	  path		string	true	"Deck ID"
// @Param		 updateDeckData body httputils.UpdateDeckRequest true "Update deck data"
// @Success      200  {object}  httputils.BasicResponse
// @Failure      400  {object}  httputils.HTTPError
// @Failure      500  {object}  httputils.HTTPError
// @Router       /folders/{folder_id}/decks/{deck_id} [put]
func (s *Server) UpdateDeck(c *gin.Context) {
	var input httputils.UpdateDeckRequest

	if err := c.ShouldBind(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	deckId := c.Param("deck_id")

	err := models.UpdateDeck(s.db, deckId, input.Name, models.Access(input.AccessType))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Cannot update deck"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Deck successfully updated",
	})
}
