package service

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"stockbit/2/constant"
	"stockbit/2/repository"
)

type Data struct {
	Search       []Detail `json:"Search"`
	TotalResults string   `json:"totalResults"`
	Response     string   `json:"Response"`
}

type Detail struct {
	Title  string `json:"Title"`
	Year   string `json:"Year"`
	ImdbID string `json:"imdbID"`
	Type   string `json:"Type"`
	Poster string `json:"Poster"`
}

func Get(pagination string, searchword string) (Data, error) {

	url := constant.OMBDAPI + "?apikey=" + constant.OMDBKey + "&s=" + searchword + "&page=" + pagination
	response, err := http.Get(url)
	if err != nil {
		return Data{}, err
	}
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return Data{}, err
	}

	var data Data
	err = json.Unmarshal(body, &data)
	if err != nil {
		return Data{}, err
	}

	err = repository.InsertLog(pagination, searchword)
	if err != nil {
		return Data{}, err
	}

	return data, nil
}
