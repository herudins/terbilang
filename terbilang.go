package terbilang

import "fmt"

func Terbilang(angka int) string {
	terbilangDesk := "Mohon maaf angka terlalu besar..."
	var bilangan = [...]string{"", "satu", "dua", "tiga", "empat", "lima", "enam", "tujuh", "delapan", "sembilan", "sepuluh", "sebelas"}

	if angka < 12 {
		terbilangDesk = bilangan[angka]
	} else if angka < 20 {
		terbilangDesk = bilangan[angka-10] + " belas"
	} else if angka < 100 {
		hasilBagi := angka / 10
		hasilMod := angka % 10
		terbilangDesk = fmt.Sprintf("%s puluh %s", bilangan[hasilBagi], bilangan[hasilMod])
	} else if angka < 200 {
		terbilangDesk = fmt.Sprintf("seratus %s", Terbilang(angka-100))
	} else if angka < 1000 {
		hasilBagi := angka / 100
		hasilMod := angka % 100
		terbilangDesk = fmt.Sprintf("%s ratus %s", bilangan[hasilBagi], Terbilang(hasilMod))
	} else if angka < 2000 {
		terbilangDesk = fmt.Sprintf("seribu %s", Terbilang(angka-1000))
	} else if angka < 1000000 {
		hasilBagi := angka / 1000
		hasilMod := angka % 1000
		terbilangDesk = fmt.Sprintf("%s ribu %s", Terbilang(hasilBagi), Terbilang(hasilMod))
	} else if angka < 1000000000 {
		hasilBagi := angka / 1000000
		hasilMod := angka % 1000000
		terbilangDesk = fmt.Sprintf("%s juta %s", Terbilang(hasilBagi), Terbilang(hasilMod))
	} else if angka < 1000000000000 {
		hasilBagi := angka / 1000000000
		hasilMod := angka % 1000000000
		terbilangDesk = fmt.Sprintf("%s milyar %s", Terbilang(hasilBagi), Terbilang(hasilMod))
	} else if angka < 1000000000000000 {
		hasilBagi := angka / 1000000000000
		hasilMod := angka % 1000000000000
		terbilangDesk = fmt.Sprintf("%s triliun %s", Terbilang(hasilBagi), Terbilang(hasilMod))
	}

	return terbilangDesk
}
