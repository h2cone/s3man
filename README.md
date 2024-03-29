# s3man

Provide file upload service.

## Configuration

Use environment variables:

Name | Required | Default | Remark
--- | --- | --- | ---
AWS_ACCESS_KEY_ID | true | | AWS access key
AWS_SECRET_ACCESS_KEY | true | | AWS secret access key
AWS_SESSION_TOKEN | false | | AWS session token
AWS_REGION | true | | [Regions and Endpoints](https://docs.aws.amazon.com/general/latest/gr/rande.html)
AWS_ENDPOINT | true | | [Regions and Endpoints](https://docs.aws.amazon.com/general/latest/gr/rande.html)
S3_FORCE_PATH_STYLE | false | false | Whether to force a request to use a path-style address
API_SERVER_ADDR | false | :8000 | Server address
API_SERVER_MULTIPART_MAX_REQUEST_SIZE | false | 10485760 | Maximum file upload request size (Byte)
API_SERVER_MULTIPART_MAX_FILE_SIZE | false | 10485760 | Maximum file size (Byte)
API_SERVER_MULTIPART_FORM_KEY | false | file | Form key of the file
API_BUCKET_DEFAULT | true | | Default bucket
API_BUCKET_GUESSED | false | false | Whether to guess the bucket, if set to true API_BUCKET_IMG and API_BUCKET_FILE are required
API_BUCKET_IMG | false | | Picture bucket
API_BUCKET_FILE | false | | File bucket
API_TIMEOUT | false | 10000 | File upload timeout (Byte)

## API

Upload file:

```http
PUT /upload HTTP/1.1
Host: example.com
User-Agent: example
Accept: */*
Cache-Control: no-cache
Accept-Encoding: gzip, deflate
Content-Type: multipart/form-data; boundary=----WebKitFormBoundary7MA4YWxkTrZu0gW
Content-Length: 64885
Connection: keep-alive
cache-control: no-cache


Content-Disposition: form-data; name="file"; filename="/test/golang.png"


------WebKitFormBoundary7MA4YWxkTrZu0gW--
```

Expected result:

```json
{
    "eTag": "\"6a2043fddc94020d0fd1c0c120ecb626\"",
    "versionId": "MTg0NDUxNzY4MzE2Nzc0MzkzODk",
    "bucket": "test",
    "key": "c3b75abc0e264fa089c181bf7ef57221.png",
    "path": "test/c3b75abc0e264fa089c181bf7ef57221.png"
}
```

You can also change the default configuration, see [config template](config.template.json), and start the program as a specified configuration file:

```shell
./s3man -c config.template.json
```

The default value of `-c` is [config.default.json](config.default.json).
