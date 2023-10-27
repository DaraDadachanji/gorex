package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
	"regexp"
	"strings"
)

func main() {
	delimiter := flag.String(
		"d",
		" ",
		"delimiter for outputs if multiple capture groups are specified",
	)
	flag.Parse()
	fmt.Println(*delimiter)

	reader := bufio.NewReader(os.Stdin)
	pattern := parsePatternArgument(flag.Args())
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
		output := strings.Join(captures[1:], *delimiter)
		fmt.Print(output)
		if readErr == io.EOF {
			return
		} else {
			fmt.Println()
		}
	}
}

// parsePatternArgument parses the regex pattern from the command line arguments
func parsePatternArgument(args []string) string {
	if len(args) > 1 {
		_, _ = fmt.Fprintf(os.Stderr, "too many arguments")
		for _, arg := range args {
			_, _ = fmt.Fprintf(os.Stderr, "\n")
			_, _ = fmt.Fprintf(os.Stderr, arg)
		}
		os.Exit(1)
	}
	if len(args) < 1 {
		_, _ = fmt.Fprintf(os.Stderr, "missing pattern")
		os.Exit(1)
	}
	return args[0]
}
