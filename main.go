package main

import (
	"fmt"
	"training/aritmatika"
)

func main(){
	// fmt.Println("Hello :)")

	// //cara 1
	// var name string

	// //cara 2
	// var name1 string = "shely"

	// //cara 3
	// var name2, gender string

	// //cara 4
	// var name3, gender1 string = "shely", "arista"

	// //bolean
	// isLogin := false

	// //numeric integer
	// var a int = 15
	// b := 20

	// //numeric float & integer
	// c, d := 2.1, 30

	// //string
	// s:= "hello!"

	// const pi = 3.14
	// fmt.Println(isLogin, a, b, c, d , s, pi)


	// //number operation
	// a := 10
	// t := 15
	// L := (a * t) / 2
	// fmt.Println(L)

	// //string operation
	// helloword := "hello" + " " + "word"
	// fmt.Println(helloword)



	// //challenge 1
	// var a, t, L int
	// fmt.Println("Masukkan alas: ")
	// fmt.Scanln (&a)

	// fmt.Println("Masukkan tinggi: ")
	// fmt.Scanln (&t)

	// L = (a * t) / 2
	// fmt.Println("Hasil: " , + L)


	// //challenge 2
	// var phi, r, t, L float32
	// fmt.Println("Masukkan r: ")
	// fmt.Scanln (&r)

	// fmt.Println("Masukkan t: ")
	// fmt.Scanln (&t)

	// phi = 3.14 
	// L = 2 * phi * r * (r + t)
	// fmt.Println("Hasil: " , + L)

		hasilTambah := aritmatika.Tambah(1,2)
		fmt.Println(hasilTambah)

		hasilKurang := aritmatika.Kurang(5, 3)
		fmt.Println(hasilKurang)

}