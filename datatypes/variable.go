package main

import "fmt"

func main() {
	// Deklarasi variabel dengan tipe data
	var name string

	// Inisialisasi isi variabel
	name = "Dimas Saputra"
	fmt.Println(name)

	// Isi variabel bisa diubah
	name = "Baby blooo"
	fmt.Println(name)

	// jika langsung mendeklarasikan variabel nya tipe data tidaklah wajib
	var friendName = "Andrian fadhilla"
	fmt.Println(friendName)

	var age = 18
	fmt.Println(age)

	// kata kunci var bisa digantikan dengan titik dua diawal pembuatan variabel, jika ingin mengganti isinya tidak perlu menggunakan titik dua

	country := "Indonesia"
	fmt.Println(country)

	// Tidak perlu menggunakan titik dua diawal, karna jika menggunakan akan error karene dianggap membuat variabel baru dengan nama yang sama
	country = "Amerika"
	fmt.Println(country)

	// deklarasi multiple variable
	var (
		firstName = "Dimas"
		lastName  = "Saputra"
	)

	fmt.Println(firstName)
	fmt.Println(lastName)
}
