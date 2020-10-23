package model

import "github.com/gofrs/uuid"

type Permission struct {
	PermissionId   uuid.UUID `json:"permission_id"`
	PermissionName string    `json:"permission_name"`
}
