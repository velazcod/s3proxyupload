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

var fileTypesMap map[string]string

func init() {
  fileTypesMap = make(map[string]string)
  fileTypesMap["png"] = "image/png"
  fileTypesMap["jpg"] =  "image/jpeg"
  fileTypesMap["jpeg"] = "image/jpeg"
  fileTypesMap["avi"] = "video/avi"
  fileTypesMap["mpeg"] = "video/mpeg"
  fileTypesMap["mp4"] = "video/mp4"
  fileTypesMap["ogg"] = "video/ogg"
  fileTypesMap["mov"] = "video/quicktime"
}
