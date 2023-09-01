package configs

import (
	"context"
	"os"

	"github.com/redis/go-redis/v9"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"github.com/iarsham/url-shortener/models"
)

var DB *gorm.DB
var RDB *redis.Client

func GetDB(lg *CustomLogger) (*gorm.DB, error) {
	var err error
	DB, err = gorm.Open(mysql.Open(os.Getenv("DB_URL")), &gorm.Config{})
	if err != nil {
		lg.Logger.Fatal(err.Error())
	}
	lg.Logger.Println("Mysql is Connected Successfuly!")
	if err := DB.AutoMigrate(&models.User{}, &models.Link{}); err != nil {
		lg.Logger.Fatal(err.Error())
	}
	lg.Logger.Println("Tables migrations was successfully")
	return DB, nil
}

func GetRedis(lg *CustomLogger) *redis.Client {
	RDB = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})
	if _, err := RDB.Ping(context.Background()).Result(); err != nil  {
		lg.Logger.Fatal(err.Error())
	}
	lg.Logger.Println("Redis is Connected Successfuly!")
	return RDB
}
