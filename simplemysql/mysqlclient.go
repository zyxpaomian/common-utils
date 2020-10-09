package simplemysq

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)



//var db *sqlx.DB

type MySQLUtil struct {
	db          *sql.DB
	initialized bool
}

var DB = MySQLUtil{db: nil, initialized: false}

func (m *MySQLUtil) DbInit(userName, userPass, addrPort, dataBase string) {
	//dsn := "user:password@tcp(127.0.0.1:3306)/sql_test?charset=utf8mb4&parseTime=True"
	connFormat := "%s:%s@tcp(%s)/%s?autocommit=0&collation=utf8_general_ci&parseTime=true"
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

	m.db = db
	m.initialized = true

	return
}



