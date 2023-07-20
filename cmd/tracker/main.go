package main

import (
	"fmt"
	"os"

	"github.com/waltervargas/issue"
)

func main() {
	tracker, err := issue.OpenTracker(".issues.gobdb")
	if err != nil {
		fmt.Fprintf(os.Stderr, "unable to open tracker: %s", err)
		os.Exit(1)
	}

	if len(os.Args) > 1 {
		_, err = tracker.CreateIssue(os.Args[1])
		if err != nil {
			fmt.Fprintf(os.Stderr, "unable to create issue: %s", err)
			os.Exit(1)
		}
	}
	
	for i, issue := range tracker.ListIssues() {
		fmt.Printf("[%d] %s\n", i, issue.Name)
	}
}
