package models

import (
	"time"
)

type User struct {
	ID               uint      `gorm:"primaryKey" json:"id"`
	Username         string    `json:"username"`
	Password         string    `json:"password"`
	PasswordRecovery string    `json:"password_recovery"`
	Email            string    `json:"email"`
	StatusId         string    `json:"status_id"`
	RoleId           string    `json:"role_id"`
	CompanyId        string    `json:"company_id"`
	ProfileId        string    `json:"profile_id"`
	Role             Role      `gorm:"foreignKey:RoleId" json:"role"`
	Company          Company   `gorm:"foreignKey:CompanyId" json:"company"`
	Profile          Profile   `gorm:"foreignKey:ProfileId" json:"profile"`
	CreatedAt        time.Time `json:"created_at"`
	UpdateAt         time.Time `json:"updated_at"`
}

func (User) TableName() string {
	return "MAIN_USER"
}
