package main

import "fmt"

type BisaBerjalan interface {
	Jalan()
}

type Manusia struct {
	Nama string
}

type Robot struct {
	Kode string
}

func (m Manusia) Jalan() {
	fmt.Println(m.Nama, "sedang berjalan")
}

func (r Robot) Jalan() {
	fmt.Println("Robot", r.Kode, "berjalan otomatis")
}

func Gerakkan(b BisaBerjalan) {
	b.Jalan()
}

func CetakData(data interface{}) {
	fmt.Println("Data diterima:", data)
}

func main() {
	m := Manusia{Nama: "Baiim"}
	r := Robot{Kode: "RX-01"}

	Gerakkan(m) 
	Gerakkan(r) 

	fmt.Println("---")

	CetakData("Halo Golang") 
	CetakData(2026)          
	CetakData(m)             
}