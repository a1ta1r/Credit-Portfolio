package services

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"os"
)

var pgsqlConnection *gorm.DB

func GetConnection() (*gorm.DB, error) {
	if pgsqlConnection == nil {
		conn, err := gorm.Open("postgres", os.Getenv("DATABASE_URL"))
		if err != nil {
			return pgsqlConnection, err
		}
		pgsqlConnection = conn
	}
	return pgsqlConnection, nil
}
