package main 

import "fmt"

func main(){
	// index di mulai dari 0, dan hanya 3 element saja yang dapat dimasukkan (tidak dapat diubah) mau ubah buat var array lagi
	var arr [3]string
	arr[0] = "Agus"
	arr[1] = "Suta"
	arr[2] = "Rom"

	fmt.Println(arr[0])
	fmt.Println(arr[1])
	fmt.Println(arr[2])

	// direct
	var values = [3]int{
		90,
		85,
		80,
	}

	fmt.Println(values)
	fmt.Println(len(arr))
	fmt.Println(len(values))
}