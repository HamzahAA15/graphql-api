package util

import (
	"database/sql"
	"fmt"
	"sirclo/gql/config"

	_ "github.com/go-sql-driver/mysql"
	"github.com/labstack/gommon/log"
)

var Db *sql.DB

func MysqlDriver(config *config.AppConfig) *sql.DB {

	uri := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?parseTime=true",
		config.Database.Username,
		config.Database.Password,
		config.Database.Address,
		config.Database.Port,
		config.Database.Name)

	db, err := sql.Open(config.Database.Driver, uri)
	if err != nil {
		log.Info("failed to connect database: ", err)
		panic(err)
	}
	Db = db
	return Db
}
