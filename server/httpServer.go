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

package server

import (
	"log"
	"net/http"
	"s3man/api"
	"s3man/config"
)

// Start start HTTP server
func Start(c *string) {
	cfg := config.Load(c)
	svc := api.S3(cfg)

	http.HandleFunc("/upload", svc.Upload)

	addr := cfg.API.Server.Addr
	log.Print("HTTP server running on " + addr)
	log.Fatal(http.ListenAndServe(addr, nil))
}
