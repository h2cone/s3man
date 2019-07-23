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

package api

import (
	"bytes"
	"context"
	"fmt"
	"log"
	"mime/multipart"
	"net/http"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"

	"s3man/config"
	"s3man/keygen"
	"s3man/result"
)

const (
	contentType     = "Content-Type"
	applicationJSON = "application/json"
)

// S3Service amazson s3 service
type S3Service struct {
	S3     *s3.S3
	Config *config.Config
}

// Media media information
type Media struct {
	ETag      *string `json:"eTag"`
	VersionID *string `json:"versionId"`
	Bucket    string  `json:"bucket"`
	Key       string  `json:"key"`
	Path      string  `json:"path"`
}

// S3 create s3 service
func S3(cfg *config.Config) *S3Service {
	config := &aws.Config{
		Credentials:      credentials.NewStaticCredentials(cfg.Aws.AccessKey, cfg.Aws.SecretKey, cfg.Aws.SessionToken),
		Region:           aws.String(cfg.Aws.Region),
		Endpoint:         aws.String(cfg.Aws.Endpoint),
		S3ForcePathStyle: aws.Bool(cfg.Aws.S3ForcePathStyle),
	}
	sess := session.Must(session.NewSession(config))
	return &S3Service{
		S3:     s3.New(sess),
		Config: cfg,
	}
}

// Upload upload file
func (svc *S3Service) Upload(w http.ResponseWriter, r *http.Request) {
	w.Header().Set(contentType, applicationJSON)

	conf := svc.Config.API
	if err := r.ParseMultipartForm(conf.Server.Multipart.MaxRequestSize); err != nil {
		log.Printf("Failed to parse multipart form, %v", err)
		w.Write(result.Ko(err.Error()))
		return
	}
	r.Body = http.MaxBytesReader(w, r.Body, conf.Server.Multipart.MaxFileSize)
	file, header, err := r.FormFile(conf.Server.Multipart.FormKey)
	if err != nil {
		log.Printf("Failed to limit the size of incoming request body, %v", err)
		w.Write(result.Ko(err.Error()))
		return
	}

	var bucket string
	if conf.Bucket.Guessed {
		bucket = svc.guessBucket(file)
	} else {
		bucket = conf.Bucket.Default
	}

	key, out, err := svc.put(&bucket, keygen.UUIDWithExt, file, header)
	if err != nil {
		log.Printf("Failed to upload file, %v", err)
		w.Write(result.Ko(err.Error()))
		return
	}
	w.Write(result.Ok(Media{
		ETag:      out.ETag,
		VersionID: out.VersionId,
		Bucket:    bucket,
		Key:       key,
		Path:      path(bucket, key),
	}))
}

func (svc *S3Service) put(bucket *string, keyGen func(string) string,
	file multipart.File, header *multipart.FileHeader) (string, *s3.PutObjectOutput, error) {
	filename := header.Filename
	key := keyGen(filename)
	contentDisposition := fmt.Sprintf("attachment; filename=\"%s\"", filename)

	ctx := context.Background()
	timeout := time.Duration(svc.Config.API.Timeout) * time.Millisecond
	ctx, cancelFn := context.WithTimeout(ctx, timeout)
	defer cancelFn()

	defer file.Close()
	out, err := svc.S3.PutObjectWithContext(ctx, &s3.PutObjectInput{
		Bucket:             bucket,
		Key:                aws.String(key),
		Body:               file,
		ContentDisposition: aws.String(contentDisposition),
	})
	return key, out, err
}

func (svc *S3Service) guessBucket(file multipart.File) string {
	bucketConf := svc.Config.API.Bucket

	data := make([]byte, 512)
	if _, err := file.Read(data); err != nil {
		log.Print(err)
	}
	if _, err := file.Seek(0, 0); err != nil {
		log.Print(err)
	}
	contentType := http.DetectContentType(data)
	imageTypes := []string{
		"image/png",
		"image/jpg",
		"image/jpeg",
	}
	for _, t := range imageTypes {
		if contentType == t {
			return bucketConf.Img
		}
	}
	return bucketConf.File
}

func path(bucket, key string) string {
	buf := new(bytes.Buffer)
	buf.WriteString(bucket)
	buf.WriteString("/")
	buf.WriteString(key)
	return buf.String()
}
