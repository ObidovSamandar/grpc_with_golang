package controllers

import (
	"context"
	"fmt"
	"strings"

	"github.com/obidovsamandar/go-crud-with-grpc/db"
)

type Server struct {
}

type UserInfo struct {
	ID        string `db:"id"`
	FirstName string `db:"first_name"`
	Email     string `db:"email"`
	LastName  string `db:"last_name"`
}

func (s *Server) AddUser(ctx context.Context, user *User) (*OutPut, error) {
	userData := UserInfo{}

	row := db.DBPsql.QueryRow("SELECT * FROM users WHERE email=$1", user.Email)

	err := row.Scan(&userData.ID, &userData.FirstName, &userData.LastName, &userData.Email)

	if err == nil {
		return &OutPut{Message: "Email already exists!"}, nil
	}

	_, err = db.DBPsql.Exec("INSERT INTO users(first_name, last_name,email) VALUES($1,$2,$3) RETURNING id, first_name", user.FirstName, user.LastName, user.Email)

	if err != nil {
		return &OutPut{Message: "Failed to save user data!"}, nil
	}

	getUserRow := db.DBPsql.QueryRow("SELECT * FROM users WHERE email=$1", user.Email)

	_ = getUserRow.Scan(&userData.ID, &userData.FirstName, &userData.LastName, &userData.Email)

	response := fmt.Sprintf("%v added successfully. ID=%v", user.FirstName, userData.ID)

	return &OutPut{Message: response}, nil
}

func (s *Server) GetUser(ctx context.Context, user *GetDeleteUserRequest) (*OutPut, error) {
	userData := UserInfo{}

	row := db.DBPsql.QueryRow("SELECT * FROM users WHERE id=$1", user.Id)

	err := row.Scan(&userData.ID, &userData.FirstName, &userData.LastName, &userData.Email)

	if err != nil {
		return &OutPut{Message: "User not found!"}, nil
	}

	response := fmt.Sprintf("id=%v,first_name=%v,last_name=%v,email=%v", userData.ID, userData.FirstName, userData.LastName, userData.Email)

	return &OutPut{Message: response}, nil
}

func (s *Server) GetAllUser(ctx context.Context, allUser *GetAllUserRequest) (*OutPut, error) {
	userData := []UserInfo{}

	err := db.DBPsql.Select(&userData, "SELECT id,first_name,last_name, email FROM users;")

	fmt.Println(err)

	if err != nil {
		return &OutPut{Message: "User not found!"}, nil
	}

	var usersStr []string

	for _, value := range userData {
		individualInfo := fmt.Sprintf("{id=%v first_name=%v last_name=%v email=%v}", value.ID, value.FirstName, value.LastName, value.Email)
		usersStr = append(usersStr, individualInfo)
	}

	return &OutPut{Message: strings.Join(usersStr, " ")}, nil
}

func (s *Server) DeleteUser(ctx context.Context, user *GetDeleteUserRequest) (*OutPut, error) {
	userData := UserInfo{}

	row := db.DBPsql.QueryRow("SELECT * FROM users WHERE id=$1", user.Id)

	err := row.Scan(&userData.ID, &userData.FirstName, &userData.LastName, &userData.Email)

	if err != nil {
		return &OutPut{Message: "User not found!"}, nil
	}

	_, err = db.DBPsql.Exec("DELETE FROM users WHERE id=$1", user.Id)

	if err != nil {
		return &OutPut{Message: "Failed to delete user data!"}, nil
	}

	return &OutPut{Message: "User successfully deleted!"}, nil
}

func (s *Server) UpdateUser(ctx context.Context, user *UpdateUserRequest) (*OutPut, error) {
	userData := UserInfo{}

	row := db.DBPsql.QueryRow("SELECT * FROM users WHERE id=$1", user.UserIdForUpdate)

	err := row.Scan(&userData.ID, &userData.FirstName, &userData.LastName, &userData.Email)

	if err != nil {
		return &OutPut{Message: "User not found!"}, nil
	}

	_, err = db.DBPsql.Exec("UPDATE users SET  first_name=$1 , last_name=$2, email=$3 WHERE id=$4", user.FirstName, user.LastName, user.Email, userData.ID)

	if err != nil {
		return &OutPut{Message: "Failed to update user data!"}, nil
	}

	return &OutPut{Message: "User data successfully updated!"}, nil
}
