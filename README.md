# s3man

Provide file upload service

## Configuration

Use environment variables

Name | Required | Default | Remark
--- | --- | --- | ---
AWS_ACCESS_KEY_ID | 是 | | AWS access key
AWS_SECRET_ACCESS_KEY | 是 | | AWS secret access key
AWS_SESSION_TOKEN | 否 | | AWS session token
AWS_REGION | 是 | | [Regions and Endpoints](https://docs.aws.amazon.com/general/latest/gr/rande.html)
AWS_ENDPOINT | 是 | | [Regions and Endpoints](https://docs.aws.amazon.com/general/latest/gr/rande.html)
S3_FORCE_PATH_STYLE | 否 | false | Whether to force a request to use a path-style address
API_SERVER_ADDR | 否 | :8000 | Server address
API_SERVER_MULTIPART_MAX_REQUEST_SIZE | 否 | 10485760 | Maximum file upload request size (Byte)
API_SERVER_MULTIPART_MAX_FILE_SIZE | 否 | 10485760 | Maximum file size (Byte)
API_SERVER_MULTIPART_FORM_KEY | 否 | file | Form key of the file
API_BUCKET_DEFAULT | 是 | | Default bucket
API_BUCKET_GUESSED | 否 | false | Whether to guess the bucket, if set to true API_BUCKET_IMG and API_BUCKET_FILE are required
API_BUCKET_IMG | 否 | | Picture bucket
API_BUCKET_FILE | 否 | | File bucket
API_TIMEOUT | 否 | 10000 | File upload timeout (Byte)

## API

Upload file

```http
POST /upload HTTP/1.1
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

Return result

```json
{
    "code": 1,
    "message": "ok",
    "data": {
        "eTag": "\"6a2043fddc94020d0fd1c0c120ecb626\"",
        "versionId": "MTg0NDUxNzk1MTU2NjQ5NDE1OTk",
        "bucket": "test",
        "key": "d2a04fffff874fa2863680214582c3d6.png",
        "path": "test/d2a04fffff874fa2863680214582c3d6.png"
    }
}
```
