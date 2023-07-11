package issue_test

import (
	"testing"

	"github.com/waltervargas/issue"
)

func TestCreateIssue(t *testing.T) {
	t.Parallel()

	tmp := t.TempDir() + "/createIssue.gobdb"
	tracker, err := issue.OpenTracker(tmp)
	if err != nil {
		t.Fatalf("unable to open tracker: %s", err)
	}

	// how does create change the world?
	// world before change
	  // list issues, not issues yet
	  // fail if issues are present
	issues := tracker.ListIssues()
	if len(issues) > 0 {
		t.Fatalf("db is not empty")
	}
	
	// world after change
	  // created issue is present in the list
	issueName := "name of the issue"
	issue, err := tracker.CreateIssue(issueName)
	if err != nil {
		t.Fatalf("unable to create issue: %s", err)
	}
	if issue.Name != issueName {
		t.Fatalf("want: %q, got: %q", issueName, issue.Name)
	}

	issues = tracker.ListIssues()
	if len(issues) != 1 {
		t.Fatalf("want: 1 issues, got: %d issues after calling tracker.CreateIssue()", len(issues))
	}
}
