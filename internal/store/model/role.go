package model

import "github.com/gofrs/uuid"

type Role struct {
	RoleId      uuid.UUID    `json:"role_id"`
	RoleName    string       `json:"role_name"`
	Permissions []Permission `json:"permissions"`
}
