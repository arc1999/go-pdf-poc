package transformer

import (
	"go-pdf-poc/beans"
	"go-pdf-poc/model"
)

func TransformPrescriptionRequest(input beans.PrescriptionRequest) *model.Prescription {
	var output model.Prescription
	output.Medicines=input.Medicines
	output.PatientDetails=input.PatientDetails
	output.DateCreated=input.DateCreated
	output.ID=input.ID
	output.DoctorDetails=input.DoctorDetails
	output.Status=input.Status
	output.ChiefComplaints=input.ChiefComplaints
	output.ConsultationDate=input.ConsultationDate
	output.Diagnoses=input.Diagnoses
	output.Instructions=input.Instructions
	output.Investigations=input.Investigations
	output.Reports=input.Reports
	output.DateUpdated=input.DateUpdated
	output.Etag=input.Etag
	return &output
}
func TransformPrescription(input model.Prescription) *beans.PrescriptionResponse{
	var output beans.PrescriptionResponse
	output.Medicines=input.Medicines
	output.PatientDetails=input.PatientDetails
	output.DateCreated=input.DateCreated
	output.ID=input.ID
	output.DoctorDetails=input.DoctorDetails
	output.Status=input.Status
	output.ChiefComplaints=input.ChiefComplaints
	output.ConsultationDate=input.ConsultationDate
	output.Diagnoses=input.Diagnoses
	output.Instructions=input.Instructions
	output.Investigations=input.Investigations
	output.Reports=input.Reports
	output.PdfUrl=input.PdfUrl
	output.DateUpdated=input.DateUpdated
	output.Etag=input.Etag
	return &output
}
