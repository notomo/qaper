package datastore

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"

	"github.com/notomo/qaper/domain/model"
	"github.com/notomo/qaper/internal"
	"github.com/notomo/qaper/internal/datastore"
)

// PaperRepositoryImpl implements paper repository
type PaperRepositoryImpl struct {
	Port string
}

// Add adds a paper
func (repo *PaperRepositoryImpl) Add(bookID string) (model.Paper, error) {
	u := fmt.Sprintf("http://localhost:%s/book/%s/paper", repo.Port, bookID)
	res, err := http.PostForm(u, url.Values{})
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	if res.StatusCode != http.StatusOK {
		return nil, errors.New(string(body))
	}

	var paper datastore.PaperImpl
	if err := json.Unmarshal(body, &paper); err != nil {
		return nil, err
	}

	return &paper, nil
}

// Get gets a paper
func (repo *PaperRepositoryImpl) Get(paperID string) (model.Paper, error) {
	return nil, internal.ErrNotImplemented
}
