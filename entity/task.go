package entity

import "time"

type Task struct {
	ID_task   uint64    `gorm:"primary_key:auto_incremnet" json:"id_task"`
	Name_task string    `gorm:"type:varchar(255)" json:"name_task"`
	Datetime  time.Time `json:"datetime"`
	Note      string    `gorm:"type:varchar(255)" json:"note"`
	Done      *bool     `gorm:"default:false" json:"done"`
	UserID    uint64    `gorm:"not null" json:"-"`
	User      User      `gorm:"foreignkey:UserID;constraint:onUpdate:CASCADE,onDelete:CASCADE" json:"-"`
	LabelID   uint64    `gorm:"not null" json:"-"`
	Label     Label     `gorm:"foreignkey:LabelID;constraint:onUpdate:CASCADE,onDelete:CASCADE" json:"label"`
}
