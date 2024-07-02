package repository

import (
	"database/sql"
	"log"
	model2 "user/model"
)

type UserRepository interface {
	GetAllUsers() ([]*model2.User, error)
	GetUserByID(id int64) (*model2.User, error)
	GetUserByUserName(username string) (*model2.User, error)
	CreateUser(user *model2.User) error
	UpdateUser(user *model2.User) error
	DeleteUser(id int64) error
}

type userRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) UserRepository {
	return &userRepository{db}
}

func (repo *userRepository) GetAllUsers() ([]*model2.User, error) {
	rows, err := repo.db.Query("SELECT * FROM users")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []*model2.User
	for rows.Next() {
		user := &model2.User{}
		if err := rows.Scan(&user.ID, &user.Email, &user.Username); err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	return users, nil
}

func (repo *userRepository) GetUserByID(id int64) (*model2.User, error) {
	row := repo.db.QueryRow("SELECT * FROM users WHERE id = ?", id)
	user := &model2.User{}

	if err := row.Scan(&user.ID, &user.Email, &user.Username); err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}

		return nil, err
	}

	return user, nil
}

func (repo *userRepository) GetUserByUserName(username string) (*model2.User, error) {
	row := repo.db.QueryRow("SELECT * FROM users WHERE username = ?", username)
	user := &model2.User{}

	if err := row.Scan(&user.ID, &user.Email, &user.Username); err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}

		return nil, err
	}

	return user, nil
}

func (repo *userRepository) CreateUser(user *model2.User) error {
	_, err := repo.db.Exec("INSERT INTO users (username, email, password) VALUES (?, ?, ?)", user.Username, user.Email, user.Password)
	if err != nil {
		return err
	}

	return nil
}

func (repo *userRepository) UpdateUser(user *model2.User) error {
	updateUserName := user.Username != ""
	updateEmail := user.Email != ""
	updatePassword := user.Password != ""

	if updateEmail {
		_, err := repo.db.Query("UPDATE users SET email = ? WHERE id = ?", user.Email, user.ID)
		if err != nil {

			return err
		}
	}

	if updateUserName {
		_, err := repo.db.Query("UPDATE users SET username = ? WHERE id = ?", user.Username, user.ID)
		if err != nil {

			return err
		}
	}

	if updatePassword {
		_, err := repo.db.Query("UPDATE users SET password = ? WHERE id = ?", user.Password, user.ID)
		if err != nil {

			return err
		}
	}

	log.Printf("updated note: %d", user.ID)

	return nil
}

func (repo *userRepository) DeleteUser(id int64) error {
	_, err := repo.db.Exec("DELETE FROM users WHERE id = ?", id)
	if err != nil {
		return err
	}

	log.Printf("deleted note: %d", id)

	return nil
}
