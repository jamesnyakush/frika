package model

import (
	"github.com/gofrs/uuid"
	"golang.org/x/crypto/bcrypt"
	"time"
)

type User struct {
	UserId               uuid.UUID `json:"user_id"`
	FirstName            string    `json:"first_name"`
	LastName             string    `json:"last_name"`
	Username             string    `json:"username"`
	Email                string    `json:"email"`
	Password             string    `json:"password"`
	Phone                string    `json:"phone"`
	CountyCode           string    `json:"county_code"`
	ImageUrl             string    `json:"image_url"`
	PushNotificationKey  string    `json:"push_notification_key"`
	InviteCode           string    `json:"invite_code"`
	Verified             bool      `json:"verified"`
	VerificationExpires  string    `json:"verification_expires"`
	VerificationToken    uuid.UUID `json:"verification_token"`
	ResetPasswordExpires time.Time `json:"reset_password_expires"`
	ResetPasswordToken   string    `json:"reset_password_token"`
	CreatedAt            time.Time `json:"created_at"`
	UpdatedAt            time.Time `json:"updated_at"`
	DeletedAt            time.Time `json:"deleted_at"`
	Roles                Role      `json:"roles"`
}


//HashPassword hashes the user password using bcrypt hash function
func (u *User) HashPassword() error {

	pwd := []byte(u.Password)

	hash, err := bcrypt.GenerateFromPassword(pwd, bcrypt.DefaultCost)

	if err != nil {
		return err
	}

	u.Password = string(hash)

	return nil
}

//Compare compares the password hash against the passed in password string
func (u User) Compare(hash, password string) (bool, error) {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	if err != nil {
		return false, err
	}
	return true, nil
}
