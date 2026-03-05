# Contributing to njchilds90/go-zod-port

This guide outlines the process for contributing to the njchilds90/go-zod-port Go module. It covers coding standards, testing requirements, and the pull request process.

## Coding Standards

* All Go files must be formatted using `gofmt`. This can be done automatically with `go fmt ./...`.
* Code must pass `go vet` checks. Run `go vet ./...` to check your code.
* The `context.Context` type should be used as the first parameter in functions where appropriate.
* Error handling should use `fmt.Errorf` with `%w` for wrapping errors. Exported sentinel errors should be defined for error types that can be recovered from or need special handling.
* All exported symbols must have godoc comments starting with the symbol name.

## Testing Requirements

* All code must be covered by unit tests. Use table-driven tests with `t.Run()` to test different cases.
* Test files must have a `_test.go` suffix and be in the same package as the code they're testing.
* Example usage should be provided in a separate file with an `_example_test.go` suffix. This file should have an `Example<FunctionName>()` function with a `// Output:` comment.
* Tests should cover success cases, error cases, and edge cases (e.g., nil, empty, boundary values).

## Pull Request Process

1. Fork the njchilds90/go-zod-port repository.
2. Make your changes, following the coding standards and testing requirements above.
3. Run `make test` and `make lint` to ensure your changes pass all checks.
4. Commit your changes with a meaningful commit message.
5. Open a pull request against the njchilds90/go-zod-port repository.
6. Wait for review and address any feedback or issues that arise during the review process.

## Additional Tools and Checks

* This repository uses GitHub Actions to automate testing and linting. PRs will automatically be checked for formatting, linting, and test coverage.
* The `.golangci.yml` file configures golangci-lint to enforce coding standards.
* The `Makefile` provides standard Go targets (build, test, lint, fmt, vet, bench, coverage).
