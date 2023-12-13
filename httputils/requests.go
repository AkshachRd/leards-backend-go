package httputils

type CreateUserRequest struct {
	Username string `json:"username" binding:"required" example:"Bob"`
	Email    string `json:"email" binding:"required" example:"bob@leards.space"`
	Password string `json:"password" binding:"required" example:"123"`
} // @name CreateUserRequest

type CreateDeckRequest struct {
	Name           string `json:"name" binding:"required" example:"My new deck"`
	ParentFolderId string `json:"parentFolderId" binding:"required" example:"72a30ffb-1896-48b1-b006-985fb055db0f"`
	UserId         string `json:"userId" binding:"required" example:"72a30ffb-1896-48b1-b006-985fb055db0f"`
} // @name CreateDeckRequest

type CreateFolderRequest struct {
	Name           string `json:"name" binding:"required" example:"My new folder"`
	ParentFolderId string `json:"parentFolderId" example:"72a30ffb-1896-48b1-b006-985fb055db0f"`
	UserId         string `json:"userId" binding:"required" example:"72a30ffb-1896-48b1-b006-985fb055db0f"`
} // @name CreateFolderRequest

type UpdateDeckRequest struct {
	Name string `json:"name" binding:"required" example:"My new deck"`
} // @name UpdateDeckRequest

type UpdateUserRequest struct {
	Name     string `json:"name,omitempty" example:"Ivan"`
	Email    string `json:"email,omitempty" example:"rostislav.glizerin@ispring.com"`
	Password string `json:"password,omitempty" example:"qwerty123"`
} // @name UpdateUserRequest

type SyncCardsRequest struct {
	Cards []Card `json:"cards" binding:"required"`
} // @name SyncCardsRequest

type UpdateFolderRequest struct {
	Name string `json:"name" binding:"required" example:"My new folder"`
} // @name UpdateFolderRequest

type UpdateUserSettingsRequest struct {
	Settings Settings `json:"settings" binding:"required"`
} // @name UpdateUserSettingsRequest

type TagsRequest struct {
	Tags []string `json:"tags" binding:"required"`
} // @name TagsRequest

type SetStorageAccessRequest struct {
	Type string `json:"type" binding:"required" example:"public" enums:"shared,public,private"`
} // @name SetStorageAccessRequest
