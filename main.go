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

	// 250 and 13 obtained via grid search generates 592904kb in go14.2 on linux
	_ = gen.GenerateThrift(sandboxDir, "thrift", 250, 13)
	_ = gen.GenerateMain(sandboxDir, genImport, "thrift", 250)
}
