package main

import (
  "os"
)

// App config
var MEDIA_API_PORT = os.Getenv("MEDIA_API_PORT")
var REQUIRES_AUTH = os.Getenv("REQUIRES_AUTH")
var AUTH_BASE_URL = os.Getenv("AUTH_BASE_URL")

// AWS S3 Configuration
var AWS_S3_ACCESS_KEY = os.Getenv("AWS_S3_ACCESS_KEY")
var AWS_S3_SECRET_KEY = os.Getenv("AWS_S3_SECRET_KEY")
var AWS_S3_BUCKET_NAME = os.Getenv("AWS_S3_BUCKET_NAME")
var AWS_S3_FOLDER = os.Getenv("AWS_S3_FOLDER")
