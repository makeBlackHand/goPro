package utily

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

var (
	Db  *sql.DB
	err error
)

const (
	dbUser  = "go"
	dbPass  = "PassCode987!"
	dbProto = "tcp"
	dbHost  = "api.ikanned.com"
	dbPort  = "4407"
	dbName  = "db1"
	dbTable = "user1"
)

func init() {
	Db, err = sql.Open("mysql", fmt.Sprintf("%s:%s@%s(%s:%s)/%s", dbUser, dbPass, dbProto, dbHost, dbPort, dbName))
	if err != nil {
		panic(err.Error())
	}
}
