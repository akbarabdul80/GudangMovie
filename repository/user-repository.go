package repository

import (
	"fmt"
	"log"

	"github.com/zerodev/golang_api/entity"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type UserRepository interface {
	InsertUser(user entity.User) entity.User
	UpdateUser(user entity.User) (entity.User, error)
	VerifyCredential(email string, password string) interface{}
	IsDuplicateEmail(email string) (ctx *gorm.DB)
	FindByEmail(email string) entity.User
	ProfileUser(UserID string) entity.User
}

type userConnection struct {
	connection *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userConnection{
		connection: db,
	}
}

func hashAndSalt(pwd []byte) string {
	hash, err := bcrypt.GenerateFromPassword(pwd, bcrypt.DefaultCost)
	if err != nil {
		log.Println(err.Error())
		panic("Failed to hash a password ")
	}

	return string(hash)
}

func (db *userConnection) InsertUser(user entity.User) entity.User {
	user.Password = hashAndSalt([]byte(user.Password))
	fmt.Println(user)
	db.connection.Save(&user)
	return user
}

func (db *userConnection) UpdateUser(user entity.User) (entity.User, error) {
	if user.Password != "" {
		user.Password = hashAndSalt([]byte(user.Password))
	} else {
		var tmpUser entity.User
		db.connection.Find(&tmpUser, user.ID)
		user.Password = tmpUser.Password
	}
	err := db.connection.Save(&user)
	return user, err.Error
}

func (db *userConnection) VerifyCredential(email string, password string) interface{} {
	var user entity.User
	res := db.connection.Where("email = ?", email).Take(&user)
	if res.Error == nil {
		return user
	}
	return nil
}

func (db *userConnection) IsDuplicateEmail(email string) (ctx *gorm.DB) {
	var user entity.User
	return db.connection.Where("email = ?", email).Take(&user)
}

func (db *userConnection) FindByEmail(email string) entity.User {
	var user entity.User
	db.connection.Where("email = ?", email).Take(&user)
	return user
}

func (db *userConnection) ProfileUser(userID string) entity.User {
	var user entity.User
	db.connection.Find(&user, userID)
	return user
}
