package models

type ConnectionString struct {
	User     string
	DBName   string
	Password string
	SslMode  string
	Port     string
}

type Users struct {
	RollNumber      uint            `json:"roll_number" gorm:"primaryKey"`
	Name            string          `json:"name"`
	Dob             string          `json:"dob"`
	Address         string          `json:"address"`
	UserRoleMapping UserRoleMapping `gorm:"foreignKey:RollNumber;references:RollNumber;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}

type UserRoles struct {
	RoleID          uint            `json:"role_id" gorm:"primaryKey"`
	Role            string          `json:"role"`
	Permission      string          `json:"permission"`
	UserRoleMapping UserRoleMapping `json:"user_role_mapping" gorm:"foreignKey:RoleID;references:RoleID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}

type UserRoleMapping struct {
	ID         uint `json:"id" gorm:"primaryKey"`
	RollNumber uint `json:"roll_number"`
	RoleID     uint `json:"role_id"`
}

type Error struct {
	Message    string
	Error      error
	StatusCode int
}
