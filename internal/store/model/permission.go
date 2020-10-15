package model

type Permission struct {
	PermissionId   int    `json:"permission_id"` // uuid
	PermissionName string `json:"permission_name"`
}
