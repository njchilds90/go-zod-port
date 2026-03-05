package zod

import (
    "errors"
    "reflect"
    "strings"
)

// Object represents an object schema.
type Object struct {
    fields map[string]*Field
    strict bool
}

// NewObject returns a new object schema.
func Object() *Object {
    return &Object{
        fields: make(map[string]*Field),
    }
}

// Field represents a field in an object schema.
type Field struct {
    name  string
    rules []*Rule
}

// NewField returns a new field.
func NewField(name string) *Field {
    return &Field{
        name: name,
    }
}

// Rule represents a validation rule.
type Rule struct {
    kind string
    args []interface{}
}

// NewRule returns a new rule.
func NewRule(kind string, args ...interface{}) *Rule {
    return &Rule{
        kind: kind,
        args: args,
    }
}

// String returns a new string field.
func String() *Field {
    return NewField("string")
}

// ref sets the reference of a field.
func (f *Field) ref(name string) *Field {
    f.name = name
    return f
}

// min sets the minimum length of a string field.
func (f *Field) min(length int, message string) *Field {
    f.rules = append(f.rules, NewRule("min", length, message))
    return f
}

// max sets the maximum length of a string field.
func (f *Field) max(length int, message string) *Field {
    f.rules = append(f.rules, NewRule("max", length, message))
    return f
}

// email sets a field to be an email address.
func (f *Field) email(message string) *Field {
    f.rules = append(f.rules, NewRule("email", message))
    return f
}

// Validate validates the given data against the schema.
func (o *Object) Validate(data interface{}) error {
    rv := reflect.ValueOf(data)
    if rv.Kind() != reflect.Struct {
        return errors.New("data is not a struct")
    }

    for fieldName, field := range o.fields {
        fv := rv.FieldByName(fieldName)
        if !fv.IsValid() {
            return errors.New("field not found")
        }

        for _, rule := range field.rules {
            switch rule.kind {
            case "min":
                if fv.Len() < rule.args[0].(int) {
                    return errors.New(rule.args[1].(string))
                }
            case "max":
                if fv.Len() > rule.args[0].(int) {
                    return errors.New(rule.args[1].(string))
                }
            case "email":
                if !strings.Contains(fv.String(), "@") {
                    return errors.New(rule.args[0].(string))
                }
            }
        }
    }

    return nil
}

// strict sets the schema to strict mode.
func (o *Object) strict() *Object {
    o.strict = true
    return o
}