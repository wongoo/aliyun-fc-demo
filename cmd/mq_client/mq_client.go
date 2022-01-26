/*
 * Licensed to the Apache Software Foundation (ASF) under one or more
 * contributor license agreements.  See the NOTICE file distributed with
 * this work for additional information regarding copyright ownership.
 * The ASF licenses this file to You under the Apache License, Version 2.0
 * (the "License"); you may not use this file except in compliance with
 * the License.  You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

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
