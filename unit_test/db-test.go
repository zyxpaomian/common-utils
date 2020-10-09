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
    var db *simplemysql.DB.db

    sql := "select agentip, alive, dpswitch from agent_status;"

    var a Agent
    err := db.Get(&a, sql, 1)
	if err != nil {
		fmt.Printf("get failed, err:%v\n", err)
		return
	}
    fmt.Printf("ip:%s alive:%s Dpswitch:%d\n", a.Agentip, a.Alive, a.Dpswitch)

    /*agentObject := &Agent{}
    _, err := simplemysql.DB.Query(sql, &agentObject.Agentip, &agentObject.Alive, &agentObject.Dpswitch)
    if err != nil {
        fmt.Println(err)
    }

    fmt.Println(agentObject.Agentip)
    //for _, v := range(agentObject){
        //fmt.Println(v)
    //}

    //AgentList := []*Agent{}
    /*Agents := []Agent{}
    AgentList := make([]interface{}, 0)
    
    for i, v := range Agents {
        AgentList[i] = v
    }
    fmt.Println(AgentList)
    agentObject := &Agent{}
    fmt.Printf("%T\n",&agentObject.Agentip)
    _, err := simplemysql.DB.Query(sql, &AgentList, &agentObject.Agentip, &agentObject.Alive, &agentObject.Dpswitch)
    if err != nil {
        fmt.Println(err)
    }
    for _, v := range AgentList {
        fmt.Printf("%T\n",v)
        op, _ := v.(Agent)
        fmt.Println(op.Agentip)
    }

    //fmt.Println(AgentList)
    //fmt.Println(cnt)
    /*fmt.Printf("%T\n",aa)
    for _, v := range aa {
        fmt.Printf("%T\n",v)
        op, _ := v.(Agent)
        fmt.Println(op.Agentip)
        /*p, ok := (v.Value).(Agent)
        if ok {
            fmt.Println(p)
        } else {            
            fmt.Println("err")
        }*/
       // fmt.Printf("%T",v)
    //}*/
    //fmt.Println(agentObject.Agentip)
}