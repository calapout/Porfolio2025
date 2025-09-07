package Controllers

import (
	"Backend/Models"
	"encoding/json"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/gorilla/mux"
)

type ProjectsController struct{}

func singleProjectHandler(res http.ResponseWriter, req *http.Request) {
	// CHECK WE HAVE A VALID ID
	vars := mux.Vars(req)
	slug := vars["slug"]
	if slug == "" {
		res.WriteHeader(http.StatusBadRequest)
		return
	}

	res.WriteHeader(http.StatusOK)
	_, err := res.Write(make([]byte, 0))
	if err != nil {
		return
	}
}

func multipleProjectsHandler(res http.ResponseWriter, req *http.Request) {
	locale := req.URL.Query().Get("locale")

	if locale == "" {
		res.WriteHeader(http.StatusBadRequest)

		_, err := res.Write(CreateErrorResponse("locale is missing"))
		if err != nil {
			log.Println("Error writing error message:")
			return
		}
		return
	}

	baseUrl := os.Getenv("BASE_URL")
	client := &http.Client{}

	query := "?" + strings.Join([]string{
		"locale=" + locale,
		"populate=*",
	}, "&")

	request, err := http.NewRequest("GET", baseUrl+"/api/projects"+query, nil)
	if err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		return
	}
	request.Header.Add("Accept", "application/json")
	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("Authorization", "Bearer "+os.Getenv("API_TOKEN"))

	clientRes, err := client.Do(request)
	if err != nil {
		log.Println(err.Error())
		res.WriteHeader(http.StatusInternalServerError)
		return
	}

	resBuffer := StrapiResponse[[]Models.ProjectModel]{}
	err = json.NewDecoder(clientRes.Body).Decode(&resBuffer)
	if err != nil {
		log.Println(err.Error())
		res.WriteHeader(http.StatusInternalServerError)
		return
	}

	responseBuffer, err := json.Marshal(resBuffer.Data)
	if err != nil {
		log.Println(err.Error())
		res.WriteHeader(http.StatusInternalServerError)
		return
	}
	res.WriteHeader(http.StatusOK)
	_, err = res.Write(responseBuffer)
	if err != nil {
		return
	}
}

func (ProjectsController) Setup(r *mux.Router) {
	r.HandleFunc("/api/v1/projects", multipleProjectsHandler).Methods(http.MethodGet)
	r.HandleFunc("/api/v1/projects/{slug:[A-Za-z]+(?:-[A-Za-z]+)*}", singleProjectHandler).Methods(http.MethodGet)
}
