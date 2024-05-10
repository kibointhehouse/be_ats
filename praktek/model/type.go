package _714220052

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type NamaPemain struct {
	ID           primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	Nama         string             `bson:"nama,omitempty" json:"nama,omitempty"`
	Nomor_hp 	 string             `bson:"nomor_hp,omitempty" json:"nomor_hp,omitempty"`
	Posisi      string             `bson:"posisi,omitempty" json:"posisi,omitempty"`
	Umur 		string				`bson:"umur,omitempty" json:"umur,omitempty"`
	Kewarganegaraan 	string		`bson:"kewarganegaraan,omitempty" json:"kewarganegaraan,omitempty"`
}

type Statistik struct {
	Tinggi_badan     int      `bson:"tinggi_badan,omitempty" json:"tinggi_badan,omitempty"`
	berat_badan  string   `bson:"berat_badan,omitempty" json:"berat_badan,omitempty"`
	Point_per_game string   `bson:"point_per_game,omitempty" json:"point_per_game,omitempty"`
	Assist_per_game      string      `bson:"assist_per_game,omitempty" json:"assist_per_game,omitempty"`
	Rebound_per_game     string `bson:"rebound_per_game,omitempty" json:"hari,omitempty"`
}

type Presensi struct {
	ID           primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	Tempat_Tinggal     string             `bson:"tempat_tinggal,omitempty" json:"tempat_tinggal,omitempty"`
	Nomor_hp	 string             `bson:"nomor_hp,omitempty" json:"nomor_hp,omitempty"`
	Waktu_Presensi     primitive.DateTime `bson:"waktu_presensi,omitempty" json:"waktu_presensi,omitempty"`
	Checkin      string             `bson:"checkin,omitempty" json:"checkin,omitempty"`
	Biodata      NamaPemain           `bson:"biodata,omitempty" json:"biodata,omitempty"`
}
