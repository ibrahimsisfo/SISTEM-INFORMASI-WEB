package main

import "fmt"

func ubahNilai(x int) {
	x = 100
}

func ubahPointer(x *int) {
	*x = 100 
}

func main() {
	var angka int = 10
	var ptr *int = &angka

	fmt.Println("Nilai angka:", angka)
	fmt.Println("Alamat angka:", &angka)
	fmt.Println("Pointer ptr:", ptr)

	fmt.Println("Nilai melalui pointer:", *ptr)

	a := 10
	ubahNilai(a)
	fmt.Println("Setelah ubahNilai (tanpa pointer):", a) 

	ubahPointer(&a)                                       
	fmt.Println("Setelah ubahPointer (pakai pointer):", a) 

	angkaBaru := new(int)
	*angkaBaru = 50
	fmt.Println("Nilai dari new(int):", *angkaBaru) 
}