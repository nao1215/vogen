// Package vogen provides a code generator for Value Objects in Go.
// Value Objects are immutable objects that represent a value.
package vogen

import (
	"errors"
	"testing"
)

func TestNew(t *testing.T) {
	t.Parallel()

	t.Run("error occurs when the file path is not set", func(t *testing.T) {
		t.Parallel()

		_, err := New(WithFilePath(""))
		if !errors.Is(err, ErrInvalidFilePath) {
			t.Errorf("want: %v, got: %v", ErrInvalidFilePath, err)
		}
	})

	t.Run("error occurs when the package name is not set", func(t *testing.T) {
		t.Parallel()

		_, err := New(WithPackageName(""))
		if !errors.Is(err, ErrInvalidPackageName) {
			t.Errorf("want: %v, got: %v", ErrInvalidPackageName, err)
		}
	})
}

func TestVogen_AppendValueObjects(t *testing.T) {
	t.Parallel()

	t.Run("error occurs when the StructName is not set", func(t *testing.T) {
		t.Parallel()

		gen, _ := New()
		err := gen.AppendValueObjects(ValueObject{
			StructName: "",
			Fields: []Field{
				{Name: "Name", Type: "string"},
			},
			Comments: []string{"test comment"},
		})
		if !errors.Is(err, ErrStructNameEmpty) {
			t.Errorf("want: %v, got: %v", ErrStructNameEmpty, err)
		}
	})

	t.Run("error occurs when the Field is not set", func(t *testing.T) {
		t.Parallel()

		gen, _ := New()
		err := gen.AppendValueObjects(ValueObject{
			StructName: "Person",
			Fields:     nil,
			Comments:   []string{"test comment"},
		})
		if !errors.Is(err, ErrInvalidField) {
			t.Errorf("want: %v, got: %v", ErrInvalidField, err)
		}
	})

	t.Run("error occurs when the Field 'Name' is not set", func(t *testing.T) {
		t.Parallel()

		gen, _ := New()
		err := gen.AppendValueObjects(ValueObject{
			StructName: "Person",
			Fields: []Field{
				{Name: "", Type: "string"},
			},
			Comments: []string{"test comment"},
		})
		if !errors.Is(err, ErrInvalidFieldName) {
			t.Errorf("want: %v, got: %v", ErrInvalidFieldName, err)
		}
	})

	t.Run("error occurs when the Field 'Type' is not set", func(t *testing.T) {
		t.Parallel()

		gen, _ := New()
		err := gen.AppendValueObjects(ValueObject{
			StructName: "Person",
			Fields: []Field{
				{Name: "Name", Type: ""},
			},
			Comments: []string{"test comment"},
		})
		if !errors.Is(err, ErrInvalidFieldType) {
			t.Errorf("want: %v, got: %v", ErrInvalidFieldType, err)
		}
	})
}
