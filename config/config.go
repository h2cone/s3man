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

package config

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
	"strconv"
)

var (
	awsAccessKey     = os.Getenv("AWS_ACCESS_KEY_ID")
	awsSecretKey     = os.Getenv("AWS_SECRET_ACCESS_KEY")
	awsSessionToken  = os.Getenv("AWS_SESSION_TOKEN")
	awsRegion        = os.Getenv("AWS_REGION")
	awsEndpoint      = os.Getenv("AWS_ENDPOINT")
	s3ForcePathStyle = os.Getenv("S3_FORCE_PATH_STYLE")

	apiServerAddr                    = os.Getenv("API_SERVER_ADDR")
	apiServerMultipartMaxRequestSize = os.Getenv("API_SERVER_MULTIPART_MAX_REQUEST_SIZE")
	apiServerMultipartMaxFileSize    = os.Getenv("API_SERVER_MULTIPART_MAX_FILE_SIZE")
	apiServerMultipartFormKey        = os.Getenv("API_SERVER_MULTIPART_FORM_KEY")
	apiBucketDefault                 = os.Getenv("API_BUCKET_DEFAULT")
	apiBaseURLImg                    = os.Getenv("API_BASE_URL_IMG")
	apiBaseURLOds                    = os.Getenv("API_BASE_URL_ODS")
)

// Config .
type Config struct {
	Aws AwsConfig
	API APIConfig
}

// AwsConfig .
type AwsConfig struct {
	AccessKey        string
	SecretKey        string
	SessionToken     string
	Region           string
	Endpoint         string
	S3ForcePathStyle bool
}

// APIConfig .
type APIConfig struct {
	Server  ServerConfig
	Bucket  BucketConfig
	BaseURL BaseURLConfig
}

// ServerConfig .
type ServerConfig struct {
	Addr      string
	Multipart MultipartConfig
}

// MultipartConfig .
type MultipartConfig struct {
	FormKey        string
	MaxRequestSize int64
	MaxFileSize    int64
}

// BucketConfig .
type BucketConfig struct {
	Default string
}

// BaseURLConfig .
type BaseURLConfig struct {
	Img string
	Ods string
}

// Load load config file by filename
func Load(filename string) *Config {
	log.Printf("Loading config file: %s", filename)
	bytes, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}
	config := &Config{
		Aws: AwsConfig{
			AccessKey:    awsAccessKey,
			SecretKey:    awsSecretKey,
			SessionToken: awsSessionToken,
			Region:       awsRegion,
			Endpoint:     awsEndpoint,
		},
		API: APIConfig{
			Bucket: BucketConfig{
				Default: apiBucketDefault,
			},
			BaseURL: BaseURLConfig{
				Img: apiBaseURLImg,
				Ods: apiBaseURLOds,
			},
		},
	}
	if err := json.Unmarshal(bytes, config); err != nil {
		log.Fatal(err)
	}
	override(config)
	return config
}

func override(config *Config) {
	if len(s3ForcePathStyle) > 0 {
		pathStyleEnabled, err := strconv.ParseBool(s3ForcePathStyle)
		if err != nil {
			log.Print(err)
		}
		config.Aws.S3ForcePathStyle = pathStyleEnabled
	}
	if len(apiServerAddr) > 0 {
		config.API.Server.Addr = apiServerAddr
	}
	if len(apiServerMultipartFormKey) > 0 {
		config.API.Server.Multipart.FormKey = apiServerMultipartFormKey
	}
	if len(apiServerMultipartMaxRequestSize) > 0 {
		maxReqSize, err := strconv.ParseInt(apiServerMultipartMaxRequestSize, 10, 0)
		if err != nil {
			log.Print(err)
		}
		config.API.Server.Multipart.MaxRequestSize = maxReqSize
	}
	if len(apiServerMultipartMaxFileSize) > 0 {
		maxFileSize, err := strconv.ParseInt(apiServerMultipartMaxFileSize, 10, 0)
		if err != nil {
			log.Print(err)
		}
		config.API.Server.Multipart.MaxFileSize = maxFileSize
	}
}
