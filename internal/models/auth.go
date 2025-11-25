package models

import "time"

// User represents a system user
type User struct {
	BaseModel

	Username     string `gorm:"size:100;uniqueIndex;not null" json:"username"`
	Email        string `gorm:"size:255;uniqueIndex;not null" json:"email"`
	PasswordHash string `gorm:"size:255;not null" json:"-"` // Never expose in JSON
	FirstName    string `gorm:"size:100" json:"first_name"`
	LastName     string `gorm:"size:100" json:"last_name"`

	Active   bool `gorm:"default:true" json:"active"`
	Verified bool `gorm:"default:false" json:"verified"`

	LastLoginAt *time.Time `json:"last_login_at,omitempty"`

	Roles []Role `gorm:"many2many:user_roles;" json:"roles,omitempty"`
}

// Role represents a user role
type Role struct {
	BaseModel

	Name        string `gorm:"size:100;uniqueIndex;not null" json:"name"`
	Description string `gorm:"type:text" json:"description"`

	Permissions []Permission `gorm:"many2many:role_permissions;" json:"permissions,omitempty"`
}

// Permission represents a system permission
type Permission struct {
	BaseModel

	Name        string `gorm:"size:100;uniqueIndex;not null" json:"name"`
	Resource    string `gorm:"size:100;not null" json:"resource"` // e.g., "patient", "employee"
	Action      string `gorm:"size:50;not null" json:"action"`    // e.g., "read", "write", "delete"
	Description string `gorm:"type:text" json:"description"`
}

// Session represents an active user session
type Session struct {
	BaseModel

	UserID    uint      `gorm:"index;not null" json:"user_id"`
	Token     string    `gorm:"size:500;uniqueIndex;not null" json:"token"`
	ExpiresAt time.Time `gorm:"index;not null" json:"expires_at"`
	IPAddress string    `gorm:"size:50" json:"ip_address"`
	UserAgent string    `gorm:"size:500" json:"user_agent"`
}

// TableName overrides
func (User) TableName() string {
	return "users"
}

func (Role) TableName() string {
	return "roles"
}

func (Permission) TableName() string {
	return "permissions"
}

func (Session) TableName() string {
	return "sessions"
}
