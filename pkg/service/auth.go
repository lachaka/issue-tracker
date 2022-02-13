package service

import (
	"fmt"
	"issue-tracker/cmd/utils"
	"issue-tracker/pkg/models"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/spf13/viper"
)

type JWTService interface {
	GenerateToken(user models.User) (string, error)
	ValidateToken(token string) (*jwt.Token, error)
}

type authCustomClaims struct {
	Id   string `json:"id"`
	Role string `json:"role"`
	jwt.StandardClaims
}

type jwtServices struct {
	secretKey string
}

func JWTAuthService() JWTService {
	return &jwtServices{
		secretKey: getSecretKey("."),
	}
}

func getSecretKey(path string) string {
	viper.AddConfigPath(path)
	viper.SetConfigName(".env")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	err := viper.ReadInConfig()
	if err != nil {
		utils.Logger.ErrorLog.Fatal(err)
	}

	var secret utils.JWTSecret
	err = viper.Unmarshal(&secret)
	if err != nil {
		utils.Logger.ErrorLog.Fatal(err)
	}

	return secret.Secret
}

func (service *jwtServices) GenerateToken(user models.User) (string, error) {
	secret := service.secretKey // getSecretKey(".")
	claims := &authCustomClaims{
		user.Id.String(),
		user.Role,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 48).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
	}

	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString([]byte(secret))

	if err != nil {
		return "", err
	}

	return token, nil
}

func (service *jwtServices) ValidateToken(encodedToken string) (*jwt.Token, error) {
	return jwt.Parse(encodedToken, func(token *jwt.Token) (interface{}, error) {
		if _, isvalid := token.Method.(*jwt.SigningMethodHMAC); !isvalid {
			return nil, fmt.Errorf("Invalid token %s", token.Header["alg"])
		}
		return []byte(service.secretKey), nil
	})

}
