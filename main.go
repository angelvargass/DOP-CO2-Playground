package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/ssm"
)

func main() {
	cfg, err := config.LoadDefaultConfig(context.TODO(), config.WithDefaultRegion("us-east-1"))
	if err != nil {
		log.Fatal(err)
	}

	client := ssm.NewFromConfig(cfg)
	res, err := client.GetParameter(context.TODO(), &ssm.GetParameterInput{
		Name: aws.String("/devopsplayground/hello-world"),
	})

	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	fmt.Println(*res.Parameter.Value)
}
