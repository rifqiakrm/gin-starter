// Package cache provides caching utilities.
package cache

const (
	// FiveMinutes is ttl for the cache
	FiveMinutes = 300
	// OneHour is ttl for the cache
	OneHour = 3600
	// OneMonth is ttl for the cache
	OneMonth = 2592000
)

const (
	prefix = "gin-starter"
	// AuthToken is a redis key for find auth token of the user
	AuthToken = prefix + ":auth:token:%v"
	// UserRoleByUserID is a redis key for find cms auth role by auth id.
	UserRoleByUserID = prefix + ":auth-role:find-by-auth-id:%v"
	// UserPermissionByUserID is a redis key for storing all user permission by user ID
	UserPermissionByUserID = prefix + ":auth:permission:%v"
	// PermissionFindByName is a redis key for find cms permission by name.
	PermissionFindByName = prefix + ":permission:find-by-name:%v"
	// RolePermissionFindByRoleIDAndPermissionID is a redis key for find cms role permission by role id and permission id.
	RolePermissionFindByRoleIDAndPermissionID = prefix + ":role-permission:find-by-role-id-and-permission-id:%v:%v"
	// RoleFindByID is a redis key for find cms role by id.
	RoleFindByID = prefix + ":role:find-by-id:%v"
)
