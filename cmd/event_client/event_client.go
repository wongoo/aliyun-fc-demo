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
	"fmt"
	"net/http"
	"os"

	"github.com/aliyun/fc-go-sdk"
)

func main() {
	serviceName := "helloworld"

	// endpoint: https://help.aliyun.com/document_detail/52984.html
	client, err := fc.NewClient(os.Getenv("ENDPOINT"),
		"2016-08-15",
		os.Getenv("ACCESS_KEY_ID"),
		os.Getenv("ACCESS_KEY_SECRET"),
		fc.WithTransport(&http.Transport{MaxIdleConnsPerHost: 100}))
	if err != nil {
		_, _ = fmt.Fprintln(os.Stderr, err)
		return
	}

	// InvokeFunction
	fmt.Println("Invoking function")
	payload := `{
	"name":"test"
}`
	invokeInput := fc.NewInvokeFunctionInput(serviceName, "event").WithPayload([]byte(payload))
	invokeOutput, err := client.InvokeFunction(invokeInput)
	if err != nil {
		_, _ = fmt.Fprintln(os.Stderr, err)
		return
	}

	fmt.Printf("InvokeFunction response: %s \n", invokeOutput)
	logResult, err := invokeOutput.GetLogResult()
	if err != nil {
		fmt.Printf("Failed to get LogResult due to %v\n", err)
	} else {
		fmt.Printf("Invoke function LogResult %s \n", logResult)
	}

	fmt.Printf("result: %s\n", invokeOutput.Payload)
}
