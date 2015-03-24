# S3 Proxy upload

## Description


## Steps to build

* Open s3uploader.go and set your access key, secret, upload folder and bucket
* Build the app like this:

  ```
  $ go build
  $ ./go-image-proxy
  ```

* Deploy


## Endpoints available

See routes.go to configure them:

* GET /status
  Expected response:
  200 OK
  { "status": "ok" }

* POST /upload (Param: "file")
  Expected responses:

  202 ACCEPTED
  { "checksum": "sha1checksum" }

  500 INTERNAL SERVER ERROR
  { "error": "Error parsing uploaded file" }
