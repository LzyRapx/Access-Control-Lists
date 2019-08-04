package control

import (
	"github.com/TuSimple/Role-based-access-control/resource"
	err "github.com/TuSimple/Role-based-access-control/errors"
)
type Perm struct {
	name     string
	resource resource.Resource
}

func NewPerm(permName string, res resource.Resource) (*Perm, error) {
	if _, exist := egn.GetPermission(permName, res.String(), true); exist {
		return &Perm{permName, res}, err.ErrDupPerm
	} else {
		return &Perm{permName, res}, nil
	}
}

func NewGlobalPerm(permName string) (*Perm, error) {
	return NewPerm(permName, Res(""))
}

func (p *Perm) Resource() resource.Resource {
	return p.resource
}

func (p *Perm) Name() string {
	return p.name
}

func (p *Perm) Drop() error {
	return egn.DropPermission(p.Name(), p.Resource().String())
}
