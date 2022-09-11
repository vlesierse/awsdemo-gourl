package store

import (
	"context"
	"encoding/base64"
	"fmt"
	"hash/fnv"
	"log"
	"os"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
)

type UrlItem struct {
	Slug  		string    `json:"slug" dynamodbav:"slug"`
	OriginalUrl	string    `json:"original" dynamodbav:"original"`
}

var svc 		*dynamodb.Client
var tableName 	string

func init() {
	cfg, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		log.Fatalf("unable to load SDK config, %v", err)
	}
	svc = dynamodb.NewFromConfig(cfg)
	tableName = os.Getenv("DYNAMODB_TABLENAME")
}

func generateSlug(url string) string  {
	h := fnv.New64a()
	ts := time.Now().Unix()
	h.Write([]byte(fmt.Sprintf("%s%d", h, ts)))
	bs := h.Sum(nil)
	return base64.URLEncoding.EncodeToString(bs)
}

func GetUrlItem(slug string) *UrlItem {
	key := struct {Slug string `dynamodbav:"slug"`}{slug}
	keyMap, _ := attributevalue.MarshalMap(key)
	input := &dynamodb.GetItemInput{
		TableName: aws.String(tableName),
		Key: keyMap,
	}
	res, err := svc.GetItem(context.TODO(), input)
	if err != nil {
		fmt.Printf("Unable to get item: %v\n", err.Error())
	}
	if res.Item == nil {
		return nil
	}
	item := UrlItem{}
	attributevalue.UnmarshalMap(res.Item, &item)
	return &item
}

func CreateUrlItem(url string) *UrlItem {
	slug := generateSlug(url)
	item := UrlItem{slug, url}
	itemMap, err := attributevalue.MarshalMap(item)
	if err != nil {
		panic("Cannot marshal url into AttributeValue map")
	}

	// create the api params
	input := &dynamodb.PutItemInput{
		TableName: aws.String(tableName),
		Item:      itemMap,
	}

	// put the item
	_, err = svc.PutItem(context.TODO(), input)
	if err != nil {
		fmt.Printf("Unable to add item: %v\n", err.Error())
	}
	return &item
}


