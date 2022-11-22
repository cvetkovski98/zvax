package dto

type User struct {
	Email string  `json:"email"`
	Name  string  `json:"name,omitempty"`
	Roles []*Role `json:"roles"`

	Timestamped `json:"timestamped"`
}

type Role struct {
	Name        string        `json:"name"`
	Permissions []*Permission `json:"permissions"`

	Timestamped `json:"timestamped"`
}

type Permission struct {
	Name string `json:"name"`

	Timestamped `json:"timestamped"`
}
