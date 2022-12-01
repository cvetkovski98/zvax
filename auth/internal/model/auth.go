package model

import (
	"time"

	"github.com/uptrace/bun"
)

type User struct {
	bun.BaseModel `bun:"table:users"`

	Id       int     `bun:"id,pk,nullzero"`
	Email    string  `bun:"email,unique,nullzero,notnull"`
	Name     string  `bun:"name,nullzero,notnull"`
	Password string  `bun:"password,nullzero,notnull"`
	Roles    []*Role `bun:"m2m:user_roles,join:User=Role"`

	CreatedAt time.Time `bun:"created_at,nullzero,notnull,default:current_timestamp"`
	UpdatedAt time.Time `bun:"updated_at,nullzero,notnull,default:current_timestamp"`
}

type Role struct {
	bun.BaseModel `bun:"roles"`

	Id          int           `bun:"id,pk"`
	Name        string        `bun:"name,unique,notnull,nullzero"`
	Permissions []*Permission `bun:"m2m:role_permissions,join:Role=Permission"`

	CreatedAt time.Time `bun:"created_at,nullzero,notnull,default:current_timestamp"`
	UpdatedAt time.Time `bun:"updated_at,nullzero,notnull,default:current_timestamp"`
}

type Permission struct {
	bun.BaseModel `bun:"permissions"`

	Id   int    `bun:"id,pk"`
	Name string `bun:"name,unique,notnull,nullzero"`

	CreatedAt time.Time `bun:"created_at,nullzero,notnull,default:current_timestamp"`
	UpdatedAt time.Time `bun:"updated_at,nullzero,notnull,default:current_timestamp"`
}

type UserRole struct {
	bun.BaseModel `bun:"user_roles"`

	UserId int   `bun:"user_id,pk"`
	RoleId int   `bun:"role_id,pk"`
	User   *User `bun:"rel:belongs-to,join:user_id=id"`
	Role   *Role `bun:"rel:belongs-to,join:role_id=id"`

	CreatedAt time.Time `bun:"created_at,nullzero,notnull,default:current_timestamp"`
	UpdatedAt time.Time `bun:"updated_at,nullzero,notnull,default:current_timestamp"`
}

type RolePermission struct {
	bun.BaseModel `bun:"role_permissions"`

	RoleId       int         `bun:"role_id,pk"`
	PermissionId int         `bun:"permission_id,pk"`
	Role         *Role       `bun:"rel:belongs-to,join:role_id=id"`
	Permission   *Permission `bun:"rel:belongs-to,join:permission_id=id"`

	CreatedAt time.Time `bun:"created_at,nullzero,notnull,default:current_timestamp"`
	UpdatedAt time.Time `bun:"updated_at,nullzero,notnull,default:current_timestamp"`
}
