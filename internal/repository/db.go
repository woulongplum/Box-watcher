package repository

import (
	"fmt"
	"os"

	"github.com/woulongplum/Box-watcher/internal/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)


func InitDB()(*gorm.DB, error) {
		dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_SSLMODE"),
	)	

	db, err := gorm.Open(postgres.Open(dsn),&gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("DBへの接続に失敗しました: %w", err)
	}

	err = db.AutoMigrate(&model.Item{})
	if err != nil {
		return  nil , fmt.Errorf("テーブルの作成に失敗しました: %w", err)
	}
	
	fmt.Println("DBに接続しました")
	return db, nil
}
