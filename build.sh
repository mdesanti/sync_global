go get -d ./...
go build -o bin/config config/*.go
go build -o bin/ControlGas ControlGas/*.go
go build -o bin/application sync_global.go
