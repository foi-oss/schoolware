package scenarios

type Scenario struct {
  Name, Description string
  Run               func()
}

var All []*Scenario
