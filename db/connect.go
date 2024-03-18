package db

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

const (
	HOST     = "localhost"
	USER     = "prime"
	PASSWORD = "belindat2014"
	DBNAME   = "luxury"
	PORT     = "5433"
)

const (
	TEST_HOST     = "postgres"
	TEST_USER     = "prime"
	TEST_PASSWORD = "belindat2014"
	TEST_DBNAME   = "test_luxury"
	TEST_PORT     = "5432"
)

var Db *gorm.DB

func Connect() {
	db, err := gorm.Open(postgres.New(postgres.Config{
		DSN:                  fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s", HOST, USER, PASSWORD, DBNAME, PORT),
		PreferSimpleProtocol: true,
	}), &gorm.Config{})

	if err != nil {
		db, err = gorm.Open(postgres.New(postgres.Config{
			DSN:                  fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s", TEST_HOST, TEST_USER, TEST_PASSWORD, TEST_DBNAME, TEST_PORT),
			PreferSimpleProtocol: true,
		}), &gorm.Config{})
	}
	Db = db
}

func ConnectTestDB() {

	db, err := gorm.Open(postgres.New(postgres.Config{
		DSN:                  fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s", TEST_HOST, TEST_USER, TEST_PASSWORD, TEST_DBNAME, TEST_PORT),
		PreferSimpleProtocol: true,
	}), &gorm.Config{})

	if err != nil {
		panic("Failed to connect to database")
	}
	Db = db
}
