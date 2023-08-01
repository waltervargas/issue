package issue

import (
	"fmt"

	"github.com/waltervargas/gobdb"
)

type Tracker struct {
	db gobdb.Gobdb[Issue]
	nextID int
}

type Issue struct {
	Name string
	ID string
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
	// calculate the ID
	// ID depends on the storage engine
	// scope for this ID is the tracker
	// tracker + issueID
	// how can we generate IDs that are unique for this tracker? 
	// based on ids that we know?
	
	issue := Issue{Name: name, ID: fmt.Sprintf("%d", t.nextID)}
	t.nextID++
	err := t.db.Add(issue)
	if err != nil {
		return Issue{}, err
	}
	
	return issue, nil
}
