## Test Commands 🛠️

Here are the essential commands for running and analyzing Go test cases:

### Basic Testing

```bash
go test -v -cover ./...
```

Runs all tests recursively with verbose output and displays coverage in terminal.

### Race Condition Testing

```bash
go test -race .
```

Executes test cases with race condition detection in the current directory.

```bash
go test -race -v .
```

Same as above but with verbose output for detailed test execution information.

### Coverage Analysis

```bash
go test -coverprofile=coverage.out
```

Generates a coverage profile file (`coverage.out`) for detailed analysis.

```bash
go tool cover -html=coverage.out
```

Creates an interactive HTML report from the coverage profile, allowing visual exploration of test coverage in your browser.

## Project Contents 🧭

This repository contains the following implementations:

- [Producer Consumer Pattern](./01-producer-consumer/README.md)
  - A classic concurrent programming pattern implementation in Go
- [Dining Philosophers Problem](./02-dining-philosophers/README.md)
  - An implementation of Dijkstra's famous dining philosophers problem demonstrating deadlock prevention
- [Sleeping Barber Problem](./03-sleeping-barber/README.md)
  - A synchronization problem that demonstrates resource management and scheduling in concurrent programming
