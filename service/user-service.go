package service

import (
	"log"

	"github.com/mashingan/smapping"
	"github.com/zerodev/golang_api/dto"
	"github.com/zerodev/golang_api/entity"
	"github.com/zerodev/golang_api/repository"
)

type UserService interface {
	Update(user dto.UserUpdateDTO) (entity.User, error)
	Profile(userID string) entity.User
}

type userService struct {
	userRepository repository.UserRepository
}

func NewUserService(userRepo repository.UserRepository) UserService {
	return &userService{
		userRepository: userRepo,
	}
}

func (service *userService) Update(user dto.UserUpdateDTO) (entity.User, error) {
	userToUpdate := entity.User{}
	err := smapping.FillStruct(&userToUpdate, smapping.MapFields((&user)))
	if err != nil {
		log.Fatalf("Failed mapping %v", err.Error())
	}
	res, err := service.userRepository.UpdateUser(userToUpdate)
	return res, err
}

func (service *userService) Profile(userID string) entity.User {
	res := service.userRepository.ProfileUser(userID)
	return res
}
