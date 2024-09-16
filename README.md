# Go Interpreter Written in Go

Welcome to the **gogo** project! This is an interpreter for the Go programming language, written in Go itself. The project aims to execute Go code by parsing and evaluating the Abstract Syntax Tree (AST) generated from the source code.

---

## Table of Contents

- [Introduction](#introduction)
- [How It Works](#how-it-works)
- [Project Structure](#project-structure)
- [Features](#features)
  - [Implemented Features](#implemented-features)
  - [Planned Features](#planned-features)
- [Feature Summary](#feature-summary)
- [Getting Started](#getting-started)
  - [Prerequisites](#prerequisites)
  - [Building the Interpreter](#building-the-interpreter)
  - [Running Go Programs](#running-go-programs)
- [Testing](#testing)
- [Contributing](#contributing)
- [License](#license)

---

## Introduction

This project is an interpreter for a subset of the Go language, written entirely in Go. The primary goal is to interpret Go source code by leveraging existing Go compiler packages, such as `go/parser`, `go/ast`, and `go/token`. By reusing these packages, we avoid reimplementing parsing and focus on interpreting the AST to execute Go code.

---

## How It Works

The interpreter works by:

1. **Parsing**: Using `go/parser` to parse Go source code into an AST.
2. **Evaluation**: Traversing the AST and evaluating nodes to execute the program logic.
3. **Environment Management**: Keeping track of variable scopes and values during execution.
4. **Reusing Go Compiler Code**: Leveraging Go's standard library packages for parsing and AST manipulation.

---

## Project Structure

The project follows idiomatic Go project structure and organizes code into packages for clarity and maintainability.

```go
GoGo/
├── cmd/
│   └── main.go
├── internal/
│   ├── evaluator/
│   │   ├── evaluator.go
│   │   ├── statements.go
│   │   ├── expressions.go
│   │   ├── operators.go
│   │   ├── declarations.go
│   │   └── helpers.go
│   ├── interpreter/
│   │   └── interpreter.go
│   └── runtime/
│       ├── environment.go
│       └── object.go
├── testdata/
│   ├── test1.go
│   ├── test2.go
│   └── test3.go
├── go.mod
└── README.md
```

- **`cmd/main.go`**: Entry point of the interpreter.
- **`internal/evaluator/`**: Contains the evaluation logic, split into multiple files for readability.
- **`internal/interpreter/interpreter.go`**: Coordinates parsing and evaluation.
- **`internal/runtime/`**: Manages variable scopes (`environment.go`) and defines runtime objects (`object.go`).
- **`testdata/`**: Sample Go programs for testing the interpreter.
- **`go.mod`**: Go module file.
- **`README.md`**: Project documentation.

---

## Features

### Implemented Features

- **Parsing Go Code**: Using `go/parser` to parse Go source files.
- **Evaluating Basic Expressions**: Handling integer arithmetic and basic literals.
- **Variable Declarations and Assignments**: Supporting `var` declarations and shorthand assignments.
- **Basic Control Structures**:
  - **If Statements**: Evaluating `if` and `else` blocks.
  - **For Loops**: Supporting `for` loops with initialization, condition, and post statements.
- **Function Calls**: Interpreting calls to built-in functions like `println`.
- **Environment Management**: Tracking variable scopes and values during execution.
- **Error Handling**: Reporting unsupported features and runtime errors.

### Planned Features

- **Advanced Expressions**: Support for floating-point numbers, strings, and boolean expressions.
- **Composite Types**: Handling arrays, slices, maps, and structs.
- **User-Defined Functions**: Defining and calling user functions.
- **Methods and Interfaces**: Implementing method calls and interface types.
- **Concurrency Primitives**: Supporting goroutines and channels.
- **Standard Library Integration**: Access to Go's standard library packages.
- **Type Checking**: Implementing type checking using `go/types`.
- **Improved Error Messages**: Providing detailed and user-friendly error messages.

### Feature Summary

| Feature                         | Status        |
| ------------------------------- | ------------- |
| Parsing Go Code                 | Implemented   |
| Evaluating Basic Expressions    | Implemented   |
| Variable Declarations           | Implemented   |
| If Statements                   | Implemented   |
| For Loops                       | Implemented   |
| Function Calls (`println`)      | Implemented   |
| Environment Management          | Implemented   |
| Error Handling                  | Implemented   |
| Floating-Point Numbers          | Planned       |
| Strings and Booleans            | Planned       |
| Composite Types                 | Planned       |
| User-Defined Functions          | Planned       |
| Methods and Interfaces          | Planned       |
| Concurrency (Goroutines)        | Planned       |
| Standard Library Integration    | Planned       |
| Type Checking                   | Planned       |
| Improved Error Messages         | Planned       |

---

## Getting Started

### Prerequisites

- **Go 1.20 or higher**: Make sure you have Go installed. You can download it from the [official website](https://golang.org/dl/).

### Building the Interpreter

Clone the repository and navigate to the project directory:

```bash
git clone https://github.com/AlexanderStocks/gogo.git
cd gogo
```

Build the interpreter executable:

```bash
go build -o gogo cmd/main.go
```

### Running Go Programs

You can run Go programs using the `gogo` interpreter. For example, to run the sample program `testdata/test1.go`, use:

```bash
./gogo testdata/test1.go
```

This will interpret the Go code and execute the program logic.

---

## Testing

The project includes sample Go programs in the `testdata/` directory for testing the interpreter. You can run these tests using the `gogo` interpreter:

```bash
go test ./...
```

---

## Contributing

Contributions are welcome! Feel free to open issues or submit pull requests to help improve the interpreter.

---

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.

---

