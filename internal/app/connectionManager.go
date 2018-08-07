package app

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mssql"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"os"
)

var sqlConn *gorm.DB

func GetConnection() (gorm.DB, error) {
	if sqlConn == nil {
		if conn, err := getPostgres(); err == nil {
			sqlConn = &conn
		} else if conn, err := getMssql(); err == nil {
			sqlConn = &conn
		} else {
			return gorm.DB{}, err
		}
	}
	return *sqlConn, nil
}

func getPostgres() (gorm.DB, error) {
	//"host=myhost port=myport user=gorm dbname=gorm password=mypassword"
	conn, err := gorm.Open("postgres", os.Getenv("POSTGRES_URL"))
	if err != nil {
		return *conn, err
	}
	return *conn, nil
}

func getMssql() (gorm.DB, error) {
	//"sqlserver://username:password@localhost:1433?database=dbname"
	conn, err := gorm.Open("mssql", os.Getenv("MSSQL_URL"))
	if err != nil {
		return *conn, err
	}
	return *conn, nil
}
