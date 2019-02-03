package datastore

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/notomo/qaper/domain/model"
	cmodel "github.com/notomo/qaper/internal/client/model"
)

// StateRepositoryImpl implements state repository
type StateRepositoryImpl struct {
}

var fileName = "qaper"

// Save state
func (repo *StateRepositoryImpl) Save(paper model.Paper) error {
	// TODO: get the file lock

	state := &StateImpl{StatePaperID: paper.ID()}
	content, err := json.Marshal(state)
	if err != nil {
		return err
	}

	return ioutil.WriteFile(filePath(), content, os.FileMode(0600))
}

// Load state
func (repo *StateRepositoryImpl) Load() (cmodel.State, error) {
	content, err := ioutil.ReadFile(filePath())
	if err != nil {
		return nil, err
	}

	var state StateImpl
	if err := json.Unmarshal(content, &state); err != nil {
		return nil, err
	}

	return &state, nil
}

func filePath() string {
	return filepath.Join(os.TempDir(), fileName)
}

// StateImpl implements state
type StateImpl struct {
	StatePaperID string `json:"paperId"`
}

// PaperID returns a paper id
func (p *StateImpl) PaperID() string {
	return p.StatePaperID
}
