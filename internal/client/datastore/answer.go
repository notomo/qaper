package datastore

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/notomo/qaper/domain/model"
	"github.com/notomo/qaper/internal/client/datastore/httpc"
)

// AnswerRepositoryImpl implements answer repository
type AnswerRepositoryImpl struct {
	Client *httpc.Client
}

// Set gets the current answer
func (repo *AnswerRepositoryImpl) Set(paperID string, answer model.Answer) error {
	path := fmt.Sprintf("/paper/%s/answer", paperID)
	res, body, err := repo.Client.Post(path, answer)
	if err != nil {
		return err
	}

	if res.StatusCode != http.StatusOK {
		return errors.New(string(body))
	}

	return nil
}
