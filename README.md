# README

This is a hello world for learning the first steps of Golang.

## Importing Packages
Go comes with built in packages for common asks. Check the [Go Standard Library](https://pkg.go.dev/std)

### Third Party Packages
- refer to third party packages with their fully-qualified names, ie: `github.com/google/uuid`
- before downloading a package, you need to initialize a module
- then you can add modules with `go get <module-name>`
  ie: `go get github.com/google/uuid`
- After adding the import, run `go mod tidy` to update the `go.mod` file to reflect the dependencies in the codebase

## Initializing Go Modules
- [Tutorial: Create a Go Module](https://go.dev/doc/tutorial/create-module)
- A fundamental feature of Go

### Naming a Module
[Managing Dependencies Naming](https://go.dev/doc/modules/managing-dependencies#naming_module)

- Use the pattern: `<prefix>/<descriptive-tex>`
  - **prefix**: a string that partially descripbes the module, such as a origin location:
    - location of the repository where go tools can find the module's source code
      - ie: `github.com/<project-name>`
  - [Guidelines for Developing and Publishing Modules](https://go.dev/doc/modules/developing)
  - avoid common names like: widgets, utilities, or app
  - **descriptive-text** should use the project name, package names carry most of the weight of describing functionality.

### Local Modules
- Move module file into it's own directory and set up `go.mod` file
  - in this repo it's the greeting/ directory
- In the caller directory (hello/), initialize as normal `go mod init example.com/hello`  
- run this command from the hello directory:
  - `go mod edit -replace example.com/greetings=../greetings`
    - this will get it to read from the greetings/ dir and not from example.com/greetings
  - then run `go mod tidy` to sync the changes
- then you can run the hello module and successfuly use your greetings module with `go run .`