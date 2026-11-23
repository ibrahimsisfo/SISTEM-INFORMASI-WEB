package main

import "fmt"

type Mahasiswa struct {
	Nama string
	Umur int
}

func (m Mahasiswa) Perkenalan() {
	fmt.Printf("Halo, nama saya %s dan umur saya %d\n", m.Nama, m.Umur)
}

func (m Mahasiswa) TahunLahir(tahunSekarang int) int {
	return tahunSekarang - m.Umur
}

func (m Mahasiswa) UbahNamaBaru(nama string) {
	m.Nama = nama
}

func (m *Mahasiswa) UbahNama(nama string) {
	m.Nama = nama
}

func main() {

	mhs1 := Mahasiswa{Nama: "Baiim", Umur: 20}
	mhs1.Perkenalan()
	fmt.Println("Tahun lahir Baiim:", mhs1.TahunLahir(2026))

	fmt.Println("---")

	mhs2 := Mahasiswa{Nama: "Aurelia", Umur: 20}

	mhs2.UbahNamaBaru("Bolesin")
	fmt.Println("Setelah UbahNamaBaru (tanpa pointer):", mhs2.Nama) 

	mhs2.UbahNama("Rain")
	fmt.Println("Setelah UbahNama (pakai pointer):", mhs2.Nama) 
}