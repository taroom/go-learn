package main 

import "fmt"

func main(){

	var namaLengkap string

	// jika tidak digunakan maka variabel akan mengeluarkan error variabel tidak pernah digunakan

	namaLengkap = "Agus Sutarom"
	fmt.Println(namaLengkap)

	namaLengkap = "Hilya FG"
	fmt.Println(namaLengkap)

	// cara membuat variabel tanpa harus menyebutkan tipe data
	var namaDepan = "Agus"
	fmt.Println(namaDepan)

	var umur = 30
	fmt.Println(umur)

	// jika tanpa var maka harus menggunakan := dan langsung diisikan variablenya
	alamat := "Jl. Meraki Sendang"
	fmt.Println(alamat)
	
	alamat = "Jl. Merdeka Indah"
	fmt.Println(alamat)

	// deklarasi multiple variable
	var (
		firstName = "Agus"
		lastName = "Sutarom"
	)

	fmt.Println("hallo nama saya ", firstName, lastName)
	
	// https://www.youtube.com/watch?v=IO_vkyJnMas&list=PL-CtdCApEFH-0i9dzMzLw6FKVrFWv3QvQ teruskan disini

	// https://www.youtube.com/watch?v=tCmTX7__S84&list=PL5dTjWUk_cPbExff-KfZ18abspIgfcfmt - todolist
	// https://www.youtube.com/watch?v=PxlHx6Whcic


	// DOCKER
	// https://www.youtube.com/watch?v=26O6Ke03j3Y
	// https://www.youtube.com/watch?v=C5y-14YFs_8
}