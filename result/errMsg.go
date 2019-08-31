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
	"s3man/codec"
	"time"
)

// ErrMsg error message
type ErrMsg struct {
	Timestamp int64  `json:"timestamp"`
	Message   string `json:"message"`
}

// NewErrMsg .
func NewErrMsg(message string) []byte {
	errMsg := &ErrMsg{
		Timestamp: time.Now().Unix(),
		Message:   message,
	}
	return codec.Encode(errMsg)
}