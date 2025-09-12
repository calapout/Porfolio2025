package Models

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
)

type CompanyModel struct {
	Name    string `json:"Name"`
	Locale  string `json:"locale"`
	Website string `json:"Website"`
}

func GetCompanies(locale string) ([]CompanyModel, error) {
	var err error

	baseUrl := os.Getenv("BASE_URL")
	client := &http.Client{}

	query := "?" + strings.Join([]string{
		"locale=" + locale,
		"populate=*",
	}, "&")

	request, err := http.NewRequest("GET", fmt.Sprintf("%s/api/companies%s", baseUrl, query), nil)
	if err != nil {
		return []CompanyModel{}, err
	}
	request.Header.Add("Accept", "application/json")
	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("Authorization", "Bearer "+os.Getenv("API_TOKEN"))

	clientRes, err := client.Do(request)
	if err != nil {
		log.Println(err.Error())
		return []CompanyModel{}, err
	}

	if clientRes.StatusCode != http.StatusOK {
		log.Println("Error getting companies: " + clientRes.Status)
		return []CompanyModel{}, err
	}

	var decodedResponse StrapiResponse[[]CompanyModel]
	err = json.NewDecoder(clientRes.Body).Decode(&decodedResponse)
	if err != nil {
		log.Println(err.Error())
		return []CompanyModel{}, err
	}

	return decodedResponse.Data, nil
}

func GetCompany(id int, locale string) (CompanyModel, error) {
	var err error

	baseUrl := os.Getenv("BASE_URL")
	client := &http.Client{}

	query := "?" + strings.Join([]string{
		"locale=" + locale,
		"populate=*",
	}, "&")

	request, err := http.NewRequest("GET", fmt.Sprintf("%s/api/companies/%d%s", baseUrl, id, query), nil)
	if err != nil {
		return CompanyModel{}, err
	}
	request.Header.Add("Accept", "application/json")
	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("Authorization", "Bearer "+os.Getenv("API_TOKEN"))

	clientRes, err := client.Do(request)
	if err != nil {
		log.Println(err.Error())
		return CompanyModel{}, err
	}

	if clientRes.StatusCode != http.StatusOK {
		log.Println("Error getting companies: " + clientRes.Status)
		return CompanyModel{}, err
	}

	var decodedResponse StrapiResponse[CompanyModel]
	err = json.NewDecoder(clientRes.Body).Decode(&decodedResponse)
	if err != nil {
		log.Println(err.Error())
		return CompanyModel{}, err
	}

	return decodedResponse.Data, nil
}
