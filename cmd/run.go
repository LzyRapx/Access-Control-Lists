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
		logger.Errorf("err = %v.", err)
	}
	rbac.NewUser("zhaoyang.liang")
	rbac.NewUser("pony.ma")
	rbac.NewUser("zeming.jiang")
	rbac.NewUser("jack.ma")

	rbac.GrantRole("zhaoyang.liang","super admin", "cluster admin", "namespace admin", "user") // zhaoyang.liang has admin permission
	rbac.GrantRole("pony.ma","cluster admin", "namespace admin", "user")
	rbac.GrantRole("zeming.jiang", "namesapce admin", "user")
	rbac.GrantRole("jack.ma", "user")
	

	rbac.GrantRole("super admin", "cluster admin")
	rbac.GrantRole("cluster admin", "namespace admin")
	rbac.GrantRole("namespace admin", "user")


	task := "www.baidu.com"
	rbac.GrantPermission("zhaoyang.liang", task, "delete", "update")
	rbac.GrantPermission("pony.ma", task, "insert", "read", "readlist")
	rbac.GrantPermission("jack.ma", task, "create")

	rbac.GrantGlobalPermission("super admin", "read")

	if rbac.HasRole("zhaoyang.liang","cluster admin") == false {
		logger.Info("zhaoyang.liang is a super admin, should has cluster admin's perimission")
	}
	if rbac.HasRole("zhaoyang.liang","super admin") == true {
		logger.Info("zhaoyang.liang is a super admin, should has super admin's perimission")
	}
	if rbac.HasRole("zhaoyang.liang","user") == true {
		logger.Info("zhaoyang.liang is a super admin, should has user's perimission")
	}
	if rbac.HasRole("zhaoyang.liang","user") == true {
		logger.Info("zhaoyang.liang is a super admin, should has user's perimission")
	}
	if rbac.HasRole("jack.ma", "delete") == false {
		logger.Info("Jack.ma is just a user, should not has delete permission")
	}
    if rbac.Decision("jack.ma", task, "delete") == false {
		logger.Info("jack.ma is user and should not have delete task permission")
	}
	if rbac.DecisionEx("jack.ma", "abc", "create") == false {
		logger.Info("jack.ma should has all create permission on all target")
	}
	rbac.RevokeRole("zhaoyang.liang","super admin")
	if rbac.HasRole("zhaoyang.liang","cluster admin") == false {
		logger.Info("zhaoyang.liang is not a super admin now, should not has cluster admin's perimission")
	}
	db.DropDatabase()
}

