package main

import (
  "bytes"
  "encoding/json"
  "fmt"
  "log/slog"
  "net/http"
)

// Update response
type UpdateResponse struct {
  Status struct {
    Code int `json:"code"`
    Message string `json:"message"`
  } `json:"status"`
}

// default thermostat selection for update requests
var defaultSelection = UpdateParamsSelection {
  SelectionType: "registered",
  SelectionMatch: "",
}

// Thermostat mode URL.
var updateUrl = `https://api.ecobee.com/1/thermostat?format=json`

// Send update parameters.  Used by doSetTemperature() and doSetMode().
func update(config *Config, params UpdateParams) error {
  // JSON-encode update parameters into buffer
  var buf bytes.Buffer
  if err := json.NewEncoder(&buf).Encode(params); err != nil {
    return err
  }

  // build request
  req, err := http.NewRequest("POST", updateUrl, &buf)
  if err != nil {
    return err
  }

  // add headers
  req.Header.Add("Content-Type", "application/json;charset=UTF-8")
  req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", config.LastResponse.AccessToken))

  // log request url and headers
  slog.Debug("update", "url", updateUrl, "headers", req.Header, "body", buf.String())

  // send request, get response
  resp, err := http.DefaultClient.Do(req)
  if err != nil {
    return err
  }
  defer resp.Body.Close()

  // decode JSON from response body
  var r UpdateResponse
  if err = json.NewDecoder(resp.Body).Decode(&r); err != nil {
    return err
  }

  // check response code
  if r.Status.Code != 0 {
    return fmt.Errorf("%s", r.Status.Message)
  }

  // return success
  return nil
}
