package main

import (
  "encoding/json"
  "fmt"
  "log/slog"
  "net/http"
)

// Response from status URL.
type StatusResponse struct {
  Page struct {
    Page int `json:"page"` // page number
    TotalPages int `json:"totalPages"` // total number of pages
    PageSize int `json:"pageSize"` // items per page
    Total int `json:"total"` // total number of items
  } `json:"page"`

  // Array of thermstat statuses
  Thermostats []Thermostat `json:"thermostatList"`

  Status struct {
    Code int `json:"code"`
    Message string `json:"message"`
  } `json:"status"`
}

// Thermostat status URL.
var statusUrl = `https://api.ecobee.com/1/thermostat?format=json&body={"selection":{"selectionType":"registered","selectionMatch":"","includeSettings":true,"includeRuntime":true}}`

// Get system status.
func getStatus(accessToken string) (*StatusResponse, error) {
  // build request
  req, err := http.NewRequest("GET", statusUrl, nil)
  if err != nil {
    return nil, err
  }

  // add headers
  req.Header.Add("Content-Type", "text/json")
  req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", accessToken))

  // log request url and headers
  slog.Debug("getStatus", "url", statusUrl, "headers", req.Header)

  // send request, get response
  resp, err := http.DefaultClient.Do(req)
  if err != nil {
    return nil, err
  }
  defer resp.Body.Close()

  // decode JSON from response body
  var r StatusResponse
  err = json.NewDecoder(resp.Body).Decode(&r)

  // return result
  return &r, err
}
