// package main

// import (
//   "net/http"
//   "github.com/labstack/echo"
// )

// // user controller
// func UserSearchController(c echo.Context) error {
// 	match := c.QueryParam("match")
// 	return c.JSON(http.StatusOK, map[string]interface{}{
// 	  "match" : match,
// 	  "result" : []string{"adi", "aan", "asif"},
// 	})
//   }

// func main() {
//   e := echo.New()
//   //routing
//   e.GET("/users", UserSearchController)
//   e.Start(":8000")
// }