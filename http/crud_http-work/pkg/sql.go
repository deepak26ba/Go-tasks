package pkg

import (
	"fmt"
	"http/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitDB(connectionKey string) (*gorm.DB, error) {

	db, err := gorm.Open(postgres.Open(connectionKey), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("Failed connecting to DB : %v", err)
	}

	err = db.AutoMigrate(&models.Users{}, &models.UserRoles{}, &models.UserRoleMapping{})
	if err != nil {
		return nil, fmt.Errorf("Failed migrate : %v", err)
	}

	return db, nil

}
