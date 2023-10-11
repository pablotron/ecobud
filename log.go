package main

import (
  "log/slog"
  "os"
)

// map of level to slog.Level
var logLevelStrs = map[string]slog.Level {
  "debug": slog.LevelDebug,
  "info": slog.LevelInfo,
  "warn": slog.LevelWarn,
  "error": slog.LevelError,
}

// global log level
var logLevel = new(slog.LevelVar)

// Initialize logger
func logInit() {
  opts := &slog.HandlerOptions { Level: logLevel }
  handler := slog.NewJSONHandler(os.Stderr, opts)
  slog.SetDefault(slog.New(handler))

  // get log level from environment
  if s := os.Getenv("ECOBUD_LOG_LEVEL"); s != "" {
    // map string value to log level
    level, ok := logLevelStrs[s]
    if !ok {
      // fail with error
      slog.Error("unknown log level", "level", s)
      os.Exit(-1)
    }

    // set log level
    logLevel.Set(level)
  }
}
