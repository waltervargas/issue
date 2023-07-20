package issue

import (
	"fmt"

	"github.com/waltervargas/gobdb"
)

type Tracker struct {
	db gobdb.Gobdb[Issue]
}

type Issue struct {
	Name string
}

func OpenTracker(path string) (*Tracker, error){
	db, err := gobdb.Open[Issue](path)
	if err != nil {
		return nil, fmt.Errorf("unable to open DB: %w", err)
	}

	return &Tracker{db: db}, nil
}

func (t *Tracker) ListIssues() ([]Issue){
	return t.db.List()
}

func (t *Tracker) CreateIssue(name string) (Issue, error){
	issue := Issue{Name: name}
	err := t.db.Add(issue)
	if err != nil {
		return Issue{}, err
	}
	return issue, nil
}
