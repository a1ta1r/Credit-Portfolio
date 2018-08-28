package app

import (
	"os"

	"github.com/jinzhu/gorm"

	//Required by GORM to run over MSSQL database
	_ "github.com/jinzhu/gorm/dialects/mssql"
)

var sqlConn *gorm.DB

// GetConnection returns Connection to an MSSQL database using GORM
func GetConnection() (gorm.DB, error) {
	if sqlConn == nil {
		if conn, err := getMssql(); err == nil {
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
