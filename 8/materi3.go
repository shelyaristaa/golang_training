// package main

// import (
//   "net/http"
//   "github.com/labstack/echo"
//   "strconv"
// )

// type User struct{
//   Id int
//   Name string
//   Email string

// }

// // user controller
// func GetUserController(c echo.Context) error{
//   id, _ := strconv.Atoi(c.Param("id"))
//   user := User{Id: id, Name: "Ismail", Email: "ismail@alterra.id"}
//   // render data - JSON response
//   return c.JSON(http.StatusOK, map[string]interface{}{
// 	"user": user,
//   })
// }

// func main()  {
//   e := echo.New()
//   //routing
//   e.GET("/user/:id", GetUserController)
//   e.Start(":8000")
// }