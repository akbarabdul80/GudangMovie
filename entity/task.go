package entity

import "time"

type Task struct {
	Id_task   uint64     `gorm:"primary_key:auto_incremnet" json:"id_task"`
	Name_task string     `gorm:"type:varchar(255)" json:"name_task"`
	Datetime  *time.Time `gorm:"uniqueIndex;type:varchar(255)" json:"datetime"`
	Note      string     `gorm:"type:varchar(255)" json:"note"`
	Done      *bool      `gorm:"default:false" json:"done"`
}
