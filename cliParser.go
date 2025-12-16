package main

import (
	"fmt"
	"os"
	"strings"
)

type CLIParser struct {
	flags    map[string]string
	defaults map[string]string
	aliases  map[string]string
}

func NewCLIParser(args []string) *CLIParser {
	parser := &CLIParser{
		flags:    make(map[string]string),
		defaults: make(map[string]string),
		aliases:  make(map[string]string),
	}
	parser.parse(args)
	return parser
}

// ----------------------
// Alias support
// ----------------------

func (c *CLIParser) SetAlias(alias, flag string) {
	c.aliases[alias] = flag
}

func (c *CLIParser) normalize(flag string) string {
	if canonical, exists := c.aliases[flag]; exists {
		return canonical
	}
	return flag
}

// ----------------------
// Argument parsing
// ----------------------

func (c *CLIParser) parse(args []string) {
	for i := 0; i < len(args); i++ {
		arg := args[i]

		// Support --flag=value or -f=value
		if strings.Contains(arg, "=") {
			parts := strings.SplitN(arg, "=", 2)
			flag := parts[0]
			value := parts[1]

			if strings.HasPrefix(flag, "--") {
				key := c.normalize(flag[2:])
				c.flags[key] = value
			} else if strings.HasPrefix(flag, "-") {
				key := c.normalize(flag[1:])
				c.flags[key] = value
			}
			continue
		}

		// Long flag: --verbose
		if strings.HasPrefix(arg, "--") && len(arg) > 2 {
			key := c.normalize(arg[2:])
			if i+1 < len(args) && args[i+1][0] != '-' {
				c.flags[key] = args[i+1]
				i++
			} else {
				c.flags[key] = "true"
			}
			continue
		}

		// Short flags (supports bundling): -abc or -n 10
		if strings.HasPrefix(arg, "-") && len(arg) > 1 {
			shorts := arg[1:]

			// Single short flag may take a value
			if len(shorts) == 1 {
				key := c.normalize(shorts)
				if i+1 < len(args) && args[i+1][0] != '-' {
					c.flags[key] = args[i+1]
					i++
				} else {
					c.flags[key] = "true"
				}
				continue
			}

			// Bundled short flags: -abc → -a -b -c (all boolean)
			for _, ch := range shorts {
				key := c.normalize(string(ch))
				c.flags[key] = "true"
			}
		}
	}
}

// ----------------------
// Default values support
// ----------------------

func (c *CLIParser) SetDefault(flag, value string) {
	flag = c.normalize(flag)
	c.defaults[flag] = value

	if _, exists := c.flags[flag]; !exists {
		c.flags[flag] = value
	}
}

// ----------------------
// Flag accessors
// ----------------------

func (c *CLIParser) HasFlag(flag string) bool {
	flag = c.normalize(flag)
	_, exists := c.flags[flag]
	return exists
}

func (c *CLIParser) GetFlagValue(flag string) string {
	flag = c.normalize(flag)
	return c.flags[flag]
}

func (c *CLIParser) GetBoolFlag(flag string) bool {
	flag = c.normalize(flag)
	val, exists := c.flags[flag]
	if !exists {
		return false
	}
	val = strings.ToLower(val)
	return val == "true" || val == "1"
}

// ----------------------
// Example usage
// ----------------------

func main() {
	parser := NewCLIParser(os.Args[1:])

	// Aliases
	parser.SetAlias("v", "verbose")
	parser.SetAlias("f", "file")
	parser.SetAlias("d", "debug")

	// Defaults
	parser.SetDefault("verbose", "false")
	parser.SetDefault("debug", "false")
	parser.SetDefault("file", "input.txt")
	parser.SetDefault("n", "10")

	if parser.GetBoolFlag("v") {
		fmt.Println("Verbose mode is ON")
	}

	if parser.HasFlag("file") {
		fmt.Println("File:", parser.GetFlagValue("f"))
	}

	if parser.HasFlag("n") {
		fmt.Println("Number:", parser.GetFlagValue("n"))
	}

	if parser.GetBoolFlag("debug") {
		fmt.Println("Debug mode is ON")
	} else {
		fmt.Println("Debug mode is OFF")
	}
}
