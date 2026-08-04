package seeder

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"

	"wedding/internal/infrastructure/dbmodel"
)

func SeedAdminUser(db *gorm.DB) error {
	fmt.Println("Seeding admin account...")

	email := "admin@gmail.com"
	password := "123456"

	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return fmt.Errorf("failed to hash password: %w", err)
	}

	var existing dbmodel.User
	err = db.Where("email = ?", email).First(&existing).Error
	if err == nil {
		fmt.Println("Admin account already exists.")
		return nil
	}
	if err != gorm.ErrRecordNotFound {
		return fmt.Errorf("failed to check existing admin: %w", err)
	}

	user := dbmodel.User{
		UUID:         "00000000-0000-0000-0000-000000000001",
		Name:         "Administrator",
		Email:        email,
		PasswordHash: string(hash),
		Role:         "admin",
	}

	if err := db.Create(&user).Error; err != nil {
		return fmt.Errorf("failed to seed admin account: %w", err)
	}

	fmt.Println("Seeded admin account successfully.")
	return nil
}

func Seed() error {
	return fmt.Errorf("not implemented")
}
