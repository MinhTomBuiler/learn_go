package controllers

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
)

type Item struct {
	Id   string `json:"id"`
	Path string `json:"path"`
}

func GetS3Path(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	videoId := vars["videoId"]
	table := "video"
	cfg, err := config.LoadDefaultConfig(context.TODO(), func(o *config.LoadOptions) error {
		o.Region = "ap-northeast-1"
		return nil
	})
	if err != nil {
		panic(err)
	}

	client := dynamodb.NewFromConfig(cfg)

	resp, err := client.GetItem(context.TODO(), &dynamodb.GetItemInput{
		TableName: aws.String(table),
		Key: map[string]types.AttributeValue{
			"Id": &types.AttributeValueMemberS{Value: videoId},
		},
	})

	if err != nil {
		panic(err)
	}

	item := Item{}

	err = attributevalue.UnmarshalMap(resp.Item, &item)
	if err != nil {
		panic(fmt.Sprintf("Failed to unmarshal Record, %v", err))
	}

	jsonResp, err := json.Marshal(item)

	if err != nil {
		panic(fmt.Sprintf("Failed to marshal Json, %v", err))
	}

	w.Write(jsonResp)
}
