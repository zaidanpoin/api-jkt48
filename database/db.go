package database

import "strings"

type Member struct {
	ID   int
	Name string
}

// Inisialisasi slice names dengan data member beserta ID-nya
var names = []Member{
	{1, "Amanda Sukma"},
	{2, "Angelina Christy"},
	{3, "Aurellia Azizi"},
	{4, "Asadel Callista"},
	{5, "Alifia Cornelia"},
	{6, "Vanisa FebriolaSinambela"},
	{7, "Feni Fitriyanti"},
	{8, "Fiony Alveria"},
	{9, "Flora Shafiq"},
	{10, "Freya Jayawardana"},
	{11, "Gabriel Abigail"},
	{12, "Gita Sekar Andarini"},
	{13, "Grace Octaviani"},
	{14, "Greesella Adhalia"},
	{15, "Helisma Putri"},
	{16, "Indah Cahya"},
	{17, "Indira Seruni"},
	{18, "Jessica Chandra"},
	{19, "Jesslyn Elly"},
	{20, "Kathrina Irene"},
	{21, "Lulu Salsabila"},
	{22, "Marsha Lenathea"},
	{23, "Mutiara Azzahra"},
	{24, "Raisha Syifa"},
	{25, "Reva Fidela"},
	{26, "Shania Gracia"},
}

// Fungsi untuk mendapatkan semua data
func GetAllData() []Member {
	return names
}

func GetName(nama string) string {
	var mem string
	for _, data := range names {
		// Mengonversi kedua string menjadi huruf kecil untuk membandingkan tanpa memperhatikan case
		if strings.Contains(strings.ToLower(data.Name), strings.ToLower(nama)) {
			mem = data.Name
			break
		}
	}

	return mem
}
