package app

import (
	"github.com/go-sql-driver/mysql"
	"os"
)

type Config struct {
	dbConfig mysql.Config
}

func LoadConfig() Config {
	config := Config{}
	dbConfig := mysql.Config{
		User:      os.Getenv("MYSQL_USER"),
		Passwd:    os.Getenv("MYSQL_PASS"),
		Addr:      os.Getenv("MYSQL_ADDR"),
		Net:       os.Getenv("MYSQL_NET"),
		DBName:    os.Getenv("MYSQL_DBNAME"),
		ParseTime: true,
	}
	config.dbConfig = dbConfig
	return config
}

func (config Config) GetDBConfig() mysql.Config {
	return config.dbConfig
}
