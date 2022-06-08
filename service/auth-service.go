package service

import (
	"log"

	"github.com/mashingan/smapping"
	"github.com/zerodev/golang_api/dto"
	"github.com/zerodev/golang_api/entity"
	"github.com/zerodev/golang_api/repository"
	"golang.org/x/crypto/bcrypt"
)

type AuthService interface {
	VerifyCredential(email string, password string) interface{}
	CreateUser(user dto.UserCreateDTO) entity.User
	RegisterUser(user dto.RegisterDTO) entity.User
	FindByEmail(email string) entity.User
	FindByID(userID string) interface{}
	IsDuplicateEmail(email string) bool
}

type authService struct {
	userRepository repository.UserRepository
}

func NewAuthService(userRep repository.UserRepository) AuthService {
	return &authService{
		userRepository: userRep,
	}
}

func (service *authService) VerifyCredential(email string, password string) interface{} {

	res := service.userRepository.VerifyCredential(email, password)
	if v, ok := res.(entity.User); ok {
		comparedPassword := comparePassword(v.Password, []byte(password))
		if v.Email == email && comparedPassword {
			return res
		}
		return false
	}
	return false
}

func (service *authService) CreateUser(user dto.UserCreateDTO) entity.User {
	userToCreate := entity.User{}
	err := smapping.FillStruct(&userToCreate, smapping.MapFields((&user)))
	if err != nil {
		log.Fatalf("Failed mapping %v", err.Error())
	}

	res := service.userRepository.InsertUser(userToCreate)

	return res
}

func (service *authService) RegisterUser(user dto.RegisterDTO) entity.User {
	userToCreate := entity.User{}
	err := smapping.FillStruct(&userToCreate, smapping.MapFields((&user)))
	if err != nil {
		log.Fatalf("Failed mapping %v", err.Error())
	}

	res := service.userRepository.InsertUser(userToCreate)

	return res
}

func (service *authService) FindByEmail(email string) entity.User {
	return service.userRepository.FindByEmail(email)
}

func (service *authService) IsDuplicateEmail(email string) bool {
	res := service.userRepository.IsDuplicateEmail(email)
	return !!(res.Error == nil)
}

func (service *authService) FindByID(userID string) interface{} {
	return service.userRepository.ProfileUser(userID)
}

func comparePassword(hashedPwd string, plainPassword []byte) bool {
	byteHash := []byte(hashedPwd)
	err := bcrypt.CompareHashAndPassword(byteHash, plainPassword)

	if err != nil {
		log.Println(err.Error())
		return false
	}
	return true
}
