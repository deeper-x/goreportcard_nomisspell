[![Go Report Card](https://goreportcard.com/badge/gojp/goreportcard)](https://goreportcard.com/report/gojp/goreportcard) [![License](https://img.shields.io/badge/License-Apache%202.0-blue.svg)](https://github.com/gojp/goreportcard/blob/master/LICENSE)

# Go Report Card

A web application that generates a report on the quality of an open source Go project. It uses several measures, including `gofmt`, `go vet`, `go lint` and `gocyclo`. 

### Installation

```
git clone https://github.com/gojp/goreportcard.git
cd goreportcard
make install
```

Now run:

```
GRC_DATABASE_PATH=./db make start
```

and you should see

```
Running on :8000...
```

Navigate to `localhost:8000` and you should see the Go Report Card front page.

### Command Line Interface

There is also a CLI available for grading applications on your local machine.

Example usage:
```
git clone https://github.com/gojp/goreportcard.git
cd goreportcard
make install
go install ./cmd/goreportcard-cli
goreportcard-cli
```

```
Grade: A+ (99.9%)
Files: 362
Issues: 2
gofmt: 100%
go_vet: 99%
gocyclo: 99%
golint: 100%
ineffassign: 100%
license: 100%
misspell: 100%
```

Verbose output:

```
goreportcard-cli -v
```

```
Grade: A+ (99.9%)
Files: 332
Issues: 2
gofmt: 100%
go_vet: 99%
go_vet  vendor/github.com/prometheus/client_golang/prometheus/desc.go:25
        error: cannot find package "github.com/prometheus/client_model/go" in any of: (vet)

gocyclo: 99%
gocyclo download/download.go:22
        warning: cyclomatic complexity 17 of function download() is high (> 15) (gocyclo)

golint: 100%
ineffassign: 100%
license: 100%
misspell: 100%
```
