package main

import (
	logger "github.com/sirupsen/logrus"
	rbac "github.com/TuSimple/Role-based-access-control"
	// "github.com/TuSimple/Role-based-access-control/pkg"
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
	rbac.Init(db)
	rbac.NewUser("zhaoyang.liang")

}

