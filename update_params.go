package main

// Thermostat selection.
type UpdateParamsSelection struct {
  SelectionType string `json:"selectionType"`
  SelectionMatch string `json:"selectionMatch"`
}

// Thermostat settings.
type UpdateParamsThermostatSettings struct {
  // HVAC mode
  HvacMode HvacMode `json:"hvacMode"`
}

// Thermostat settings.
type UpdateParamsThermostat struct {
  Settings UpdateParamsThermostatSettings `json:"settings"`
}

// Parameters for function.
type UpdateParamsFunctionParams struct {
  HoldType string `json:"holdType"`
  HeatHoldTemp Temperature `json:"heatHoldTemp"`
  CoolHoldTemp Temperature `json:"coolHoldTemp"`
}

// Function for update.
type UpdateParamsFunction struct {
  Type string `json:"type"`
  Params UpdateParamsFunctionParams `json:"params"`
}

// Update parameters
type UpdateParams struct {
  Selection UpdateParamsSelection `json:"selection"`
  Thermostat *UpdateParamsThermostat `json:"thermostat,omitempty"`
  Functions []UpdateParamsFunction `json:"functions,omitempty"`
}
