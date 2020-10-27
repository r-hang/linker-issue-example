package main

import (
	"os"
	"path/filepath"

	"github.com/r-hang/linker-issue-example/gen"
)

func main() {
	path, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	sandboxDir := filepath.Join(path, "experiment-sandbox")
	genImport := filepath.Join("github.com/r-hang/linker-issue-example", "experiment-sandbox", "thriftgen")

	_ = gen.GenerateThrift(sandboxDir, "thrift", 250, 10)
	_ = gen.GenerateMain(sandboxDir, genImport, "thrift", 250)
}
