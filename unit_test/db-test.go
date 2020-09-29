package main

import (
    "github.com/zyxpaomian/common-utils/mysqlclient"
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
    mysqlclient.DbInit(userName, userPass, addPort, dataBase)
    sql = "select * from agent_status;"
    cnt, err := mysqlclient.db.Query(sql, []&Agents, &Agents)
    if err != nil {
        fmt.Println(err)
    }
    fmt.Println(cnt)
}
