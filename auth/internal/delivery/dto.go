package delivery

import (
	"github.com/cvetkovski98/zvax-auth/internal/dto"
	"github.com/cvetkovski98/zvax-common/gen/pbauth"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func PermissionDtoToResponse(dto *dto.Permission) *pbauth.PermissionResponse {
	return &pbauth.PermissionResponse{
		Name:      dto.Name,
		CreatedAt: timestamppb.New(dto.CreatedAt),
		UpdatedAt: timestamppb.New(dto.UpdatedAt),
	}
}

func PermissionDtosToResponses(dtos []*dto.Permission) []*pbauth.PermissionResponse {
	var responses []*pbauth.PermissionResponse
	for _, dto := range dtos {
		responses = append(responses, PermissionDtoToResponse(dto))
	}
	return responses
}

func RoleDtoToResponse(dto *dto.Role) *pbauth.RoleResponse {
	return &pbauth.RoleResponse{
		Name:        dto.Name,
		Permissions: PermissionDtosToResponses(dto.Permissions),
		CreatedAt:   timestamppb.New(dto.CreatedAt),
		UpdatedAt:   timestamppb.New(dto.UpdatedAt),
	}
}

func RoleDtosToResponses(dtos []*dto.Role) []*pbauth.RoleResponse {
	var responses []*pbauth.RoleResponse
	for _, dto := range dtos {
		responses = append(responses, RoleDtoToResponse(dto))
	}
	return responses
}

func UserDtoToResponse(dto *dto.User) *pbauth.AuthResponse {
	return &pbauth.AuthResponse{
		Email:     dto.Email,
		Name:      dto.Name,
		Roles:     RoleDtosToResponses(dto.Roles),
		CreatedAt: timestamppb.New(dto.CreatedAt),
		UpdatedAt: timestamppb.New(dto.UpdatedAt),
	}
}

func LoginRequestToDto(request *pbauth.LoginRequest) *dto.Login {
	return &dto.Login{
		Email:    request.Email,
		Password: request.Password,
	}
}

func RegisterRequestToDto(request *pbauth.RegisterRequest) *dto.Register {
	return &dto.Register{
		Email:           request.Email,
		Name:            request.Name,
		Password:        request.Password,
		ConfirmPassword: request.ConfirmPassword,
	}
}
