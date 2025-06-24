package usecase

import (
	"errors"
	"go-rest-api/internal/domain"
	"time"

	"github.com/golang-jwt/jwt/v5" // Importing jwt package for token generation
	"golang.org/x/crypto/bcrypt"   // Importing bcrypt package for password hashing
)

var jwtSecret = []byte("haki-vnese") // Secret key for signing JWT tokens, should be kept secure

type UserUsecase struct {
	repo domain.UserRepository // UserUsecase is a struct that implements the UserUsecase interface
}

func NewUserUsecase(repo domain.UserRepository) *UserUsecase {
	return &UserUsecase{
		repo: repo,
	}
}

func (u *UserUsecase) Create(user *domain.User) error {
	return u.repo.Create(user)
}
func (u *UserUsecase) GetByID(id int) (*domain.User, error) {
	return u.repo.GetByID(id)
}
func (u *UserUsecase) GetByUsername(username string) (*domain.User, error) {
	return u.repo.GetByUsername(username)
}
func (u *UserUsecase) GetAll() ([]domain.User, error) {
	return u.repo.GetAll()
}
func (u *UserUsecase) Update(user *domain.User) error {
	return u.repo.Update(user)
}
func (u *UserUsecase) Delete(id int) error {
	return u.repo.Delete(id)
}
func (u *UserUsecase) Login(username, password string) (string, error) {
	user, err := u.repo.GetByUsername(username) // Retrieve user by username from the repository

	if err != nil {
		return "", err
	}

	// Compare hashed password with the provided password
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return "", errors.New("invalid credentials")
	}

	// Generate JWT token
	claims := &jwt.RegisteredClaims{
		Issuer:    user.Username,                                      // Set the issuer to the username
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)), // Set the token to expire in 24 hours
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtSecret)

	if err != nil {
		return "", err // Return an error if token signing fails
	}
	return tokenString, nil // Return the generated token string
}
