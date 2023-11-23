package database

import (
	"job-portal/internal/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Open() (*gorm.DB, error) {
	dsn := "host=postgres user=postgres password=admin dbname=job-portal-api port=5432 sslmode=disable TimeZone=Asia/Shanghai"
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
