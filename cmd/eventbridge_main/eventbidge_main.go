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
	"context"
	"fmt"

	"github.com/aliyun/fc-runtime-go-sdk/fc"
)

func main() {
	fc.Start(eventHandler)
}

func eventHandler(ctx context.Context, event map[string]interface{}) (string, error) {
	fmt.Printf("context: %v\n", ctx)
	fmt.Printf("event: %v\n", event)

	data := event["data"].(map[string]interface{})
	fmt.Printf("data: %v\n", data)

	body := data["body"]
	fmt.Printf("body: %v\n", body)

	return "hello world! 你好，世界!", nil
}
