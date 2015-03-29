# S3 Proxy upload

## Description

Using the AWS SDK on Mobile apps is usually overkill and bloated, when in most cases, it's only needed to upload to S3. Doing a direct request to S3 is also a bit painful due to headers signing, so this is a solution for that problem.

## Steps to build

* Set up your environment

  ```
  $ . environment_setup/local.sh
  ```

* Build the app like this:

  ```
  $ go build
  $ ./go-image-proxy
  ```

* Deploy

* For development, install gin:
  ```
  $ go get github.com/codegangsta/gin
  ```

* Listen for changes and build:
  ```
  $ gin -p '8080' run *.go
  ```

## Endpoints available

See routes.go to configure them:

* GET /status
  Expected response:

  ```
  200 OK
  { "status": "ok" }
  ```

* POST /upload (Param: "file")
  Expected responses:

  ```
  202 ACCEPTED
  { "checksum": "sha1checksum" }
  ```

  ```
  400 BAD REQUEST
  { "error": "Unauthorized" }
  ```

  ```
  400 BAD REQUEST
  { "error": "Error parsing uploaded file" }
  ```

  ```
  401 UNAUTHORIZED
  { "error": "Unauthorized" }
  ```
