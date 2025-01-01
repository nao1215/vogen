package vogen

import "fmt"

// Validator represents a validator.
type Validator interface {
	write(vo *Vogen, structName string, field Field)
}

// PositiveValueValidator is a Positive Value Validator.
type PositiveValueValidator struct{}

// NewPositiveValueValidator returns a new Positive Value Validator.
func NewPositiveValueValidator() Validator {
	return &PositiveValueValidator{}
}

// write writes the Positive Value Validator to the code.
func (pvv *PositiveValueValidator) write(vo *Vogen, structName string, field Field) {
	vo.code = append(vo.code, "\tif o."+field.lowerCamelCase()+" < 0 {\n")
	vo.code = append(vo.code, "\t\treturn fmt.Errorf(\"struct '"+structName+"' field '"+field.Name+"' value is negative: %d\", "+field.lowerCamelCase()+")\n")
	vo.code = append(vo.code, "\t}\n")
}

// NegativeValueValidator is a Negative Value Validator.
type NegativeValueValidator struct{}

// NewNegativeValueValidator returns a new Negative Value Validator.
// NegativeValueValidator checks if the value is negative.
func NewNegativeValueValidator() Validator {
	return &NegativeValueValidator{}
}

// write writes the Negative Value Validator to the code.
func (nvv *NegativeValueValidator) write(vo *Vogen, structName string, field Field) {
	vo.code = append(vo.code, "\tif o."+field.lowerCamelCase()+" >= 0 {\n")
	vo.code = append(vo.code, "\t\treturn fmt.Errorf(\"struct '"+structName+"' field '"+field.Name+"' value is positive: %d\", "+field.lowerCamelCase()+")\n")
	vo.code = append(vo.code, "\t}\n")
}

// MaxValueValidator is a Max Value Validator.
type MaxValueValidator struct {
	maxValue int
}

// NewMaxValueValidator returns a new Max Value Validator.
// MaxValueValidator checks if the value exceeds the maximum value.
func NewMaxValueValidator(m int) Validator {
	return &MaxValueValidator{maxValue: m}
}

// write writes the Max Value Validator to the code.
func (mvv *MaxValueValidator) write(vo *Vogen, structName string, field Field) {
	vo.code = append(vo.code, "\tif o."+field.lowerCamelCase()+" > "+fmt.Sprintf("%d", mvv.maxValue)+" {\n")
	vo.code = append(vo.code, "\t\treturn fmt.Errorf(\"struct '"+structName+"' field '"+field.Name+"' value exceeds the maximum value: %d\", "+field.lowerCamelCase()+")\n")
	vo.code = append(vo.code, "\t}\n")
}

// MinValueValidator is a Min Value Validator.
type MinValueValidator struct {
	minValue int
}

// NewMinValueValidator returns a new Min Value Validator.
// MinValueValidator checks if the value is less than the minimum value.
func NewMinValueValidator(m int) Validator {
	return &MinValueValidator{minValue: m}
}

// write writes the Min Value Validator to the code.
func (miv *MinValueValidator) write(vo *Vogen, structName string, field Field) {
	vo.code = append(vo.code, "\tif o."+field.lowerCamelCase()+" < "+fmt.Sprintf("%d", miv.minValue)+" {\n")
	vo.code = append(vo.code, "\t\treturn fmt.Errorf(\"struct '"+structName+"' field '"+field.Name+"' value is less than the minimum value: %d\", "+field.lowerCamelCase()+")\n")
	vo.code = append(vo.code, "\t}\n")
}

// RangeValueValidator is a Range Value Validator.
type RangeValueValidator struct {
	minValue int
	maxValue int
}

// NewRangeValueValidator returns a new Range Value Validator.
// RangeValueValidator checks if the value is out of range.
func NewRangeValueValidator(minV, maxV int) Validator {
	return &RangeValueValidator{minValue: minV, maxValue: maxV}
}

// write writes the Range Value Validator to the code.
func (rvv *RangeValueValidator) write(vo *Vogen, structName string, field Field) {
	vo.code = append(vo.code, "\tif o."+field.lowerCamelCase()+" < "+fmt.Sprintf("%d", rvv.minValue)+" || o."+field.lowerCamelCase()+" > "+fmt.Sprintf("%d", rvv.maxValue)+" {\n")
	vo.code = append(vo.code, "\t\treturn fmt.Errorf(\"struct '"+structName+"' field '"+field.Name+"' value is out of range: %d\", "+field.lowerCamelCase()+")\n")
	vo.code = append(vo.code, "\t}\n")
}

// StringLengthValidator is a String Length Validator.
type StringLengthValidator struct {
	minLength int
	maxLength int
}

// NewStringLengthValidator returns a new String Length Validator.
// StringLengthValidator checks if the string length is out of range.
func NewStringLengthValidator(minL, maxL int) Validator {
	return &StringLengthValidator{minLength: minL, maxLength: maxL}
}

// write writes the String Length Validator to the code.
func (slv *StringLengthValidator) write(vo *Vogen, structName string, field Field) {
	vo.code = append(vo.code, "\tif len(o."+field.lowerCamelCase()+") < "+fmt.Sprintf("%d", slv.minLength)+" || len(o."+field.lowerCamelCase()+") > "+fmt.Sprintf("%d", slv.maxLength)+" {\n")
	vo.code = append(vo.code, "\t\treturn fmt.Errorf(\"struct '"+structName+"' field '"+field.Name+"' length is out of range: %d\", len(o."+field.lowerCamelCase()+"))\n")
	vo.code = append(vo.code, "\t}\n")
}
