package main

import (
  "encoding/json"
  "os"
  "path/filepath"
  "time"
)

// Application configuration.
type Config struct {
  // Expiration time for access token.
  ExpiresAt time.Time `json:"expires_at"`

  // last refreshToken response
  LastResponse *RefreshResponse `json:"last_response"`
}

// Get absolute path to config file.
func getConfigPath() (string, error) {
  // check for environment variable
  if s := os.Getenv("ECOBUD_CONFIG_PATH"); s != "" {
    return s, nil
  }

  // get user config directory (e.g. ~/.config)
  userConfigDir, err := os.UserConfigDir()
  if err != nil {
    return "", err
  }

  // build absolute path to config.json
  configPath := filepath.Join(userConfigDir, "ecobud", "config.json")

  // return result
  return configPath, nil
}

// Create new Config from file.
func newConfigFromFile(path string) (*Config, error) {
  f, err := os.Open(path)
  if err != nil {
    return nil, err
  }
  defer f.Close()

  // parse config as json
  var r Config
  err = json.NewDecoder(f).Decode(&r)
  if err != nil {
    return nil, err
  }

  // does access token need refresh?
  if time.Now().After(r.ExpiresAt) {
    // refresh access token
    t, err := refreshToken(r.LastResponse.RefreshToken)
    if err != nil {
      return nil, err
    }

    // get expiration time
    expiresAt := time.Now().Add(time.Duration(t.ExpiresIn * int(time.Second)))

    // create new config with updated token
    r = Config { ExpiresAt: expiresAt, LastResponse: t }

    // save config
    if err = r.Save(path); err != nil {
      return nil, err
    }
  }

  return &r, err
}

// save config to file as JSON
func (c Config) Save(path string) error {
  // open file for writing
  f, err := os.Create(path)
  if err != nil {
    return err
  }
  defer f.Close()

  // encode as json, write to file
  return json.NewEncoder(f).Encode(c)
}

// Get path to config file, load it, and return config object.
func getConfig() (*Config, error) {
  // get path to config.json
  configPath, err := getConfigPath()
  if err != nil {
    return nil, err
  }

  // load config
  config, err := newConfigFromFile(configPath)
  if err != nil {
    return nil, err
  }

  // return result
  return config, nil
}
