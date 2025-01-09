package main

import (
	"context"
	"fmt"
	"log"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/ssm"
)

func main() {
	cfg, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		log.Fatal(err)
	}

	client := ssm.NewFromConfig(cfg)
	res, _ := client.GetParameter(context.TODO(), &ssm.GetParameterInput{
		Name: aws.String("/devopsplayground/hello-world"),
	})

	fmt.Println(*res.Parameter.Value)
}
