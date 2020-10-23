package permissions

import "net/http"

type permissionRequest struct {
	PermissionName string `json:"permission_name"`
}

func (body *permissionRequest) Bind(r http.Request) error {
	return nil
}

func (rs Resource) handleCreatePermission(rw http.ResponseWriter, r *http.Request) {

}

func (rs Resource) handleGetPermissions(rw http.ResponseWriter, r *http.Request) {

}
