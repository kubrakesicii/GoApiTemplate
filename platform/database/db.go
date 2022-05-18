package database

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	db *gorm.DB
)

// func New(db *gorm.DB) *DB{
// 	return &DB{
// 		db:db
// 	}
// }

// type DB struct{
// 	db *gorm.DB
// }

func Connect() {
	err := godotenv.Load()
	if err != nil {
		panic("Failed to load env file")
	}

	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASS")
	dbHost := os.Getenv("DB_HOST")
	dbName := os.Getenv("DB_NAME")

	connString := fmt.Sprintf("%s:%s@tcp(%s:3306)/%s?charset=utf8&parseTime=True&loc=Local", dbUser, dbPass, dbHost, dbName)

	d, err := gorm.Open(mysql.Open(connString), &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: true,
	})

	if err != nil {
		panic("Failed to create connection to database")
	}
	db = d
}

// func (d DB) Test(ctx context.Context){
// 	d.db.ConnPool.QueryContext(ctx, "SELECT * FROM users")
// }

func CloseDb(db *gorm.DB) {
	dbSQL, err := db.DB()
	if err != nil {
		panic("Failed to close connection from database")
	}
	dbSQL.Close()
}

func GetDb() *gorm.DB {
	Connect()
	//db.AutoMigrate(&entity.User{}, entity.Test{})
	return db
}
