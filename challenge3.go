package main

import (
	"fmt"
	// "strconv"
)

func main(){
	// //challenge 1
	// var a , b int
	// fmt.Printf("Masukkan bilangan: ")
	// fmt.Scanln (&a)
	// for b = 1; b <= a; b++{
	// 	if a % b == 0 {
	// 		fmt.Println("Bilangan " , a , " bisa dibagi ", b)
	// 	}
	// }

	// //challenge 2
	// var a , b int
	// fmt.Printf("Masukkan bilangan: ")
	// fmt.Scanln (&a)
	// for b = 2; b <= a; b++{
	// 	if a%b == 0 && a != b {
	// 		fmt.Println("Bilangan " , a , " adalah bukan bilangan prima ")
	// 		break
	// 	} else {
	// 		fmt.Println("Bilangan " , a , " adalah bilangan prima ")
	// 		break
	// 	}
	// }

	// //challenge 3
	// var a , b string
	// fmt.Printf("Masukkan kata: ")
	// fmt.Scanln (&a)
	// 	for i := len(a); i > 0; i-- {
	// 		b += string(a[i-1])		
	// 	}
	// 	if b == a {
	// 		fmt.Println("Kata ", a, " adalah Palindrom")
	// 	} else {
	// 		fmt.Println("Kata ", a, " bukan Palindrom")
	// 	}

	//challenge 4
	var a, b, c, d int
	fmt.Printf("Masukkan kata: ")
	fmt.Scanf("%v", &a)

	for b=1; b<=a; b++{
		for c=a; c>=b; c--{
			fmt.Printf(" ")
		}
		for d=1; d<=b; d++{
			fmt.Printf("* ")
		}
	fmt.Println(" ")
	}
}