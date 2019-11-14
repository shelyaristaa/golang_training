// package main

// import (
//   "net/http"
//   "github.com/labstack/echo"
// )

// type User struct{
//   Name string
//   Email string

// }

// func GetUser(c echo.Context)error{
//   user := User{Name:"Ismail", Email:"ismail@alterra.id"}
//   return c.JSON(http.StatusOK, user)

// }

// func main()  {
//   e := echo.New()
//   e.GET("/user", GetUser)
//   e.Start(":8000")

// }