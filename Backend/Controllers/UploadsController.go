package Controllers

import (
	"io"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

type UploadsController struct{}

func UploadsHandler(res http.ResponseWriter, req *http.Request) {
	// CHECK WE HAVE A VALID path
	vars := mux.Vars(req)
	path := vars["path"]
	if path == "" {
		log.Println("Invalid path")
		res.WriteHeader(http.StatusBadRequest)
		res.Write(CreateErrorResponse("Invalid path"))
		return
	}

	baseUrl := os.Getenv("BASE_URL")
	client := &http.Client{}
	request, err := http.NewRequest("GET", baseUrl+"/uploads/"+path, nil)
	if err != nil {
		log.Println(err.Error())
		res.WriteHeader(http.StatusInternalServerError)
		res.Write(make([]byte, 0))
		return
	}
	request.Header.Add("Accept", "application/json")
	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("Authorization", "Bearer "+os.Getenv("API_TOKEN"))

	clientRes, err := client.Do(request)
	if err != nil {
		log.Println(err.Error())
		res.WriteHeader(http.StatusInternalServerError)
		res.Write(make([]byte, 0))
		return
	}

	bytes, err := io.ReadAll(clientRes.Body)
	if err != nil {
		log.Println(err.Error())
		res.WriteHeader(http.StatusInternalServerError)
		res.Write(make([]byte, 0))
		return
	}

	res.WriteHeader(http.StatusOK)
	write, err := res.Write(bytes)
	if err != nil {
		log.Println(err.Error())
	}

	if write != len(bytes) {
		log.Println("Incomplete write")
		return
	}
}

func (u UploadsController) Setup(r *mux.Router) {
	r.HandleFunc("/api/v1/uploads/{path}", UploadsHandler)
}
