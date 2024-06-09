package main 

import "fmt"

func main(){
	var nilai32 int32 = 100001
	var nilai64 int64 = int64(nilai32)
	var nilai8 int8 = int8(nilai32) // hati hati jika nilai tidak cukup akan jadi minus atau nilai tak terduga

	fmt.Println(nilai32)
	fmt.Println(nilai64)
	fmt.Println(nilai8)

	var nama = "Agus"
	var g byte = nama[1]//jadi byte / uint8
	var gString = string(g)

	fmt.Println(gString)
}