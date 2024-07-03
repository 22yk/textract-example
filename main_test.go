package main

import (
	"io/ioutil"
	"testing"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/textract"
)

func init() {
	textractSession = textract.New(session.Must(session.NewSession(
		&aws.Config{
			Region: aws.String("us-east-1"),
		},
	)))
}

func BenchmarkProcessFile(b *testing.B) {
	filePath := "./test.jpg"
	file, err := ioutil.ReadFile(filePath)
	if err != nil {
		b.Fatal(err)
	}

	for i := 0; i < b.N; i++ {
		_, err := textractSession.DetectDocumentText(&textract.DetectDocumentTextInput{
			Document: &textract.Document{
				Bytes: file,
			},
		})
		if err != nil {
			b.Fatal(err)
		}
	}
}
