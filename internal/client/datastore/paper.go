package datastore

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/notomo/qaper/domain/model"
	"github.com/notomo/qaper/internal"
	"github.com/notomo/qaper/internal/client/datastore/httpc"
	"github.com/notomo/qaper/internal/datastore"
)

// PaperRepositoryImpl implements paper repository
type PaperRepositoryImpl struct {
	Client *httpc.Client
}

// Add adds a paper
func (repo *PaperRepositoryImpl) Add(bookID string) (model.Paper, error) {
	path := fmt.Sprintf("/book/%s/paper", bookID)
	res, body, err := repo.Client.Post(path, "")
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
