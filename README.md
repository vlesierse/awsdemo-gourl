# GoURL
GoURL is a URL shortner written in Go to demonstrate the building and deploying of Go applications in AWS using AWS Lambda and AWS CDK

## Prerequisites 

- Go 1.16+
- AWS CDK v2

## Deployment

```sh
git clone https://github.com/vlesierse/awsdemo-gourl
cdk deploy
```

## Usage

```sh
curl -H "Content-Type: application/json" --data '{"url":"https://aws.amazon.com"}' $SERVICE_URL
```
