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
		"F",
		" ",
		"Field separator for outputs if multiple capture groups are specified",
	)
	flag.Parse()

	reader := bufio.NewReader(os.Stdin)
	pattern := parsePatternArgument(flag.Args())
	rex := regexp.MustCompile(pattern)
	for {
		line, readErr := reader.ReadString('\n')
		if readErr != io.EOF {
			//trim newline from end of line
			line = line[0 : len(line)-1]
		}
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
		_, _ = fmt.Fprintf(os.Stderr, "too many arguments (flags must come befoer pattern)")
		os.Exit(1)
	}
	if len(args) < 1 {
		_, _ = fmt.Fprintf(os.Stderr, "missing pattern")
		os.Exit(1)
	}
	return args[0]
}
