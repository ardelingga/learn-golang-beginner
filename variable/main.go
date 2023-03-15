package main

import "fmt"

// CONSTANT VARIABLE 
const VaribleConstant string = "Tentang Kode"

// Notes Variable Naming
/*
	Kalau nama variable di awali dengan huruf besar itu artinya dia bisa di export, bisa digunakan di file file yang lainnya

	Namun jika hurufnya kecil didepannya itu cuman bisa di gunakan di file itu sendiri
*/

func main() {
	
	// EXPLICIT TIPE -> Membuat sebuah variable dengan menyertakan type variable ------------------------------------------------------------------------------
	var youtube string = "ardelingga" // Variable String
	fmt.Println(youtube)

	// To find type data of variable
	fmt.Printf("Tipe data variable youtube adalah: %T \n", youtube)

	var isHuman bool = true
	fmt.Println(isHuman)
	fmt.Printf("Tipe data isHuman adalah: %T \n", isHuman)

	var angkaKecil uint8 = 255
	fmt.Println(angkaKecil)
	fmt.Printf("Tipe data angkaKecil adalah : %T \n", angkaKecil)

	var angkaDecimal1 float64 = 255.234986198
	fmt.Println(angkaDecimal1)
	fmt.Printf("Tipe data angkaDecimal1 adalah : %T \n", angkaDecimal1)

	var angkaLain int
	fmt.Println(angkaLain)
	fmt.Printf("Tipe data variable angkaLain adalah : %T", angkaLain)

	var stringLain string
	fmt.Println(stringLain)
	fmt.Printf("Tipe data variable stringLain adalah : %T \n", stringLain)

	// IMPLICIT TIPE -> Membuat sebuah variable tanpa menyertakan type variable ------------------------------------------------------------------------------
	var myName = "Ardelingga"
	fmt.Println(myName)
	fmt.Printf("Tipe data variable myName adalah : %T \n", myName)



	// MEMBUAT VARIABLE DENGAN METODE CEPAT
	myLongName := "Ardelingga Pramesta Kusuma"
	fmt.Println(myLongName)
	fmt.Printf("Tipe data variable myName adalah : %T \n", myLongName)


	// CONSTANT VARIABLE 
	const NamaLengkap string = "Tentang Kode"
}
