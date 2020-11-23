package service

import (
	"bytes"
	"fmt"
	"github.com/SebastiaanKlippert/go-wkhtmltopdf"
	log "github.com/sirupsen/logrus"
	"go-pdf-poc/dao"
	"go-pdf-poc/model"
	"html/template"
	"io/ioutil"
	"os"
	"strconv"
	"time"
//	"github.com/go-resty/resty/v2"

)

var PrescriptionDao dao.PrescriptionDaoImpl
var SequenceDao dao.SequenceDaoImpl

type PrescriptionServiceImpl struct {
}

func (i PrescriptionServiceImpl) Save(prescription model.Prescription) (*model.Prescription, error) {
	save, err := PrescriptionDao.Save(prescription)
	if err != nil {
		return nil, err
	}
	//log.Println(save)
	ok, err := i.GeneratePdf(prescription)
	if err != nil {
		return nil, err
	}
	log.Println("pdf generaed %v",ok)
	return save, nil
}
func (i PrescriptionServiceImpl) UpdatePrescription(prescriptionRequest model.Prescription) (*model.Prescription, error) {
	prescriptionRequest.DateUpdated = time.Now()
	prescription, err := PrescriptionDao.Update(prescriptionRequest)
	if err != nil {
		return nil, err
	}
	ok, err := i.GeneratePdf(*prescription)
	if err != nil {
		return nil, err
	}
	log.Println("pdf generaed %v",ok)
	return prescription, nil
}
func (i PrescriptionServiceImpl) GetDraftByPatientId(patientID int64) (*model.Prescription, error) {
	prescription, err := PrescriptionDao.GetAllDraftByPatientId(patientID)
	if err != nil {
		return nil, err
	}
	return prescription, nil
}
func (i PrescriptionServiceImpl) GetPrescriptionsByPatientId(patientID int64) (*model.Prescription, error) {
	prescription, err := PrescriptionDao.GetAllDraftByPatientId(patientID)
	if err != nil {
		return nil, err
	}
	return prescription, nil
}
func (i PrescriptionServiceImpl) GeneratePdf(prescription model.Prescription) (bool, error) {

	t, _ := template.ParseFiles("./utils/htmlsample.html")
	file, err := os.Create("./utils/test.html")
	err = t.Execute(file, prescription)
	if err != nil {
		log.Println(err)
		return false, err
	}
	// Create new PDF generator
	pdfg, err := wkhtmltopdf.NewPDFGenerator()
	if err != nil {
		log.Println(err)
		return false, err
	}

	// Add one page from an URL
	htmlfile, err := ioutil.ReadFile("./utils/test.html")

	if err != nil {
		log.Println(err)
		return false, err
	}
	page:=wkhtmltopdf.NewPageReader(bytes.NewReader(htmlfile))
	//page.FooterSpacing.Set(1.5)
	page.FooterFontSize.Set(20)
	page.FooterSpacing.Set(5)
	//page.FooterLeft.Set("\n")
	page.FooterRight.Set("| [page]")

	page.FooterFontSize.Set(6)
	page.FooterLeft.Set("The scope of this document is limited to the information provided by the patient to the Zyla team over the phone or Zyla app and should not be used for medico legal purposes.")
	//page.FooterLeft.Set("")
	//page.FooterLeft.Parse()
	//page.FooterHTML.Set("./utils/footer.html")
	page.FooterLine.Set(true)
	pdfg.AddPage(page)
	pdfg.MarginBottom.Set(15)
	pdfg.Dpi.Set(300)

	// Create PDF document in internal buffer
	err = pdfg.Create()
	if err != nil {
		log.Println(err)
		return false, err
	}
	sequence, err := SequenceDao.GetSequence()
	log.Println(&sequence)
	log.Println(strconv.FormatInt(*sequence,10))
	// Write buffer contents to file on disk
	err = pdfg.WriteFile("./sample.pdf")
	if err != nil {
		log.Println(err)
		return false, err
	}
	fmt.Println("Done")
	return true,nil
}
func (i PrescriptionServiceImpl) UpdateTracking(prescription model.Prescription) (bool, error) {
	//url := os.Getenv("SERVICE_PA_POST_URL")
	//
	//restClient := resty.New()
	//res, err := restClient.R().
	//	SetHeader("Content-Type", "application/json").
	//	SetHeader("Accept", "application/json").
	//	SetBody(json).
	//	SetResult(&response).Post(authenticationUrl.String())
	//if err != nil {
	//	log.Println(err)
	//	return nil, err
	//} else {
	//	log.Println(res)
	//}
return true,nil
}
func (i PrescriptionServiceImpl) UploadPdf(prescription model.Prescription) (bool, error) {
	return true,nil
}