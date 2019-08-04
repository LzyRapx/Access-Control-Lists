package pkg

type RbacInterface interface {
	GrantRole(grante string, granted ...string) error
	RevokeRole(revoke string, revoked ...string) error
	GrantPermission(roleName string, res string, perms ...string) error
	RevokePermission(roleName string, res string, perms ...string) error
	GetRole(roleName string, create bool) (id int, rbacType int, exist bool)
	DropRole(roleName string) error
	GetPermission(permName, resString string, create bool) (id int, exist bool)
	DropPermission(permName, resString string) error
	SetDesc(id int, desc string) (exist bool)
	GetDesc(id int) string
	SetRoleType(roleName string, rbacType int) error
	HasAllRole(roleName string, hasRoleNames ...string) bool
	HasAnyRole(roleName string, hasRoleNames ...string) bool
	Decision(roleName string, resource string, permission ...string) bool
	DecisionEx(roleName string, resource string, permission ...string) bool
}