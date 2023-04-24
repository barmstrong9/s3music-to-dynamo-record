package main

import (
	"log"

	"github.com/aws/aws-lambda-go/lambda"
)

func main() {
	lambda.Start(Handler)
}

func Handler() {
	log.Println("Testing")
}

//Remember to do
// GOOS=linux go build -o build/main cmd/main.go
// zip -jrm build/main.zip build/main
// Then upload to s3 if CI/CD has not been set up.