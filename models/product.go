package models

import (
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
)

type Product struct {
	gorm.Model
	UUID        uuid.UUID `gorm:"size:36;not null;unique" json:"uuid"`
	UserUuid    uuid.UUID `gorm:"size:36;not null" json:"user_uuid"`
	Name        string    `gorm:"size:255;not null;" json:"name"`
	Description string    `gorm:"size:255;" json:"description"`
	User        *User     `gorm:"foreignKey:UserUuid"`
}

func (product *Product) BeforeCreate(tx *gorm.DB) (err error) {
	// Generate UUID v4 and assign it to ID field before creating the record
	product.UUID = uuid.New()
	return nil
}

func SeedProducts() error {
	startTime := time.Now()
	products := []Product{}
	for i := 1; i <= 100; i++ {
		product := Product{
			Name:        fmt.Sprintf("product%d", i),
			Description: fmt.Sprintf("product%d@example.com", i),
		}
		products = append(products, product)
	}

	for _, product := range products {
		err := DB.Create(&product).Error
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
