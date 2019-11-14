package main

import (
	"net/http"
	"strconv"
	"github.com/labstack/echo"
)

type user struct {
	Id			int		`json:"id" from:"id"`
	Name		string	`json:"name" from:"name"`
	Email		string	`json:"email" from:"email"`
	Password	string	`json:"password" from:"password"`
}

var user []User

// ------------------------ controller ------------------------

// get all users
func GetUsersController(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]interface{}{
		"messages": "success get all users",
		"users": users,
	})
}

//get user by id
func GetUserController(c echo.Context) error {
	
}

//delete user by id
func DeleteUserController(c echo.Context) error {
	
}

//update user by id
func UpdateUserController(c echo.Context) error {
	
}

//create new user
func CreateUserController(c echo.Context) error {
	//binding data
	user := User{}
	c.Bind(&user)

	if len(users) == 0 {
		user.Id = 1
	} else {
		newId := users[len(users)-1].Id + 1
		user.Id = newId
	}
	users = append(users, user)
	return c.JSON(http.StatusOK, map[string]interface{}{
		"messages": "success create user",
		"user": user,
	})
}
// --------------------------------------------------------------
func main() {
	e := echo.New()
	// routing with query parameter
	e.GET("/users", GetUsersController)
	e.POST("/users", GetUserController)

	// start the server, and log if it fails
	e.Logger.Fatal(e.Start(:8000))
}