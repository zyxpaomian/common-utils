package simplemysql

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)


type MySQLUtil struct {
	Con         *sqlx.DB
}

var DB = MySQLUtil{Con: nil}

func (m *MySQLUtil) DbInit(userName, userPass, addrPort, dataBase string) {
	connFormat := "%s:%s@tcp(%s)/%s?autocommit=1&collation=utf8_general_ci&parseTime=true"
	connStr := fmt.Sprintf(
		connFormat,
        userName,
        userPass,
        addrPort,
        dataBase,
	)

	db, err := sqlx.Connect("mysql", connStr)
	if err != nil {
		panic("打开mysql 连接失败")
	}

	db.SetMaxOpenConns(20)
	db.SetMaxIdleConns(10)


	err = db.Ping()
	if err != nil {
		panic("mysql初始化失败,ping失败")
	}

	m.Con = db
	return
}