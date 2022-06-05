package service

import (
	"fmt"
	"math/rand"
	"os"
	"time"

	"github.com/golang-jwt/jwt"
)

type JWTService interface {
	GenerateToken(UserID string) (string, string)
	ValidateToken(token string) (*jwt.Token, error)
	ValidateRefreshToken(token string) (*jwt.Token, error)
	ValidatePlayload(token jwt.Token, tokenRefersh jwt.Token) (bool, string)
}

type jwtCustomClaim struct {
	UserID   string `json:"user_id"`
	RandUUID string `json:"uuid"`
	jwt.StandardClaims
}

type jwtCustomClaimRefresh struct {
	UserID   string `json:"user_id"`
	RandUUID string `json:"uuid"`
	jwt.StandardClaims
}

type jwtService struct {
	secretKey       string
	secretKeyRefesh string
	issuer          string
}

func NewJWTService() JWTService {
	return &jwtService{
		issuer:          "jdasdad",
		secretKey:       getSecretKey(),
		secretKeyRefesh: getSecretKeyRefresh(),
	}
}

func getSecretKey() string {
	secretKey := os.Getenv("JWT_SECRET")
	if secretKey == "" {
		panic("Set secret key in env")
	}
	return secretKey
}

func getSecretKeyRefresh() string {
	secretKey := os.Getenv("JWT_SECRET_REFRESH")
	if secretKey == "" {
		panic("Set secret key in env")
	}
	return secretKey
}

func (j *jwtService) GenerateToken(UserID string) (string, string) {
	UUID := RandStringBytesRmndr(10)
	claims := &jwtCustomClaim{
		UserID,
		UUID,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Minute * 15).Unix(),
			Issuer:    j.issuer,
			IssuedAt:  time.Now().Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	t, err := token.SignedString([]byte(j.secretKey))
	if err != nil {
		panic(err.Error())
	}

	claimsRt := &jwtCustomClaimRefresh{
		UserID,
		UUID,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 300).Unix(),
			Issuer:    j.issuer,
			IssuedAt:  time.Now().Unix(),
		},
	}

	tokenRt := jwt.NewWithClaims(jwt.SigningMethodHS256, claimsRt)
	rt, err := tokenRt.SignedString([]byte(j.secretKeyRefesh))
	if err != nil {
		panic(err.Error())
	}
	return t, rt
}

func RandStringBytesRmndr(n int) string {
	letterBytes := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

	b := make([]byte, n)
	for i := range b {
		b[i] = letterBytes[rand.Int63()%int64(len(letterBytes))]
	}
	return string(b)
}

func (j *jwtService) ValidateToken(token string) (*jwt.Token, error) {
	return jwt.Parse(token, func(t_ *jwt.Token) (interface{}, error) {
		if _, ok := t_.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method %s", t_.Header["alg"])
		}
		return []byte(j.secretKey), nil
	})
}

func (j *jwtService) ValidateRefreshToken(token string) (*jwt.Token, error) {
	return jwt.Parse(token, func(t_ *jwt.Token) (interface{}, error) {
		if _, ok := t_.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method %s", t_.Header["alg"])
		}
		return []byte(j.secretKeyRefesh), nil
	})
}

func (j *jwtService) ValidatePlayload(token jwt.Token, tokenRefersh jwt.Token) (bool, string) {
	claims, ok := token.Claims.(*jwtCustomClaim)
	claimsRt, okRt := tokenRefersh.Claims.(*jwtCustomClaimRefresh)
	if !ok || !okRt {
		return false, ""
	}

	if claims.UserID != claimsRt.UserID || claims.RandUUID != claimsRt.RandUUID {
		return false, ""
	}
	return true, claims.UserID
}
