package main

import (
	"fmt"
	"os"
	"strconv"
)

type TemanKls struct {
	Nama      string
	Alamat    string
	Pekerjaan string
	Alasan    string
}

func main() {

	teman := []TemanKls{
		{"Afi", "Padang", "Mahasiswa", "Tambah Ilmu"},
		{"Baba", "Jakarta", "Freelancer", "Isi waktu luang"},
		{"Cece", "Bandung", "Programer", "Tambah skill"},
	}

	args := os.Args[1:]

	for _, arg := range args {
		num, err := strconv.Atoi(arg)
		if err != nil {
			fmt.Printf("%s is not a valid integer, use a number!\n", arg)
			continue
		}

		if num == 1 {
			printPeople(teman, 0)

		} else if num == 2 {
			printPeople(teman, 1)

		} else if num == 3 {
			printPeople(teman, 2)

		} else {
			fmt.Println("no data found, use other number")
		}
	}
}
func printPeople(teman []TemanKls, index int) {

	if index < 0 || index >= len(teman) {
		fmt.Println("Invalid index")
		return
	}
	t := teman[index]

	fmt.Printf("{\nNama : %q\nAlamat : %q\nPekerjaan : %q\nAlasan : %q\n}", t.Nama, t.Alamat, t.Pekerjaan, t.Alasan)

}
