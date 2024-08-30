# Advent of Code Template

This repository serves as a template for solving challenges from the 
[Advent of Code](https://adventofcode.com/) using Go. It leverages 
the [goaoc](https://github.com/hvpaiva/goaoc) library as the execution framework and is best 
used in conjunction with the command-line interface [goaoc-cli](https://github.com/hvpaiva/goaoc-cli).

## Project Structure

The core structure of this project is organized as follows:

- `/internal/<year>/<day>/main.go`: Contains solutions for a specific day's challenge.
- `/internal/<year>/<day>/main_test.go`: Holds unit tests for the day's solutions.
- `/internal/<year>/<day>/input.txt`: The input data for the day's challenge, provided by Advent of Code.
- `/pkg/*`: Useful utility functions for your solutions.

## Installation and Setup

To get started with this template, clone the repository and ensure you have Go installed on your machine.

```bash
git clone https://github.com/yourusername/aoc-template.git
cd aoc-template
go mod tidy
```

Install the goaoc library and CLI:

```bash
go install github.com/hvpaiva/goaoc-cli
```

## Usage

### Setting Up a New Challenge

Utilize the `goaoc-cli` to auto-generate the necessary structure for each day's challenge:

```bash
goaoc-cli init --year=<year> --day=<day>
```

This command will create the basic file structure and can fetch the input directly from the Advent of Code website, 
given that your session cookie is configured in `~/.config/goaoc/config.toml`.

### Implementing Solutions

Implement your solutions in the `main.go` file under the corresponding year/day directory. The example solutions include
two parts:

```go
func partOne(input string) int {
	// Implement solution for Part One
}

func partTwo(input string) int {
	// Implement solution for Part Two
}
```

### Running Solutions and Tests

To execute your solution, use the `goaoc-cli`:

```bash
goaoc-cli run --year=<year> --day=<day>
```

For running tests, leverage Go's testing tool:

```bash
goaoc-cli test --year=<year> --day=<day>
```

## Example

The day one of the year 2015 is already implemented as an example.

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## Contributing

Contributions are welcome! Please open an issue or submit a pull request.
