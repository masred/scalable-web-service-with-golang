package main

import (
	"fmt"
	"os"
	"strconv"
)

type Biodata struct {
	Nama      string
	Alamat    string
	Pekerjaan string
	Alasan    string
}

func (biodata *Biodata) LihatAlasan() string {
	return biodata.Nama + " memilih kelas Golang karena " + biodata.Alasan
}

func main() {
	biodata := []Biodata{
		{
			Nama:      "Aldi",
			Alamat:    "Bandung",
			Pekerjaan: "Fullstack Engineer",
			Alasan:    "ingin belajar hal baru",
		},
		{
			Nama:      "Dwi Fitriyanto",
			Alamat:    "Tasikmalaya",
			Pekerjaan: "Backend Engineer",
			Alasan:    "ingin hijrah dari bahasa PHP",
		},
		{
			Nama:      "Clara",
			Alamat:    "Medan",
			Pekerjaan: "Frondend Engineer",
			Alasan:    "ingin belajar backend",
		},
	}

	arg, _ := strconv.Atoi(os.Args[1])

	fmt.Println("Nama : ", biodata[arg].Nama)
	fmt.Println("Alamat : ", biodata[arg].Alamat)
	fmt.Println("Pekerjaan : ", biodata[arg].Pekerjaan)
	fmt.Println("Alasan : ", biodata[arg].LihatAlasan())
}
