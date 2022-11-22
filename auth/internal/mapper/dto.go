package mapper

import (
	"github.com/cvetkovski98/zvax-auth/internal/dto"
	"github.com/cvetkovski98/zvax-auth/internal/model"
)

func PermissionModelToDto(from *model.Permission) *dto.Permission {
	var permission = &dto.Permission{
		Name: from.Name,
	}
	permission.CreatedAt = from.CreatedAt
	permission.UpdatedAt = from.UpdatedAt
	return permission
}

func PermissionModelsToDtos(from []*model.Permission) []*dto.Permission {
	var permissions []*dto.Permission
	for _, permission := range from {
		permissions = append(permissions, PermissionModelToDto(permission))
	}
	return permissions
}

func RoleModelToDto(from *model.Role) *dto.Role {
	var role = &dto.Role{
		Name:        from.Name,
		Permissions: PermissionModelsToDtos(from.Permissions),
	}
	role.CreatedAt = from.CreatedAt
	role.UpdatedAt = from.UpdatedAt
	return role
}

func RoleModelsToDtos(from []*model.Role) []*dto.Role {
	var roles []*dto.Role
	for _, role := range from {
		roles = append(roles, RoleModelToDto(role))
	}
	return roles
}

func UserModelToDto(from *model.User) *dto.User {
	var user = &dto.User{
		Email: from.Email,
		Name:  from.Name,
		Roles: RoleModelsToDtos(from.Roles),
	}
	user.CreatedAt = from.CreatedAt
	user.UpdatedAt = from.UpdatedAt
	return user
}
