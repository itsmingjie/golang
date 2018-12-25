# Hello World Demo
12/24/2018

## go-cli commands

* Difference between `go run` and `go build`
    - `go run` immediately executes the file(s) passed in
    - `go build` creates and executable, does not run

* `go fmt` automatically formats all code
* Commands dealing with dependencies:
    - `go install` compiles and installs a package
    - `go get` downloads the raw source code of someone else's pkg
* `go test` runs test files

## Package
* Definition: A collection of common source code files
* Types of packages
    - Executable ("doing something")
        * Generates a file that we can run
    - Reusable
        * Helpers, somewhere we put reusable logic in
* Only `package main` would create executable files, all others will create reusable files.
    - Building a reusable file won't create anything
* Executable files must have a main function - `func main()`

## File Structure
* Package declaration
* Import statements
    - Give current package access to code in some other package, or libraries
* Declaring functions, doing stuff, etc.