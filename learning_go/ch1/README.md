### Chapter 1 Summary

#### Environment Variables

- GOPATH should be set to $HOME/go
- PATH should be set to $GOPATH/bin

#### Commands

`go run` is used to run the Go script on the fly

```
go run hello.go
```

If you want to treat Go as a script then use `go run`. This command will compile the source code in a temp dir, run the executable and then remove it.

`go build` is used to compile the source code. The executable name use the filename, unless specified with `-o` argument.

```
go build hello.go 
# produces hello
```

```
go build -o hello_world hello.go 
# produces hello_world
```

`go install` is used to download certain Go project from a Git repository, build it and install the binary into the $GOPATH/bin directory.

```
go install github.com/rakyll/hey@latest
```

#### Formatting and Linting

`go fmt` is used to autoformat the source files.

`go imports -l -w .` is an enhanced version of `go fmt` which also fix import statements.

`go lint` is the Go linter

`go vet` is the tool to detect programming mistakes

`golangci-lint run` is used to run all above commands (and more)

**Note**: Make sure the team agree on the rules to apply!

#### Makefile

Go developers have adopted `make` as their build tool.