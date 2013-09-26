package main

import (
  "bitbucket.org/jol/service/stdservice"
  "flag"
  "fmt"
  "github.com/foi-oss/schoolware/scenarios"
  "log"
  "os"
  "strings"
)

var (
  // command line flag
  scenario = flag.String("scenario", "", "scenario to run")

  // command line flag
  serviceCommand = flag.String("service", "start|stop", "control background service")
)

func main() {
  flag.Usage = usage
  flag.Parse()

  stdservice.Run(&stdservice.Config{
    Name:            "schoolware",
    DisplayName:     "Schoolware",
    LongDescription: "School maleware service",
    Start:           start,
    Stop:            stop,
    Args:            strings.Join(os.Args[1:], " "),
  })

  /*
     s, err := service.NewService("schoolware",
       "Schoolware",
       "School maleware service")
     if err != nil {
       fmt.Fprintf(os.Stderr, "unable to create service: %s", err)
     }

     if len(flag.Args()) >= 1 {
       switch flag.Arg(0) {
       case "install":
         if err := s.Install(); err != nil {
           fmt.Fprintf(os.Stderr, "failed to install: %s", err)
           return
         }
       case "remove":
         s.Remove()
       case "run":
         run()
       case "start":
         if err := s.Start(); err != nil {
           fmt.Fprintf(os.Stderr, "failed to start: %s", err)
           return
         }
       case "stop":
         s.Stop()
       }

       return
     }

     err = s.Run(func() error {
       run()
       return nil
     }, func() error {
       return nil
     })
     if err != nil {
       s.Error(err.Error())
     }
  */
}

func start(c *stdservice.Config) {
  if len(*scenario) == 0 {
    fmt.Fprintf(os.Stderr, "no scenario specified")
    return
  }

  for _, s := range scenarios.All {
    if s.Name == *scenario {
      log.Println("Scenario", *scenario, "started")
      s.Run()
      return
    }
  }

  fmt.Fprintf(os.Stderr, "unknown scenario specified")
}

func stop(c *stdservice.Config) {

}

// usage prints list of known command-line options and scenarion descriptions
func usage() {
  fmt.Fprintf(os.Stderr, "Usage of %s:\n", os.Args[0])
  flag.PrintDefaults()

  fmt.Fprintf(os.Stderr, "\nKnown scenarios:\n")
  for _, s := range scenarios.All {
    fmt.Fprintf(os.Stderr, "  %s: %s\n", s.Name, s.Description)
  }
}
