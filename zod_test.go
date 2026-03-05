
// Package zod provides a simple validation library.
package zod

import (
    "context"
    "errors"
    "testing"
)

// ErrValidation indicates a validation error.
var ErrValidation = errors.New("validation error")

// Object represents a validation object.
func Object() *Object {
    return &Object{}
}

// String represents a string validation.
func String() *String {
    return &String{}
}

// TestObjectValidate tests object validation.
func TestObjectValidate(t *testing.T) {
    tests := []struct {
        name     string
        schema  *Object
        user     interface{}
        wantErr  bool
        wantErrMsg string
    }{
        {
            name: "valid user",
            schema: func() *Object {
                schema := Object().strict()
                schema.fields["name"] = String().ref("name").min(1, "Name is required").max(50, "Name is too long")
                schema.fields["email"] = String().ref("email").email("Invalid email")
                return schema
            }(),
            user: User{
                Name:  "John Doe",
                Email: "john@example.com",
            },
            wantErr: false,
        },
        {
            name: "invalid email",
            schema: func() *Object {
                schema := Object().strict()
                schema.fields["email"] = String().ref("email").email("Invalid email")
                return schema
            }(),
            user: User{
                Email: "john",
            },
            wantErr:  true,
            wantErrMsg: "Invalid email",
        },
        {
            name: "nil user",
            schema: Object().strict(),
            user:     nil,
            wantErr:  true,
            wantErrMsg: "validation error",
        },
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            ctx := context.Background()
            err := tt.schema.Validate(ctx, tt.user)
            if (err != nil) != tt.wantErr {
                t.Errorf("Validate() error = %v, wantErr %v", err, tt.wantErr)
            }
            if tt.wantErr && err.Error() != tt.wantErrMsg {
                t.Errorf("Validate() error message = %v, want %v", err.Error(), tt.wantErrMsg)
            }
        })
    }
}

// TestObjectValidateMinLength tests object validation min length.
func TestObjectValidateMinLength(t *testing.T) {
    schema := Object().strict()
    schema.fields["name"] = String().ref("name").min(5, "Name is too short")

    type User struct {
        Name string `json:"name""
    }

    user := User{
        Name: "John",
    }

    ctx := context.Background()
    err := schema.Validate(ctx, user)
    if err == nil {
        t.Errorf("Validate() error = nil, want %v", "Name is too short")
    }
}

// TestObjectValidateMaxLength tests object validation max length.
func TestObjectValidateMaxLength(t *testing.T) {
    schema := Object().strict()
    schema.fields["name"] = String().ref("name").max(10, "Name is too long")

    type User struct {
        Name string `json:"name""
    }

    user := User{
        Name: "John Doe John",
    }

    ctx := context.Background()
    err := schema.Validate(ctx, user)
    if err == nil {
        t.Errorf("Validate() error = nil, want %v", "Name is too long")
    }
}

// TestObjectValidateEmail tests object validation email.
func TestObjectValidateEmail(t *testing.T) {
    schema := Object().strict()
    schema.fields["email"] = String().ref("email").email("Invalid email")

    type User struct {
        Email string `json:"email""
    }

    user := User{
        Email: "john",
    }

    ctx := context.Background()
    err := schema.Validate(ctx, user)
    if err == nil {
        t.Errorf("Validate() error = nil, want %v", "Invalid email")
    }
}

// BenchmarkObjectValidate benchmarks object validation.
func BenchmarkObjectValidate(b *testing.B) {
    schema := Object().strict()
    schema.fields["name"] = String().ref("name").min(1, "Name is required").max(50, "Name is too long")
    schema.fields["email"] = String().ref("email").email("Invalid email")

    type User struct {
        Name  string `json:"name""
        Email string `json:"email""
    }

    user := User{
        Name:  "John Doe",
        Email: "john@example.com",
    }

    ctx := context.Background()
    for i := 0; i < b.N; i++ {
        schema.Validate(ctx, user)
    }
}
