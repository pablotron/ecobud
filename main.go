package main

import (
  "errors"
  "fmt"
  "log/slog"
  "os"

  "github.com/urfave/cli/v2"
)

// application-specific API key
var apiKey = `gpAtMluDyuanOApnQpW1uHGTITnQvgZ3`

// TODO: init func which uses api key and presents user with registration pin
// then writes initial values to config

// Convert thermostat status to brief status line.
func briefStatusLine(t Thermostat) string {
  return fmt.Sprintf("mode: %s, target: %s°F, current: %s°F", t.Settings.HvacMode, t.Runtime.DesiredHeat, t.Runtime.ActualTemperature)
}

// no thermostats found
var errNoThermostats = errors.New("no thermostats found")

// Entry point for "status" command-line subcommand.
func doStatus(config *Config) error {
  // get full system status
  s, err := getStatus(config.LastResponse.AccessToken)
  if err != nil {
    return err
  }

  // check for at least one thermostat
  if len(s.Thermostats) == 0 {
    return errNoThermostats
  }

  // print status of first thermostat
  fmt.Println(briefStatusLine(s.Thermostats[0]))

  // return success
  return nil
}


// Set hold temperature.
//
// Reference: https://www.ecobee.com/home/developer/api/examples/ex5.shtml
func doSetTemperature(config *Config, temp Temperature) error {
  return update(config, UpdateParams {
    Selection: defaultSelection,

    Functions: []UpdateParamsFunction {
      UpdateParamsFunction {
        Type: "setHold",
        Params: UpdateParamsFunctionParams {
          HoldType: "nextTransition",
          HeatHoldTemp: temp,
          CoolHoldTemp: temp,
        },
      },
    },
  })
}

// Set HVAC mode.
//
// Reference: https://www.ecobee.com/home/developer/api/examples/ex8.shtml
func doSetMode(config *Config, mode HvacMode) error {
  return update(config, UpdateParams {
    Selection: defaultSelection,
    Thermostat: &UpdateParamsThermostat {
      Settings: UpdateParamsThermostatSettings {
        HvacMode: mode,
      },
    },
  })
}

// Set temperature from string if string is not empty.
func doMaybeSetTemperature(config *Config, s string) error {
  // get temperature from args (optional)
  if s == "" {
    // optional temperature argument is not set, do nothing
    return nil
  }

  // parse temperature from string
  temp, err := ParseTemperature(s)
  if err != nil {
    return err
  }

  // set temperature
  return doSetTemperature(config, temp)
}

// command-line application
var app = &cli.App {
  Name: "ecobud",
  Usage: "control ecobee thermostat",

  Commands: []*cli.Command {
    &cli.Command {
      Name: "status",
      Usage: "get thermostat status",
      Aliases: []string { "st", "stat" },

      Action: func(cctx *cli.Context) error {
        // load config
        config, err := getConfig()
        if err != nil {
          return nil
        }

        // print status
        return doStatus(config)
      },
    },

    &cli.Command {
      Name: "set",
      Usage: "Set HVAC mode and (optionally) temperature.",
      ArgsUsage: "[auto|auxheatonly|cool|heat|off] <temperature>",

      Action: func(cctx *cli.Context) error {
        // load config
        config, err := getConfig()
        if err != nil {
          return err
        }

        // parse hvac mode from args
        mode, err := ParseHvacMode(cctx.Args().First())
        if err != nil {
          return err
        }

        // set thermostat mode
        if err = doSetMode(config, mode); err != nil {
          return err
        }

        // set temperature (optional)
        return doMaybeSetTemperature(config, cctx.Args().Get(1))
      },
    },

    &cli.Command {
      Name: "auto",
      Usage: `Set mode to "auto" and optionally set hold temperature.`,
      ArgsUsage: "<temperature>",

      Action: func(cctx *cli.Context) error {
        // load config
        config, err := getConfig()
        if err != nil {
          return err
        }

        // set mode
        if err = doSetMode(config, Auto); err != nil {
          return err
        }

        // set temperature (optional)
        return doMaybeSetTemperature(config, cctx.Args().First())
      },
    },

    &cli.Command {
      Name: "heat",
      Usage: `Set mode to "heat" and optionally set hold temperature.`,
      ArgsUsage: "<temperature>",

      Action: func(cctx *cli.Context) error {
        // load config
        config, err := getConfig()
        if err != nil {
          return err
        }

        // set mode
        if err = doSetMode(config, Heat); err != nil {
          return err
        }

        // set temperature (optional)
        return doMaybeSetTemperature(config, cctx.Args().First())
      },
    },

    &cli.Command {
      Name: "cool",
      Usage: `Set mode to "cool" and optionally set hold temperature.`,
      ArgsUsage: "<temperature>",

      Action: func(cctx *cli.Context) error {
        // load config
        config, err := getConfig()
        if err != nil {
          return err
        }

        // set mode
        if err = doSetMode(config, Cool); err != nil {
          return err
        }

        // set temperature (optional)
        return doMaybeSetTemperature(config, cctx.Args().First())
      },
    },

    &cli.Command {
      Name: "off",
      Usage: `Set mode to "off" and optionally set hold temperature.`,
      ArgsUsage: "<temperature>",

      Action: func(cctx *cli.Context) error {
        // load config
        config, err := getConfig()
        if err != nil {
          return err
        }

        // set mode
        if err = doSetMode(config, Off); err != nil {
          return err
        }

        // set temperature (optional)
        return doMaybeSetTemperature(config, cctx.Args().First())
      },
    },

    &cli.Command {
      Name: "temp",
      Usage: "set thermostat temperature",
      Aliases: []string { "t", "temperature" },
      ArgsUsage: "[temp in degrees fahrenheit]",

      Action: func(cctx *cli.Context) error {
        // load config
        config, err := getConfig()
        if err != nil {
          return err
        }

        // parse temperature from args
        temp, err := ParseTemperature(cctx.Args().First())
        if err != nil {
          return err
        }

        // set temperature
        return doSetTemperature(config, temp)
      },
    },
  },
}

func main() {
  // init logger
  logInit()

  // get args, set default command
  args := os.Args
  if len(args) < 2 {
    // default to "status" command
    args = append(args, "status")
  }

  // run app
  if err := app.Run(args); err != nil {
    slog.Error("run", "err", err)
    os.Exit(-1)
  }
}
