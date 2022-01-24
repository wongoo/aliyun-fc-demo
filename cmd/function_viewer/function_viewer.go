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

	client, err := fc.NewClient(os.Getenv("ENDPOINT"),
		"2016-08-15",
		os.Getenv("ACCESS_KEY_ID"),
		os.Getenv("ACCESS_KEY_SECRET"),
		fc.WithTransport(&http.Transport{MaxIdleConnsPerHost: 100}))
	if err != nil {
		_, _ = fmt.Fprintln(os.Stderr, err)
		return
	}

	// GetService
	fmt.Println("Getting service")
	getServiceOutput, err := client.GetService(fc.NewGetServiceInput(serviceName))
	if err != nil {
		_, _ = fmt.Fprintln(os.Stderr, err)
		return
	}

	fmt.Printf("GetService response: %s \n", getServiceOutput)

	// ListServices
	fmt.Println("Listing services")
	listServicesOutput, err := client.ListServices(fc.NewListServicesInput().WithLimit(100))
	if err != nil {
		_, _ = fmt.Fprintln(os.Stderr, err)
		return
	}
	fmt.Printf("ListServices response: %s \n", listServicesOutput)

	// ListFunctions
	fmt.Println("Listing functions")
	listFunctionsOutput, err := client.ListFunctions(fc.NewListFunctionsInput(serviceName))
	if err != nil {
		_, _ = fmt.Fprintln(os.Stderr, err)
		return
	}

	fmt.Printf("ListFunctions response: %s \n", listFunctionsOutput)
}
