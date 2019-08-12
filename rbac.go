package control

import (
	logger "github.com/sirupsen/logrus"
	interFace "github.com/TuSimple/Role-based-access-control/pkg"
)
var (
	egn interFace.RbacInterface
)

func Init(conn interface{}) error {
	logger.Info("Rbac Init...")
	e, err := interFace.Factory(conn)
	if err != nil {
		logger.Errorf("%v.", err)
	}
	egn = e
	logger.Info("engine = ", egn)
	return err
}

func HasRole(roleName string, hasRoleName string) bool {
	return egn.HasAllRole(roleName, hasRoleName)
}

func HasAllRole(roleName string, hasRoleNames ...string) bool {
	return egn.HasAllRole(roleName, hasRoleNames...)
}

func HasAnyRole(roleName string, hasRoleNames ...string) bool {
	return egn.HasAnyRole(roleName, hasRoleNames...)
}

func GrantRole(grantee string, granted ...string) error {
	return egn.GrantRole(grantee, granted...)
}
// 撤销权限
func RevokeRole(revokee string, revoked ...string) error {
	return egn.RevokeRole(revokee, revoked...)
}

func GrantPermission(roleName string, res string, perm ...string) error {
	return egn.GrantPermission(roleName, res, perm...)
}

func RevokePermission(roleName, res string, perm ...string) error {
	return egn.RevokePermission(roleName, res, perm...)
}

func GrantGlobalPermission(roleName string, perm ...string) error {
	return egn.GrantPermission(roleName, "", perm...)
}

func RevokeGLobalPermission(roleName string, perm ...string) error {
	return egn.RevokePermission(roleName, "", perm...)
}

func Decision(roleName string, res string, perm ...string) bool {
	return egn.Decision(roleName, res, perm...)
}

func DecisionEx(roleName string, res string, perm ...string) bool {
	return egn.DecisionEx(roleName, res, perm...)
}