package models

type Role struct {
	ID       string `gorm:"primaryKey" json:"ID"`
	Nama     string `json:"nama"`
	StatusId string `json:"status_id"`
	Status   Status `gorm:"foreignKey:StatusId" json:"status"`
}

func (Role) TableName() string {
	return "MAIN_ROLE"
}
