package gen

import (
	"fmt"
	"html/template"
	"log"
	"os"
	"path/filepath"
)

type Main struct {
	Imports []string
	Main    string
}

func CreateMain(prefix string, n int) Main {
	var imports []string
	for i := 0; i < n; i++ {
		path := fmt.Sprintf("%v/thrift%v", prefix, i)
		imports = append(imports, path)
	}

	mainFunc := `func main() {
		// large go program
	}`

	return Main{
		Imports: imports,
		Main:    mainFunc,
	}
}

func GenerateMain(repoRoot, genImport, prefix string, n int) error {
	data := CreateMain(genImport, n)

	// Create a new template and parse the letter into it.
	files := []string{"templates/main.tmpl"}
	t := template.Must(template.New("main.tmpl").ParseFiles(files...))

	// Execute the template for each recipient.
	err := t.Execute(os.Stdout, data)
	if err != nil {
		log.Println("executing template:", err)
		return err
	}

	if _, err := os.Stat(repoRoot); os.IsNotExist(err) {
		os.MkdirAll(repoRoot, os.ModePerm)
	}

	writeFile := filepath.Join(repoRoot, "main.go")

	f, err := os.Create(writeFile)
	if err != nil {
		panic(err)
		return fmt.Errorf("create file: %v", err)
	}
	defer f.Close()

	// Execute the template for each recipient.
	err = t.Execute(f, data)
	if err != nil {
		log.Println("executing template:", err)
		return err
	}

	return nil
}
