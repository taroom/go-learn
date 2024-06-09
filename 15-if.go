package main

import "fmt"

func main(){
	name := "Taroom"

	// IF
	if name == "Taroom" {
		fmt.Println("Hallo Taroom")
	}

	// IF ELSE
	jabatan := "manager" 

	if jabatan == "karyawan" {
		fmt.Println("Saya Karyawan")
	} else {
		fmt.Println("Saya Manager")
	}

	// IF ELSE IF
	gaji := 4

	if(gaji > 5){
		fmt.Println("Gaji Tinggi Tapera Tinggi")
	} else if(gaji > 2) {
		fmt.Println("Gaji Sedang")
	} else {
		fmt.Println("Gaji Rendah")
	}

	// SHORTIF
	if length := len(name); length > 3 {
		fmt.Println("Nama terlalu panjang")
	} else {
		fmt.Println("Nama ok")
	}
}