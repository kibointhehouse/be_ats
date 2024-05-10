package _714220052

import (
	"fmt"
	"testing"
)

func TestInsertPresensi(t *testing.T) {
	tempattinggal := "Jl.Sariasih.54"
	nomorhp := "085781085560"
	checkin := "masuk"
	biodata := NamaPemain{
		Nama : "Lebron James",
		Nomor_hp : "085763681513",
		Posisi : "Small Forward",
		Umur : "38",
		Kewarganegaraan : "Amerika",
	}
	hasil:=InsertPresensi(tempattinggal , nomorhp , checkin , biodata )
	fmt.Println(hasil)
}

func TestGetNamaPemainFromNomor_Hp(t *testing.T) {
	nomorhp := "08564567987"
	biodata:=GetNamaPemainFromNomor_Hp(nomorhp)
	fmt.Println(biodata)
}

func TestGetAll(t *testing.T) {
	data := GetAllPresensi()
	fmt.Println(data)
}

