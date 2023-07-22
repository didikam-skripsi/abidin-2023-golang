package models

import (
	"html"
	"strings"

	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	gorm.Model
	UUID     uuid.UUID `gorm:"size:36;not null;unique" json:"uuid"`
	Username string    `gorm:"size:255;not null;unique" json:"username"`
	Password string    `gorm:"size:255;not null;" json:"password"`
	Role     RoleType  `gorm:"not null;default:'user'" json:"role"`
}

type RoleType string

const (
	RoleAdmin RoleType = "admin"
	RoleUser  RoleType = "user"
)

func (u *User) PrepareGive() {
	u.Password = ""
}

func (user *User) BeforeCreate(tx *gorm.DB) (err error) {
	// Generate UUID v4 and assign it to ID field before creating the record
	user.UUID = uuid.New()
	return nil
}

func (u *User) BeforeSave() error {

	//turn password into hash
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	u.Password = string(hashedPassword)

	//remove spaces in username
	u.Username = html.EscapeString(strings.TrimSpace(u.Username))

	return nil

}
