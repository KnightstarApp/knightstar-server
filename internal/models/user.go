package models

import (
	"knightstar/pkg/util"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Email         string `json:"email" validate:"required,email" gorm:"unique;not null"`
	Username      string `json:"username" validate:"required" gorm:"unique;not null"`
	PublicKey     string `json:"publicKey" gorm:"default:''"`
	SecurityToken string `json:"securityToken" gorm:"default:''"`
	OtpVerified   bool   `json:"otpVerified" gorm:"default:false"`
}

func (u *User) ToJSON() util.JSON {
	return util.JSON{
		"id":        u.ID,
		"email":     u.Email,
		"username":  u.Username,
		"publicKey": u.PublicKey,
	}
}
