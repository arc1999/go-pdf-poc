package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type Medicine struct {
	MedicineName string        `json:"medicineName"`
	Frequency    FrequencyMed  `json:"frequency"`
	Instruction  CodeValuePair `json:"instruction"`
	Duration     CodeValuePair `json:"duration"`
	Formulation  CodeValuePair  `json:"formulation"`
}
type FrequencyMed struct {
	Morning   string `json:"morning"`
	Afternoon string `json:"afternoon"`
	Evening   string `json:"evening"`
}
type PatientDetails struct {
	PatientID            int64     `json:"patientId"`
	Name                 string    `json:"name"`
	Age                  int       `json:"age"`
	Gender               int       `json:"gender"`    //:”1 female”||”2 male”
	HeightUOM            string    `json:"heightUom"` //cms/ft
	Height               string    `json:"height"`    //”str”,
	Weight               float64   `json:"weight"`
	DateOfWeightRecorded time.Time `json:"dateOfWeightRecorded"`
	Bmi                  float64   `json:"bmi"`
}
type DoctorDetails struct {
	Name          string `json:"name"`
	Qualification string `json:"qualification"`
	Experience    string `json:"experience"`
	LicenseNumber string `json:"licenseNumber"`
	Address       string `json:"address"`
}
type Report struct {
	ReportName   string        `json:"reportName"`
	DateRecorded time.Time     `json:"dateRecorded"`
	Value        float64       `json:"value"`
	ValueUom     CodeValuePair `json:"valueUom"`
}
type CodeValuePair struct {
	Label string `json:"label"`
	Value string `json:"value"`
}
type Investigation struct {
	TestName    CodeValuePair `json:"testName"`
	Instruction string        `json:"instruction"`
}
type Prescription struct {
	ID               primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	DoctorDetails    DoctorDetails      `json:"doctorDetails"`
	PatientDetails   PatientDetails     `json:"patientDetails"`
	ConsultationDate time.Time          `json:"consultationDate"`
	Diagnoses        []CodeValuePair    `json:"diagnoses"`
	ChiefComplaints  []CodeValuePair    `json:"chiefComplaints"`
	Reports          []Report           `json:"reports"`
	Medicines        []Medicine         `json:"medicines"`
	Investigations   []Investigation    `json:"investigations"`
	Instructions     []string           `json:"instructions"`
	DateCreated      time.Time          `json:"dateCreated"`
	DateUpdated      time.Time          `json:"dateUpdated"`
	Status           bool               `json:"status"`
	PdfUrl 			*string              `json:"pdfUrl"`
	Etag 			string				`json:"etag"`
}
