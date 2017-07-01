# Stackpathcdnbeat

Welcome to Stackpathcdnbeat.

Ensure that this folder is at the following location:
`${GOPATH}/github.com/bmconklin/stackpathcdnbeat`

## Getting Started with Stackpathcdnbeat

### Requirements

* [Golang](https://golang.org/dl/) 1.7

### Init Project
To get running with Stackpathcdnbeat and also install the
dependencies, run the following command:

```
make setup
```

It will create a clean git history for each major step. Note that you can always rewrite the history if you wish before pushing your changes.

To push Stackpathcdnbeat in the git repository, run the following commands:

```
git remote set-url origin https://github.com/bmconklin/stackpathcdnbeat
git push origin master
```

For further development, check out the [beat developer guide](https://www.elastic.co/guide/en/beats/libbeat/current/new-beat.html).

### Build

To build the binary for Stackpathcdnbeat run the command below. This will generate a binary
in the same directory with the name stackpathcdnbeat.

```
make
```


### Run

To run Stackpathcdnbeat with debugging output enabled, run:

```
./stackpathcdnbeat -c stackpathcdnbeat.yml -e -d "*"
```


### Test

To test Stackpathcdnbeat, run the following command:

```
make testsuite
```

alternatively:
```
make unit-tests
make system-tests
make integration-tests
make coverage-report
```

The test coverage is reported in the folder `./build/coverage/`

### Update

Each beat has a template for the mapping in elasticsearch and a documentation for the fields
which is automatically generated based on `etc/fields.yml`.
To generate etc/stackpathcdnbeat.template.json and etc/stackpathcdnbeat.asciidoc

```
make update
```


### Cleanup

To clean  Stackpathcdnbeat source code, run the following commands:

```
make fmt
make simplify
```

To clean up the build directory and generated artifacts, run:

```
make clean
```


### Clone

To clone Stackpathcdnbeat from the git repository, run the following commands:

```
mkdir -p ${GOPATH}/github.com/bmconklin/stackpathcdnbeat
cd ${GOPATH}/github.com/bmconklin/stackpathcdnbeat
git clone https://github.com/bmconklin/stackpathcdnbeat
```


For further development, check out the [beat developer guide](https://www.elastic.co/guide/en/beats/libbeat/current/new-beat.html).


## Packaging

The beat frameworks provides tools to crosscompile and package your beat for different platforms. This requires [docker](https://www.docker.com/) and vendoring as described above. To build packages of your beat, run the following command:

```
make package
```

This will fetch and create all images required for the build process. The hole process to finish can take several minutes.