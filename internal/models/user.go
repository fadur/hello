package models

import (
	"github.com/google/uuid"
)

// User is a struct that represents a user
type User struct {
	ID       string `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (u *User) Validate() error {
	return nil
}

func (u *User) Save() error {
	// if user id is not set, set new uuid
	id := uuid.New()
	if u.ID == "" {
		u.ID = id.String()
	}
	// save user to db
	_, err := DB.Exec("INSERT INTO users (id, username, email, password) VALUES (?, ?, ?, ?)", u.ID, u.Username, u.Email, u.Password)
	if err != nil {
		return err
	}
	return nil

}

func (u *User) Update() error {
	user := User{}
	// check if user exists
	err := DB.QueryRow("SELECT * FROM users WHERE id = ?", u.ID).Scan(&user.ID, &user.Username, &user.Email, &user.Password)
	if err != nil {
		return err
	}
	// update user
	_, err = DB.Exec("UPDATE users SET username = ?, email = ?, password = ? WHERE id = ?", u.Username, u.Email, u.Password, u.ID)
	if err != nil {
		return err
	}
	return nil
}

func (u *User) Delete() error {
	_, err := DB.Exec("DELETE FROM users WHERE id = ?", u.ID)
	if err != nil {
		return err
	}
	return nil

}

func (u *User) FindByID(id int) error {
	err := DB.QueryRow("SELECT * FROM users WHERE id = ?", id).Scan(&u.ID, &u.Username, &u.Email, &u.Password)
	if err != nil {
		return err
	}
	return nil
}

// list all users
func ListUsers() ([]User, error) {
	rows, err := DB.Query("SELECT * FROM users")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	users := []User{}
	for rows.Next() {
		var user User
		rows.Scan(&user.ID, &user.Username, &user.Email, &user.Password)
		users = append(users, user)
	}
	return users, nil
}
