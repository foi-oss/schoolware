package scenarios

import (
  "bitbucket.org/jol/service/stdservice"
  "flag"
  "io/ioutil"
  "os"
  "strings"
  "time"
)

var (
  // text to be written to files
  CONTENTS = []byte("<3")

  // extra files to which data should be written, passed in via command-line
  FILES = flag.String("files-paths", "/schoolware.txt", "comma-separated list of paths relative to ~ (files scenario)")
)

func init() {
  All = append(All, &Scenario{"files", "Write data to files", runLocalfiles})
}

func runLocalfiles(s *stdservice.Config) error {
  l := s.Logger()
  files := strings.Split(*FILES, ",")
  l.Info("Writing to: " + HomeDir + "{" + strings.Join(files, ", ") + "}")

  for {
    for _, path := range files {
      path = strings.TrimSpace(path)
      ioutil.WriteFile(HomeDir+path, CONTENTS, os.ModePerm)
      ioutil.WriteFile("D:"+path, CONTENTS, os.ModePerm)
    }

    time.Sleep(5 * time.Second)
  }

  return nil
}
