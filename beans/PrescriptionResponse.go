package beans

import (
	"go-pdf-poc/model"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type PrescriptionResponse struct {
	ID               primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	DoctorDetails    model.DoctorDetails      `json:"doctorDetails"`
	PatientDetails   model.PatientDetails     `json:"patientDetails"`
	ConsultationDate time.Time          `json:"consultationDate"`
	Diagnoses        []model.CodeValuePair    `json:"diagnoses"`
	ChiefComplaints  []model.CodeValuePair    `json:"chiefComplaints"`
	Reports          []model.Report           `json:"reports"`
	Medicines        []model.Medicine         `json:"medicines"`
	Investigations   []model.Investigation    `json:"investigations"`
	Instructions     []string           `json:"instructions"`
	DateCreated      time.Time          `json:"dateCreated"`
	DateUpdated      time.Time          `json:"dateUpdated"`
	Status           bool               `json:"status"`
	PdfUrl 			*string              `json:"pdfUrl"`
	Etag 			string					`json:"etag"`

}

