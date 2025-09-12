package Models

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
)

type PrizeModel struct {
	Name   string `json:"Name"`
	Locale string `json:"locale"`
}

func GetPrizes(locale string) ([]PrizeModel, error) {
	var err error

	baseUrl := os.Getenv("BASE_URL")
	client := &http.Client{}

	query := "?" + strings.Join([]string{
		"locale=" + locale,
		"populate=*",
	}, "&")

	request, err := http.NewRequest("GET", fmt.Sprintf("%s/api/prizes%s", baseUrl, query), nil)
	if err != nil {
		return []PrizeModel{}, err
	}
	request.Header.Add("Accept", "application/json")
	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("Authorization", "Bearer "+os.Getenv("API_TOKEN"))

	clientRes, err := client.Do(request)
	if err != nil {
		log.Println(err.Error())
		return []PrizeModel{}, err
	}

	if clientRes.StatusCode != http.StatusOK {
		log.Println("Error getting prizes: " + clientRes.Status)
		return []PrizeModel{}, err
	}

	var decodedResponse StrapiResponse[[]PrizeModel]
	err = json.NewDecoder(clientRes.Body).Decode(&decodedResponse)
	if err != nil {
		log.Println(err.Error())
		return []PrizeModel{}, err
	}

	return decodedResponse.Data, nil
}

func GetPrize(id int, locale string) (PrizeModel, error) {
	var err error

	baseUrl := os.Getenv("BASE_URL")
	client := &http.Client{}

	query := "?" + strings.Join([]string{
		"locale=" + locale,
		"populate=*",
	}, "&")

	request, err := http.NewRequest("GET", fmt.Sprintf("%s/api/prizes/%d%s", baseUrl, id, query), nil)
	if err != nil {
		return PrizeModel{}, err
	}
	request.Header.Add("Accept", "application/json")
	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("Authorization", "Bearer "+os.Getenv("API_TOKEN"))

	clientRes, err := client.Do(request)
	if err != nil {
		log.Println(err.Error())
		return PrizeModel{}, err
	}

	if clientRes.StatusCode != http.StatusOK {
		log.Println("Error getting prizes: " + clientRes.Status)
		return PrizeModel{}, err
	}

	var decodedResponse StrapiResponse[PrizeModel]
	err = json.NewDecoder(clientRes.Body).Decode(&decodedResponse)
	if err != nil {
		log.Println(err.Error())
		return PrizeModel{}, err
	}

	return decodedResponse.Data, nil
}
