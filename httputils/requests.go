package httputils

type CreateUserRequest struct {
	Username string `json:"username" binding:"required" example:"Bob"`
	Email    string `json:"email" binding:"required" example:"bob@leards.space"`
	Password string `json:"password" binding:"required" example:"123"`
}
