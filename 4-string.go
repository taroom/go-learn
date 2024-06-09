package main 

import "fmt"

func main(){
	fmt.Println("Benar ");

	// fx
	fmt.Println("Jumlah huruf taroom", len("taroom"))
	
	//mengembalikan r itu dalam 111 (byte) bukan string r
	fmt.Println("ambil huruf ke-3 taroom", "taroom"[3])
}