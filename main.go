package main

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sqs"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

const sqsQueueEnv string = "SQS_QUEUE_URL"

func send(message string, queue string) {
	sess := session.Must(session.NewSession())
	svc := sqs.New(sess, &aws.Config{})
	params := &sqs.SendMessageInput{
		MessageBody:  aws.String(message),
		QueueUrl:     aws.String(queue),
		DelaySeconds: aws.Int64(1),
	}
	resp, err := svc.SendMessage(params)
	if err != nil {
		log.Println(err.Error())
		return
	}
	fmt.Println(resp)
}

func process(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Fatal("ioutil.ReadAll: ", err)
	}
	bodyString := string(body[:])
	log.Printf("Received message: \"%s\".", bodyString)
	send(bodyString, os.Getenv(sqsQueueEnv))
	fmt.Fprintf(w, "OK")
}

func serve() {
	http.HandleFunc("/", process)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

func main() {
	if os.Getenv(sqsQueueEnv) == "" {
		log.Fatalf("%s not set.", sqsQueueEnv)
	}
	log.Printf("SQS Queue: \"%s\"", os.Getenv(sqsQueueEnv))
	serve()
}
