package entity

type MovieUser struct {
	ID_movie_user  uint64 `gorm:"primary_key:auto_incremnet" json:"id_movie_user"`
	Movie_title    string `gorm:"type:varchar(255)" json:"movie_title"`
	Movie_overview string `gorm:"type:text" json:"movie_overview"`
	Movie_image    string `gorm:"type:varchar(255)" json:"movie_image"`
	Release_date   string `gorm:"type:varchar(255)" json:"release_date"`
	Status         uint64 `json:"status"`
	UserID         uint64 `gorm:"not null" json:"-"`
	User           User   `gorm:"foreignkey:UserID;constraint:onUpdate:CASCADE,onDelete:CASCADE" json:"-"`
}

type MovieUserGet struct {
	ID_movie_user  uint64 `gorm:"primary_key:auto_incremnet" json:"id_movie_user"`
	Movie_title    string `gorm:"type:varchar(255)" json:"movie_title"`
	Movie_overview string `gorm:"type:text" json:"movie_overview"`
	Movie_image    string `gorm:"type:varchar(255)" json:"movie_image"`
	Release_date   string `gorm:"type:varchar(255)" json:"release_date"`
	Status         uint64 `json:"status"`
}
