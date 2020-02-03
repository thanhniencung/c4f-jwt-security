package model

type User struct {
	UserId    string    `json:"-" db:"user_id, omitempty"`
	FullName  string    `json:"fullName,omitempty" db:"full_name, omitempty"`
	Role      string    `json:"role,omitempty" db:"role, omitempty"`
	Token     string    `json:"token,omitempty"`
}