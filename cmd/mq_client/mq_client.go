package main

import (
	"encoding/json"
	"log"
	"os"

	"github.com/wongoo/aliyun-fc-demo/mq_http_publisher"
)

func main() {
	endpoint, _ := os.LookupEnv("MQ_ENDPOINT")
	accessId, _ := os.LookupEnv("MQ_ACCESS_ID")
	accessSecret, _ := os.LookupEnv("MQ_ACCESS_SECRET")
	instanceId, _ := os.LookupEnv("MQ_INSTANCE_ID")

	log.Println(endpoint)
	log.Println(accessId)
	log.Println(accessSecret)
	log.Println(instanceId)

	topic := "test"
	messageKey := ""
	tag := "fc"
	body := `{"message":"hello world"}`

	ret, err := mq_http_publisher.Publish(endpoint, accessId, accessSecret, instanceId, topic,
		messageKey, tag, body, nil)
	if err != nil {
		log.Panicln(err)
		return
	}

	result, _ := json.Marshal(ret)
	log.Printf("result: %s\n", result)

}
