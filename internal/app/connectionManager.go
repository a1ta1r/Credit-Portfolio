package app

import (
	ae "github.com/a1ta1r/Credit-Portfolio/internal/components/advertisements/entities"
	fe "github.com/a1ta1r/Credit-Portfolio/internal/components/finance/entities"
	le "github.com/a1ta1r/Credit-Portfolio/internal/components/loans/entities"
	"os"

	"github.com/jinzhu/gorm"

	_ "github.com/jinzhu/gorm/dialects/postgres"
	_ "github.com/jinzhu/gorm/dialects/mssql"

)

var sqlConn *gorm.DB

// GetConnection returns Connection to an MSSQL database using GORM
func GetConnection() (gorm.DB, error) {
	if sqlConn == nil {
		if conn, err := getMssql(); err == nil {
			sqlConn = &conn
		} else if conn, err = getPostgres(); err == nil {
			sqlConn = &conn
		} else {
			return gorm.DB{}, err
		}
	}
	return *sqlConn, nil
}

func getMssql() (gorm.DB, error) {
	conn, err := gorm.Open("mssql", os.Getenv("CREDIT_API_MSSQL"))
	if err != nil {
		println(err.Error())
		return *conn, err
	}
	return *conn, nil
}

func getPostgres() (gorm.DB, error) {
	conn, err := gorm.Open("postgres", os.Getenv("CREDIT_API_POSTGRES"))
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
		&ae.Advertiser{},
		&ae.Advertisement{},
		&ae.BannerPlace{},
		&ae.Banner{},
	)
}

func DropAllTables() {
	db, err := GetConnection()
	if err != nil {
		panic(err)
	}
	db.DropTable(
		&fe.Bank{},
		&fe.Currency{},
		&le.User{},
		&le.PaymentPlan{},
		&le.Payment{},
		&le.Income{},
		&le.Expense{},
		&ae.Advertiser{},
		&ae.Advertisement{},
		&ae.BannerPlace{},
		&ae.Banner{},
	)
}