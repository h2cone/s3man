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
	req, err := http.NewRequest("POST", "http://localhost:8000/upload", body)
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
	respBody := &result.RespBody{}
	if err := json.NewDecoder(resp.Body).Decode(respBody); err != nil {
		t.Fatal(err)
	}
	if respBody.Code <= 0 {
		t.Error(respBody)
	}
	data := respBody.Data
	if data == nil {
		t.Error(respBody)
	}
	media := data.(map[string]interface{})
	t.Log(media)

	eTag := media["eTag"].(string)
	versionID := media["versionId"].(string)
	imgURL := media["imgUrl"].(string)
	fileURL := media["fileUrl"].(string)
	if len(eTag) == 0 || len(versionID) == 0 ||
		len(imgURL) == 0 || len(fileURL) == 0 {
		t.Error()
	}
}
