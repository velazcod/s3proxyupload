package main

type StatusResponse struct {
  Status    string      `json:"status"`
}

type ErrorResponse struct {
  Error    string      `json:"error"`
}

type MediaUploadResponse struct {
  Checksum string       `json:"checksum"`
}
