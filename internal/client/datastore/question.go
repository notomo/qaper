package datastore

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/notomo/qaper/domain/model"
	"github.com/notomo/qaper/internal/client/datastore/httpc"
	"github.com/notomo/qaper/internal/datastore"
)

// QuestionRepositoryImpl implements question repository
type QuestionRepositoryImpl struct {
	Client *httpc.Client
}

// GetCurrent gets the current question
func (repo *QuestionRepositoryImpl) GetCurrent(paperID string) (model.Question, error) {
	path := fmt.Sprintf("/paper/%s/question", paperID)
	res, body, err := repo.Client.Get(path)
	if err != nil {
		return nil, err
	}

	if res.StatusCode != http.StatusOK {
		return nil, errors.New(string(body))
	}

	var question datastore.QuestionImpl
	if err := json.Unmarshal(body, &question); err != nil {
		return nil, err
	}

	return &question, nil
}
