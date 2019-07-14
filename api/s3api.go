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
	"log"
	"mime/multipart"
	"net/http"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"

	"s3uploader/config"
	"s3uploader/keygen"
	"s3uploader/result"
)

const (
	formKey         = "file"
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
	ImgURL    string  `json:"imgUrl"`
	OdsURL    string  `json:"odsUrl"`
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
	conf := svc.Config.API
	if err := r.ParseMultipartForm(conf.Server.Multipart.MaxRequestSize); err != nil {
		handleErr(err, w)
		return
	}
	w.Header().Set(contentType, applicationJSON)

	r.Body = http.MaxBytesReader(w, r.Body, conf.Server.Multipart.MaxFileSize)
	file, header, err := r.FormFile(formKey)
	if err != nil {
		handleErr(err, w)
		return
	}
	// TODO: timout setting
	defaultBucket := conf.Bucket.Default
	key, out, err := svc.put(defaultBucket, keygen.UUIDWithExt, file, header)
	if err != nil {
		handleErr(err, w)
		return
	}
	baseURL := conf.BaseURL
	w.Write(result.Ok(Media{
		ETag:      out.ETag,
		VersionID: out.VersionId,
		ImgURL:    url(baseURL.Img, defaultBucket, key),
		OdsURL:    url(baseURL.Ods, defaultBucket, key),
	}))
}

func (svc *S3Service) put(bucket string, keyGen func(string) string,
	file multipart.File, header *multipart.FileHeader) (string, *s3.PutObjectOutput, error) {
	key := keyGen(header.Filename)

	defer file.Close()
	out, err := svc.S3.PutObject(&s3.PutObjectInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(key),
		Body:   file,
	})
	return key, out, err
}

func url(prefix, bucket, key string) string {
	buf := new(bytes.Buffer)
	buf.WriteString(prefix)
	buf.WriteString(bucket)
	buf.WriteString("/")
	buf.WriteString(key)
	return buf.String()
}

func handleErr(err error, w http.ResponseWriter) {
	log.Print(err)
	w.Write(result.Ko(err.Error()))
}
