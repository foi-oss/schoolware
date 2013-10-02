package scenarios

import (
  "bitbucket.org/jol/service/stdservice"
)

type Scenario struct {
  Name, Description string
  Run               func(*stdservice.Config) error
}

var All []*Scenario
var HomeDir string
