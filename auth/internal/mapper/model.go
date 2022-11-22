package mapper

import (
	"github.com/cvetkovski98/zvax-auth/internal/dto"
	"github.com/cvetkovski98/zvax-auth/internal/model"
)

func PermissionDtoToModel(from *dto.Permission) *model.Permission {
	return &model.Permission{
		Name:      from.Name,
		CreatedAt: from.CreatedAt,
		UpdatedAt: from.UpdatedAt,
	}
}

func PermissionDtosToModels(from []*dto.Permission) []*model.Permission {
	var permissions []*model.Permission
	for _, permission := range from {
		permissions = append(permissions, PermissionDtoToModel(permission))
	}
	return permissions
}

func RoleDtoToModel(from *dto.Role) *model.Role {
	return &model.Role{
		Name:        from.Name,
		Permissions: PermissionDtosToModels(from.Permissions),
		CreatedAt:   from.CreatedAt,
		UpdatedAt:   from.UpdatedAt,
	}
}

func RoleDtosToModels(from []*dto.Role) []*model.Role {
	var roles []*model.Role
	for _, role := range from {
		roles = append(roles, RoleDtoToModel(role))
	}
	return roles
}

func UserDtoToModel(from *dto.User) *model.User {
	return &model.User{
		Email:     from.Email,
		Roles:     RoleDtosToModels(from.Roles),
		CreatedAt: from.CreatedAt,
		UpdatedAt: from.UpdatedAt,
	}
}
