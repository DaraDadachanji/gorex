package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"regexp"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	pattern := parsePatternArgument()
	rex := regexp.MustCompile(pattern)
	for {
		line, readErr := reader.ReadString('\n')
		captures := rex.FindStringSubmatch(line)
		if len(captures) == 0 {
			// skip line if not matched
			if readErr == io.EOF {
				return
			}
			continue
		}
		output := strings.Join(captures[1:], " ")
		fmt.Print(output)
		if readErr == io.EOF {
			return
		} else {
			fmt.Println()
		}
	}
}

// parsePatternArgument parses the regex pattern from the command line arguments
func parsePatternArgument() string {
	if len(os.Args) > 2 {
		_, _ = fmt.Fprintf(os.Stderr, "too many arguments")
		os.Exit(1)
	}
	if len(os.Args) < 2 {
		_, _ = fmt.Fprintf(os.Stderr, "missing pattern")
		os.Exit(1)
	}
	return os.Args[1]
}
