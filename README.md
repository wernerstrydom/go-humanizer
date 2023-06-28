# Go-Humanizer

Go-Humanizer is a Go module designed to simplify the process of text 
manipulation in various common formats. It provides functionality to convert 
singular words to plural and vice versa, and to convert text to Pascal, 
camel, snake, or kebab case.

The module is heavily inspired by the [Humanizer](https://github.com/Humanizr/Humanizer) project for .NET. Some parts
have been ported and adapted to Go, while other parts have been written from scratch.

## Features

- **Singular to Plural and Plural to Singular:** This feature allows you to 
  easily convert words between singular and plural forms.
- **Case Conversion:** Go-Humanizer can convert text between various cases, 
  including Pascal, camel, snake, and kebab case.

## Usage

To use the `go-humanizer` module in your Go application, follow these steps:

1. Install the Go-Humanizer module by running the following command in your terminal:
   
   ```bash
   go get github.com/wernerstrydom/go-humanizer
   ```
   
2. Reference the module in your application by adding the following import statement:
   
   ```go
   import (
      "github.com/wernerstrydom/go-humanizer/en"
   )
   ```
   
3. Use the API to perform the desired text manipulation.

   ```go
    func main() {
      value := en.Plural("Person")
      println(value)
    }
   ```

## Building the Codebase

To build the Go-Humanizer module, follow these steps:

1. Clone the repository to your local machine.
2. Navigate to the directory containing the `go-humanizer` module.
3. Run `go build` to compile the module.

## Testing

To run the tests for the Go-Humanizer module, use the `go test` command in 
the directory containing the go-humanizer module.

## Contributing

Contributions to Go-Humanizer are welcome! To contribute:

1. Fork the repository.
2. Create a new branch for your feature or bug fix.
3. Make your changes in the new branch.
4. Submit a pull request to merge your branch into the main repository.

Before submitting your pull request, please ensure that your code adheres 
to the existing style conventions of the `go-humanizer` codebase and that 
all tests pass.

## Release Process

To release a new version of Go-Humanizer, follow these steps:

1. Create a new tag
   
   ```sh
    git tag -a v0.1.5 -m "Version 0.1.5"
   ```

2. Push the tag to GitHub

   ```sh
   git push origin v0.1.5
   ```

3. Run go list

    ```sh
    go list -m github.com/wernerstrydom/go-humanizer@v0.1.5
    ```
   
And then verify that the version is correct, and working.

## License

Go-Humanizer is licensed under the [MIT License](LICENSE).

## Acknowledgements

Go-Humanizer is inspired by the [Humanizer](https://github.com/Humanizr/Humanizer)