package httputils

type CreateUserRequest struct {
	Username string `json:"username" binding:"required" example:"Bob"`
	Email    string `json:"email" binding:"required" example:"bob@leards.space"`
	Password string `json:"password" binding:"required" example:"123"`
}

type CreateDeckRequest struct {
	Name           string `json:"name" binding:"required" example:"My new deck"`
	ParentFolderId string `json:"parentFolderId" binding:"required" example:"72a30ffb-1896-48b1-b006-985fb055db0f"`
}

type UpdateDeckRequest Deck
