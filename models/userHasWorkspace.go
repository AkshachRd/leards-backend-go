package models

import "github.com/google/uuid"

type UserHasWorkspace struct {
	UserID      uuid.UUID `gorm:"primaryKey"`
	User        User
	WorkspaceID uuid.UUID `gorm:"primaryKey"`
	Workspace   Workspace
	RoleID      uint8
	Role        UserWorkspaceRole
}
