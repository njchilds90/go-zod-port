// Package zod provides a schema validation library.
// 
// This package provides a simple way to define schemas for validating data.
package zod

import (
    "errors"
    "reflect"
    "strings"
    "context"
)

// ErrInvalidData is returned when the data is not valid.
var ErrInvalidData = errors.New("data is not valid")

// ErrFieldNotFound is returned when a field is not found in the data.
var ErrFieldNotFound = errors.New("field not found")

// ErrDataNotStruct is returned when the data is not a struct.
var ErrDataNotStruct = errors.New("data is not a struct")

// Object represents an object schema.
// 
// Object is used to define a schema for validating data.
type Object struct {
    fields map[string]*Field
    strict bool
}

// NewObject returns a new object schema.
// 
// NewObject creates a new Object instance.
func NewObject() *Object {
    return &Object{
        fields: make(map[string]*Field),
    }
}

// Field represents a field in an object schema.
// 
// Field is used to define a field in a schema.
type Field struct {
    name  string
    rules []*Rule
}

// NewField returns a new field.
// 
// NewField creates a new Field instance.
func NewField(name string) *Field {
    return &Field{
        name: name,
    }
}

// Rule represents a validation rule.
// 
// Rule is used to define a validation rule for a field.
type Rule struct {
    kind string
    args []interface{}
}

// NewRule returns a new rule.
// 
// NewRule creates a new Rule instance.
func NewRule(kind string, args ...interface{}) *Rule {
    return &Rule{
        kind: kind,
        args: args,
    }
}

// String returns a new string field.
// 
// String creates a new Field instance with the name "string".
func String() *Field {
    return NewField("string")
}

// Ref sets the reference of a field.
// 
// Ref sets the name of the field to the given name.
func (f *Field) Ref(name string) *Field {
    if name == "" {
        return f
    }
    f.name = name
    return f
}

// Min sets the minimum length of a string field.
// 
// Min adds a new rule to the field to validate the minimum length.
func (f *Field) Min(length int, message string) *Field {
    if length < 0 {
        return f
    }
    f.rules = append(f.rules, NewRule("min", length, message))
    return f
}

// Max sets the maximum length of a string field.
// 
// Max adds a new rule to the field to validate the maximum length.
func (f *Field) Max(length int, message string) *Field {
    if length < 0 {
        return f
    }
    f.rules = append(f.rules, NewRule("max", length, message))
    return f
}

// Email sets a field to be an email address.
// 
// Email adds a new rule to the field to validate the email address.
func (f *Field) Email(message string) *Field {
    f.rules = append(f.rules, NewRule("email", message))
    return f
}

// Validate validates the given data against the schema.
// 
// Validate checks if the data is valid according to the schema.
func (o *Object) Validate(ctx context.Context, data interface{}) error {
    if data == nil {
        return ErrInvalidData
    }
    rv := reflect.ValueOf(data)
    if rv.Kind() != reflect.Struct {
        return fmt.Errorf("%w: data is not a struct", ErrDataNotStruct)
    }

    for fieldName, field := range o.fields {
        fv := rv.FieldByName(fieldName)
        if !fv.IsValid() {
            return fmt.Errorf("%w: field %s not found", ErrFieldNotFound, fieldName)
        }

        for _, rule := range field.rules {
            switch rule.kind {
            case "min":
                if fv.Len() < rule.args[0].(int) {
                    return fmt.Errorf("%w: field %s has length %d, expected at least %d", ErrInvalidData, fieldName, fv.Len(), rule.args[0].(int))
                }
            case "max":
                if fv.Len() > rule.args[0].(int) {
                    return fmt.Errorf("%w: field %s has length %d, expected at most %d", ErrInvalidData, fieldName, fv.Len(), rule.args[0].(int))
                }
            case "email":
                if !strings.Contains(fv.String(), "@") {
                    return fmt.Errorf("%w: field %s is not an email address", ErrInvalidData, fieldName)
                }
            }
        }
    }
    return nil
}