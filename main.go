package main

import (
	"fmt"
)

func main() {
	//controller.CariBuku("Harry potter")
	//CobaVariabel()
	//Percabangan("Buku 2")
	SwitchCase()
}

func CobaVariabel() {
	//var NamaBukuNya = 2
	//fmt.Println("Variabel 1 :", NamaBukuNya)
	//
	//javaString := 3 // Short definition variabel
	//fmt.Println(javaString)
	//
	//var CobaString string
	//fmt.Println(CobaString)

	for i := 0; i < 10; i++ {
		fmt.Println("Looping ", i)
	}

	j := 1
	for j <= 10 {
		fmt.Println(j)
		j++
	}


}

func Percabangan(namaBuku string) {

	if namaBuku == "Buku 2" {
		fmt.Println("Bukan buku 1")
	} else {
		fmt.Println("Buku 1")

	}
}

func SwitchCase() {
	i := 5
	j := 6

	switch i {
	case 1 :
		fmt.Println("one")
	case 2 :
		fmt.Println("two")
	case 3 :
		fmt.Println("three")
	default:
		fmt.Print("Variabel 1:",i)

	}
}




