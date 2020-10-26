package gen

import (
	"fmt"
	"html/template"
	"log"
	"os"
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

	for i := 1; i <= n; i++ {
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
		for i := 1; i <= n; i++ {
			structTypeKey := fmt.Sprintf("Structure%v", i)
			sm := Method{
				ReturnType: structTypeKey,
				Name:       fmt.Sprintf("SMethod%v", i),
				Params:     fmt.Sprintf("%v: i32 iparam", 1),
			}
			im := Method{
				ReturnType: fmt.Sprintf("IntType%v", i),
				Name:       fmt.Sprintf("IMethod%v", i),
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

func Generate(_ string) error {
	data := CreateThrift(2)

	// Create a new template and parse the letter into it.
	files := []string{"templates/thrift.tmpl"}
	t := template.Must(template.New("thrift.tmpl").ParseFiles(files...))

	// Execute the template for each recipient.
	err := t.Execute(os.Stdout, data)
	if err != nil {
		log.Println("executing template:", err)
		return err
	}
	return nil
}

/*
// Prepare some data to insert into the template.
type Recipient struct {
	Name, Gift string
	Attended   bool
}
var recipients = []Recipient{
	{"Aunt Mildred", "bone china tea set", true},
	{"Uncle John", "moleskin pants", false},
	{"Cousin Rodney", "", false},
}
*/
