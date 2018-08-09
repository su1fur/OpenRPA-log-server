package models

import (
	"fmt"
	"os"
	"time"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

var DB *gorm.DB

type Model struct {
	ID        uint64     `gorm:"primary_key" json:"id"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at"`
}

func InitDB() {

	user := getParamFromEnv("OPEN_RPA_LOG_DB_USER", "")
	host := getParamFromEnv("OPEN_RPA_LOG_DB_HOST", "")
	password := getParamFromEnv("OPEN_RPA_LOG_DB_PASSWORD", "")
	name := getParamFromEnv("OPEN_RPA_LOG_DB_NAME", "")
	protocol := getParamFromEnv("OPEN_RPA_LOG_DB_PROTOCOL", "tcp")
	port := getParamFromEnv("OPEN_RPA_LOG_DB_PORT", "3306")
	args := getParamFromEnv("OPEN_RPA_LOG_DB_ARGS", "parseTime=true&loc=Asia%2FTokyo")

	connectionString := fmt.Sprintf("%s:%s@%s([%s]:%s)/%s?%s",
		user, password, protocol, host, port, name, args)

	db, err := gorm.Open("mysql", connectionString)
	if err != nil {
		fmt.Println("Failed to connect to database: %v\n", err)
	}

	db.DB()

	// モデルに応じてデータベースのテーブルを作成(マイグレーション処理)
	db.AutoMigrate(&Task{})

	DB = db.Debug()
}

func getParamFromEnv(env, defaultValue string) string {
	param := os.Getenv(env)

	if param == "" {
		param = defaultValue
	}

	return param
}
