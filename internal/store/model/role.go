package model

type Role struct {
	RoleId      int          `json:"role_id"` // uuid
	RoleName    string       `json:"role_name"`
	Permissions []Permission `json:"permissions"`
}
