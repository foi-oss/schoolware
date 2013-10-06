package main

import (
  "bitbucket.org/jol/service/stdservice"
  "flag"
  "fmt"
  "github.com/foi-oss/schoolware/scenarios"
  "os"
  "os/user"
  "strings"
)

var (
  // command line flag
  scenario = flag.String("scenario", "", "scenario to run")
  homedir  = flag.String("homedir", "~", "home directory")
)

func main() {
  if len(os.Args) == 1 {
    fmt.Fprintf(os.Stderr, "Missing arguments. Seek --help")
    return
  }

  flag.Usage = usage
  flag.Parse()

  u, _ := user.Current()
  args := strings.Join(os.Args[1:len(os.Args)-1], " ") + " -homedir=\"" + u.HomeDir + "\""

  stdservice.Run(&stdservice.Config{
    Name:            "schoolware",
    DisplayName:     "Schoolware",
    LongDescription: "School malware service",
    Start:           start,
    Stop:            stop,
    Args:            args,
  })
}

func start(c *stdservice.Config) {
  l := c.Logger()
  l.Info("schoolware started")

  if len(*scenario) == 0 {
    l.Error("no scenario specified")
    return
  }

  if *homedir == "~" {
    u, _ := user.Current()
    scenarios.HomeDir = u.HomeDir
  } else {
    scenarios.HomeDir = *homedir
  }

  for _, s := range scenarios.All {
    if s.Name == *scenario {
      l.Info(fmt.Sprintf("scenario %s started", *scenario))

      err := s.Run(c)
      if err != nil {
        l.Error("scenario failed with: " + err.Error())
      }

      return
    }
  }

  l.Error("unknown scenario specified")
}

func stop(c *stdservice.Config) {
  l := c.Logger()
  l.Info("schoolware is shutting down")
}

// usage prints list of known command-line options and scenarion descriptions
func usage() {
  fmt.Fprintf(os.Stderr, "Usage of %s:\n", os.Args[0])
  fmt.Fprintf(os.Stderr, "  %s [--scenario=net|files|... OPTIONS] COMMAND\n\n", os.Args[0])

  fmt.Fprintln(os.Stderr, "Options:")
  flag.PrintDefaults()

  fmt.Fprintln(os.Stderr, "\nScenarios:")
  for _, s := range scenarios.All {
    fmt.Fprintf(os.Stderr, "  %s: %s\n", s.Name, s.Description)
  }
  fmt.Fprintln(os.Stderr, "\nOptions for each scenario are prefixed with its name.")
  fmt.Fprintln(os.Stderr, "\nCommands:")
  fmt.Fprintln(os.Stderr, ` 
  run     immediately run specified scenario
  install install background service
  start   start previously installed service
  stop    stops the service
  remove  removes schoolware service from the systems`)
}
