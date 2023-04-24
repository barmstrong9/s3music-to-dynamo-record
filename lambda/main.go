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
