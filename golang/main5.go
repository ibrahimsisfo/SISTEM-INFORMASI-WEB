package main

import "fmt"

const phi = 3.14
const aplikasi = "Belajar Golang"
const angka = 10    
const umur int = 20 
const (
	satu = 1
	dua  = 2
	tiga = 3
)

const (
	A = iota 
	B        
	C        
)
func main() {
	fmt.Println("Aplikasi:", aplikasi, "| Phi:", phi)
	fmt.Println("Angka:", angka, "| Umur:", umur)
	fmt.Println("Kelompok angka:", satu, dua, tiga)
	fmt.Println("Hasil iota:", A, B, C)
}