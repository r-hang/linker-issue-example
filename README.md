# linker-issue-example

$ export GOPATH=~/gocode
$ go mod init
$ go mod tidy
$ go run main.go
$ ./gen_go_from_thrift.sh
$ cd experiment-sandbox
$ go build .
