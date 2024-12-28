package vogen_test

import (
	"fmt"
	"path/filepath"

	"github.com/nao1215/vogen"
)

func ExampleVogen_Generate() {
	// Step 1: Create a Vogen instance with custom file path and package name
	gen, err := vogen.New(
		vogen.WithFilePath(filepath.Join("testdata", "example_output.go")),
		vogen.WithPackageName("vo_example"),
	)
	if err != nil {
		fmt.Printf("Failed to create Vogen instance: %v\n", err)
		return
	}

	// Step 2: Append the ValueObject definition
	if err := gen.AppendValueObjects(
		vogen.ValueObject{
			StructName: "Person",
			Fields: []vogen.Field{
				{Name: "Name", Type: "string", Comments: []string{"Name is the name of the person."}},
				{Name: "Age", Type: "int", Comments: []string{"Age is the age of the person."}},
			},
			Comments: []string{
				"Person is a Value Object to describe the feature of vogen.",
				"This is sample comment.",
			},
		},
		// Use auto generated comments.
		vogen.ValueObject{
			StructName: "Address",
			Fields: []vogen.Field{
				{Name: "City", Type: "string"},
			},
		},
	); err != nil {
		fmt.Printf("Failed to append ValueObject: %v\n", err)
		return
	}

	// Step 3: Generate the code
	if err := gen.Generate(); err != nil {
		fmt.Printf("Failed to generate code: %v\n", err)
		return
	}

	fmt.Println("Code generated successfully. Check 'example_output.go' for the output.")
	// Output:
	// Code generated successfully. Check 'example_output.go' for the output.
}
