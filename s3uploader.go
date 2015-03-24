package main

import (
  "path/filepath"
  "strings"
  "launchpad.net/goamz/s3"
  "launchpad.net/goamz/aws"
)

var AWS_S3_ACCESS_KEY = ""
var AWS_S3_SECRET_KEY = ""
var AWS_S3_BUCKET_NAME = ""
var AWS_S3_FOLDER = "";

func uploadToS3(path string, hash string, data []byte) {

    auth := aws.Auth{
    AccessKey: AWS_S3_ACCESS_KEY,
    SecretKey: AWS_S3_SECRET_KEY,
  }

  ext := strings.TrimLeft(filepath.Ext(path), ".")

  var contType string

  if ext == "jpg" {
    contType = "image/jpeg"
  } else if ext == "png" {
    contType = "image/png"
  }

  connection := s3.New(auth, aws.USEast)
  esBucket := connection.Bucket(AWS_S3_BUCKET_NAME)
  esBucket.Put(path, data, contType, s3.BucketOwnerFull)
}


