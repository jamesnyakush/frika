package auth

import "github.com/go-chi/chi"

func (rs Resource) Router() *chi.Mux {
	r := chi.NewRouter()

	r.Post("/signup", rs.HandleSignUp)
	r.Post("/login", rs.HandleLogin)

	return r
}
