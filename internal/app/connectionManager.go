package app

import (
	ae "github.com/a1ta1r/Credit-Portfolio/internal/components/advertisements/entities"
	fe "github.com/a1ta1r/Credit-Portfolio/internal/components/finance/entities"
	le "github.com/a1ta1r/Credit-Portfolio/internal/components/loans/entities"
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
		&fe.Bank{},
		&fe.Currency{},
		&le.User{},
		&le.PaymentPlan{},
		&le.Payment{},
		&le.Income{},
		&le.Expense{},
		&ae.BannerPlace{},
		&ae.Banner{},
		&ae.Advertiser{},
		&ae.Advertisement{},
	)
}
