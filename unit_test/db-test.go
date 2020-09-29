package main

import (
    "github.com/zyxpaomian/common-utils/simplemysql"
    "fmt"
)

type Agent struct {
    Agentip string
    Alive   int
    Dpswitch    int
}

func main() {
    userName := "root"
    userPass := "root"
    addPort := "192.168.159.133:3306"
    dataBase := "rinck"
    simplemysql.DB.DbInit(userName, userPass, addPort, dataBase)
    sql := "select agentip, alive, dpswitch from agent_status;"
    //AgentList := []*Agent{}
    Agents := []Agent{}
    AgentList := make([]interface{}, 2)
    for i, v := range Agents {
        AgentList[i] = v
    }
    agentObject := &Agent{}
    cnt, err := simplemysql.DB.Query(sql, AgentList, &agentObject.Agentip, &agentObject.Alive, &agentObject.Dpswitch)
    if err != nil {
        fmt.Println(err)
    }
    fmt.Println(cnt)
    fmt.Println(agentObject.Agentip)
}
