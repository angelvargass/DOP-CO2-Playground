package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/sqs"
	"github.com/aws/aws-sdk-go-v2/service/ssm"
)

func main() {
	log.Println("Starting queue script")
	cfg, err := config.LoadDefaultConfig(context.TODO(), config.WithRegion("us-east-1"))
	if err != nil {
		log.Fatalf("Failed to load configuration, %v", err)
	}

	ssmClient := ssm.NewFromConfig(cfg)
	ssmParam, err := ssmClient.GetParameter(context.TODO(), &ssm.GetParameterInput{
		Name: aws.String("/devopsplayground/sqs-url"),
	})

	if err != nil {
		log.Fatal("Could not retrieve SQS URL from SSM Parameter Store")
		os.Exit(1)
	}

	sqsClient := sqs.NewFromConfig(cfg)
	sqsURL := *ssmParam.Parameter.Value
	for {
		output, err := sqsClient.ReceiveMessage(context.TODO(), &sqs.ReceiveMessageInput{
			QueueUrl:            aws.String(sqsURL),
			MaxNumberOfMessages: 5,
			WaitTimeSeconds:     10,
		})

		if err != nil {
			log.Printf("Error receiving messages: %v", err)
			continue
		}
		if len(output.Messages) > 0 {
			fmt.Printf("Received %d message(s):\n", len(output.Messages))
			for _, message := range output.Messages {
				fmt.Printf("Message ID: %s\n", *message.MessageId)
				fmt.Printf("Body: %s\n", *message.Body)

				_, err := sqsClient.DeleteMessage(context.TODO(), &sqs.DeleteMessageInput{
					QueueUrl:      aws.String(sqsURL),
					ReceiptHandle: message.ReceiptHandle,
				})
				if err != nil {
					log.Printf("Error deleting message ID %s: %v", *message.MessageId, err)
				} else {
					fmt.Printf("Message ID %s deleted successfully.\n", *message.MessageId)
				}
			}
		} else {
			fmt.Println("No messages received in this polling cycle.")
		}
	}
}
