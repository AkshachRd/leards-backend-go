package models

type UserHasWorkspace struct {
	UserID      string `gorm:"primaryKey; size:36;"`
	User        User
	WorkspaceID string `gorm:"primaryKey; size:36"`
	Workspace   Workspace
	RoleID      uint8
	Role        UserWorkspaceRole
}
