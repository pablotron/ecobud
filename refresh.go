package main

import (
  "encoding/json"
  "net/http"
  net_url "net/url"
)

// Response from token refresh URL.
type RefreshResponse struct {
  // Limited-duration access token.
  AccessToken string `json:"access_token"`

  // Token type (example: "Bearer")
  TokenType string `json:"token_type"`

  // Refresh token (may be updated from request).
  RefreshToken string `json:"refresh_token"`

  // Access token duration, in seconds.
  ExpiresIn int `json:"expires_in"`

  // Comma-delimited scope list.
  Scope string `json:"scope"`
}

// Token refresh URL
var refreshUrl = `https://api.ecobee.com/token`

// Send refresh token and get response with fresh access token and
// (possibly) updated refresh token.
func refreshToken(token string) (*RefreshResponse, error) {
  // build form values
  vals := net_url.Values {}
  vals.Add("grant_type", "refresh_token")
  vals.Add("code", token)
  vals.Add("client_id", apiKey)

  // post form, get response
  resp, err := http.PostForm(refreshUrl, vals)
  if err != nil {
    return nil, err
  }
  defer resp.Body.Close()

  // decode JSON from response body
  var r RefreshResponse
  err = json.NewDecoder(resp.Body).Decode(&r)

  // return result
  return &r, err
}
