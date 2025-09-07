package Controllers

import (
	"encoding/json"
	"log"

	"github.com/gorilla/mux"
)

type Controller interface {
	Setup(r *mux.Router)
}

type ErrorStruct struct {
	Error string
}

type StrapiResponse[T any] struct {
	Data T
	Meta struct {
		Pagination struct {
			Page      int
			PageSize  int
			PageCount int
			Total     int
		}
	}
}

func CreateErrorResponse(ErrorMessage string) []byte {
	bytes, err := json.Marshal(ErrorStruct{Error: ErrorMessage})
	if err != nil {
		log.Println(err.Error())
		return make([]byte, 0)
	}
	return bytes
}
