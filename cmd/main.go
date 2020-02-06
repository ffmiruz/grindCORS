package main

import (
	"bytes"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"net/http"
)

func handler(r events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	rbody := bytes.NewReader([]byte{r.Body})
	req, err := http.NewRequest(r.HTTPMethod, r.QueryStringParameters["u"], rbody)
	if err != nil {
		log.Fatal(err)
	}

	req.Header.Add("Origin", req.URL.Host)
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	buf := new(bytes.Buffer)
	buf.ReadFrom(resp.Body)
	return events.APIGatewayProxyResponse{
		StatusCode: 200,
		headers: {
			"Access-Control-Allow-Origin": "*",
		},
		Body: buf.String(),
	}, nil
}

func main() {
	// Make the handler available for Remote Procedure Call by AWS Lambda
	lambda.Start(handler)
}
