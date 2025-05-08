package tests

import (
	"fmt"
	"github.com/joho/godotenv"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var (
	db *gorm.DB
)

func connect(driver string) *gorm.DB {
	if err := godotenv.Load(".env"); err != nil {
		panic("`.env` file does not exist, please copy `.env.example` file to `.env` file")
	}
	var dia gorm.Dialector
	switch driver {
	case "mysql":
		dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", os.Getenv("MySQL_DB_USERNAME"), os.Getenv("MySQL_DB_PASSWORD"), os.Getenv("MySQL_DB_HOST"), os.Getenv("MySQL_DB_PORT"), os.Getenv("MySQL_DB_DATABASE"))
		dia = mysql.Open(dsn)
	case "pgsql":
		dsn := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable TimeZone=Asia/Shanghai", os.Getenv("PgSQL_DB_HOST"), os.Getenv("PgSQL_DB_PORT"), os.Getenv("PgSQL_DB_USERNAME"), os.Getenv("PgSQL_DB_DATABASE"), os.Getenv("PgSQL_DB_PASSWORD"))
		dia = postgres.Open(dsn)
	}
	db, err := gorm.Open(dia, &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		panic("failed to connect database")
	}
	return db
}
