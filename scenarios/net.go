package scenarios

import (
  "io/ioutil"
  "log"
  "net"
  "net/http"
  "strconv"
  "time"
)

var (
  GETS      = []string{"http://google.com", "http://foi.hr"}
  DNSREQS   = []string{"bing.com", "fer.hr"}
  OPENPORTS = []int{12300, 45600}
)

func init() {
  All = append(All, &Scenario{"net", "Network tests", runNet})
}

func runNet() {
  for _, port := range OPENPORTS {
    go openPort(port)
  }

  c := time.Tick(5 * time.Second)
  for _ = range c {
    // HTTP GETs
    for _, url := range GETS {
      http.Get(url)
    }

    // DNS requests
    for _, name := range DNSREQS {
      net.LookupCNAME(name)
    }
  }
}

func openPort(port int) {
  ln, err := net.Listen("tcp", ":"+strconv.FormatInt(int64(port), 10))
  if err != nil {
    log.Println("Failed to open port", port, err)
    return
  }

  log.Println("Listening on TCP port", port)

  for {
    conn, err := ln.Accept()
    if err != nil {
      continue
    }

    go suckData(conn)
  }
}

func suckData(conn net.Conn) {
  log.Println("New connection from", conn.RemoteAddr())
  ioutil.ReadAll(conn)
  conn.Close()
}
