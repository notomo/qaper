package datastore

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"

	"github.com/notomo/qaper/domain/model"
)

// PaperRepositoryImpl implements paper repository
type PaperRepositoryImpl struct {
	Port string
}

// Add adds a paper
func (repo *PaperRepositoryImpl) Add() (model.Paper, error) {
	u := fmt.Sprintf("http://localhost:%s/paper", repo.Port)
	res, err := http.PostForm(u, url.Values{})
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	var paper *PaperImpl
	if err := json.Unmarshal(body, paper); err != nil {
		return nil, err
	}

	return paper, nil
}

// PaperImpl implements paper
type PaperImpl struct {
	PaperID string `json:"id"`
}

// ID returns an id
func (p *PaperImpl) ID() string {
	return p.PaperID
}
