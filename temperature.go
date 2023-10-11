package main

import (
  "encoding/json"
  "fmt"
  "strconv"
)

// Temperature value, (represented as 1/10° Fahrenheit).
type Temperature int

// Unmarshal integer from JSON as Temperature.
func (t *Temperature) UnmarshalJSON(b []byte) error {
  // unmarshal int
  var i int
  if err := json.Unmarshal(b, &i); err != nil {
    return err
  }

  *t = Temperature(i)
  return nil
}

// Unmarshal integer from JSON as Temperature.
func (t Temperature) MarshalJSON() ([]byte, error) {
  return json.Marshal(int(t))
}

// Parse string as temperature
func ParseTemperature(s string) (Temperature, error) {
  val, err := strconv.ParseInt(s, 10, 32)
  if err != nil {
    return Temperature(0), err
  } else {
    return Temperature(int(val*10)), nil
  }
}

// Print temperature as string.
func (t Temperature) String() string {
  return fmt.Sprintf("%d.%d", int(t)/10, int(t)%10)
}
