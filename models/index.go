package models

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Config struct {
	Host     string
	Port     string
	User     string
	Passwrod string
	DBName   string
	SSLMode  string
}

var DB *gorm.DB

func InitDB(cnf Config) {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s", 
							cnf.Host, cnf.User,
							cnf.Passwrod, cnf.DBName,
							cnf.Port, cnf.SSLMode)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		fmt.Println("Error Opening the databse")
		panic(err)
	}

	if err := db.AutoMigrate((&User{})); err != nil {
		fmt.Println("Error migrating the database")
		panic(err)
	}

	fmt.Println("database was migrated")

	DB = db
}