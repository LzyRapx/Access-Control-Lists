package control

import (
	logger "github.com/sirupsen/logrus"
	err "github.com/TuSimple/Role-based-access-control/errors"
	// client "github.com/TuSimple/Role-based-access-control/pkg/mongo"
	Type "github.com/TuSimple/Role-based-access-control/pkg"
	
)
type Role struct {
	rbacType int
	name     string
	desc     string
}

func NewRole(roleName string) (*Role, error) {
	logger.Info("enter newRole...", egn)
	if id, rbacType, exist := egn.GetRole(roleName, true); !exist { // 存在
		return &Role{rbacType: Type.ROLE, name: roleName}, nil
	} else {
		return &Role{rbacType: rbacType, name: roleName, desc: egn.GetDesc(id)}, err.ErrDupRole
	}
}

func DropRole(roleName string) error {
	return egn.DropRole(roleName)
}

func NewUser(userName string) (*Role, error) {
	logger.Info("Add userName:", userName)
	r, err := NewRole(userName)
	r.SetAsUser()
	return r, err
}

func DropUser(userName string) error {
	return DropRole(userName)
}

func (r *Role) Name() string {
	return r.name
}

func (r *Role) Desc() string {
	return r.desc
}

func (r *Role) SetDesc(desc string) {
	id, _, _ := egn.GetRole(r.Name(), false)
	egn.SetDesc(id, desc)
	r.desc = desc
}

func (r *Role) SetAsUser() {
	r.rbacType = Type.USER
	egn.SetRoleType(r.Name(), Type.USER)
}

func (r *Role) GrantRole(grantedroles ...*Role) error {
	for _, gr := range grantedroles {
		if err := GrantRole(r.Name(), gr.Name()); err != nil {
			return err
		}
	}
	return nil
}

func (r *Role) RevokeRole(revokedRoles ...*Role) error {
	for _, rr := range revokedRoles {
		if err := RevokeRole(r.Name(), rr.Name()); err != nil {
			return err
		}
	}
	return nil
}

func (r *Role) GrantPermission(grantedPerms ...*Perm) error {
	for _, gp := range grantedPerms {
		if err := GrantPermission(r.Name(), gp.Resource().String(), gp.Name()); err != nil {
			return err
		}
	}
	return nil
}

func (r *Role) RevokePermission(revokedPerms ...*Perm) error {
	for _, rp := range revokedPerms {
		if err := RevokePermission(r.Name(), rp.Resource().String(), rp.Name()); err != nil {
			return err
		}
	}
	return nil
}

func (r *Role) Drop() error {
	return egn.DropRole(r.Name())
}

func (r *Role) HasRole(roles ...*Role) bool {
	rl := []string{}
	for _, role := range roles {
		rl = append(rl, role.Name())
	}
	return egn.HasAllRole(r.Name(), rl...)
}

func (r *Role) HasPermission(perms ...*Perm) bool {
	for _, perm := range perms {
		if !egn.Decision(r.Name(), perm.Resource().Name(), perm.Name()) {
			return false
		}
	}
	return true
}

func (r *Role) HasPermission2(perms ...*Perm) bool {
	for _, perm := range perms {
		if !egn.DecisionEx(r.Name(), perm.Resource().Name(), perm.Name()) {
			return false
		}
	}
	return true
}