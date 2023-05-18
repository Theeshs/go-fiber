package models

import (
	"errors"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID            int    `gorm:"primaryKey"`
	FirstName     string `gorm:"column:first_name"`
	LastName      string `gorm:"column:first_name"`
	UserName      string `gorm:"column:username"`
	Password      string `gorm:"column:password"`
	Email         string `gorm:"column:email"`
	AddressLine1  string `gorm:"column:address_line_1"`
	AddressLine2  string `gorm:"column:address_line_2"`
	DOB           string `gorm:"column:date_of_birth"`
	Image         string `gorm:"column:image"`
	DeviceID      string `gorm:"column:device_id"`
	DeviceType    string `gorm:"column:device_type"`
	Language      string `gorm:"column:language"`
	IsVerified    bool   `gorm:"column:is_verified"`
	IsPremiumUser bool   `gorm:"column:is_permium_user"`
}

func (u *User) PasswordMatches(passwordString string) (bool, error) {
	err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(passwordString))

	if err != nil {
		switch {
		case errors.Is(err, bcrypt.ErrMismatchedHashAndPassword):
			return false, nil
		default:
			return false, err
		}
	}
	return true, nil
}
