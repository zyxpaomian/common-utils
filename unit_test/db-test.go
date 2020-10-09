package main

import (
    "github.com/zyxpaomian/common-utils/simplemysql"
    "fmt"
    "github.com/jmoiron/sqlx"
)

type Agent struct {
    Agentip string
    Alive   int
    Dpswitch    int
}

func SingleQuery(db *sqlx.DB) {
    sql := "select agentip, alive, dpswitch from agent_status where agentip = ?"
    var a Agent
    err := db.Get(&a, sql, "1.1.1.1")

	if err != nil {
		fmt.Printf("get failed, err: %v \n", err)
		return
	}
    fmt.Printf("Single Data Query --- ip: %s alive: %d Dpswitch: %d\n", a.Agentip, a.Alive, a.Dpswitch)    
}

func MultiQuery(db *sqlx.DB) {
    sql := "select agentip, alive, dpswitch from agent_status where alive = ?"
    var as []Agent
    err := db.Select(&as, sql, 1)
    if err != nil {
        fmt.Printf("get failed, err: %v \n", err)
        return
    }
    for k, v := range(as){
        fmt.Printf("Multi Data Query row line: %d --- ip: %s alive: %d Dpswitch: %d\n", k, v.Agentip, v.Alive, v.Dpswitch)
    }
}

func SingleInsert(db *sqlx.DB) {
	sql := "insert into agent_status(agentip, alive, dpswitch) values (?, ?, ?)"
	ret, err := db.Exec(sql, "2.2.2.7", 0, 1)
	if err != nil {
		fmt.Printf("insert failed, err: %v \n", err)
		return
	}
	theID, err := ret.LastInsertId()
	if err != nil {
		fmt.Printf("get lastinsert ID failed, err: %v \n", err)
		return
	}
	fmt.Printf("insert success, the id is %d. \n", theID)  
}

func Update(db *sqlx.DB) {
	sql := "update agent_status set alive= ? where dpswitch = ?"
	ret, err := db.Exec(sql, 1, 1)
	if err != nil {
		fmt.Printf("update failed, err: %v \n", err)
		return
	}
	n, err := ret.RowsAffected()
	if err != nil {
		fmt.Printf("get RowsAffected failed, err: %v \n", err)
		return
	}
	fmt.Printf("update success, affected rows: %d \n", n)
}

func Delete(db *sqlx.DB) {
	sql := "delete from agent_status where alive = ?"
	ret, err := db.Exec(sql, 1)
	if err != nil {
		fmt.Printf("delete failed, err: %v \n", err)
		return
	}
	n, err := ret.RowsAffected()
	if err != nil {
		fmt.Printf("get RowsAffected failed, err: %v \n", err)
		return
	}
	fmt.Printf("delete success, affected rows: %d \n", n)
}

func main() {
    userName := "root"
    userPass := "root"
    addPort := "192.168.159.133:3306"
    dataBase := "rinck"
    simplemysql.DB.DbInit(userName, userPass, addPort, dataBase)

    db := simplemysql.DB.Con
    SingleQuery(db)
    MultiQuery(db)
    SingleInsert(db)
    Update(db)
    Delete(db)


  


}