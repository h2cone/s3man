// Copyright 2019 h2cone https://github.com/h2cone

// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at

//     http://www.apache.org/licenses/LICENSE-2.0

// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package result

import (
	"encoding/json"
	"log"
)

// RespBody JSON response body
type RespBody struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

// Ok create a JSON response body that represents success
func Ok(data interface{}) []byte {
	body := &RespBody{
		Code:    1,
		Message: "Ok",
		Data:    data,
	}
	return encode(body)
}

// Ko create a JSON response body that represents failure
func Ko(message string) []byte {
	body := &RespBody{
		Code:    0,
		Message: message,
	}
	return encode(body)
}

func encode(body *RespBody) []byte {
	bytes, err := json.Marshal(body)
	if err != nil {
		log.Print(err)
		return nil
	}
	return bytes
}
