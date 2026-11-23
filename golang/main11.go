package main

import "fmt"

type Mahasiswa struct {
	Nama  string
	Umur  int
	Aktif bool
}

type Alamat struct {
	Kota  string
	Jalan string
}

type MahasiswaLengkap struct {
	Nama   string
	Umur   int
	Alamat Alamat 
}

func ubahUmur(m *Mahasiswa) {
	
	m.Umur = 25
}

func main() {

	var mhs1 Mahasiswa
	mhs1.Nama = "Baiim"
	mhs1.Umur = 20
	mhs1.Aktif = true
	fmt.Println("Data mhs1:", mhs1)

	mhs2 := Mahasiswa{
		Nama:  "Aurelia",
		Umur:  22,
		Aktif: false,
	}
	fmt.Println("Nama mhs2:", mhs2.Nama)
	fmt.Println("Umur mhs2:", mhs2.Umur)

	mhs := MahasiswaLengkap{
		Nama: "Bolesin",
		Umur: 21,
		Alamat: Alamat{
			Kota:  "Balanipa",
			Jalan: "Jl. Merdeka",
		},
	}
	fmt.Println("Kota Bolesin:", mhs.Alamat.Kota) 

	m := Mahasiswa{Nama: "Rain", Umur: 20}
	ubahUmur(&m)                                 
	fmt.Println("Umur Rain setelah diubah:", m.Umur) 
}