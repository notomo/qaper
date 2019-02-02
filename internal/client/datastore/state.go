package datastore

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/notomo/qaper/domain/model"
	"github.com/notomo/qaper/internal/datastore"
)

// StateRepositoryImpl implements state repository
type StateRepositoryImpl struct {
}

var fileName = "qaper"

// Save state
func (repo *StateRepositoryImpl) Save(paper model.Paper) error {
	// TODO: get the file lock

	state := &datastore.StateImpl{StatePaperID: paper.ID()}
	content, err := json.Marshal(state)
	if err != nil {
		return err
	}

	return ioutil.WriteFile(filePath(), content, os.FileMode(0600))
}

// Load state
func (repo *StateRepositoryImpl) Load() (model.State, error) {
	content, err := ioutil.ReadFile(filePath())
	if err != nil {
		return nil, err
	}

	var state datastore.StateImpl
	if err := json.Unmarshal(content, &state); err != nil {
		return nil, err
	}

	return &state, nil
}

func filePath() string {
	return filepath.Join(os.TempDir(), fileName)
}
