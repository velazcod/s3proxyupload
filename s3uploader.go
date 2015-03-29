package main

import (
  "path/filepath"
  "strings"
  "launchpad.net/goamz/s3"
  "launchpad.net/goamz/aws"
)

func uploadToS3(path string, hash string, data []byte) {
  // Setup the AWS authentication
  auth := aws.Auth {
    AccessKey: AWS_S3_ACCESS_KEY,
    SecretKey: AWS_S3_SECRET_KEY,
  }

  // Use the extension type to get the content type
  ext := strings.ToLower(strings.TrimLeft(filepath.Ext(path), "."))
  contType := fileTypesMap[ext]

  // Setup the connection to S3 and bucket, and upload the file
  connection := s3.New(auth, aws.USEast)
  esBucket := connection.Bucket(AWS_S3_BUCKET_NAME)
  esBucket.Put(path, data, contType, s3.BucketOwnerFull)
}


