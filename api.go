package main

import (
  "net/http"
  "encoding/json"
  "io/ioutil"
  "strings"
  "path/filepath"
  "github.com/zenazn/goji/web"
)

// Configuration variables
var AUTH_ENDPOINT_URL = AUTH_BASE_URL + "/oauth/token/info"

// Status endpoint just used to verify that the server is up and running
func status(c web.C, w http.ResponseWriter, r *http.Request) {
  status := StatusResponse{"ok"}
  js, err := json.Marshal(status)
  if err != nil {
    http.Error(w, err.Error(), http.StatusInternalServerError)
    return
  }
  w.Header().Set("Content-Type", "application/json")
  w.WriteHeader(http.StatusOK)
  w.Write(js)
}

// Method as a proxy to upload media to S3, supports verifying auth tokens
func uploadMedia(c web.C, w http.ResponseWriter, r *http.Request) {

  if REQUIRES_AUTH == "true" {
    // Get the bearer auth token from the current request
    receivedBearerToken := r.Header.Get("Authorization")
    if len(receivedBearerToken) == 0 {
      // If there is no token, respond with a 401
      errorResponse := ErrorResponse{"Unauthorized"}
      js, jsonError := json.Marshal(errorResponse)
      if jsonError != nil {
        http.Error(w, "", http.StatusInternalServerError)
        return
      }
      w.Header().Set("Content-Type", "application/json")
      w.WriteHeader(http.StatusUnauthorized)
      w.Write(js)
      return
    }

    // Create the HTTP client that will be used to verify the token
    client := &http.Client{}
    req, httpClientError := http.NewRequest("GET", AUTH_ENDPOINT_URL, nil)
    if httpClientError != nil {
      http.Error(w, httpClientError.Error(), http.StatusInternalServerError)
      return
    }

    // Add the auth token to the verification request and make the request
    req.Header.Add("Authorization", receivedBearerToken)
    apiAuthResponse, verificationReqError := client.Do(req)
    if verificationReqError != nil {
      http.Error(w, verificationReqError.Error(), http.StatusInternalServerError)
      return
    }

    // The verification request should be 200, otherwise we should return an 401 to original request
    if apiAuthResponse.StatusCode != http.StatusOK {
      errorResponse := ErrorResponse{"Unauthorized"}
      js, jsonError := json.Marshal(errorResponse)
      if jsonError != nil {
        http.Error(w, "", http.StatusInternalServerError)
        return
      }
      w.Header().Set("Content-Type", "application/json")
      w.WriteHeader(http.StatusUnauthorized)
      w.Write(js)
      return
    }
  }

  // Get the file from the request
  file, fileHeader, err := r.FormFile("file")
  if err != nil {
    errorResponse := ErrorResponse{"Error parsing uploaded file"}
    js, jsonError := json.Marshal(errorResponse)
    if jsonError != nil {
      http.Error(w, "", http.StatusInternalServerError)
      return
    }
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusBadRequest)
    w.Write(js)
    return
  }

  // Get the file extension
  ext := strings.ToLower(strings.TrimLeft(filepath.Ext(fileHeader.Filename), "."))

  _, fileTypeSupported := fileTypesMap[ext]
  if !fileTypeSupported {
    errorResponse := ErrorResponse{"File type not supported"}
    js, jsonError := json.Marshal(errorResponse)
    if jsonError != nil {
      http.Error(w, "", http.StatusInternalServerError)
      return
    }
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusBadRequest)
    w.Write(js)
    return
  }

  // Read the file from the request
  fileBuf, _ := ioutil.ReadAll(file)

  // Calculate hash of the file
  fileHash := keyOf(fileBuf)

  // Set the path of the file by using the configured "folder", the SHA1 hash and extension
  path := AWS_S3_FOLDER + fileHash + "." + ext

  // Upload the file to S3
  uploadToS3(path, fileHash, fileBuf)

  // Generate the success response
  mediaUploadResponse := MediaUploadResponse{fileHash}
  js, jsonError := json.Marshal(mediaUploadResponse)
  if jsonError != nil {
    http.Error(w, "", http.StatusInternalServerError)
    return
  }

  w.Header().Set("Content-Type", "application/json")
  w.WriteHeader(http.StatusAccepted)
  w.Write(js)
}


