package models

import (
	"fmt"
	"html"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	gorm.Model
	UUID     uuid.UUID `gorm:"size:36;not null;unique" json:"uuid"`
	Name     string    `gorm:"size:255;not null" json:"name"`
	Username string    `gorm:"size:255;not null;unique" json:"username"`
	Password string    `gorm:"size:255;not null;" json:"password"`
	Role     RoleType  `gorm:"not null;default:'user'" json:"role"`
}

type RoleType string

const (
	RoleAdmin    RoleType = "admin"
	RoleOperator RoleType = "operator"
	RoleUser     RoleType = "user"
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

func SeedUsers() error {
	startTime := time.Now()
	users := []User{}
	for i := 1; i <= 100; i++ {
		user := User{
			Name:     fmt.Sprintf("Name User %d", i),
			Username: fmt.Sprintf("user%d@example.com", i),
			Password: "secret",
			Role:     RoleUser,
		}
		users = append(users, user)
	}

	for _, user := range users {
		err := DB.Create(&user).Error
		if err != nil {
			return err
		}
	}
	endTime := time.Now()

	// Hitung durasi waktu (lama waktu eksekusi loop)
	duration := endTime.Sub(startTime)

	// Tampilkan informasi waktu
	fmt.Printf("Start Time: %s\n", startTime.Format("2006-01-02 15:04:05"))
	fmt.Printf("End Time: %s\n", endTime.Format("2006-01-02 15:04:05"))
	fmt.Printf("Duration: %s\n", duration)
	return nil
}
