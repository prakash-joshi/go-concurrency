## Commands to run go test cases 🛠️

go test -v -cover ./... :: This command shows the coverage in terminal, to show coverage highlighting in vs code I did the following.

go test -race . :: This command executes the test cases with race condition.

go test -race -v . :: This command will execute your tests in the current directory with the race detector enabled and verbose output.

go test -coverprofile=coverage.out :: This command runs tests and generates a coverage profile file named coverage.out to analyze code coverage.

go tool cover -html=coverage.out :: This command converts the coverage profile file named coverage.out into an HTML report, which you can open in your browser to visually explore your test coverage.
