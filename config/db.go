package config

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/spf13/viper"
)

var db *sql.DB

// ConnectSql is Singleton
func ConnectSql() *sql.DB {
	if db == nil {
		username := viper.GetString("db.username")
		password := viper.GetString("db.password")
		host := viper.GetString("db.host")
		port := viper.GetString("db.port")
		dbname := viper.GetString("db.name")

		var err error
		db, err = sql.Open("mysql", username+":"+password+"@tcp("+host+":"+port+")/"+dbname+"?parseTime=true")
		// db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3307)/zonart?parseTime=true")

		if err != nil {
			panic(fmt.Errorf("fatal error db is not connected: %w", err))
		}
	}

	return db
}
