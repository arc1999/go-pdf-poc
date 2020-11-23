package dao

import (
	"context"
	"errors"
	log "github.com/sirupsen/logrus"
	"go-pdf-poc/db"
	"go-pdf-poc/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"os"
	"strconv"
	"time"
)

type PrescriptionDaoImpl struct {
}

func (mi PrescriptionDaoImpl) Save(request model.Prescription) (*model.Prescription, error) {
	db := db.GetDb()
	request.Etag= strconv.FormatInt(time.Now().UnixNano()/int64(time.Millisecond), 10)
	//log.Println(time.Now().UnixNano() / int64(time.Millisecond))
	result, err := db.Collection(os.Getenv("PRESCRIPTION_COLLECTION")).InsertOne(context.TODO(), request)
	if err != nil {
		log.Printf("error in saving Prescription")
		return nil, err
	}
	return mi.FindById(result.InsertedID.(primitive.ObjectID))
}
func (mi PrescriptionDaoImpl) FindById(prescriptionId primitive.ObjectID) (*model.Prescription, error) {
	db := db.GetDb()
	var prescription model.Prescription
	result := db.Collection(os.Getenv("PRESCRIPTION_COLLECTION")).FindOne(context.TODO(), bson.M{"_id": prescriptionId})
	if result.Err() != nil {
		log.Printf("error in Finding Prescription")
		return nil, result.Err()
	}
	err := result.Decode(&prescription)
	if err != nil {
		return nil, err
	}
	return &prescription, nil
}
func (mi PrescriptionDaoImpl) Update(prescription model.Prescription) (*model.Prescription, error) {
	db := db.GetDb()
	log.Println(prescription.ID)
	var presc model.Prescription
	err := db.Collection(os.Getenv("PRESCRIPTION_COLLECTION")).FindOne(context.TODO(), bson.M{"_id": prescription.ID}).Decode(&presc)
	if err != nil {
		log.Printf("error in Finding Prescription")
		return nil, err
	}
	log.Println(presc.Etag)
	if prescription.Etag!=presc.Etag{
		log.Println(prescription.Etag)
		log.Println(presc.Etag)
		return nil,errors.New("Etag Mismatch")
	}
	prescription.Etag=strconv.FormatInt(time.Now().UnixNano()/int64(time.Millisecond), 10)
	result, err := db.Collection(os.Getenv("PRESCRIPTION_COLLECTION")).ReplaceOne(context.TODO(), bson.M{"_id": prescription.ID}, prescription)
	if err != nil {
		log.Printf("error in Finding Prescription")
		return nil, err
	}
	//err := result.Decode(&prescription)
	log.Println(result)
	return &prescription, nil
}
func (mi PrescriptionDaoImpl) GetAllDraftByPatientId(patientId int64) (*model.Prescription, error) {
	db := db.GetDb()
	//log.Println(prescription.ID)
	var prescription model.Prescription
	err := db.Collection(os.Getenv("PRESCRIPTION_COLLECTION")).FindOne(context.TODO(),bson.D{{"patientdetails.patientid",patientId},{"status",false}}).Decode(&prescription)
	if err != nil {
		log.Printf("error in Finding Prescription")
		return nil, err
	}
	//err := result.Decode(&prescription)
	log.Println(prescription)
	return &prescription, nil
}
func (mi PrescriptionDaoImpl) GetAllPrescriptionByPatientId(patientId int64) ([]model.Prescription, error) {
	db := db.GetDb()
	//log.Println(prescription.ID)
	cur,err := db.Collection(os.Getenv("PRESCRIPTION_COLLECTION")).Find(context.TODO(),bson.D{{"patientdetails.patientid",patientId},{"status",true}})
	if err != nil {
		return nil, err
	}
	defer cur.Close(context.TODO())
	var prescriptions []model.Prescription
	for cur.Next(context.TODO()) {
		var prescription model.Prescription

		err := cur.Decode(&prescription)
		if err != nil {
			return prescriptions, nil
		}
		prescriptions = append(prescriptions, prescription)
	}
		//err := result.Decode(&prescription)
	return prescriptions, nil
}