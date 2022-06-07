package entity

type Label struct {
	ID_label    uint64 `gorm:"primary_key:auto_incremnet" json:"id_label"`
	Name_label  string `gorm:"type:varchar(255)" json:"name_label"`
	Color_label string `gorm:"type:varchar(255)" json:"color_label"`
	UserID      uint64 `gorm:"not null" json:"-"`
	User        User   `gorm:"foreignkey:UserID;constraint:onUpdate:CASCADE,onDelete:CASCADE" json:"-"`
}

type Label_get struct {
	ID_label    uint64 `gorm:"primary_key:auto_incremnet" json:"id_label"`
	Name_label  string `gorm:"type:varchar(255)" json:"name_label"`
	Color_label string `gorm:"type:varchar(255)" json:"color_label"`
	NumTask     uint64 `json:"num_task"`
}
