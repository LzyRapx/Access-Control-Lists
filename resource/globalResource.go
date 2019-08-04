package resource

import (
	err "github.com/TuSimple/Role-based-access-control/errors"
)

type GlobalRes struct{}

func ParseGlobalRes(resString, name string) (Resource, error) {
	if resString=="" && name=="" {
		return &GlobalRes{}, nil
	} else {
		return nil, err.ErrParseRes
	}
}

func (g *GlobalRes) Name() string {
	return "GLOBAL"
}

func (g *GlobalRes) Equals(resource Resource) bool {
	return false
}

//global resource contains all resource
func (g *GlobalRes) Contains(resl ...Resource) bool {
	return true
}

func (g *GlobalRes) String() string {
	return ""
}