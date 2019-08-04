package control

import (
	"github.com/TuSimple/Role-based-access-control/resource"
)
func Res(resString string) resource.Resource {
	res, _ := resource.Parse(resString, "")
	return res
}