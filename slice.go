package main

import "fmt"

func main() {
	var month = [...]string{
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

	var slice1 = month[4:7]
	fmt.Println(slice1)
	fmt.Println(len(slice1))
	fmt.Println(cap(slice1))

	month[5] = "Diubah"
	fmt.Println(slice1)

	slice1[0] = "Mei diubah"
	fmt.Println(slice1)

	days := [...]string{"Senin", "Selasa", "Rabu", "Kamis", "Jumat", "Sabtu", "Minggu"}
	daySlice1 := days[5:]
	daySlice1[0] = "Sabtu Baru"
	daySlice1[1] = "Minggu Baru"
	fmt.Println(days) // [Senin Selasa Rabu Kamis Jumat Sabtu Baru Minggu Baru]

	daySlice2 := append(daySlice1, "Libur baru")
	daySlice2[0] = "Ups"
	fmt.Println(daySlice1)
	fmt.Println(days)

	newSlice := make([]string, 2, 5)
	newSlice[0] = "Rezki"
	newSlice[1] = "Ananda"

	fmt.Println(newSlice)
	fmt.Println(newSlice)

	copySlice := make([]string, len(newSlice), cap(newSlice))
	copy(copySlice, newSlice)
	fmt.Println(copySlice)

	iniArray := [...]int{1, 2, 3, 4}
	iniSlice := []int{1, 2, 3, 4}

	fmt.Println(iniArray)
	fmt.Println(iniSlice)
}
