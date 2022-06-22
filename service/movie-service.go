package service

import (
	"log"

	"github.com/mashingan/smapping"
	"github.com/zerodev/golang_api/dto"
	"github.com/zerodev/golang_api/entity"
	"github.com/zerodev/golang_api/repository"
)

type MovieService interface {
	GetMovie(userID uint64) ([]entity.MovieUserGet, error)
	GetMovieByID(userID uint64, movieID uint64) (entity.MovieUserGet, error)
	CreateMovie(movie dto.MovieCreateDTO) (entity.MovieUser, error)
	DeleteMovie(userID uint64, movieID uint64) error
	WatchMovie(userID uint64, movieID uint64) error
}

type movieService struct {
	movieRepository repository.MovieRepository
}

func NewMovieService(movieRepo repository.MovieRepository) MovieService {
	return &movieService{
		movieRepository: movieRepo,
	}
}

func (service *movieService) GetMovie(userID uint64) ([]entity.MovieUserGet, error) {
	res, err := service.movieRepository.GetMovie(userID)
	return res, err
}

func (service *movieService) GetMovieByID(userID uint64, movieID uint64) (entity.MovieUserGet, error) {
	res, err := service.movieRepository.GetMovieByID(userID, movieID)
	return res, err
}

func (service *movieService) WatchMovie(userID uint64, movieID uint64) error {
	err := service.movieRepository.WatchMovie(userID, movieID)
	return err
}

func (service *movieService) CreateMovie(movie dto.MovieCreateDTO) (entity.MovieUser, error) {
	movieToCreate := entity.MovieUser{}
	err := smapping.FillStruct(&movieToCreate, smapping.MapFields((&movie)))
	if err != nil {
		log.Fatalf("Failed mapping %v", err.Error())
	}

	res, err := service.movieRepository.InsertMovie(movieToCreate)

	return res, err
}

func (service *movieService) DeleteMovie(userID uint64, movieID uint64) error {
	err := service.movieRepository.DeleteMovie(userID, movieID)
	return err
}
