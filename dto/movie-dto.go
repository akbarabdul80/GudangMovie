package dto

type MovieIDDTO struct {
	ID_movie_user uint64 `json:"id_movie_user" form:"id_movie_user" binding:"required"`
	UserID        uint64 `json:"user_id,omitempty" form:"user_id,omitempty"`
}

type MovieCreateDTO struct {
	Movie_title    string `json:"movie_title" form:"movie_title" binding:"required"`
	Movie_overview string `json:"movie_overview" form:"movie_overview" binding:"required"`
	Movie_image    string `json:"movie_image" form:"movie_image" binding:"required"`
	Release_date   string `json:"release_date" form:"release_date" binding:"required"`
	Status         uint64 `json:"status" form:"status" binding:"required"`
	UserID         uint64 `json:"user_id,omitempty" form:"user_id,omitempty"`
}
