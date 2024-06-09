package main 

import "fmt"

func main(){
	// kita bisa membuat alias
	type sudahMenikah bool
	type namaLengkap string

	var apaNamaLengkapAnda namaLengkap = "Agus Sutarom"

	fmt.Println(apaNamaLengkapAnda)

	var sudahMenikahStatus sudahMenikah = false

	fmt.Println(sudahMenikahStatus)
}