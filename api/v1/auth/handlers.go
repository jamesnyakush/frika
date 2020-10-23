package auth

import (
	"fmt"
	"github.com/go-chi/render"
	"log"
	"net/http"
)

type signupRequest struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type loginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (body *signupRequest) Bind(r *http.Request) error {
	return nil
}

func (body *loginRequest) Bind(r *http.Request) error {
	return nil
}

func (rs Resource) HandleSignUp(rw http.ResponseWriter, r *http.Request) {
	body := signupRequest{}

	if err := render.Bind(r, &body); err != nil {
		log.Println(err)
		return
	}

	usr, err := rs.Store.User().New(body.Name, body.Email, body.Password)

	if err != nil {
		log.Println(err)
		return
	}

	fmt.Println(usr)
}

func (rs Resource) HandleLogin(rw http.ResponseWriter, r *http.Request) {
	body := loginRequest{}

	if err := render.Bind(r, &body); err != nil {
		log.Println(err)
		//render.Render(rw,r,http.StatusBadRequest)
		return
	}
}
