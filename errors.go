package vogen

import "errors"

var (
	// ErrStructNameEmpty is an error that occurs when the StructName is not set.
	ErrStructNameEmpty = errors.New("vogen: ValueObject 'StructName() is not set")
	// ErrInvalidField is an error that occurs when the Field is not set.
	ErrInvalidField = errors.New("vogen: ValueObject 'Field' is not set")
	// ErrInvalidFieldName is an error that occurs when the Field 'Name' is not set.
	ErrInvalidFieldName = errors.New("vogen: ValueObject 'Field.Name' is not set")
	// ErrInvalidFieldType is an error that occurs when the Field 'Type' is not set.
	ErrInvalidFieldType = errors.New("vogen: ValueObject 'Field.Type' is not set")
	// ErrInvalidFilePath is an error that occurs when the file path is not set.
	ErrInvalidFilePath = errors.New("vogen: file path is not set")
	// ErrInvalidPackageName is an error that occurs when the package name is not set.
	ErrInvalidPackageName = errors.New("vogen: package name is not set")
)
