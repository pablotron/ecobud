package main

import (
  "fmt"
)

// Hvac mode.
//
// One of "auto", "auxHeatOnly", "cool", "heat", or "off".
type HvacMode uint8

const (
  UnknownHvacMode HvacMode = iota
  Auto
  AuxHeatOnly
  Cool
  Heat
  Off
)

// Map of string to HvacMode.
//
// Used by ParseHvacMode() and HvacMode.UnmarshalText().
var hvacModeStrs = map[string]HvacMode {
  "auto": Auto,
  "auxHeatOnly": AuxHeatOnly,
  "auxheatonly": AuxHeatOnly,
  "aux": AuxHeatOnly,
  "cool": Cool,
  "heat": Heat,
  "off": Off,
}

// Parse string as HVAC mode.
func ParseHvacMode(s string) (HvacMode, error) {
  if m, ok := hvacModeStrs[s]; ok {
    return m, nil
  } else {
    return UnknownHvacMode, fmt.Errorf("unknown mode: %s", s)
  }
}

// Unmarshal text as mode.
func (m *HvacMode) UnmarshalText(b []byte) error {
  nm, err := ParseHvacMode(string(b))
  if err != nil {
    return err
  }

  *m = nm
  return nil
}

// Map of HvacMode to string.
//
// Used by HvacMode.MarshalText() and HvacMode.String().
var hvacModes = map[HvacMode]string {
  Auto: "auto",
  AuxHeatOnly: "auxHeatOnly",
  Cool: "cool",
  Heat: "heat",
  Off: "off",
}

// Marshal mode as text.
func (m HvacMode) MarshalText() ([]byte, error) {
  if s, ok := hvacModes[m]; ok {
    return []byte(s), nil
  } else {
    return []byte(""), fmt.Errorf("unknown hvac mode: %d", m)
  }
}

func (m HvacMode) String() string {
  return hvacModes[m]
}

