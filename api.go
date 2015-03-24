package main

import (
  "net/http"
  "encoding/json"
  "io/ioutil"
  "path/filepath"
  "github.com/zenazn/goji/web"
)

func status(c web.C, w http.ResponseWriter, r *http.Request) {
  status := StatusResponse{"ok"}
  js, err := json.Marshal(status)
  if err != nil {
    http.Error(w, err.Error(), http.StatusInternalServerError)
    return
  }

  w.Header().Set("Content-Type", "application/json")
  w.Write(js)
}

func uploadMedia(c web.C, w http.ResponseWriter, r *http.Request) {

  // TODO: authenticate user against the proper API

  file, fileHeader, err := r.FormFile("file")
  if err != nil {
    errorResponse := ErrorResponse{"Error parsing uploaded file"}
    js, err := json.Marshal(errorResponse)
    if err != nil {
      http.Error(w, err.Error(), http.StatusInternalServerError)
      return
    }

    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusBadRequest)
    w.Write(js)
    return
  }

  fileBuf, _ := ioutil.ReadAll(file)
  fileHash := keyOf(fileBuf)

  ext := filepath.Ext(fileHeader.Filename)

  path := AWS_S3_FOLDER + fileHash + ext
  uploadToS3(path, fileHash, fileBuf)

  mediaUploadResponse := MediaUploadResponse{fileHash}
  js, err := json.Marshal(mediaUploadResponse)
  if err != nil {
    http.Error(w, err.Error(), http.StatusInternalServerError)
    return
  }

  w.Header().Set("Content-Type", "application/json")
  w.WriteHeader(http.StatusAccepted)
  w.Write(js)
}


