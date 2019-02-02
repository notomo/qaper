package datastore

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/notomo/qaper/domain/model"
)

// AnswerRepositoryImpl implements answer repository
type AnswerRepositoryImpl struct {
	Port string
}

// Set gets the current answer
func (repo *AnswerRepositoryImpl) Set(paperID string, answer model.Answer) error {
	content, err := json.Marshal(answer)
	if err != nil {
		return err
	}

	u := fmt.Sprintf("http://localhost:%s/paper/%s/answer", repo.Port, paperID)
	req, err := http.NewRequest("POST", u, bytes.NewBuffer(content))
	if err != nil {
		return err
	}

	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return err
	}

	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return err
	}

	if res.StatusCode != http.StatusOK {
		return errors.New(string(body))
	}

	return nil
}
