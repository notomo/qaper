package datastore

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/notomo/qaper/domain/model"
)

// StateRepositoryImpl implements state repository
type StateRepositoryImpl struct {
}

var fileName = "qaper"

// Save state
func (repo *StateRepositoryImpl) Save(paper model.Paper) error {
	state := map[string]string{"id": paper.ID()}
	content, err := json.Marshal(state)
	if err != nil {
		return err
	}

	path := getFilePath()
	return ioutil.WriteFile(path, content, os.FileMode(0600))
}

// Load state
func (repo *StateRepositoryImpl) Load() (string, error) {
	path := getFilePath()
	file, err := ioutil.ReadFile(path)
	if err != nil {
		return "", err
	}

	var content map[string]string
	if err := json.Unmarshal(file, &content); err != nil {
		return "", err
	}

	return content["id"], nil
}

func getFilePath() string {
	return filepath.Join(os.TempDir(), fileName)
}
