package issue

import (
	"fmt"
	"strconv"

	"github.com/waltervargas/gobdb"
	"golang.org/x/exp/slices"
)

type Tracker struct {
	db gobdb.Gobdb[Issue]
	Path string
	nextID int
}

type Issue struct {
	ID string
	Name string
	Description string
}

func OpenTracker(path string) (*Tracker, error){
	db, err := gobdb.Open[Issue](path)
	if err != nil {
		return nil, fmt.Errorf("unable to open DB: %w", err)
	}
	// calculate the nextID
	issues := db.List()
	var nextID int
	if len(issues) > 0 {
		issue := slices.MaxFunc(issues, func(a, b Issue) int { 
			var max int 
			if a.ID > b.ID {
				max, _ = strconv.Atoi(a.ID)
				return max
			}
			max, _ = strconv.Atoi(b.ID)
			return max
		})
		maxID, _ := strconv.Atoi(issue.ID)
		nextID = maxID + 1
	}

	return &Tracker{db: db, Path: path, nextID: nextID}, nil
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

func (t *Tracker) GetIssue(ID string) (Issue, bool) {
	issues := t.db.List()
	index := slices.IndexFunc(issues, func(i Issue) bool {
		return i.ID == ID
	})
	if index < 0 {
		return Issue{}, false
	}

	return issues[index], true
}

func (t *Tracker) UpdateIssue(issue Issue) (error) {
	// get old issue
	oldIssue, ok := t.GetIssue(issue.ID)
	if !ok {
		return fmt.Errorf("unable to get the issue: %s", issue.ID)
	}
	err := t.db.Delete(oldIssue)
	if err != nil {
		return fmt.Errorf("unable to delete issue: %w", err)
	}
	t.db.Add(issue)
	if err != nil {
		return fmt.Errorf("unable to add issue: %w", err)
	}
	return nil
}
