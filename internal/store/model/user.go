package model

import "time"

type User struct {
	UserId               int       `json:"user_id"` //uuid
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
	VerificationToken    int       `json:"verification_token"` // uuid
	ResetPasswordExpires time.Time `json:"reset_password_expires"`
	ResetPasswordToken   string    `json:"reset_password_token"`
	CreatedAt            time.Time `json:"created_at"`
	UpdatedAt            time.Time `json:"updated_at"`
	DeletedAt            time.Time `json:"deleted_at"`
	Roles                Role      `json:"roles"`
}
