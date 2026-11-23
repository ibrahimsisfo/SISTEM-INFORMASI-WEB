package main

import "fmt"

func main() {
	
	nilai := 85

	if nilai >= 90 {
		fmt.Println("Nilai A")
	} else if nilai >= 80 {
		fmt.Println("Nilai B") 
	} else {
		fmt.Println("Nilai C atau di bawahnya")
	}

	hari := "Senin"

	switch hari {
	case "Senin":
		fmt.Println("Awal minggu") 
	case "Sabtu", "Minggu": 
		fmt.Println("Weekend!")
	default:
		fmt.Println("Hari biasa")
	}

	for i := 0; i < 5; i++ {
		fmt.Println("Loop ke:", i)
	}

	for {
		fmt.Println("Terus jalan...")
		break 
	}

	angka := []int{1, 2, 3, 4}

	for index, value := range angka {
		fmt.Println("Index:", index, "Value:", value)
	}

	for _, value := range angka {
		fmt.Println("Hanya value:", value)
	}
}