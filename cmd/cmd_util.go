package cmd

import (
	"fmt"
	"os"

	"github.com/MakeNowJust/heredoc/v2"
)

// CheckErr will print error message to stderr and exit with code 1
func CheckErr(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err.Error())
		os.Exit(1)
	}
}

// LongUsage is helper for print long usage
func LongUsage(s string) string {
	if len(s) == 0 {
		return s
	}
	return heredoc.Doc(s)
}
