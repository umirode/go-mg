# Go Microservice Generator with gRPC and net package

This package for generate base code for microservice.

## Install
`go get github.com/umirode/go-mg-net`

## Usage 
`go-mg-net -name=greeter -network=tcp -address=:56001 -output=/home/user/projects`

For help: `go-mg-net -help`

## Protobuf files
Add your repositories with protobuf files in `.modules` file (one line - one repository).
Then run `make install-modules`, this command load repositories from `.modules` file and generate Golang code based on `.proto` files.
