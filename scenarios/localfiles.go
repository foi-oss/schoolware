package scenarios

import (
  "io/ioutil"
  "log"
  "os"
  "os/user"
  "time"
)

var (
  CONTENTS = []byte("<3")
  // root directory is user's home directory/drive D
  FILES = []string{"/test.txt", "/love.txt"}
)

func init() {
  All = append(All, &Scenario{"files", "Write data to files", runLocalfiles})
}

func runLocalfiles() {
  u, _ := user.Current()

  for {
    for _, path := range FILES {
      log.Println("Writing to", u.HomeDir+path)
      ioutil.WriteFile(u.HomeDir+path, CONTENTS, os.ModePerm)
      ioutil.WriteFile("D:"+path, CONTENTS, os.ModePerm)
    }

    time.Sleep(5 * time.Second)
  }
}
