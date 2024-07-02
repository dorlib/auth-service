package tests

import (
	"github.com/DATA-DOG/go-sqlmock"
	"testing"
	model2 "user/model"
	"user/repository"
	service2 "user/service"
)

func TestGetAllUsers(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	mockRepo := repository.NewUserRepository(db)
	mockService := service2.NewUserService(mockRepo)

	rows := sqlmock.NewRows([]string{"id", "username"}).
		AddRow(1, "user1").
		AddRow(2, "user2")

	mock.ExpectQuery("SELECT id, username FROM users").WillReturnRows(rows)

	users, err := mockService.GetAllUsers()
	if err != nil {
		t.Errorf("expected no error, but got %v", err)
	}

	if len(users) != 2 {
		t.Errorf("expected 2 users, but got %v", len(users))
	}
}

func TestGetUserByID(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	mockRepo := repository.NewUserRepository(db)
	mockService := service2.NewUserService(mockRepo)

	mockUser := &model2.User{
		ID:       1,
		Username: "user1",
	}

	mock.ExpectQuery("SELECT * FROM users WHERE id = ?").
		WithArgs(mockUser.ID).
		WillReturnRows(sqlmock.NewRows([]string{"id", "username"}).
			AddRow(mockUser.ID, mockUser.Username))

	user, err := mockService.GetUserByID(mockUser.ID)
	if err != nil {
		t.Errorf("expected no error, but got %v", err)
	}

	if user.Username != mockUser.Username {
		t.Errorf("expected username %v, but got %v", mockUser.Username, user.Username)
	}

	if user.Email != mockUser.Email {
		t.Errorf("expected email %v, but got %v", mockUser.Email, user.Email)
	}
}

func TestCreateUser(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	mockRepo := repository.NewUserRepository(db)
	mockService := service2.NewUserService(mockRepo)

	mockUser := &model2.User{
		Username: "user1",
		Email:    "example@gmail.com",
		Password: "password",
	}

	mock.ExpectExec("INSERT INTO users").
		WithArgs(mockUser.Username, mockUser.Password).
		WillReturnResult(sqlmock.NewResult(1, 1))

	err = mockService.CreateUser(mockUser)
	if err != nil {
		t.Errorf("expected no error, but got %v", err)
	}
}

func TestUpdateUser(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	mockRepo := repository.NewUserRepository(db)
	mockService := service2.NewUserService(mockRepo)

	mockUser := &model2.User{
		ID:       1,
		Username: "user1",
		Email:    "example@gmail.com",
		Password: "password",
	}

	mock.ExpectExec("INSERT INTO users").
		WithArgs(mockUser.Username, mockUser.Email, mockUser.Password).
		WillReturnResult(sqlmock.NewResult(1, 1))

	err = mockService.CreateUser(mockUser)
	if err != nil {
		t.Errorf("expected no error, but got %v", err)
	}

	newUser := &model2.User{
		Email: "example2@gmail.com",
	}

	err = mockService.UpdateUser(newUser)
	if err != nil {
		t.Errorf("expected no error, but got %v", err)
	}

	user, err := mockService.GetUserByID(mockUser.ID)
	if err != nil {
		t.Errorf("expected no error, but got %v", err)
	}

	if user.Email != newUser.Email {
		t.Errorf("expected email %v, but got %v", user.Email, newUser.Username)
	}
}

func TestDeleteUser(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	mockRepo := repository.NewUserRepository(db)
	mockService := service2.NewUserService(mockRepo)

	mockUser := &model2.User{
		ID:       1,
		Username: "user1",
		Email:    "example@gmail.com",
		Password: "password",
	}

	mock.ExpectExec("INSERT INTO users").
		WithArgs(mockUser.Username, mockUser.Email, mockUser.Password).
		WillReturnResult(sqlmock.NewResult(1, 1))

	err = mockService.CreateUser(mockUser)
	if err != nil {
		t.Errorf("expected no error, but got %v", err)
	}

	err = mockService.DeleteUser(mockUser.ID)
	if err != nil {
		t.Errorf("expected no error, but got %v", err)
	}

	user, err := mockService.GetUserByID(mockUser.ID)
	if err != nil {
		t.Errorf("expected no error, but got %v", err)
	}

	if user != nil {
		t.Error("user should not exists")
	}
}
