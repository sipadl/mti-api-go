package models

type Status struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	ParentId string `json:"parent_id"`
	Desc     string `json:"description"`
}

func (Status) TableName() string {
	return "MAIN_STATUS"
}
