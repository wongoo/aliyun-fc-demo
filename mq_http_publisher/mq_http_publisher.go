package mq_http_publisher

import "github.com/aliyunmq/mq-http-go-sdk"

func Publish(endpoint, accessId, accessSecret, instanceId, topic, messageKey,
	tag, body string, properties map[string]string) (interface{}, error) {

	client := mq_http_sdk.NewAliyunMQClient(endpoint, accessId, accessSecret, "")

	mqProducer := client.GetProducer(instanceId, topic)

	msg := mq_http_sdk.PublishMessageRequest{
		MessageBody: body,
		MessageKey:  messageKey,
		MessageTag:  tag,
		Properties:  properties,
	}

	return mqProducer.PublishMessage(msg)
}
