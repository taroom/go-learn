package main 

import "fmt"

func main(){
	// saat array tetap, slice dinamis (tapi berdasarkan array karena slice berdasarkan array)
	// untuk membuat slice maka buat dulu array
	var bulan = [...]string {
		"Januari",
		"Februari",
		"Maret",
		"April",
		"Mei",
		"Juni",
		"Juli",
		"Agustus",
		"September",
		"Oktober",
		"November",
		"Desember",
	}

	/* buat slice menggunakan format berikut
	[low:high]
	[low:]
	[:high] ->
	[:] -> semua,
	slice mengacu ke array yang terslice
	*/

	// slice 1
	var sliceBulan1 = bulan[3:6]
	fmt.Println("sliceBulan1")
	fmt.Println(sliceBulan1)
	fmt.Println(len(sliceBulan1))
	fmt.Println(cap(sliceBulan1))

	// slice 2
	var sliceBulan2 = bulan[9:]
	fmt.Println("sliceBulan2")
	fmt.Println(sliceBulan2)
	fmt.Println(len(sliceBulan2))
	fmt.Println(cap(sliceBulan2))

	// slice 3
	var sliceBulan3 = bulan[:4]
	fmt.Println("sliceBulan3")
	fmt.Println(sliceBulan3)
	fmt.Println(len(sliceBulan3))
	fmt.Println(cap(sliceBulan3))

	// slice 4
	var sliceBulan4 = bulan[:]
	fmt.Println("sliceBulan4")
	fmt.Println(sliceBulan4)
	fmt.Println(len(sliceBulan4))
	fmt.Println(cap(sliceBulan4))

	// ingat kalau sliceBulan1 - 4 itu mengambil data yang sama yakni bulan, jadi jika ada yang berubah pada bulan maka semua berubah,
	// begitu juga jika anda mengganti salah satu sliceBulan.. maka akan mengganti semua
	// capacity adalah jumlah sisa yang tersedia

	/* percobaan append */
	var sliceBulan5 = bulan[10:]
	fmt.Println("slice bulan 5", sliceBulan5)
	sliceBulan5[0] = "November lagi" // mengganti bulan november yang indexnya 10 di bulan, tapi 0 di slice 5 mengapa pointer akan menjadi 10 -> 0 di slice5 walaupun begitu ketika 0 dislice 5 diubah, maka akan mengubah bulan di index 10 (november)
	fmt.Println(bulan)

	var appendBulan = append(sliceBulan5, "Lipsember") //bulan tetap, karena append akan membuat array baru, jika capacity terlampaui

	fmt.Println("bulan", bulan)
	fmt.Println("append bulan", appendBulan)

	// membuat slice baru dengan cara direct
	motor := make([]string, 2, 5) // 2 adalah jumlah item, 5 adalah capacity
	fmt.Println("motor belum init", motor)
	motor[0] = "Harley"
	motor[1] = "Supra"
	fmt.Println("motor sudah init", motor)
	fmt.Println("len motor sudah init", len(motor))
	fmt.Println("motor sudah init", cap(motor))

	// copy
	copyMotor := make([]string, len(motor), cap(motor)) //ini untuk memastikan agar jumlah len dan capacity sama
	// lalu ketika copyMotor masih kosong kita gunakan
	copy(copyMotor, motor)

	// warning hati hati jika salah definisi akan jadi slice
	iniArr := [5]int { 1,2,3,4,5}
	iniArrJuga := [...]int { 1,2,3,4,5}
	iniSlice := []int { 1,2,3,4,5}

}