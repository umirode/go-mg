# {{.Name | ToUpper}} microservice

## Protobuf files
Add your repositories with protobuf files in `.modules` file (one line - one repository).
Then run `make install-modules`, this command load repositories from `.modules` file and generate Golang code based on `.proto` files.
