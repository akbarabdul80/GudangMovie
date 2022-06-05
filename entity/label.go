package entity

type Label struct {
	Id_label    uint64 `gorm:"primary_key:auto_incremnet" json:"id_label"`
	Name_label  string `gorm:"type:varchar(255)" json:"name_label"`
	Color_label string `gorm:"uniqueIndex;type:varchar(255)" json:"color_label"`
}
