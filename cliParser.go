package main

import (
	"fmt"
	"os"
	"strings"
)

type CLIParser struct {
	flags map[string]string
}

func NewCLIParser(args []string) *CLIParser {
	parser := &CLIParser{flags: make(map[string]string)}
	parser.parse(args)
	return parser
}

func (c *CLIParser) parse(args []string) {
	for i := 0; i < len(args); i++ {
		arg := args[i]

		// --- NEW FEATURE: support --flag=value or -f=value ---
		if strings.Contains(arg, "=") {
			parts := strings.SplitN(arg, "=", 2)
			flag := parts[0]
			value := parts[1]

			if strings.HasPrefix(flag, "--") {
				c.flags[flag[2:]] = value
			} else if strings.HasPrefix(flag, "-") {
				c.flags[flag[1:]] = value
			}
			continue
		}

		// Long flag: --verbose
		if len(arg) > 2 && arg[:2] == "--" {
			key := arg[2:]
			if i+1 < len(args) && args[i+1][0] != '-' {
				c.flags[key] = args[i+1]
				i++
			} else {
				c.flags[key] = "true"
			}
			continue
		}

		// Short flag: -v
		if len(arg) > 1 && arg[0] == '-' {
			key := arg[1:]
			if i+1 < len(args) && args[i+1][0] != '-' {
				c.flags[key] = args[i+1]
				i++
			} else {
				c.flags[key] = "true"
			}
		}
	}
}

func (c *CLIParser) HasFlag(flag string) bool {
	_, exists := c.flags[flag]
	return exists
}

func (c *CLIParser) GetFlagValue(flag string) string {
	return c.flags[flag]
}

func main() {
	parser := NewCLIParser(os.Args[1:]) // skip program name

	if parser.HasFlag("verbose") || parser.HasFlag("v") {
		fmt.Println("Verbose mode is ON")
	}

	if parser.HasFlag("file") || parser.HasFlag("f") {
		fmt.Println("File:", parser.GetFlagValue("file"))
	}

	if parser.HasFlag("n") {
		fmt.Println("Number:", parser.GetFlagValue("n"))
	}
}
