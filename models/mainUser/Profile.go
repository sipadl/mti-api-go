package models

import "time"

type Profile struct {
	Id          string    `json:"id" gorm:"primaryKey"`
	Name        string    `json:"name"`
	FullName    string    `json:"full_name"`
	FirstName   string    `json:"first_name"`
	LastName    string    `json:"last_name"`
	Email       string    `json:"email"`
	Phone       string    `json:"phone"`
	Address     string    `json:"address"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

func (Profile) TableName() string {
	return "MAIN_PROFILE"
}
