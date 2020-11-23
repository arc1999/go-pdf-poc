package service

import "go-pdf-poc/model"

type PrescriptionService interface {
	GeneratePdf(prescription model.Prescription)
}
