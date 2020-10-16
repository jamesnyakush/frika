package store

import "github.com/jamesnyakush/frika/internal/store/model"

type Store interface {
	User() UserStore
}

type UserStore interface {
	New(name, email, password string)
	Get(id string) (model.User, error)
	GetByEmail(email string) (model.User, error)
	SetAPIKey(id string, key int) (model.User, error)
}
