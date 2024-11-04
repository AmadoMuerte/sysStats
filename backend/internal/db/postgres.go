package db

import (
	"fmt"
	"log"

	"github.com/AmadoMuerte/FlickSynergy/internal/config"
	"github.com/AmadoMuerte/FlickSynergy/internal/db/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Storage struct {
	DB *gorm.DB
}

func New(cfg *config.Config) *Storage {

	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s",
		cfg.DB.Host, cfg.DB.User, cfg.DB.Pass, cfg.DB.Name, cfg.DB.Port,
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{DisableForeignKeyConstraintWhenMigrating: false})
	if err != nil {
		panic(err)
	}
	autoMigrateModels(db)

	return &Storage{DB: db}
}

func autoMigrateModels(db *gorm.DB) {
	modelsSlice := []interface{}{}

	modelsSlice = append(
		modelsSlice,
		&models.User{},
	)

	// Выполнение миграции
	if err := db.AutoMigrate(modelsSlice...); err != nil {
		log.Fatalf("Ошибка миграции: %v", err)
	}
}
