package datastore

import (
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/notomo/qaper/domain/model"
)

// StateRepositoryImpl implements state repository
type StateRepositoryImpl struct {
}

var dirName = "qaper"

// Save state
func (repo *StateRepositoryImpl) Save(paper model.Paper) error {
	var content []byte
	dirPath := filepath.Join(os.TempDir(), dirName)

	if err := os.MkdirAll(dirPath, os.ModePerm); err != nil {
		return err
	}

	path := filepath.Join(dirPath, paper.ID())
	return ioutil.WriteFile(path, content, os.FileMode(0600))
}
