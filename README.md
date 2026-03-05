# Go-Zod-Port

[![Build Status](https://travis-ci.org/travis-ci/travis-web.svg?branch=master)](https://travis-ci.org/travis-ci/travis-web)
[![Go Report Card](https://goreportcard.com/badge/github.com/golang-standards/project-layout)](https://goreportcard.com/report/github.com/golang-standards/project-layout)
[![GoDoc](https://godoc.org/github.com/golang-standards/project-layout?status.svg)](https://godoc.org/github.com/golang-standards/project-layout)

## Overview

Go-Zod-Port is a Go port of JavaScript's Zod library. It provides a simple way to define and validate data structures, including nested objects and arrays.

## Installation

To install Go-Zod-Port, run the following command:

```go
import (
    "github.com/go-zod-port/zod"
)
```

Then, run `go get` to fetch the dependencies:

```bash
go get github.com/go-zod-port/zod
```

## Usage

Here's an example of how to use Go-Zod-Port to define and validate a simple data structure:

    package main

    import (
        "fmt"
        "github.com/go-zod-port/zod"
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
        err := schema.Validate(user)
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
*   `zod.Object().min(length int, message string)` - Sets the minimum length of a string field.
*   `zod.Object().max(length int, message string)` - Sets the maximum length of a string field.
*   `zod.Object().email(message string)` - Sets a field to be an email address.
*   `zod.Object().Validate(data interface{}) error` - Validates the given data against the schema.