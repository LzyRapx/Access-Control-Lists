package pkg

import (
	"fmt"
	"reflect"
)

type InitFunc func(interface{}) (RbacInterface, error)

type Registry map[string]InitFunc

var registry = make(Registry)

func Register(conn interface{}, f InitFunc) {
	fmt.Println("register. type = ", reflect.TypeOf(conn).String())
	registry[reflect.TypeOf(conn).String()] = f
}

func Factory(conn interface{}) (RbacInterface, error){
	if f, flag := registry[reflect.TypeOf(conn).String()]; !flag {
		return nil, fmt.Errorf("Erroe. conn type %T does not registered\n", conn)
	} else {
		return f(conn)
	}
}




