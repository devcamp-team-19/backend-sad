package module

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"

	"github.com/devcamp-team-19/backend-sad/core/entity"
	"github.com/devcamp-team-19/backend-sad/core/repository"
)

type UserUsecase interface {
	GetUsers(c *gin.Context) ([]entity.User, error)
	GetUser(c *gin.Context) (entity.User, error)
	CreateUser(c *gin.Context, user entity.User) error
	UpdateUser(c *gin.Context, user entity.User) error
	DeleteUser(c *gin.Context) error
}

type userUsecase struct {
	userRepo repository.UserRepository
}

// NewUserUsecase use for initiate new user usecase
func NewUserUsecase(repo repository.UserRepository) UserUsecase {
	return &userUsecase{repo}
}

var (
	secretkey       = "asdfasdfasdf"
	ErrUserNotFound = errors.New("user error: ")
)

type Error struct {
	IsError bool   `json:"isError"`
	Message string `json:"message"`
}

func GenerateJWT(email string) (string, error) {
	var mySigningKey = []byte(secretkey)
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)

	claims["authorized"] = true
	claims["email"] = email
	claims["exp"] = time.Now().Add(time.Minute * 30).Unix()

	tokenString, err := token.SignedString(mySigningKey)

	if err != nil {
		fmt.Errorf("Something Went Wrong: %s", err.Error())
		return "", err
	}
	return tokenString, nil
}

// GeneratehashPassword take password as input and generate new hash password from it
func GeneratehashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

// CheckPasswordHash compare plain password with hash password
func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

// SetError set error message in Error struct
func SetError(err Error, message string) Error {
	err.IsError = true
	err.Message = message
	return err
}

// IsAuthorized check whether user is authorized or not
func IsAuthorized(handler http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		if r.Header["Token"] == nil {
			var err Error
			err = SetError(err, "No Token Found")
			json.NewEncoder(w).Encode(err)
			return
		}

		var mySigningKey = []byte(secretkey)

		token, err := jwt.Parse(r.Header["Token"][0], func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("There was an error in parsing token.")
			}
			return mySigningKey, nil
		})

		if err != nil {
			var err Error
			err = SetError(err, "Your Token has been expired.")
			json.NewEncoder(w).Encode(err)
			return
		}

		if _, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			r.Header.Set("Role", "user")
			handler.ServeHTTP(w, r)
			return
		}

		var reserr Error
		reserr = SetError(reserr, "Not Authorized.")
		json.NewEncoder(w).Encode(err)
	}
}

func (em *userUsecase) GetUsers(c *gin.Context) ([]entity.User, error) {
	data, err := em.userRepo.FindAll(c)
	if err != nil {
		if errors.Is(err, repository.ErrRecordUserNotFound) {
			return nil, fmt.Errorf("%w.", ErrUserNotFound)
		}
		return nil, fmt.Errorf("%w: %v", ErrUserNotFound, err)
	}
	return data, nil
}

func (em *userUsecase) GetUser(c *gin.Context) (entity.User, error) {
	data, err := em.userRepo.FindSingle(c)
	if err != nil {
		if errors.Is(err, repository.ErrRecordUserNotFound) {
			return entity.User{}, fmt.Errorf("%w.", ErrUserNotFound)
		}
		return entity.User{}, fmt.Errorf("%w: %v", ErrUserNotFound, err)
	}
	return data, nil
}

func (em *userUsecase) CreateUser(c *gin.Context, user entity.User) error {
	hashedPass, err := GeneratehashPassword(user.Password)
	if err != nil {
		log.Fatalln("error in password hash")
	}
	user.Password = hashedPass

	err = em.userRepo.Create(c, user)
	if err != nil {
		if errors.Is(err, repository.ErrRecordUserNotFound) {
			return fmt.Errorf("%w.", ErrUserNotFound)
		}
		return fmt.Errorf("%w: %v", ErrUserNotFound, err)
	}
	return nil
}

func (em *userUsecase) UpdateUser(c *gin.Context, user entity.User) error {
	err := em.userRepo.Update(c, user)
	if err != nil {
		if errors.Is(err, repository.ErrRecordUserNotFound) {
			return fmt.Errorf("%w.", ErrUserNotFound)
		}
		return fmt.Errorf("%w: %v", ErrUserNotFound, err)
	}
	return nil
}

func (em *userUsecase) DeleteUser(c *gin.Context) error {
	err := em.userRepo.Delete(c)
	if err != nil {
		if errors.Is(err, repository.ErrRecordUserNotFound) {
			return fmt.Errorf("%w.", ErrUserNotFound)
		}
		return fmt.Errorf("%w: %v", ErrUserNotFound, err)
	}
	return nil
}
