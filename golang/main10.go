package main

import "fmt"

func main() {
	
	var angka [3]int
	angka[0] = 10
	angka[1] = 20
	angka[2] = 30
	fmt.Println("Isi array angka:", angka)

	angkaArray := [3]int{10, 20, 30}
	fmt.Println("Elemen index ke-1:", angkaArray[1]) 

	
	angkaSlice := []int{10, 20, 30}
	fmt.Println("Elemen index ke-0:", angkaSlice[0]) 

	angkaSlice = append(angkaSlice, 40)
	fmt.Println("Setelah di-append:", angkaSlice) 

	arr := [5]int{1, 2, 3, 4, 5}
	slice := arr[1:4]                      
	fmt.Println("Hasil slicing:", slice)   

	nilai := map[string]int{
		"Baiim": 90,
		"Aurelia": 85,
	}

	fmt.Println("Nilai Baiim:", nilai["Baiim"]) 

	nilai["Bolesin"] = 88
	delete(nilai, "Aurelia")
	fmt.Println("Isi map sekarang:", nilai)

	val, ok := nilai["Rain"]
	if ok {
		fmt.Println("Ada:", val)
	} else {
		fmt.Println("Tidak ditemukan") 
	}
}