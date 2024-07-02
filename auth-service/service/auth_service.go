package service

import (
	model2 "auth/model"
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
)

type AuthService struct {
	userServiceURL string
	jwtSecret      string
}

func NewAuthService(userServiceURL, jwtSecret string) *AuthService {
	return &AuthService{userServiceURL: userServiceURL, jwtSecret: jwtSecret}
}

func (s *AuthService) Signup(user model2.User) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	user.Password = string(hashedPassword)

	// Send user data to user service
	userJson, err := json.Marshal(user)
	if err != nil {
		return err
	}

	resp, err := http.Post(s.userServiceURL+"/users", "application/json", bytes.NewBuffer(userJson))
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusCreated {
		return errors.New("failed to create user in user service")
	}

	return nil
}

func (s *AuthService) Login(username, password string) (string, error) {
	user, err := s.getUserFromUserService(username)
	if err != nil {
		return "", err
	}

	if user == nil {
		return "", errors.New("user not found")
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return "", errors.New("invalid password")
	}

	token, err := s.generateJWT(user)
	if err != nil {
		return "", err
	}
	return token, nil
}

func (s *AuthService) getUserFromUserService(username string) (*model2.User, error) {
	resp, err := http.Get(s.userServiceURL + "/users?username=" + username)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, errors.New("user not found in user service")
	}

	var user model2.User
	if err := json.NewDecoder(resp.Body).Decode(&user); err != nil {
		return nil, err
	}

	return &user, nil
}

func (s *AuthService) generateJWT(user *model2.User) (string, error) {
	claims := jwt.MapClaims{
		"sub":  user.ID,
		"name": user.Username,
		"exp":  time.Now().Add(time.Hour * 72).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(s.jwtSecret))
}
