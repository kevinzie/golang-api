package models

import (
	"time"
)

type Users struct {
	//gorm.Model
	ID              uint       `db:"id" json:"-" validate:"" gorm:"primary_key"` // column name is `id`
	FirstName       string     `db:"first_name" json:"first_name" validate:"required, lte=255"`
	LastName        *string    `db:"last_name" json:"last_name" validate:"lte=255"`
	Email           *string    `db:"email" json:"email" validate:"lte=255"`
	EmailVerifiedAt *time.Time `db:"email_verified_at" json:"email_verified_at" validate:"lte=255"`
	Password        *string    `db:"password" json:"-" validate:"required, lte=255"`
	RememberToken   *string    `db:"remember_token" json:"remember_token" validate:""`
	Activated       uint       `db:"activated" json:"activated" validate:""`
	Status          string     `db:"status" json:"status" validate:"required"`
	SignupIpAddress *string    `db:"signup_ip_address" json:"signup_ip_address" validate:""`
	Profiles        Profiles   `json:"profile" gorm:"foreignKey:UserId"`
	CreatedAt       time.Time  `db:"created_at" json:"created_at" validate:""`
	UpdatedAt       *time.Time `db:"updated_at" json:"updated_at" validate:""`
	DeletedAt       *time.Time `db:"deleted_at" json:"-" gorm:"index"`
}
