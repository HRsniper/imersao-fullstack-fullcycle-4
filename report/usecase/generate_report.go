package usecase

import (
	"encoding/json"
	"html/template"
	"os"
	"strconv"
	"time"

	"github.com/HRsniper/imersao-fullstack-fullcycle-4/report/dto"
	"github.com/HRsniper/imersao-fullstack-fullcycle-4/report/infra/kafka"
	"github.com/HRsniper/imersao-fullstack-fullcycle-4/report/infra/repository"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
)

func GenerateReport(requestJson []byte, repository repository.TransactionElasticRepository) error {
	var requestReport dto.RequestReport

	err := json.Unmarshal(requestJson, &requestReport)
	if err != nil {
		return err
	}

	data, err := repository.Search(requestReport.ReportID, requestReport.AccountID, requestReport.InitDate, requestReport.EndDate)
	if err != nil {
		return err
	}

	result, err := generateReportFile(data)
	if err != nil {
		return err
	}

	err = publishMessage(data.ReportID, string(result), "complete")
	if err != nil {
		return err
	}

	err = os.Remove("data/" + data.ReportID + ".html")
	if err != nil {
		return err
	}

	return nil
}

func generateReportFile(data dto.SearchResponse) ([]byte, error) {
	file, err := os.Create("data/" + data.ReportID + ".html")
	if err != nil {
		return nil, err
	}

	htmlTemplate := template.Must(template.New("report.html").ParseFiles("template/report.html"))
	err = htmlTemplate.Execute(file, data)
	if err != nil {
		return nil, err
	}

	result, err := uploadReport(data)
	if err != nil {
		return nil, err
	}

	return []byte(result), nil
}

func uploadReport(data dto.SearchResponse) (string, error) {
	sess := session.Must(session.NewSession())
	svc := s3.New(sess)
	uploader := s3manager.NewUploader(sess)

	file, err := os.Open("data/" + data.ReportID + ".html")
	if err != nil {
		return "", err
	}

	_, err = uploader.Upload(&s3manager.UploadInput{
		Bucket: aws.String(os.Getenv("S3_BUCKET")),
		Key:    aws.String(data.ReportID + ".html"),
		Body:   file,
	})
	if err != nil {
		return "", err
	}

	req, _ := svc.GetObjectRequest(&s3.GetObjectInput{
		Bucket: aws.String(os.Getenv("S3_BUCKET")),
		Key:    aws.String(data.ReportID + ".html"),
	})

	reportTTL, err := strconv.ParseInt(os.Getenv("REPORT_TTL"), 10, 64)
	if err != nil {
		return "", err
	}

	urlStr, err := req.Presign(time.Duration(reportTTL) * time.Hour)
	if err != nil {
		return "", err
	}

	return urlStr, nil
}

func publishMessage(reportID string, fileURL string, status string) error {
	responseReport := dto.ResponseReport{
		ID:      reportID,
		FileURL: fileURL,
		Status:  status,
	}

	responseJson, err := json.Marshal(responseReport)
	if err != nil {
		return err
	}

	producer := kafka.NewKafkaProducer()
	producer.SetupProducer(os.Getenv("KAFKA_BOOTSTRAP_SERVERS"))

	err = producer.Publish(string(responseJson), os.Getenv("KAFKA_PRODUCER_TOPIC"))
	if err != nil {
		return err
	}

	return nil
}
