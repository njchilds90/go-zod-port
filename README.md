# Go-Zod-Port

[![PkgGoDev](https://pkg.go.dev/badge/github.com/njchilds90/go-zod-port)](https://pkg.go.dev/github.com/njchilds90/go-zod-port)
[![Go Report Card](https://goreportcard.com/badge/github.com/njchilds90/go-zod-port)](https://goreportcard.com/report/github.com/njchilds90/go-zod-port)
[![GoDoc](https://godoc.org/github.com/njchilds90/go-zod-port?status.svg)](https://godoc.org/github.com/njchilds90/go-zod-port)

## Overview

Go-Zod-Port is a Go port of JavaScript's Zod library. It provides a simple way to define and validate data structures, including nested objects and arrays.

## Installation

To install Go-Zod-Port, run the following command:
```bash
go get github.com/njchilds90/go-zod-port/zod
```

## Usage

Here's an example of how to use Go-Zod-Port to define and validate a simple data structure:
```go
package main

import (
	"context"
	"fmt"
	"github.com/njchilds90/go-zod-port/zod"
)

type User struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

func main() {
	// Define the schema
	schema := zod.Object(
		zod.String().ref("name").min(1, "Name is required").max(50, "Name is too long"),
		zod.String().ref("email").email("Invalid email"),
	).strict()

	// Create a new user
	user := User{
		Name:  "John Doe",
		Email: "john@example.com",
	}

	// Validate the user
	err := schema.Validate(context.Background(), user)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(user)
}

## API Reference

### zod.Object

*   `zod.Object()` - Creates a new object schema.
*   `zod.Object().strict()` - Sets the schema to strict mode, which means that only defined fields are allowed.
*   `zod.Object().ref(field string)` - References a field in the schema.
*   `zod.Object().min(length int, msg string)` - Sets the minimum length of a field.
*   `zod.Object().max(length int, msg string)` - Sets the maximum length of a field.
*   `zod.Object().email(msg string)` - Sets an email validator for a field.