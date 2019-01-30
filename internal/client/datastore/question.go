package datastore

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/notomo/qaper/domain/model"
	"github.com/notomo/qaper/internal/datastore"
)

// QuestionRepositoryImpl implements question repository
type QuestionRepositoryImpl struct {
	Port string
}

// GetCurrent gets the current question
func (repo *QuestionRepositoryImpl) GetCurrent(paperID string) (model.Question, error) {
	u := fmt.Sprintf("http://localhost:%s/paper/%s/question", repo.Port, paperID)
	res, err := http.Get(u)
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	if res.StatusCode == http.StatusNotFound {
		return nil, errors.New(string(body))
	}

	var question datastore.QuestionImpl
	if err := json.Unmarshal(body, &question); err != nil {
		return nil, err
	}

	return &question, nil
}
