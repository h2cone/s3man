// Copyright 2019 hehuang https://github.com/h2cone

// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at

//     http://www.apache.org/licenses/LICENSE-2.0

// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package test

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"os"
	"runtime"
	"s3man/result"
	"s3man/server"
	"testing"
)

const platform = runtime.GOOS

func TestUpload(t *testing.T) {
	c := "../config.default.json"
	go server.Start(&c)
	// test file
	dir, err := os.Getwd()
	if err != nil {
		t.Fatal(err)
	}
	var path string
	if platform == "linux" {
		path = dir + "/../test/golang.png"
	} else if platform == "windows" {
		path = dir + "\\..\\test\\golang.png"
	} else {
		path = dir
	}
	file, err := os.Open(path)
	if err != nil {
		t.Fatal(err)
	}
	content, err := ioutil.ReadAll(file)
	if err != nil {
		t.Fatal(err)
	}
	f, err := file.Stat()
	if err != nil {
		t.Fatal(err)
	}
	file.Close()

	body := new(bytes.Buffer)
	writer := multipart.NewWriter(body)
	part, err := writer.CreateFormFile("file", f.Name())
	if err != nil {
		t.Fatal(err)
	}
	part.Write(content)
	err = writer.Close()
	if err != nil {
		t.Fatal(err)
	}
	// do reqeust
	req, err := http.NewRequest("PUT", "http://localhost:8000/upload", body)
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Add("Content-Type", writer.FormDataContentType())
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		t.Fatal(err)
	}
	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		t.Fatalf("response status code: %d", resp.StatusCode)
	}
	// test response
	defer resp.Body.Close()
	media := &result.Media{}
	if err := json.NewDecoder(resp.Body).Decode(media); err != nil {
		t.Fatal(err)
	}
	t.Log(media)
	if len(*media.ETag) == 0 || len(*media.VersionID) == 0 ||
		len(media.Bucket) == 0 || len(media.Key) == 0 || len(media.Path) == 0 {
		t.Error()
	}
}
