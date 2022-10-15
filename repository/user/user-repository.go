package userrepository

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
	"gorm.io/gorm"

	"github.com/devcamp-team-19/backend-sad/core/entity"
	repository_intf "github.com/devcamp-team-19/backend-sad/core/repository"
)

var (
	secretkey = "asdfasdfasdf"
)

type repositoryUser struct {
}

type Error struct {
	IsError bool   `json:"isError"`
	Message string `json:"message"`
}

func New() repository_intf.UserRepository {
	return &repositoryUser{}
}

func GenerateJWT(email, role string) (string, error) {
	var mySigningKey = []byte(secretkey)
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)

	claims["authorized"] = true
	claims["email"] = email
	claims["role"] = role
	claims["exp"] = time.Now().Add(time.Minute * 30).Unix()

	tokenString, err := token.SignedString(mySigningKey)

	if err != nil {
		fmt.Errorf("Something Went Wrong: %s", err.Error())
		return "", err
	}
	return tokenString, nil
}

func GeneratehashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

//compare plain password with hash password
func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

//set error message in Error struct
func SetError(err Error, message string) Error {
	err.IsError = true
	err.Message = message
	return err
}

//check whether user is authorized or not
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

		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			if claims["role"] == "admin" {
				r.Header.Set("Role", "admin")
				handler.ServeHTTP(w, r)
				return

			} else if claims["role"] == "user" {
				r.Header.Set("Role", "user")
				handler.ServeHTTP(w, r)
				return

			}
		}
		var reserr Error
		reserr = SetError(reserr, "Not Authorized.")
		json.NewEncoder(w).Encode(err)
	}
}

func (r *repositoryUser) FindAll(c *gin.Context) ([]entity.User, error) {
	var users []entity.User

	db, ok := c.MustGet("db").(*gorm.DB)
	if !ok {
		return nil, errors.New("failed to parse db to gorm")
	}

	err := db.Find(&users).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, repository_intf.ErrRecordUserNotFound
		}
		return nil, err
	}

	return users, nil
}

func (r *repositoryUser) FindSingle(c *gin.Context) (entity.User, error) {
	user := entity.User{}

	db, ok := c.MustGet("db").(*gorm.DB)
	if !ok {
		return entity.User{}, errors.New("failed to parse db to gorm")
	}

	if err := db.Where("id = ?", c.Param("id")).First(&user).Error; err != nil {
		return entity.User{}, repository_intf.ErrRecordUserNotFound
	}

	return user, nil
}

func (r *repositoryUser) Create(c *gin.Context) error {
	var input entity.UserInput
	if err := c.ShouldBindJSON(&input); err != nil {
		return errors.New("failed to create user")
	}

	db, ok := c.MustGet("db").(*gorm.DB)
	if !ok {
		return errors.New("failed to parse db to gorm")
	}

	// Create User
	User := entity.User{
		FullName: input.FullName,
		NIK:      input.NIK,
		Email:    input.Email,
		Address:  input.Address,
		Password: input.Password,
	}

	hashedPass, err := GeneratehashPassword(User.Password)
	if err != nil {
		log.Fatalln("error in password hash")
	}
	User.Password = hashedPass

	if err := db.Create(&User).Error; err != nil {
		return errors.New("failed to create user")
	}

	return nil
}

func (r *repositoryUser) Update(c *gin.Context) error {
	var input entity.UserInput
	if err := c.ShouldBindJSON(&input); err != nil {
		return errors.New("failed to update user")
	}

	db, ok := c.MustGet("db").(*gorm.DB)
	if !ok {
		return errors.New("failed to parse db to gorm")
	}

	// Update User
	User := entity.User{
		FullName: input.FullName,
		NIK:      input.NIK,
		Email:    input.Email,
		Address:  input.Address,
		Password: input.Password,
	}

	if err := db.Where("id = ?", c.Param("id")).Save(&User).Error; err != nil {
		return errors.New("failed to update user")
	}

	return nil
}

func (r *repositoryUser) Delete(c *gin.Context) error {
	db, ok := c.MustGet("db").(*gorm.DB)
	if !ok {
		return errors.New("failed to parse db to gorm")
	}

	var user entity.User
	if err := db.Where("id = ?", c.Param("id")).First(&user).Error; err != nil {
		return errors.New("failed to delete user")
	}

	if err := db.Delete(&user).Error; err != nil {
		return errors.New("failed to delete user")
	}

	return nil
}
