package simplemysql

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

type MySQLUtil struct {
	db          *sql.DB
	initialized bool
}

var DB = MySQLUtil{db: nil, initialized: false}

func (m *MySQLUtil) DbInit(userName, userPass, addrPort, dataBase string) {
	m.CloseConn()
	connFormat := "%s:%s@tcp(%s)/%s?autocommit=0&collation=utf8_general_ci&parseTime=true"
	connStr := fmt.Sprintf(
		connFormat,
        userName,
        userPass,
        addrPort,
        dataBase,
	)

	db, err := sql.Open("mysql", connStr)
	if err != nil {
        panic("打开mysql 连接失败")
	}

	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(1)

	err = db.Ping()
	if err != nil {
		panic("mysql初始化失败,ping失败")
	}

	m.db = db
	m.initialized = true
}

// 关闭整个数据库连接
func (m *MySQLUtil) CloseConn() {
	if m.initialized {
		m.db.Close()
		m.db = nil
		m.initialized = false
	}
}

// 检查事务
func (m *MySQLUtil) GetTx() (*sql.Tx, error) {
	if m.initialized == false {
		return nil, fmt.Errorf("db initialized failed")
	}
	tx, err := m.db.Begin()
	if err != nil {
		return nil, fmt.Errorf("db get tx failed")
	}
	return tx, nil
}


// 基本查询,基于结构体进行查询，返回值是查询的具体行数
func (m *MySQLUtil) Query(sql string,  resultlist *[]interface{}, result ...interface{}) (int64, error) {
	tx, err := m.GetTx()
	if err != nil {
		return -1, err
	}

	stmt, err := tx.Prepare(sql)
	if err != nil {
		tx.Rollback()
		return -1, err
	}

	rows, err := stmt.Query()
	if err != nil {
		stmt.Close()
		return -1, err
	}

	var cnt int64 = 0
	for rows.Next() {
		err := rows.Scan(result...)
		if err != nil {
			rows.Close()
			stmt.Close()
			tx.Rollback()
			return -1, err
		} else {
            *resultlist = append(*resultlist, result)
			cnt += 1
			//break
		}
	}

	err = rows.Err()
	if err != nil {
		rows.Close()
		stmt.Close()
		tx.Rollback()
		return -1, err
	}
	rows.Close()
	stmt.Close()
	tx.Commit()
	return cnt, nil
}

/*
func (m *MySQLUtil) SimpleInsert(sql string, args ...interface{}) (int, error) {
	if m.initialized == false {
		log.Errorln("MySQL 还未初始化")
		return -1, ce.DBError()
	}
	tx := m.GetTx()
	if tx == nil {
		log.Errorln("MySQL 获取TX失败")
		return -1, ce.DBError()
	}
	stmt, err := tx.Prepare(sql)
	if err != nil {
		tx.Rollback()
		log.Errorln("MySQL Prepare失败: ", err.Error())
		return -1, ce.DBError()
	}
	res, err := stmt.Exec(args...)
	if err != nil {
		stmt.Close()
		tx.Rollback()
		log.Errorln("MySQL 执行Insert失败: ", err.Error())
		return -1, ce.DBError()
	}
	InsertID, _ := res.LastInsertId()
	stmt.Close()
	err = tx.Commit()
	if err != nil {
		log.Errorln("MySQL 执行Insert失败: ", err.Error())
		return -1, ce.DBError()
	}
	return int(InsertID), nil
}

func (m *MySQLUtil) SimpleUpdate(sql string, args ...interface{}) (int, error) {
	if m.initialized == false {
		log.Errorln("MySQL 还未初始化")
		return -1, ce.DBError()
	}
	tx := m.GetTx()
	if tx == nil {
		log.Errorln("MySQL 获取TX失败")
		return -1, ce.DBError()
	}
	stmt, err := tx.Prepare(sql)
	if err != nil {
		tx.Rollback()
		log.Errorln("MySQL Prepare失败: ", err.Error())
		return -1, ce.DBError()
	}
	res, err := stmt.Exec(args...)
	if err != nil {
		stmt.Close()
		tx.Rollback()
		log.Errorln("MySQL 执行Update失败: ", err.Error())
		return -1, ce.DBError()
	}
	AddectIDs, _ := res.RowsAffected()
	stmt.Close()
	err = tx.Commit()
	if err != nil {
		log.Errorln("MySQL 执行Update失败: ", err.Error())
		return -1, ce.DBError()
	}
	return int(AddectIDs), nil
}*/