package main
import (
	"fmt"
	"strings"
)

func main()  {
	var secret interface{}

	secret = 2
	var number = secret.(int) * 10
	fmt.Println(secret, "multipledby 10 is :", number)

	secret = []string{"apple", "mango", "banana"}
	var gruits = strings.Join(secret.([]string), ", ")
	fmt.Println(gruits, "is my favorite fruits")
}