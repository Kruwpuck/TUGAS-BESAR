package main

import "fmt"

const NMAX int = 50

type data_pasien struct {
	nama        string
	id, layanan int
	waktu       periode
}
type periode struct {
	tanggal, bulan, tahun int
}
type tabInt [NMAX]data_pasien

func main() {

	home()
}
func home() {
	var opsi int
	var data tabInt
	var jumlah_pasien int
	fmt.Println("Selamat datang Di Layanan Medical Check Up")
	fmt.Println("Pilih Opsi berikut:")
	fmt.Println("1. Penambahan Data Pasien")
	fmt.Println("2. Penghapusan Data Pasien")
	fmt.Println("3. Pengeditan Data Pasien")
	fmt.Println("4. Pencarian Data Pasien")
	fmt.Scan(&opsi)
	if opsi == 1 {
		tambah(&data, &jumlah_pasien)
	} else if opsi == 2 {
		hapus(&data, &jumlah_pasien)
	} else if opsi == 3 {
		edit(&data, &jumlah_pasien)
	} else if opsi == 4 {
		cari(data, jumlah_pasien)
	}
}
func tambah(A *tabInt, N *int) {
	fmt.Print("Masukkan nama pasien: ")
	fmt.Scan(&A[*N].nama)
	fmt.Println()
	A[*N].id = *N
	fmt.Print("Masukkan Jenis Layanan (Tingkat 1,2 atau 3): ")
	fmt.Scan(&A[*N].layanan)
	fmt.Println()
	fmt.Print("Masukkan Waktu (Format Angka: Tahun Bulan Tanggal): ")
	fmt.Scan(&A[*N].waktu.tahun, &A[*N].waktu.bulan, &A[*N].waktu.tanggal)
	*N++
}
func hapus(A *tabInt, N *int) {
	var opsi int
	fmt.Print("Pilih opsi hapus: ")
	fmt.Println("1. Cari Berdasarkan Nama")
	fmt.Println("2. Cari Berdasarkan ID")
	fmt.Scan(&opsi)
	if opsi == 1 {
		display_pasien(*A, cari_nama(*A, *N))
	} else if opsi == 2 {
		display_pasien(*A, cari_id(*A, *N))
	}
}
func home_edit(A *tabInt, N *int) {
	var opsi int
	fmt.Print("Pilih opsi edit: ")
	fmt.Println("1. Cari Berdasarkan Nama")
	fmt.Println("2. Cari Berdasarkan ID")
	fmt.Scan(&opsi)
	if opsi == 1 {
		display_pasien(*A, cari_nama(*A, *N))
	} else if opsi == 2 {
		display_pasien(*A, cari_id(*A, *N))
	}
	fmt.Print("Pilih opsi edit: ")
	fmt.Println("1. Edit Nama")
	fmt.Println("2. Edit ID")
	fmt.Println("3. Edit Waktu Masuk")
	fmt.Println("4. Edit Layanan")
	fmt.Scan(&opsi)
	if opsi == 1 {
		display_pasien(*A, cari_nama(*A, *N))
	} else if opsi == 2 {
		display_pasien(*A, cari_id(*A, *N))
	} else if opsi == 3 {

	} else if opsi == 4 {

	}
}
func edit(A *tabInt, N, inx int) {

}
func remove(A *tabInt, N *int, idx int) {

	if idx == -1 {
		fmt.Println("Data Tidak Ditemukan")
	} else {
		i := idx
		for i < *N {
			A[i] = A[i+1]
			i++
		}
		*N--
	}
}
func cari_nama(A tabInt, N int) int {
	var name string
	var idx int = -1
	fmt.Print("Masukkan nama yang akan dicari: ")
	fmt.Scan(&name)
	i := 0
	for i < N && idx == -1 {
		if A[i].nama == name {
			idx = i
		}
		i++
	}
	return idx
}
func cari_id(A tabInt, N int) int {
	var id int
	var idx int = -1
	fmt.Print("Masukkan nama yang akan dicari: ")
	fmt.Scan(&id)
	i := 0
	for i < N && idx == -1 {
		if A[i].id == id {
			idx = i
		}
		i++
	}
	return idx
}
func display_pasien(A tabInt, idx int) {
	if idx != -1 {
		fmt.Println("Nama :", A[idx].nama)
		fmt.Println("ID :", A[idx].id)
		fmt.Println("Layanan :", A[idx].layanan)
		fmt.Println("Tahun :", A[idx].waktu.tahun)
		fmt.Println("Bulan :", A[idx].waktu.bulan)
		fmt.Println("Tanggal :", A[idx].waktu.tanggal)
	} else {
		fmt.Println("Data Pasien Tidak Ditemukan")
	}
}
