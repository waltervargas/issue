package issue_test

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/waltervargas/issue"
)

func TestCreateIssue(t *testing.T) {
	t.Parallel()

	tracker := newTestTracker(t)

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
	myissue, err := tracker.CreateIssue(issueName)
	if err != nil {
		t.Fatalf("unable to create issue: %s", err)
	}
	if myissue.ID == "" { 
		t.Errorf("issue.ID field should be different than empty string")
	}
	if myissue.Name != issueName {
		t.Fatalf("want: %q, got: %q", issueName, myissue.Name)
	}

	issues = tracker.ListIssues()
	if len(issues) != 1 {
		t.Fatalf("want: 1 issues, got: %d issues after calling tracker.CreateIssue()", len(issues))
	}

	tracker, err = issue.OpenTracker(tracker.Path)
	if err != nil {
		t.Fatalf("unable to open tracker: %s", err)
	}
	want := []issue.Issue{ {Name: issueName, ID: myissue.ID} }
	got := tracker.ListIssues()
	if !cmp.Equal(want, got) {
		t.Fatal(cmp.Diff(want, got))
	}

	myissue2, err := tracker.CreateIssue(issueName)
	if err != nil {
		t.Fatal(err)
	}
	if myissue.ID == myissue2.ID {
		t.Errorf("issues ids are not unique: %s == %s", myissue.ID, myissue2.ID)
	}
}

func TestGetIssue(t *testing.T) {
	t.Parallel()

	tracker := newTestTracker(t)

	issueName := "fix bug in Perl"
	issue, err := tracker.CreateIssue(issueName)
	if err != nil {
		t.Fatal(err)
	}
	got, ok := tracker.GetIssue(issue.ID)
	if !ok {
		t.Fatal(err)
	}
	if !cmp.Equal(issue, got) {
		t.Error(cmp.Diff(issue, got))
	}
}

func TestUpdateIssue(t *testing.T){
	t.Parallel()

	tracker := newTestTracker(t)

	name := "my issue"
	myIssue, err := tracker.CreateIssue(name)
	if err != nil {
		t.Fatal(err)
	}

	myIssue.Description = "blah blah"
	err = tracker.UpdateIssue(myIssue)
	if err != nil {
		t.Fatalf("unable to persist issue changes: %s", err)
	}

	got, ok := tracker.GetIssue(myIssue.ID)
	if ! ok {
		t.Fatalf("issue not found: %v", myIssue.ID)
	}

	if !cmp.Equal(myIssue, got){
		t.Error(cmp.Diff(myIssue, got))
	}
}

func newTestTracker(t *testing.T) *issue.Tracker {
	t.Helper() // if the test fails, reported as the calling test, not this functions. 

	tmp := t.TempDir() + "/createIssue.gobdb"
	tracker, err := issue.OpenTracker(tmp)
	if err != nil {
		t.Fatalf("unable to open tracker: %s", err)
	}

	return tracker
}

