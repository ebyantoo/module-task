package moduletask

import "fmt"

func CekGanjilGenap1(angka int) int {
	for bilangan := 1; bilangan <= angka; bilangan++ {
		fmt.Print("Masukan Bilangan : ")
		fmt.Scan(&bilangan)
		if bilangan <= angka {
			if bilangan%2 == 0 {
				fmt.Println(bilangan, " Adalah Bilangan Genap")
			} else {
				fmt.Println(bilangan, " Adalah Bilangan Ganjil")
			}
		} else {
			fmt.Println("Hanya Sampai ", angka)
		}
	}
	return angka
}

func CekGanjilGenap2(bilangan ...int) int {
	angka := 0
	for _, i := range bilangan {
		if i%2 == 0 {
			fmt.Println(i, " Adalah Bilangan Genap")
		} else {
			fmt.Println(i, " Adalah Bilangan Ganjil")
		}
	}
	return angka
}
