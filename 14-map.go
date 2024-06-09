package main 

import "fmt"

func main(){
	orang := map[string]string {
		"namaLengkap" : "Agus Sutarom",
		"alamat" : "Ds. Temandang",
	}

	fmt.Println(orang)
	fmt.Println(orang["namaLengkap"])
	fmt.Println(orang["alamat"])
	
	// tambah data
	orang["jabatan"] = "Programmer"
	fmt.Println(orang)
	fmt.Println(orang["jabatan"])
	
	// function builtin
	fmt.Println(len(orang))
	// mengganti alamat
	orang["alamat"] = "Tuban"
	fmt.Println(orang)
	delete(orang, "alamat") // menghapus alamat di map orang
	fmt.Println(orang)

	// fungsi make
	buku := make(map[string]string)
	buku["title"] = "aku belajar go"
	buku["author"] = "agus sutarom"
	buku["upps"] = "Salah"
	fmt.Println(buku)
	delete(buku, "upps")
	fmt.Println(buku)

	// https://youtu.be/IO_vkyJnMas?si=BnYdwuhErGlVyWKR&t=7980
}