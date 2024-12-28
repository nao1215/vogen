package vogen_test

import (
	"fmt"
	"path/filepath"

	"github.com/nao1215/vogen"
)

func ExampleVogen_Generate() {
	// Step 1: Define the ValueObject configuration
	vo := vogen.ValueObject{
		StructName: "Person",
		Fields: []vogen.Field{
			{Name: "Name", Type: "string"},
			{Name: "Age", Type: "int"},
		},
		Imports: []string{"fmt"},
	}

	// Step 2: Create a Vogen instance with custom file path and package name
	gen, err := vogen.New(
		vogen.WithFilePath(filepath.Join("testdata", "example_output.go")),
		vogen.WithPackageName("vo_example"),
	)
	if err != nil {
		fmt.Printf("Failed to create Vogen instance: %v\n", err)
		return
	}

	// Step 3: Append the ValueObject definition
	if err := gen.AppendValueObjects(vo); err != nil {
		fmt.Printf("Failed to append ValueObject: %v\n", err)
		return
	}

	// Step 4: Generate the code
	if err := gen.Generate(); err != nil {
		fmt.Printf("Failed to generate code: %v\n", err)
		return
	}

	// Step 5: Output success message
	fmt.Println("Code generated successfully. Check 'example_output.go' for the output.")

	// Output:
	// Code generated successfully. Check 'example_output.go' for the output.
}
