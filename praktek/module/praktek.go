package _714220052

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var MongoString string = os.Getenv("MONGOSTRING")

func MongoConnect(dbname string) (db *mongo.Database) {
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(MongoString))
	if err != nil {
		fmt.Printf("MongoConnect: %v\n", err)
	}
	return client.Database(dbname)
}

func InsertOneDoc(db string, collection string, doc interface{}) (insertedID interface{}) {
	insertResult, err := MongoConnect(db).Collection(collection).InsertOne(context.TODO(), doc)
	if err != nil {
		fmt.Printf("InsertOneDoc: %v\n", err)
	}
	return insertResult.InsertedID
}

func InsertPresensi( tempattinggal string, nomorhp string, checkin string,  biodata NamaPemain) (InsertedID interface{}) {
	var presensi Presensi
	presensi.Tempat_Tinggal = tempattinggal
	presensi.Nomor_hp = nomorhp
	presensi.Waktu_Presensi = primitive.NewDateTimeFromTime(time.Now().UTC())
	presensi.Checkin = checkin
	presensi.Biodata = biodata
	return InsertOneDoc("tesdb2024", "presensi", presensi)
}

func GetNamaPemainFromNomor_Hp(nomor_hp string) (staf Presensi) {
	namapemain := MongoConnect("tesdb2024").Collection("presensi")
	filter := bson.M{"nomor_hp": nomor_hp}
	err := namapemain.FindOne(context.TODO(), filter).Decode(&staf)
	if err != nil {
		fmt.Printf("getNamaPemainFromNomor_Hp: %v\n", err)
	}
	return staf
}

func GetAllPresensi() (data []Presensi) {
	namapemain := MongoConnect("tesdb2024").Collection("presensi")
	filter := bson.M{}
	cursor, err := namapemain.Find(context.TODO(), filter)
	if err != nil {
		fmt.Println("GetALLData :", err)
	}
	err = cursor.All(context.TODO(), &data)
	if err != nil {
		fmt.Println(err)
	}
	return
}
