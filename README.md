## vogen - Value Object Generator in golang
<!-- ALL-CONTRIBUTORS-BADGE:START - Do not remove or modify this section -->
[![All Contributors](https://img.shields.io/badge/all_contributors-1-orange.svg?style=flat-square)](#contributors-)
<!-- ALL-CONTRIBUTORS-BADGE:END -->

[![Go Reference](https://pkg.go.dev/badge/github.com/nao1215/vogen.svg)](https://pkg.go.dev/github.com/nao1215/vogen)
![Coverage](https://raw.githubusercontent.com/nao1215/octocovs-central-repo/main/badges/nao1215/vogen/coverage.svg)
[![reviewdog](https://github.com/nao1215/vogen/actions/workflows/reviewdog.yml/badge.svg)](https://github.com/nao1215/vogen/actions/workflows/reviewdog.yml)
[![MultiPlatformUnitTest](https://github.com/nao1215/vogen/actions/workflows/unit_test.yml/badge.svg)](https://github.com/nao1215/vogen/actions/workflows/unit_test.yml)

The vogen library is to generate value objects in golang. The vogen will automatically generate files with Value Objects defined.
  
The vogen automatically generates Getter, Constructor, Constructor with Validation, and Equal() based on metadata (vogen.ValueObject).

## Supported OS and go version
- OS: Linux, macOS, Windows
- Go: 1.22 or later

## Example
### Implement a value object metadata

Firstly, write your value object metadata. Here is an example: gen/main.go

```go
package main

import (
	"fmt"
	"path/filepath"

	"github.com/nao1215/vogen"
)

//go:generate go run main.go

func main() {
	// Step 1: Create a Vogen instance with custom file path and package name.
	// By default, the file path is "value_objects.go" and the package name is "vo".
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
				{
					Name: "Name", Type: "string",
					Comments: []string{"Name is the name of the person."},
					Validators: []vogen.Validator{
						vogen.NewStringLengthValidator(0, 120),
					},
				},
				{
					Name: "Age", Type: "int",
					Comments: []string{"Age is the age of the person."},
					Validators: []vogen.Validator{
						vogen.NewPositiveValueValidator(),
						vogen.NewMaxValueValidator(120),
					}},
			},
			Comments: []string{
				"Person is a Value Object to describe the feature of vogen.",
				"This is sample comment.",
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
}
```

If you run 'go generate ./...', the following code will be generated in the `example_vo.go` file.

```go
// Code generated by vogen. DO NOT EDIT.
package vo_example

import (
	"fmt"
)

// Person is a Value Object to describe the feature of vogen.
// This is sample comment.
type Person struct {
	// Name is the name of the person.
	name string
	// Age is the age of the person.
	age int
}

// NewPerson creates a new instance of Person.
func NewPerson(name string, age int) Person {
	return Person{name: name, age: age}
}

// NewPersonStrictly creates a new instance of Person with validation.
func NewPersonStrictly(name string, age int) (Person, error) {
	o := Person{name: name, age: age}
	if len(o.name) < 0 || len(o.name) > 120 {
		return fmt.Errorf("struct 'Person' field 'Name' length is out of range: %d", len(o.name))
	}
	if o.age < 0 {
		return fmt.Errorf("struct 'Person' field 'Age' value is negative: %d", age)
	}
	if o.age > 120 {
		return fmt.Errorf("struct 'Person' field 'Age' value exceeds the maximum value: %d", age)
	}
	return o, nil
}

// Name returns the name field.
func (o Person) Name() string {
	return o.name
}

// Age returns the age field.
func (o Person) Age() int {
	return o.age
}

// Equal checks if two Person objects are equal.
func (o Person) Equal(other Person) bool {
	return o.Name() == other.Name() && o.Age() == other.Age()
}

// Address represents a value object.
type Address struct {
	city string
}

// NewAddress creates a new instance of Address.
func NewAddress(city string) Address {
	return Address{city: city}
}

// City returns the city field.
func (o Address) City() string {
	return o.city
}

// Equal checks if two Address objects are equal.
func (o Address) Equal(other Address) bool {
	return o.City() == other.City()
}
```

### Validation list

| Validator | Description |
| --- | --- |
| NewPositiveValueValidator() | Check if the value is positive. |
| NewNegativeValueValidator() | Check if the value is negative. |
| NewMaxValueValidator(max int) | Check if the value is less than to the maximum value. |
| NewMinValueValidator(min int) | Check if the value is greater than to the minimum value. |
| NewStringLengthValidator(min, max int) | Check if the length of the string is within the specified range. |

## License

[MIT License](./LICENSE)

## Contribution

First off, thanks for taking the time to contribute! See [CONTRIBUTING.md](./CONTRIBUTING.md) for more information. Contributions are not only related to development. For example, GitHub Star motivates me to develop! Please feel free to contribute to this project.


## Contributors ✨

Thanks goes to these wonderful people ([emoji key](https://allcontributors.org/docs/en/emoji-key)):

<!-- ALL-CONTRIBUTORS-LIST:START - Do not remove or modify this section -->
<!-- prettier-ignore-start -->
<!-- markdownlint-disable -->
<table>
  <tbody>
    <tr>
      <td align="center" valign="top" width="14.28%"><a href="https://debimate.jp/"><img src="https://avatars.githubusercontent.com/u/22737008?v=4?s=75" width="75px;" alt="CHIKAMATSU Naohiro"/><br /><sub><b>CHIKAMATSU Naohiro</b></sub></a><br /><a href="https://github.com/nao1215/vogen/commits?author=nao1215" title="Code">💻</a></td>
    </tr>
  </tbody>
</table>

<!-- markdownlint-restore -->
<!-- prettier-ignore-end -->

<!-- ALL-CONTRIBUTORS-LIST:END -->

This project follows the [all-contributors](https://github.com/all-contributors/all-contributors) specification. Contributions of any kind welcome!
