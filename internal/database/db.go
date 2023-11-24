package database

import (
	"fmt"
	"job-portal/config"
	"job-portal/internal/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Open() (*gorm.DB, error) {
	// dsn := os.Getenv("DB_DSN")
	cfg := config.GetConfig()
	dsn := fmt.Sprintf("%s", cfg.DbConfig.DbConn)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	// err = db.Migrator().DropTable(&models.Job{})
	// if err != nil {
	// 	return nil, err
	// }

	err = db.Migrator().AutoMigrate(&models.User{}, &models.Company{}, &models.Job{})
	if err != nil {

		return nil, err
	}
	return db, nil
}
