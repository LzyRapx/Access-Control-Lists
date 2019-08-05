package main

import (
	logger "github.com/sirupsen/logrus"
	rbac "github.com/TuSimple/Role-based-access-control"
	_ "github.com/TuSimple/Role-based-access-control/pkg/mongo"
	"gopkg.in/mgo.v2"
	"time"
	"math/rand"
	"fmt"
)

func main() {
	
	logger.Info("Start...")
	session, err := mgo.Dial("localhost")
	if err != nil {
		panic("cannot connect to localhost")
	}
	fmt.Println("session = ", session)
	db := session.DB(fmt.Sprintf("rbac_%d", rand.New(rand.NewSource(time.Now().UnixNano())).Int31n))
	fmt.Println("db = ", db)
	err = rbac.Init(db)
	if err != nil {
		logger.Errorf("test = %v.", err)
	}
	rbac.NewUser("zhaoyang.liang")
	rbac.NewUser("ming.xiao")
	rbac.NewUser("hong.xiao")
	rbac.GrantRole("zhaoyang.liang","admin")
	rbac.GrantRole("ming.xiao","user")
	rbac.GrantRole("hong.xiao", "guest")
	// db.DropDatabase()
}

