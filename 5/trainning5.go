package main

import (
	"fmt"
	"errors"
)

// func main(){
// 	// //contoh 1
// 	// var name string = "John"
// 	// var nameAddress *string = &name

// 	// fmt.Println("name (value)", name)
// 	// fmt.Println("name (address)", &name) //mengambil address dari pointer, mereturn alamat suatu pointing
// 	// fmt.Println("nameAddress (value)", *nameAddress) //mengambil value dari pointer, mendeklarasikan variabel pointer
// 	// fmt.Println("nameAddress (address)", nameAddress)

// 	// //contoh 2
// 	// name = "Doe" //mengganti value dari pointer (terllihat seperti replace saja)

// 	// fmt.Println("name (value)", name)
// 	// fmt.Println("name (address)", &name)
// 	// fmt.Println("nameAddress (value)", *nameAddress)
// 	// fmt.Println("nameAddress (address)", &nameAddress)

// 	// //contoh 3
// 	// numberA := 25
// 	// var numberB *int

// 	// if numberB == nil { //zero value dari pointer adalah nil
// 	// 	fmt.Println("numberB :", numberB)
// 	// 	numberB = &numberA
// 	// 	fmt.Println("numberB :", *numberB)
// 	// }

// 	// // contoh 4
// 	// var size = new(int)

// 	// fmt.Printf("Size (value) : %d \n", *size)
// 	// fmt.Printf("Size (type) : %T \n", size)
// 	// fmt.Printf("Size (address) : %v \n", &size)

// 	// *size = 11
// 	// fmt.Println("new Size (value) :", *size)
// }

// //contoh function 1
// //tanpa parameter
// func main(){
// 	hour := 11
// 	greeting(hour)
// }

// //dengan parameter
// func greeting(hour int){
// 	if hour < 12 {
// 		fmt.Println("Selamat jam lebih dari 12")
// 	} else {
// 		fmt.Println("Selamat")
// 	}
// }

// //contoh function 2
// func main(){
// 	p, l := 10, 5
// 	luas, ket := calculate(p, l) //bisa beda nama variabel di func lain karena membacanya berurutan dg deklarasi func tsb

// 	fmt.Println(luas, ket)
// }

// func calculate(panjang, lebar int) (luas int, ket string){ //return bisa langsung mendeklarasikan variabel
// 	luas = panjang * lebar
// 	ket = "keterangan"
// 	return luas, ket //atau return (saja)
// }

// //contoh function variadic
// func main(){
// 	hasil := sum(1,2,3,4,5)
// 	fmt.Println(hasil)
// }

// func sum(numb ...int) int {
// 	var total int = 0
// 	for _, number := range numb {
// 		total += number
// 	}
// 	return total
// }

// //contoh anonymous function == literal function
// func main(){
// 	func (){ //tanpa parameter
// 		fmt.Println("func pertama")
// 	}()

// 	value := func() {
// 		fmt.Println("func kedua")
// 	}
// 	value() //di munculkan melalui variabel

// 	func(sentenence string) { //dengan parameter
// 		fmt.Println(sentenence)
// 	}("func ketiga")
// }

// //contoh function closure
// func newCounter() func() int {
// 	count := 0
// 	return func() int {
// 		count += 1
// 		return count //ambil variabel dari luar
// 	}
// }

// func main() {
// 	counter := newCounter()
// 	fmt.Println(counter())
// 	fmt.Println(counter())
// }

// //contoh function defer
// func main() {
// 	defer func(){
// 		fmt.Println("Later")
// 	}()
// 	fmt.Println("First")
// }

// //contoh error handling
// func main(){
// 	_, err := validate("John")

// 	if err == nil {
// 			fmt.Println("Hello, John")
// 		} else {
// 			fmt.Println("Error: ", err)
// 		}
// }

// func validate(name string) (bool, error) {
// 	if name != "" {
// 		return true, nil
// 	}

// 	//error.New berfungsi sebagai pemberi pesan error
// 	return false, errors.New("Nama tidak boleh")
// }

//contoh panic and recover
func main(){
	defer catch()
	name := ""
	_, err := validate(name)

	if err == nil {
			fmt.Println("Hello, ", name)
		} else {
			fmt.Println("Error: ", err)
		}
		fmt.Println("end")
}

func catch(){
	if r := recover(); r != nil{
		fmt.Println("Error: ", r)
	}
}
func validate(name string) (bool, error) {
	if name != "" {
		return true, nil
	}

	//error.New berfungsi sebagai pemberi pesan error
	return false, errors.New("Nama tidak boleh")
}
