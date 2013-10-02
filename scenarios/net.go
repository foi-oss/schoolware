package scenarios

import (
  "bitbucket.org/jol/service/stdservice"
  "flag"
  "io/ioutil"
  "log"
  "net"
  "net/http"
  "strconv"
  "strings"
  "time"
)

var (
  GETS      = flag.String("net-httpgets", "http://foi.hr", "list of addresses to GET (net scenario)")
  DNSREQS   = flag.String("net-dns", "bing.com,unizg.hr", "domains on which CNAME lookup will be performed (net scenario)")
  OPENPORTS = flag.String("net-ports", "12300,45600", "TCP ports to open (net scenario)")
)

func init() {
  All = append(All, &Scenario{"net", "Perform metwork tests (HTTP, DNS, TCP)", runNet})
}

func runNet(s *stdservice.Config) error {
  //l := s.Logger()

  for _, port := range strings.Split(*OPENPORTS, ",") {
    p, _ := strconv.ParseInt(port, 10, 64)
    go openPort(p)
  }

  c := time.Tick(5 * time.Second)
  gets := strings.Split(*GETS, ",")
  dnsreqs := strings.Split(*DNSREQS, ",")
  for _ = range c {
    // HTTP GETs
    for _, url := range gets {
      http.Get(strings.TrimSpace(url))
    }

    // DNS requests
    for _, name := range dnsreqs {
      net.LookupCNAME(name)
    }
  }

  return nil
}

func openPort(port int64) error {
  ln, err := net.Listen("tcp", ":"+strconv.FormatInt(port, 10))
  if err != nil {
    log.Println("Failed to open port", port, err)
    return err
  }

  log.Println("Listening on TCP port", port)

  for {
    conn, err := ln.Accept()
    if err != nil {
      continue
    }

    go suckData(conn)
  }

  return nil
}

func suckData(conn net.Conn) {
  addr := conn.RemoteAddr()

  log.Println("New connection from", addr)
  ioutil.ReadAll(conn)
  conn.Close()
  log.Println("Closed connection from", addr)
}
