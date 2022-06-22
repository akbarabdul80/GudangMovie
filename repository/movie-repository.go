package repository

import (
	"github.com/zerodev/golang_api/entity"
	"gorm.io/gorm"
)

type MovieRepository interface {
	GetMovie(userID uint64) ([]entity.MovieUserGet, error)
	GetMovieByID(userID uint64, movieID uint64) (entity.MovieUserGet, error)
	InsertMovie(movie entity.MovieUser) (entity.MovieUser, error)
	DeleteMovie(userID uint64, movieID uint64) error
	WatchMovie(userID uint64, movieID uint64) error
}

type movieConnection struct {
	connection *gorm.DB
}

func NewMovieRepository(db *gorm.DB) MovieRepository {
	return &movieConnection{
		connection: db,
	}
}

func (db *movieConnection) GetMovie(userID uint64) ([]entity.MovieUserGet, error) {
	var movie []entity.MovieUserGet
	err := db.connection.Model(&entity.MovieUser{}).Where("user_id = ?", userID).Find(&movie)
	return movie, err.Error
}

func (db *movieConnection) WatchMovie(userID uint64, movieID uint64) error {
	err := db.connection.Model(&entity.MovieUser{}).Where(&entity.MovieUser{UserID: userID, ID_movie_user: movieID}).Update("status", 2)
	return err.Error
}

func (db *movieConnection) GetMovieByID(userID uint64, movieID uint64) (entity.MovieUserGet, error) {
	var movie entity.MovieUserGet
	err := db.connection.Model(&entity.MovieUser{}).Where(&entity.MovieUser{UserID: userID, ID_movie_user: movieID}).First(&movie)
	return movie, err.Error
}

func (db *movieConnection) InsertMovie(movie entity.MovieUser) (entity.MovieUser, error) {
	err := db.connection.Save(&movie)
	db.connection.Preload("User").Find(&movie)
	return movie, err.Error
}

func (db *movieConnection) UpdateMovie(movie entity.MovieUser) (entity.MovieUser, error) {
	err := db.connection.Save(&movie)
	return movie, err.Error
}

func (db *movieConnection) DeleteMovie(userID uint64, movieID uint64) error {
	err := db.connection.Delete(&entity.MovieUser{UserID: userID, ID_movie_user: movieID})
	return err.Error
}
