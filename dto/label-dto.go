package dto

type LabelUpdateDTO struct {
	ID_label    uint64 `json:"id_label" form:"id_label" binding:"required"`
	Name_label  string `json:"name_label" form:"name_label" binding:"required"`
	Color_label string `json:"color_label" form:"color_label" binding:"required"`
	UserID      uint64 `json:"user_id,omitempty" form:"user_id,omitempty"`
}

type LabelCreateDTO struct {
	Name_label  string `json:"name_label" form:"name_label" binding:"required"`
	Color_label string `json:"color_label" form:"color_label" binding:"required"`
	UserID      uint64 `json:"user_id,omitempty" form:"user_id,omitempty"`
}
