package permissions

import "github.com/go-chi/chi"

func (rs Resource) Router() *chi.Mux {
	r := chi.NewRouter()
	r.Post("/create", rs.handleCreatePermission)
	r.Get("/show", rs.handleGetPermissions)
	return r
}
