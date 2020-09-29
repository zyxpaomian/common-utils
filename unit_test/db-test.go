package main

import (
    db "github.com/zyxpaomian/common-utils/mysqlclient"
    "fmt"
)

type Agents struct {
    Agentip string
    Alive   int
    Dpswitch    int
}

func main() {
    userName := "root"
    userPass := "root"
    addPort := "192.168.159.133:3306"
    dataBase := "rinck"
    db.DbInit(userName, userPass, addPort, dataBase)
    sql = "select * from agent_status;"
    cnt, err := db.db.Query(sql, []&Agents, &Agents)
    if err != nil {
        fmt.Println(err)
    }
    fmt.Println(cnt)
}
