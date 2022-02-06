package controllers

import (
	"bytes"
	"context"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/gorilla/mux"
)

func Home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("HELLO Luan Teacher"))

}

func ServeVideo(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	w.Header().Add("Content-Type", "video/mp4") //dinh dang mp4 truoc
	getVideo(id, w, r)
}

func getVideo(videoId string, w http.ResponseWriter, r *http.Request) {
	httpClient := &http.Client{}
	httpClient.NewRequest("GET", "localhost:4001/s3"+videoId)
	fmt.Printf()
	videoKey := "1.mp4"
	cfg, err := config.LoadDefaultConfig(context.TODO(), func(o *config.LoadOptions) error {
		o.Region = "ap-northeast-1"
		return nil
	})
	if err != nil {
		panic(err)
	}

	client := s3.NewFromConfig(cfg)
	output, err := client.GetObject(context.TODO(), &s3.GetObjectInput{
		Bucket: aws.String("learn-go-video"),
		Key:    aws.String(videoKey),
	})

	if err != nil {
		panic(err)
	}

	buff, err := ioutil.ReadAll(output.Body)
	if err != nil {
		panic(err)
	}

	reader := bytes.NewReader(buff)

	http.ServeContent(w, r, "video", time.Now(), reader)

}
