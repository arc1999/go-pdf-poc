package controller

import (
	"errors"
	"github.com/labstack/echo/v4"
	log "github.com/sirupsen/logrus"
	"go-pdf-poc/beans"
	"go-pdf-poc/model"
	"go-pdf-poc/service"
	"go-pdf-poc/transformer"
	"net/http"
	"strconv"
	"strings"
)

var PrescriptionService service.PrescriptionServiceImpl

type PrescriptionController struct {
}

func (i PrescriptionController) Create() echo.HandlerFunc {
	return func(c echo.Context) (err error) {
		request := new(beans.PrescriptionRequest)
		if err := c.Bind(request); err != nil {
			return handleError(err, http.StatusUnprocessableEntity, "Error Fetching Request Body")
		}
		input := transformer.TransformPrescriptionRequest(*request)
		if request.IsMedicineChanged {
			ok, err := PrescriptionService.UpdateTracking(*input)
			if err != nil {
				log.Println("Syncing failed")
				log.Println(err)
				return handleError(err,http.StatusInternalServerError,err.Error())
			}
			log.Println("Syncing Status %v",ok)
		}
		prescription, err := PrescriptionService.Save(*input)
		if err != nil {
			return handleError(err, http.StatusInternalServerError, err.Error())
		}
		//http.ServeFile(c.Response(pubg).Writer,c.Request(),"./sample.pdf")
		return c.JSON(http.StatusOK, prescription)
	}
}
func (i PrescriptionController) Update() echo.HandlerFunc {
	return func(c echo.Context) (err error) {
		request := new(beans.PrescriptionRequest)
		if err := c.Bind(request); err != nil {
			return handleError(err, http.StatusUnprocessableEntity, "Error Fetching Request Body")
		}
		input := transformer.TransformPrescriptionRequest(*request)
		if request.Status {
			if request.IsMedicineChanged{
				ok, err := PrescriptionService.UpdateTracking(*input)
				if err != nil {
					log.Println("Syncing failed")
					log.Println(err)
					return handleError(err,http.StatusInternalServerError,err.Error())
				}
				log.Println("Syncing Status %v",ok)
			}
		}
		updated, err := PrescriptionService.UpdatePrescription(*input)
		if err != nil {
			return handleError(err, http.StatusInternalServerError, err.Error())
		}
		//	http.ServeFile(c.Response().Writer,c.Request(),"./sample.pdf")
		return c.JSON(http.StatusOK, updated)
	}
}
func (i PrescriptionController) GetByPatientID() echo.HandlerFunc {
	return func(c echo.Context) (err error) {
		patientID := c.Param("patientId")
		if len(strings.TrimSpace(patientID)) == 0 {
			return handleError(errors.New("patientID is null"), http.StatusBadRequest, "patientID cannot be null")
		}
		id,err := strconv.ParseInt(patientID,10,64)
		if err != nil {
			return handleError(err, http.StatusBadRequest, err.Error())
		}
		prescription, err := PrescriptionService.GetDraftByPatientId(id)
		if err != nil {
			return handleError(err, http.StatusInternalServerError, err.Error())
		}
		return c.JSON(http.StatusOK, prescription)
	}
}
func handleError(err error, code int, message string) error {
	log.Error(err)
	appError := model.ApplicationError{Code: code, Message: message}
	log.Error(appError)
	return echo.NewHTTPError(code, appError)
}
