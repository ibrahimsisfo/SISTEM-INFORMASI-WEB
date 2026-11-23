package main

import "fmt"

func sapa() {
	fmt.Println("Halo dari fungsi!")
}

func halo(nama string) {
	fmt.Println("Halo", nama)
}

func tambah(a, b int) {
	fmt.Println("Hasil tambah:", a+b)
}

func kali(a, b int) int {
	return a * b
}

func hitung(a, b int) (int, int) {
	return a + b, a * b
}

func bagi(a, b int) (hasil int) {
	hasil = a / b
	return 
}

func main() {
	sapa()
	halo("Golang")
	tambah(10, 5)

	hasil := kali(3, 4)
	fmt.Println("Hasil kali:", hasil) 

	jumlah, perkalian := hitung(2, 5)
	fmt.Println("Jumlah:", jumlah)       
	fmt.Println("Perkalian:", perkalian) 

	hasilBagi := bagi(20, 4)
	fmt.Println("Hasil bagi:", hasilBagi) 

	f := func(nama string) {
		fmt.Println("Hai", nama)
	}
	f("Golang")
}