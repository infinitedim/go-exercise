package main

import "fmt"

func main() {
	// constant adalah variabel yang nilai nya tidak bisa diubah
	// jika membuat variabel menggunakan konstant variabel harus langsung di inisialisasikan isinya
	// dan tidak bisa di deklarasi lalu inisialisasi
	// tipe data tidak wajib ditulis pada saat pembuatan variabel
	// jika variabel constant tidak digunakan, golang tidak akan error tidak seperti menggunakan variabel biasa yang akan langsung error jika tidak digunakan
	// Constant tidak bisa menggunakan titik dua diawal pada pembuatan variabel, karna jika menggunakan titik dua pada variabel golang akan menganggap itu sebagai variabel biasa
	// Constant juga bisa melakukan multiple declaration

	// OK
	const firstName string = "Dimas"
	// OK juga
	const lastName = "Saputra"

	fmt.Println(firstName)
	fmt.Println(lastName)

	// Error
	// firstName = "Saputra";
	// lastName = "Dimas";

	// multiple declaration pada constant
	const (
		namaDepan    = "Saputra"
		namaBelakang = "Dimas"
	)

	fmt.Println(namaDepan)
	fmt.Println(namaBelakang)
}
