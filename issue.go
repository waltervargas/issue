package issue

type Tracker struct {
	issues []Issue
}

type Issue struct {
	Name string
}

func OpenTracker(path string) (*Tracker, error){
	return &Tracker{}, nil
}

func (t *Tracker) ListIssues() ([]Issue){
	return t.issues
}

func (t *Tracker) CreateIssue(name string) (Issue, error){
	issue := Issue{Name: name}
	t.issues = append(t.issues, issue)
	return issue, nil
}
