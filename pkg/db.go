package pkg

import (
	"github.com/alecthomas/kingpin/v2"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"strconv"
)

type DBConfig struct {
	DBName       string `json:"db_name"`
	DBUser       string `json:"db_user"`
	DBPass       string `json:"db_pass"`
	DBHost       string `json:"db_host"`
	DBPort       string `json:"db_port"`
	MaxOpenConns int    `json:"max_open_conns"`
}

// ParseDBFlag -
func ParseDBFlag(dbc *DBConfig) {
	kingpin.Flag("db-host", "The host of database").Default(dbc.DBHost).Envar("MYSQL_HOST").StringVar(&dbc.DBHost)
	kingpin.Flag("db-port", "The port of database").Default(dbc.DBPort).Envar("MYSQL_PORT").StringVar(&dbc.DBPort)

	kingpin.Flag("db-user", "The user name of database").Default(dbc.DBUser).Envar("MYSQL_USERNAME").StringVar(&dbc.DBUser)
	kingpin.Flag("db-pass", "The password of database").Default(dbc.DBPass).Envar("MYSQL_PASSWORD").StringVar(&dbc.DBPass)
	kingpin.Flag("db-name", "The database name of database").Default(dbc.DBName).Envar("MYSQL_DATABASE").StringVar(&dbc.DBName)
	kingpin.Flag("db-max-open-conns", "The maximum number of open connections to the database.").Default(strconv.Itoa(dbc.MaxOpenConns)).Envar("DB_MAX_OPEN_CONNS").IntVar(&dbc.MaxOpenConns)
}
