package app

import (
	"github.com/a1ta1r/Credit-Portfolio/internal/models"
	"os"

	"github.com/jinzhu/gorm"

	//Required by GORM to run over MSSQL and Postgres databases
	_ "github.com/jinzhu/gorm/dialects/mssql"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var sqlConn *gorm.DB

// GetConnection returns Connection to an MSSQL database using GORM
func GetConnection() (gorm.DB, error) {
	if sqlConn == nil {
		if conn, err := getMssql(); err == nil {
			sqlConn = &conn
		} else if conn, err := getPostgres(); err == nil {
			sqlConn = &conn
		} else {
			return gorm.DB{}, err
		}
	}
	return *sqlConn, nil
}

func getMssql() (gorm.DB, error) {
	conn, err := gorm.Open("mssql", os.Getenv("MSSQL"))
	if err != nil {
		println(err.Error())
		return *conn, err
	}
	return *conn, nil
}

func getPostgres() (gorm.DB, error) {
	conn, err := gorm.Open("postgres", os.Getenv("POSTGRES"))
	if err != nil {
		println(err.Error())
		return *conn, err
	}
	return *conn, nil
}

func SyncModelsWithSchema() {
	db, err := GetConnection()
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(
		&models.Bank{},
		&models.Currency{},
		&models.User{},
		&models.PaymentPlan{},
		&models.Payment{},
		&models.Income{},
		&models.Expense{},
		&models.BannerPlace{},
		&models.Banner{},
		&models.Advertiser{},
		&models.Advertisement{},
	)
}
