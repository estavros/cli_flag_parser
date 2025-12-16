# CLIParser

A lightweight command-line argument parser written in Go.

Supports:

- Long flags: `--verbose`
- Short flags: `-v`
- Short-flag bundling: `-abc` → `-a -b -c`
- Flags with values:
  - `--file=input.txt`
  - `-f input.txt`
- Boolean flags (implicit `true`)
- Flag aliases (`-v` → `--verbose`)
- Default values
- Simple string and boolean accessors

## Example

```bash
./app -vd -f data.txt -n 5

Equivalent to:

./app --verbose --debug --file=data.txt --n=5

```
## Features

### ✔ Long flags

```bash
--verbose
--file test.txt
```

### ✔ Short flags

```bash
-v
-f test.txt
```

### ✔ Key/value flags

```bash
--file=test.txt
-f=test.txt
--n=42
```

### ✔ Boolean flags

If a flag is provided without a value, the parser sets it to `"true"`:

```bash
--verbose
-v
```

## Usage

### Example command

```bash
go run main.go --verbose --file=test.txt -n 5
```

### Output

```
Verbose mode is ON
File: test.txt
Number: 5
```

## How It Works

Initialize the parser by passing your program arguments:

```go
parser := NewCLIParser(os.Args[1:])
```

Check for flags:

```go
if parser.HasFlag("verbose") || parser.HasFlag("v") {
    fmt.Println("Verbose mode is ON")
}
```

Retrieve values:

```go
file := parser.GetFlagValue("file")
```

## API

### `NewCLIParser(args []string) *CLIParser`

Creates a new parser and automatically processes arguments.

### `HasFlag(flag string) bool`

Returns `true` if the flag exists.

### `GetFlagValue(flag string) string`

Returns the flag value (or empty string if none).

## Example

```go
parser := NewCLIParser([]string{"--file=data.txt", "-v"})

if parser.HasFlag("v") {
    fmt.Println("Verbose mode is ON")
}

if parser.HasFlag("file") {
    fmt.Println("File:", parser.GetFlagValue("file"))
}
```
