# s3man

Provide file upload service

## Configuration

Use environment variables

Name | Required | Default
--- | --- | ---
AWS_ACCESS_KEY_ID | true
AWS_SECRET_ACCESS_KEY | true
AWS_SESSION_TOKEN | false |
AWS_REGION | true
AWS_ENDPOINT | true
S3_FORCE_PATH_STYLE | false | false
API_SERVER_ADDR | false | :8000
API_SERVER_MULTIPART_MAX_REQUEST_SIZE | false | 10485760
API_SERVER_MULTIPART_MAX_FILE_SIZE | false | 10485760
API_SERVER_MULTIPART_FORM_KEY | false | file
API_BUCKET_DEFAULT | true |
API_RETURN_URL_IMG | false |
API_RETURN_URL_FILE | false |
API_TIMEOUT | false | 10000
