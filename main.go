package main

import (
	"fmt"
	"io/ioutil"
	"sync"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/textract"
)

var textractSession *textract.Textract

func init() {
	textractSession = textract.New(session.Must(session.NewSession(
		&aws.Config{
			Region: aws.String("us-east-1"),
		},
	)))
}

func processFile(filePath string, wg *sync.WaitGroup) {
	defer wg.Done()

	file, err := ioutil.ReadFile(filePath)
	if err != nil {
		panic(err)
	}
	resp, err := textractSession.DetectDocumentText(&textract.DetectDocumentTextInput{
		Document: &textract.Document{
			Bytes: file,
		},
	})
	if err != nil {
		panic(err)
	}
	for i := 1; i < len(resp.Blocks); i++ {
		if *resp.Blocks[i].BlockType == "WORD" {
			fmt.Println(*resp.Blocks[i].Text)
		}
	}
}

func main() {
	files := []string{"./test.jpg", "./test2.jpg"}
	var wg sync.WaitGroup

	for _, file := range files {
		wg.Add(1)
		go processFile(file, &wg)
	}

	wg.Wait()
}
