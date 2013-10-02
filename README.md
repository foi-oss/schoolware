# schoolware


*Malware for schools (and universities)*

## Usage

*Executable files can be downloaded from [schoolware/releases](https://github.com/foi-oss/schoolware/releases).*

To run a scenario, specify it's name and, optionally, any additional arguments it takes.

    C:>schoolware --scenario=files --files-paths=test.txt run
    
`run` command will immediately run the scenario in the foreground.

In order to run a scenario in background you first have to install the `schoolware`
service by registering it with [svchost](http://en.wikipedia.org/wiki/Svchost).

    C:>schoolware --scenario=net --net-ports=123456 install
    
`install` command will register the `schoolware` binary while other flags (like `--scenario` and `--net-ports` above)
will be used during service's start-up. Once the service is registered, start-up flags cannot be altered. 

To start the service in the background, issue the following command:

    C:>schoolware start
    
Please note, due to limitations of svchost, schoolware binary has to be located on start-up disk (C: drive).

In order to stop and remove the service, use the `stop` and `remove` commands respectively.
After the service has been removed, it can be registered again with different start-up parameters.

For list of all scenarios and available options, consult built-in `--help`.

## Building

* [Install Go](http://golang.org/doc/install)
* Prepare your Go environment
  - Create a directory structure: `$GOPATH/{src,pkg}`
  - Adjust your environment variables (add `$GOPATH`)
  - [Windows users can follow this tutorial](http://support.microsoft.com/kb/310519)
* With new `GOPATH` set, go-get this repository
  - `go get github.com/foi-oss/schoolware`
* Build the project
  - navigate to `$GOPATH/src/github.com/foi-oss/schoolware/bin`
  - and run `go build -o schoolware.exe`
