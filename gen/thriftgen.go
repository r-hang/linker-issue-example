package gen

import (
	"fmt"
	"html/template"
	"log"
	"os"
	"path/filepath"
)

type Thrift struct {
	Typedefs []Typedef
	Structs  []Struct
	Services []Service
}

type Typedef struct {
	Name string
	Type string
}

type Field struct {
	Num    int
	Status string
	Type   string
	Name   string
}

type Struct struct {
	Name   string
	Fields []Field
}

type Service struct {
	Name    string
	Methods []Method
}

type Method struct {
	ReturnType string
	Name       string
	Params     string
}

func CreateThrift(n int) Thrift {

	var (
		typeDefs    []Typedef
		structDefs  []Struct
		serviceDefs []Service
	)

	for i := 0; i < n; i++ {
		// Typedefs
		st := Typedef{
			Name: fmt.Sprintf("StrType%v", i),
			Type: "string",
		}
		it := Typedef{
			Name: fmt.Sprintf("IntType%v", i),
			Type: "i32",
		}
		typeDefs = append(typeDefs, st)
		typeDefs = append(typeDefs, it)

		// Structs
		structure := Struct{
			Name: fmt.Sprintf("Structure%v", i),
			Fields: []Field{
				Field{
					Num:    1,
					Status: "optional",
					Type:   fmt.Sprintf("StrType%v", i),
					Name:   "strField",
				},
				Field{
					Num:    2,
					Status: "optional",
					Type:   fmt.Sprintf("IntType%v", i),
					Name:   "intField",
				},
			},
		}
		structDefs = append(structDefs, structure)

		// Services: Code gen n**2.
		var methodDefs []Method
		for i := 0; i < n; i++ {
			structTypeKey := fmt.Sprintf("Structure%v", i)
			sm := Method{
				ReturnType: structTypeKey,
				Name:       fmt.Sprintf("sMethod%v", i),
				Params:     fmt.Sprintf("%v: i32 iparam", 1),
			}
			im := Method{
				ReturnType: fmt.Sprintf("IntType%v", i),
				Name:       fmt.Sprintf("iMethod%v", i),
				Params:     fmt.Sprintf("%v: %v sparam", 1, structTypeKey),
			}
			methodDefs = append(methodDefs, sm)
			methodDefs = append(methodDefs, im)
		}
		service := Service{
			Name:    fmt.Sprintf("Service%v", i),
			Methods: methodDefs,
		}
		serviceDefs = append(serviceDefs, service)
	}

	return Thrift{
		Typedefs: typeDefs,
		Structs:  structDefs,
		Services: serviceDefs,
	}
}

// Generate ...
func GenerateThrift(repoRoot, prefix string, numFiles, scale int) error {
	for i := 0; i < numFiles; i++ {
		data := CreateThrift(scale)

		// Create a new template and parse the letter into it.
		files := []string{"templates/thrift.tmpl"}
		t := template.Must(template.New("thrift.tmpl").ParseFiles(files...))

		/*
			// Execute the template for each recipient.
			err := t.Execute(os.Stdout, data)
			if err != nil {
				log.Println("executing template:", err)
				return err
			}
		*/

		writeDir := filepath.Join(repoRoot, "idl")
		if _, err := os.Stat(writeDir); os.IsNotExist(err) {
			err := os.MkdirAll(writeDir, os.ModePerm)
			if err != nil {
				panic(err)
			}
		}

		writeFile := fmt.Sprintf("%v/%v%v.thrift", writeDir, prefix, i)

		f, err := os.Create(writeFile)
		if err != nil {
			panic(err)
			return fmt.Errorf("create file: %v", err)
		}
		defer f.Close()

		// Execute the template for each recipient.
		err = t.Execute(f, data)
		if err != nil {
			panic(err)
			log.Println("executing template:", err)
			return err
		}
	}
	return nil
}
