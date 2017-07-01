# Stackpathcdnbeat

Welcome to Stackpathcdnbeat. Stackpathcdnbeat is a custom beat to fetch logs from the public API of Stackpath or MaxCDN and store them privately.
The code is currently beta level quality and may require some bug fixes and optimizations. Chances are, this tool will not be efficient for StackPath customers with huge amounts of requests per second, but should be sufficient for most StackPath clients.

## Getting Started with Stackpathcdnbeat

### Requirements

* [Golang](https://golang.org/dl/) 1.7

### Install

To install the Stackpathcdnbeat, run

```
go get github.com/bmconklin/stackpathcdnbeat
```

### Configs

The config file needs to be filled in with at a bare minimum the API credentials you'll use to fetch the logs, which can be gotten from the StackPath or MaxCDN control panel.

The config file is located in
```
$GOPATH/src/github.com/bmconklin/stackpathcdnbeat/stackpathcdnbeat.yml
```

### Build

To compile the Stackpathcdnbeat, run:

```
cd $GOPATH/src/github.com/bmconklin/stackpathcdnbeat
go build
```

### Run

To run Stackpathcdnbeat with debugging output enabled, run:

```
$GOPATH/src/github.com/bmconklin/stackpathcdnbeat/stackpathcdnbeat -c $GOPATH/src/github.com/bmconklin/stackpathcdnbeat/stackpathcdnbeat.yml -e -d "*"
```

### Clone

To clone Stackpathcdnbeat from the git repository, run the following commands:

```
mkdir -p ${GOPATH}/github.com/bmconklin/stackpathcdnbeat
cd ${GOPATH}/github.com/bmconklin/stackpathcdnbeat
git clone https://github.com/bmconklin/stackpathcdnbeat
```
