package cmd

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/awslabs/aws-lambda-go-api-proxy/gorillamux"
	"github.com/spf13/cobra"
	"github.com/vlesierse/awsdemo-gourl/api"
)

var muxLambda *gorillamux.GorillaMuxAdapter

var rootCmd = &cobra.Command{
	Use:   "gourl",
	Short: "Go URL Shortener",
	Long:  `URL Shortener written in Go to demonstrate a serverless application in AWS.`,
	Run: func(cmd *cobra.Command, args []string) {
		log.Println("Starting Lambda function")
		lambda.Start(lambdaHandler)
	},
}

func lambdaHandler(ctx context.Context, req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	// If no name is provided in the HTTP request body, throw an error
	return muxLambda.ProxyWithContext(ctx, req)
}

func init() {
	// Create the router and adapter
	router := api.NewRouter()
	muxLambda = gorillamux.New(router)
	//lambdaCmd.PersistentFlags().StringVarP(&port, "port", "p","8080", "This flag sets the port of our API server")
	//rootCmd.AddCommand(lambdaCmd)
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
