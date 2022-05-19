package main

import "fmt"

type mahasiswa struct {
	nama string
}

func (m mahasiswa) getNama() string {
	return m.nama
}

func (m mahasiswa) setNama(newNama string) {
	m.nama = newNama
}

type dosen struct {
	nip string
}

func (d dosen) getNip() string {
	return d.nip
}

func main() {
	mhs1 := mahasiswa{
		nama: "Koinworks",
	}

	dsn := dosen{
		nip: "111222333",
	}

	mhs1.setNama("Hacktiv8")

	fmt.Println(mhs1.getNama())
	fmt.Println(dsn.getNip())

}
