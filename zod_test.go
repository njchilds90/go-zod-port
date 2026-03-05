package zod

import (
    "testing"
)

func TestObjectValidate(t *testing.T) {
    schema := Object().strict()
    schema.fields["name"] = String().ref("name").min(1, "Name is required").max(50, "Name is too long")
    schema.fields["email"] = String().ref("email").email("Invalid email")

    type User struct {
        Name  string `json:"name"`
        Email string `json:"email"`
    }

    user := User{
        Name:  "John Doe",
        Email: "john@example.com",
    }

    err := schema.Validate(user)
    if err != nil {
        t.Errorf("Validate() error = %v", err)
    }
}

func TestObjectValidateMinLength(t *testing.T) {
    schema := Object().strict()
    schema.fields["name"] = String().ref("name").min(5, "Name is too short")

    type User struct {
        Name string `json:"name"`
    }

    user := User{
        Name: "John",
    }

    err := schema.Validate(user)
    if err == nil {
        t.Errorf("Validate() error = nil, want %v", "Name is too short")
    }
}

func TestObjectValidateMaxLength(t *testing.T) {
    schema := Object().strict()
    schema.fields["name"] = String().ref("name").max(10, "Name is too long")

    type User struct {
        Name string `json:"name"`
    }

    user := User{
        Name: "John Doe John",
    }

    err := schema.Validate(user)
    if err == nil {
        t.Errorf("Validate() error = nil, want %v", "Name is too long")
    }
}

func TestObjectValidateEmail(t *testing.T) {
    schema := Object().strict()
    schema.fields["email"] = String().ref("email").email("Invalid email")

    type User struct {
        Email string `json:"email"`
    }

    user := User{
        Email: "john",
    }

    err := schema.Validate(user)
    if err == nil {
        t.Errorf("Validate() error = nil, want %v", "Invalid email")
    }
}