package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"

	"github.com/olekukonko/tablewriter"
)

// data_base
type Mahasiswa struct {
	Npm      string
	Nama     string
	Jurusan  string
	Angkatan int
}

type Head_Mhs struct {
	data Mahasiswa
	next *Head_Mhs
}

type Dosen struct {
	Nip  int
	Nama string
}

type Head_Dsn struct {
	data Dosen
	next *Head_Dsn
}

type Peserta_Umum struct {
	No_pendaftaran int
	Nama           string
	No_hp          int
	Email          string
	Gender         string
	Asal_pt        string
}

type Head_PU struct {
	data Peserta_Umum
	next *Head_PU
}

// model / fungsi MAHASISWA
func data_mahasiswa_statis(mahasiswa *Head_Mhs) {
	data_mhs1 := Mahasiswa{Npm: "06.2022.1.07626", Nama: "Gregoria Stefania K. Siga", Jurusan: "TI", Angkatan: 22}
	data_mhs2 := Mahasiswa{Npm: "06.2022.1.07599", Nama: "Annas Tasya E.A.J", Jurusan: "SI", Angkatan: 22}
	data_mhs3 := Mahasiswa{Npm: "06.2022.1.07616", Nama: "Putri S. Syam", Jurusan: "GI", Angkatan: 22}
	insert_db_mhs(mahasiswa, data_mhs1)
	insert_db_mhs(mahasiswa, data_mhs2)
	insert_db_mhs(mahasiswa, data_mhs3)
}

func insert_db_mhs(head_mhs *Head_Mhs, mahasiswa Mahasiswa) {
	New_Head := &Head_Mhs{}
	New_Head.data = mahasiswa
	temp := head_mhs
	if temp.next == nil {
		temp.next = New_Head
	} else {
		temp := head_mhs
		for temp.next != nil {
			temp = temp.next
		}
		temp.next = New_Head
	}
}

func MenuInsertMhs(dataMhs *Head_Mhs) {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Anda masuk pilihan 2: INSERT")
	var jurusan string
	var npm string
	var nama string
	var angkatan int

	fmt.Print("masukkan nama: ")
	scanner.Scan()
	nama = scanner.Text()

	fmt.Print("masukkan NPM: ")
	fmt.Scan(&npm)
	scanner.Scan()

	fmt.Print("masukkan nama jurusan: ")
	scanner.Scan()
	jurusan = scanner.Text()

	fmt.Print("masukkan Angkatan: ")
	scanner.Scan()
	temp := scanner.Text()
	angkatan, _ = strconv.Atoi(temp)

	data := Mahasiswa{
		Nama:     nama,
		Npm:      npm,
		Jurusan:  jurusan,
		Angkatan: angkatan,
	}
	insert_db_mhs(dataMhs, data)
}

func display_mhs(head_mhs *Head_Mhs) {

	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Nama", "NPM", "Jurusan", "Angkatan"})

	temp := head_mhs
	for temp.next != nil {
		mhs := temp.next.data
		row := []string{mhs.Nama, mhs.Npm, mhs.Jurusan, fmt.Sprint(mhs.Angkatan)}
		table.Append(row)

		temp = temp.next
	}
	table.Render()
}

func Update(g *Head_Mhs, npm string) {
	temp := g

	if temp.next == nil {
		fmt.Println("Data Kosong")
		return
	} else {
		for temp.next != nil {
			if temp.next.data.Npm == npm {
				fmt.Print("Masukkan nama baru : ")
				fmt.Scan(&temp.next.data.Nama)
				fmt.Print("Masukkan jurusan baru : ")
				fmt.Scan(&temp.next.data.Jurusan)
				fmt.Print("Masukkan angkatan baru : ")
				fmt.Scan(&temp.next.data.Angkatan)
				fmt.Print("=========================")
				fmt.Println("Data berhasil diupdate")
				fmt.Print("=========================")
				return
			}
			temp = temp.next
		}
	}
	fmt.Print("============================================")
	fmt.Printf("Data dengan NPM %s tidak ditemukan\n", npm)
	fmt.Print("============================================")
}

func MenuUpdateMhs(dataMhs *Head_Mhs) {
	var npm string
	fmt.Println("Anda masuk pilihan 2: UPDATE")
	fmt.Print("Masukkan NPM Mahasiswa yang ingin diupdate: ")
	fmt.Scan(&npm)
	Update(dataMhs, npm)
}

func Delete(g *Head_Mhs, nip string) {
	temp := g
	if temp.next == nil {
		fmt.Println("Data Kosong")
		return
	}

	if temp.data.Npm == nip {
		fmt.Println("data pertama di delete")
		*temp = *temp.next
		return
	}

	for temp.next != nil {
		if temp.next.data.Npm == nip {
			temp.next = temp.next.next
			return
		}
		temp = temp.next
	}
	fmt.Print("====================================")
	fmt.Printf("Delete gagal, id %s tidak ada", nip)
	fmt.Print("====================================")

}

func MenuDeletetMhs(dataMhs *Head_Mhs) {
	var npm string
	fmt.Print("==============================================")
	fmt.Printf("Masukan NPM Mahasiswa yang ingin di Delete : ")
	fmt.Print("==============================================")
	fmt.Scan(&npm)
	Delete(dataMhs, npm)
}

func GetListMahasiswa(g *Head_Mhs, npm string) {
	current := g

	for current.next != nil {
		if current.next.data.Npm == npm {
			mhs := current.next.data
			// Inisialisasi tabel dan header kolom
			table := tablewriter.NewWriter(os.Stdout)
			table.SetHeader([]string{"Nama", "NPM", "Jurusan", "Angkatan"})
			row := []string{mhs.Nama, mhs.Npm, mhs.Jurusan, fmt.Sprint(mhs.Angkatan)}
			table.Append(row)
			// Render tabel ke terminal
			table.Render()
			return
		}
		current = current.next
	}
	fmt.Print("=====================================")
	fmt.Printf("data dengan NPM %s tidak ada\n", npm)
	fmt.Print("=====================================")
}

func MenuViewByNpm(dataMhs *Head_Mhs) {
	var npm string
	fmt.Println("anda masuk pilihan 5")
	fmt.Print("masukkan npm : ")
	fmt.Scan(&npm)
	GetListMahasiswa(dataMhs, npm)
}

// fungsi PESERTA UMUM
func insert_db_pu(head_pu *Head_PU, pu Peserta_Umum) {
	New_Head := &Head_PU{}
	New_Head.data = pu
	temp := head_pu
	if temp == nil {
		temp.next = New_Head
	} else {
		for temp.next != nil {
			temp = temp.next
		}
		temp.next = New_Head
	}

}

func MenuInsertPU(dataPU *Head_PU) {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Anda masuk pilihan 2: INSERT")
	var no_pendaftaran int
	var nama string
	var no_hp int
	var email string
	var gender string
	var asal_pt string

	fmt.Print("masukkan no Pendaftaran : ")
	scanner.Scan()
	temp := scanner.Text()
	no_pendaftaran, _ = strconv.Atoi(temp)

	fmt.Print("masukkan nama : ")
	scanner.Scan()
	nama = scanner.Text()

	fmt.Print("masukkan No Telp: ")
	scanner.Scan()
	temp = scanner.Text()
	no_hp, _ = strconv.Atoi(temp)

	fmt.Print("masukkan Email : ")
	scanner.Scan()
	email = scanner.Text()

	fmt.Print("masukkan Gender : ")
	scanner.Scan()
	gender = scanner.Text()

	fmt.Print("masukkan Asal Perguruan : ")
	scanner.Scan()
	asal_pt = scanner.Text()

	data := Peserta_Umum{
		No_pendaftaran: no_pendaftaran,
		Nama:           nama,
		No_hp:          no_hp,
		Email:          email,
		Gender:         gender,
		Asal_pt:        asal_pt,
	}
	insert_db_pu(dataPU, data)
}

func UpdatePU(g *Head_PU, no_pendaftaran int) {
	temp := g

	if temp.next == nil {
		fmt.Println("Data Kosong")
		return
	} else {
		for temp.next != nil {
			if temp.next.data.No_pendaftaran == no_pendaftaran {
				fmt.Print("Masukan nama baru : ")
				fmt.Scan(&temp.next.data.Nama)
				fmt.Print("Masukan no hp baru : ")
				fmt.Scan(&temp.next.data.No_hp)
				fmt.Print("Masukan email baru : ")
				fmt.Scan(&temp.next.data.Email)
				fmt.Print("Masukan gender baru : ")
				fmt.Scan(&temp.next.data.Gender)
				fmt.Print("Masukan asal pt baru : ")
				fmt.Scan(&temp.next.data.Asal_pt)
				fmt.Println("Data berhasil diupdate")
				return
			}
			temp = temp.next
		}
	}
	fmt.Print("==============================================")
	fmt.Printf("data dengan no %d tidak ada\n", no_pendaftaran)
	fmt.Print("==============================================")
}

func MenuUpdatePU(dataPU *Head_PU) {
	var no_pendaftaran int
	fmt.Println("anda masuk pilihan 2: UPDATE")
	fmt.Println("Masukan No Pendaftaran : ")
	fmt.Scan(&no_pendaftaran)

	UpdatePU(dataPU, no_pendaftaran)
}

func DeletePU(g *Head_PU, no_dftr int) {
	temp := g
	if temp.next == nil {
		fmt.Println("Data Kosong")
		return
	}

	if temp.data.No_pendaftaran == no_dftr {
		fmt.Println("data pertama di delete")
		*temp = *temp.next
		return
	}

	for temp.next != nil {
		if temp.next.data.No_pendaftaran == no_dftr {
			temp.next = temp.next.next
			return
		}
		temp = temp.next
	}
	fmt.Println("====================================")
	fmt.Printf("Delete gagal, id %d tidak ada", no_dftr)
	fmt.Println("====================================")

}

func MenuDeletePU(head *Head_PU) {
	var no_dftr int
	prevNode := head
	currentNode := head.next

	fmt.Print("Masukkan NIP dosen yang akan dihapus: ")
	fmt.Scan(&no_dftr)

	for currentNode != nil {
		if currentNode.data.No_pendaftaran == no_dftr {
			prevNode.next = currentNode.next
			fmt.Println("==============================================================")
			fmt.Printf("Data dengan nomor pendaftaran %d berhasil dihapus\n", no_dftr)
			fmt.Println("==============================================================")
			return
		}
	}
	DeletePU(head, no_dftr)
}

//	prevNode = currentNode
//	currentNode = currentNode.next
//	}

//	fmt.Printf("Data dengan nomor pendaftaran %d tidak ditemukan\n", no_pendaftaran)
//}

//	func menuDeletePeserta(head *Head_PU) {
// 	var noPendaftaran int

//  	fmt.Println("Anda masuk pilihan 4: DELETE")
//	fmt.Print("Masukkan nomor pendaftaran peserta: ")
//	fmt.Scan(&noPendaftaran)

//  	deletePeserta(head, noPendaftaran)
//}

// }

func display_peserta(head_pu *Head_PU) {
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"No Pendaftaran", "Nama", "No Telp", "Email", "Jenis Kelamin", "Asal Universitas"})

	temp := head_pu
	for temp.next != nil {
		pu := temp.next.data
		row := []string{fmt.Sprint(pu.No_pendaftaran), pu.Nama, fmt.Sprint(pu.No_hp), pu.Email, pu.Gender, pu.Asal_pt}
		table.Append(row)

		temp = temp.next
	}
	table.Render()
}

// Fungsi DOSEN
// func insert_db_dsn(head_dsn *Head_Dsn, dosen Dosen) {
// 	New_Head := &Head_Dsn{}
// 	New_Head.data = dosen
// 	temp := head_dsn
// 	if temp == nil {
// 		temp.next = New_Head
// 	} else {
// 		for temp.next != nil {
// 			temp = temp.next
// 		}
// 		temp.next = New_Head
// 	}
// }

func dataStatis(dosen *Head_Dsn) {
	data_dsn1 := Dosen{Nip: 0213, Nama: "Gregoria Stefania K. Siga"}
	data_dsn2 := Dosen{Nip: 5674, Nama: "Annas Tasya E.A.J"}
	data_dsn3 := Dosen{Nip: 9812, Nama: "Putri S. Syam"}
	insert_db_dsn(dosen, data_dsn1)
	insert_db_dsn(dosen, data_dsn2)
	insert_db_dsn(dosen, data_dsn3)
}

func display_dsn(head_dsn *Head_Dsn) {

	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"NIP", "Nama"})

	temp := head_dsn
	for temp.next != nil {
		dsn := temp.next.data
		row := []string{fmt.Sprint(dsn.Nip), dsn.Nama}
		table.Append(row)

		temp = temp.next
	}
	table.Render()
}
func insert_db_dsn(head_dsn *Head_Dsn, dosen Dosen) {
	New_Head := &Head_Dsn{}
	New_Head.data = dosen
	temp := head_dsn
	if temp.next == nil {
		temp.next = New_Head
	} else {
		temp := head_dsn
		for temp.next != nil {
			temp = temp.next
		}
		temp.next = New_Head
	}
}

func MenuInsertDsn(dataDsn *Head_Dsn) {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Anda masuk pilihan 2: INSERT")

	var nama string
	var nip int

	fmt.Print("masukkan nama: ")
	scanner.Scan()
	nama = scanner.Text()

	fmt.Print("masukkan NIP: ")
	scanner.Scan()
	temp := scanner.Text()
	nip, _ = strconv.Atoi(temp)

	data := Dosen{
		Nama: nama,
		Nip:  nip,
	}
	insert_db_dsn(dataDsn, data)
}

func UpdateDosen(g *Head_Dsn, nip int) {
	temp := g

	if temp.next == nil {
		fmt.Println("Data Kosong")
		return
	} else {
		for temp.next != nil {
			if temp.next.data.Nip == nip {
				fmt.Print("Masukan nama baru : ")
				fmt.Scan(&temp.next.data.Nama)
				return
			}
			temp = temp.next
		}
	}
	fmt.Println("=====================================")
	fmt.Printf("data dengan nip %d tidak ada\n", nip)
	fmt.Println("=====================================")
}

func MenuUpdateDsn(dataDsn *Head_Dsn) {
	var nip int
	fmt.Println("anda masuk pilihan 2: UPDATE")
	fmt.Println("Masukan NIP : ")
	fmt.Scan(&nip)

	UpdateDosen(dataDsn, nip)
}

func deleteDosen(head *Head_Dsn, nip int) {
	if head == nil || head.next == nil {
		fmt.Println("Data kosong")
		return
	}

	curr := head.next
	prev := head

	for curr != nil {
		if curr.data.Nip == nip {
			prev.next = curr.next
			fmt.Println("==========================================")
			fmt.Printf("Data dengan NIP %d berhasil dihapus\n", nip)
			fmt.Println("==========================================")
			return
		}

		prev = curr
		curr = curr.next
	}
	fmt.Println("===========================================")
	fmt.Printf("Data dengan NIP %d tidak ditemukan\n", nip)
	fmt.Println("===========================================")
}

func MenuDeletetDsn(head *Head_Dsn) {
	var nip int
	fmt.Print("Masukkan NIP dosen yang akan dihapus: ")
	fmt.Scan(&nip)
	deleteDosen(head, nip)
}

func MenuViewByNo(Head_PU *Head_PU) {
	var no_dftr int
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Anda masuk pilihan 3: VIEW DATA PESERTA BY NO PENDAAFTARAN")
	fmt.Print("Masukkan NO PENDAAFTARAN : ")
	scanner.Scan()
	temp := scanner.Text()
	no_dftr, _ = strconv.Atoi(temp)

	current := Head_PU

	for current.next != nil {
		if current.next.data.No_pendaftaran == no_dftr {
			pu := current.next.data
			// Inisialisasi tabel dan header kolom
			table := tablewriter.NewWriter(os.Stdout)
			table.SetHeader([]string{"No Pendaftaran", "Nama", "No Telp", "Email", "Jenis Kelamin", "Asal Universitas"})
			row := []string{fmt.Sprint(pu.No_pendaftaran), pu.Nama, fmt.Sprint(pu.No_hp), pu.Email, pu.Gender, pu.Asal_pt}
			table.Append(row)
			// Render tabel ke terminal
			table.Render()
			return
		}
		current = current.next
	}
	fmt.Println("============================================")
	fmt.Printf("Data dengan NIP %d tidak ditemukan.\n", no_dftr)
	fmt.Println("============================================")
}

func MenuViewByNip(head_dsn *Head_Dsn) {
	var nip int
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Anda masuk pilihan 3: VIEW DATA DOSEN BY NIP")
	fmt.Print("Masukkan NIP Dosen : ")
	scanner.Scan()
	temp := scanner.Text()
	nip, _ = strconv.Atoi(temp)

	current := head_dsn

	for current.next != nil {
		if current.next.data.Nip == nip {
			dsn := current.next.data
			// Inisialisasi tabel dan header kolom
			table := tablewriter.NewWriter(os.Stdout)
			table.SetHeader([]string{"NIP", "Nama"})
			row := []string{fmt.Sprint(dsn.Nip), dsn.Nama}
			table.Append(row)
			// Render tabel ke terminal
			table.Render()
			return
		}
		current = current.next
	}
	fmt.Println("============================================")
	fmt.Printf("Data dengan NIP %d tidak ditemukan.\n", nip)
	fmt.Println("============================================")
}

// views
func Menu_utama() {
	fmt.Println("\t\t~PROGRAM KOORDINASI EVENT~")
	fmt.Println("Menu Utama :")
	fmt.Println("1. Peserta Umum")
	fmt.Println("2. Mahasiswa")
	fmt.Println("3. Dosen")
	fmt.Println("4. EXIT")
	fmt.Print("Masukan Pilihan Menu: ")
}

func Menu_peserta() {
	fmt.Println("\t\tMenu Kelola Data Peserta Umum")
	fmt.Println("1. View All Data")
	fmt.Println("2. Insert")
	fmt.Println("3. Update")
	fmt.Println("4. Delete")
	fmt.Println("5. View by No Peserta")
	fmt.Println("6. Back")
}

func Menu_mahasiswa() {
	fmt.Println("\t\tMenu Kelola Data Mahasiswa")
	fmt.Println("1. View All Data")
	fmt.Println("2. Insert")
	fmt.Println("3. Update")
	fmt.Println("4. Delete")
	fmt.Println("5. View by NPM")
	fmt.Println("6. Back")
}

func Menu_Dosen() {
	fmt.Println("\t\tMenu Kelola Data Dosen")
	fmt.Println("1. View All Data")
	fmt.Println("2. Insert")
	fmt.Println("3. Update")
	fmt.Println("4. Delete")
	fmt.Println("5. View by NIP")
	fmt.Println("6. Back")
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	var pilihan1 int = 0
	var pilih_sub1 int = 0
	var pilih_sub2 int = 0
	var pilih_sub3 int = 0
	dataMhs := Head_Mhs{}
	dataDsn := Head_Dsn{}
	dataPU := Head_PU{}
	dataStatis(&dataDsn)
	data_mahasiswa_statis(&dataMhs)

	for pilihan1 != 4 {
		Menu_utama()
		fmt.Scan(&pilihan1)
		scanner.Scan()

		switch pilihan1 {
		case 1:
			for pilih_sub1 != 6 {
				Menu_peserta()
				fmt.Print("Masukan Pilihan Menu: ")
				fmt.Scan(&pilih_sub1)
				scanner.Scan()
				switch pilih_sub1 {
				case 1:
					display_peserta(&dataPU)
				case 2:
					MenuInsertPU(&dataPU)
				case 3:
					MenuUpdatePU(&dataPU)
				case 4:
					MenuDeletePU(&dataPU)
				case 5:
					MenuViewByNo(&dataPU)
				case 6:
					break
				}
			}

		case 2:
			for pilih_sub2 != 6 {
				Menu_mahasiswa()
				fmt.Print("Masukan Pilihan Menu: ")
				fmt.Scan(&pilih_sub2)
				scanner.Scan()
				switch pilih_sub2 {
				case 1:
					display_mhs(&dataMhs)
				case 2:
					MenuInsertMhs(&dataMhs)
				case 3:
					MenuUpdateMhs(&dataMhs)
				case 4:
					MenuDeletetMhs(&dataMhs)
				case 5:
					MenuViewByNpm(&dataMhs)
				case 6:
					break
				}
			}

		case 3:
			for pilih_sub3 != 6 {
				Menu_Dosen()
				fmt.Print("Masukan Pilihan Menu: ")
				fmt.Scan(&pilih_sub3)
				scanner.Scan()
				switch pilih_sub3 {
				case 1:
					display_dsn(&dataDsn)
				case 2:
					MenuInsertDsn(&dataDsn)
				case 3:
					MenuUpdateDsn(&dataDsn)
				case 4:
					MenuDeletetDsn(&dataDsn)
				case 5:
					MenuViewByNip(&dataDsn)
				case 6:
					break
				}
			}
		}
	}

	fmt.Println("============================================")
	fmt.Println("|Terimakasih Telah Menggunakan Layanan Kami|")
	fmt.Println("============================================")
}
