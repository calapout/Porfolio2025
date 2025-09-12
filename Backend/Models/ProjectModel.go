package Models

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
)

type ProjectModel struct {
	Title           string       `json:"Title"`
	Locale          string       `json:"locale"`
	Description     string       `json:"Description"`
	TaskRealized    string       `json:"TaskRealized"`
	Technologies    []string     `json:"Technologies"`
	HasPrizes       bool         `json:"HasPrizes"`
	Slug            string       `json:"Slug"`
	IsMadeInCompany bool         `json:"IsMadeInCompany"`
	Thumbnail       Media        `json:"Thumbnail"`
	Medias          []Media      `json:"Media"`
	Prizes          []PrizeModel `json:"prizes"`
	Company         CompanyModel `json:"company"`
	IsNewAndTrendy  bool         `json:"IsNewAndTrendy"`
	IsFavorite      bool         `json:"IsFavorite"`
	FavoriteText    string       `json:"FavoriteText"`
}

func GetProjects(locale string) ([]ProjectModel, error) {
	var err error

	baseUrl := os.Getenv("BASE_URL")
	client := &http.Client{}

	query := "?" + strings.Join([]string{
		"locale=" + locale,
		"populate=*",
	}, "&")

	request, err := http.NewRequest("GET", fmt.Sprintf("%s/api/projects%s", baseUrl, query), nil)
	if err != nil {
		return []ProjectModel{}, err
	}
	request.Header.Add("Accept", "application/json")
	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("Authorization", "Bearer "+os.Getenv("API_TOKEN"))

	clientRes, err := client.Do(request)
	if err != nil {
		log.Println(err.Error())
		return []ProjectModel{}, err
	}

	if clientRes.StatusCode != http.StatusOK {
		log.Println("Error getting projects: " + clientRes.Status)
		return []ProjectModel{}, err
	}

	var decodedResponse StrapiResponse[[]ProjectModel]
	err = json.NewDecoder(clientRes.Body).Decode(&decodedResponse)
	if err != nil {
		log.Println(err.Error())
		return []ProjectModel{}, err
	}

	return decodedResponse.Data, nil
}

func GetProjectById(id int, locale string) (ProjectModel, error) {
	var err error

	baseUrl := os.Getenv("BASE_URL")
	client := &http.Client{}

	query := "?" + strings.Join([]string{
		"locale=" + locale,
		"populate=*",
	}, "&")

	request, err := http.NewRequest("GET", fmt.Sprintf("%s/api/projects/%d%s", baseUrl, id, query), nil)
	if err != nil {
		return ProjectModel{}, err
	}
	request.Header.Add("Accept", "application/json")
	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("Authorization", "Bearer "+os.Getenv("API_TOKEN"))

	clientRes, err := client.Do(request)
	if err != nil {
		log.Println(err.Error())
		return ProjectModel{}, err
	}

	if clientRes.StatusCode != http.StatusOK {
		log.Println("Error getting projects: " + clientRes.Status)
		return ProjectModel{}, err
	}

	var decodedResponse StrapiResponse[ProjectModel]
	err = json.NewDecoder(clientRes.Body).Decode(&decodedResponse)
	if err != nil {
		log.Println(err.Error())
		return ProjectModel{}, err
	}

	return decodedResponse.Data, nil
}

func GetProjectBySlug(slug string, locale string) (ProjectModel, error) {
	var err error

	baseUrl := os.Getenv("BASE_URL")
	client := &http.Client{}

	query := "?" + strings.Join([]string{
		"locale=" + locale,
		"filters[Slug]=" + slug,
		"populate=*",
	}, "&")

	request, err := http.NewRequest("GET", fmt.Sprintf("%s/api/projects%s", baseUrl, query), nil)
	if err != nil {
		return ProjectModel{}, err
	}
	request.Header.Add("Accept", "application/json")
	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("Authorization", "Bearer "+os.Getenv("API_TOKEN"))

	clientRes, err := client.Do(request)
	if err != nil {
		log.Println(err.Error())
		return ProjectModel{}, err
	}

	if clientRes.StatusCode != http.StatusOK {
		log.Println("Error getting projects: " + clientRes.Status)
		return ProjectModel{}, err
	}

	var decodedResponse StrapiResponse[[]ProjectModel]
	err = json.NewDecoder(clientRes.Body).Decode(&decodedResponse)
	if err != nil {
		log.Println(err.Error())
		return ProjectModel{}, err
	}

	if len(decodedResponse.Data) == 0 {
		return ProjectModel{}, NotFoundError{}
	}

	return decodedResponse.Data[0], nil
}
